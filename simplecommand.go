// The package simplecommand reduces to amount of boilerplate code required to use [simplecobra]
// as it provides a [*Command] type that satisfies the [simplecobra.Commander] that you can embed
// within your own custom type and implement your own [*Command.Init], [*Command.PreRun] and
// [*Command.Run] methods as required.
package simplecommand

import (
	"context"
	"strings"

	"github.com/andrewheberle/simpleviper"
	"github.com/bep/simplecobra"
)

// Command is the basis for creating your own [simplecobra.Commander] quickly.
// A [*Command] satisfies the [simplecobra.Commander] interface and is best used
// by embedding it in your own struct.
type Command struct {
	// CommandName is used as the commands name for any help pages
	CommandName string

	// Short, Long and Depreciated are set to the commands short and long descriptions for help pages when using
	// the default Init method however when implementing your own Init method
	// you should set this yourself.
	Short      string
	Long       string
	Deprecated string

	// Config is used to specify a configuration file used when WithViper is passed to New.
	// Without [WithViper] being passed to New, this has no effect.
	Config string

	// Allow missing config file when Config is set and the WithViper CommandOption is
	// provided to New
	ConfigOptional bool

	// SubCommands holds the list of sub-commands for this command
	SubCommands []simplecobra.Commander

	viperEnabled        bool
	viperEnvPrefix      string
	viperEnvKeyReplacer *strings.Replacer
}

// New creates a bare minimum [*Command] with a name and a short description set
func New(name, short string, opts ...CommandOption) *Command {
	c := &Command{
		CommandName: name,
		Short:       short,
	}

	// set options
	for _, o := range opts {
		o(c)
	}

	return c
}

func (c *Command) Name() string {
	return c.CommandName
}

func (c *Command) Commands() []simplecobra.Commander {
	return c.SubCommands
}

// Init is where the short and long description of the command are set and also where command line flags can be handled.
// The default is only suitable for implementing a deprecated command (see the [Deprecated] [CommandOption]) or a command
// that does not make use of any command line flags.
//
// See [simplecobra.Commander] for more information.
func (c *Command) Init(cd *simplecobra.Commandeer) error {
	cmd := cd.CobraCommand
	cmd.Short = c.Short
	cmd.Long = c.Long
	cmd.Deprecated = c.Deprecated

	return nil
}

// PreRun is where command line flags have been parsed, so is a place for any initialisation would go for the command.
// The default is only suitable for implementing a command that has no reliance on internal state such as command line flags.
//
// An exception to this is if [New] was called with the [WithViper] option, which will bind any defined flags
// in your custom [*Command.Init] method to an internal [*viper.Viper] instance to enable those flags to be set via
// environment variables or a provided configuration file.
//
// See [simplecobra.Commander] for more information.
func (c *Command) PreRun(this, runner *simplecobra.Commandeer) error {
	if c.viperEnabled {
		opts := []simpleviper.Option{
			simpleviper.WithEnvPrefix(c.viperEnvPrefix),
			simpleviper.WithEnvKeyReplacer(c.viperEnvKeyReplacer),
		}

		// add config file if set
		if c.Config != "" {
			if c.ConfigOptional {
				opts = append(opts, simpleviper.WithOptionalConfig(c.Config))
			} else {
				opts = append(opts, simpleviper.WithConfig(c.Config))
			}
		}

		// bring in env vars and bind to flagset
		err := simpleviper.New(opts...).Init(this.CobraCommand.Flags())
		if err != nil {
			return err
		}

	}

	return nil
}

// Run is where the command actually does it's work
// The default does no actual work, so is likely not suitable for any use case except for possibly a deprecated command.
//
// See [simplecobra.Commander] for more information.
func (c *Command) Run(ctx context.Context, cd *simplecobra.Commandeer, args []string) error {
	return nil
}

// A CommandOption is passed to [New] to change the defaults of the [*Command]
type CommandOption func(*Command)

// Long sets the long description of the command when the default [*Command.Init] is used.
func Long(description string) CommandOption {
	return func(c *Command) {
		c.Long = description
	}
}

// Deprecated sets command as deprecated when the default [*Command.Init] is used.
func Deprecated(reason string) CommandOption {
	return func(c *Command) {
		c.Deprecated = reason
	}
}

// By passing WithViper to [New], this will enable binding of any defined command
// line flags to an internal [*viper.Viper] instance using [simpleviper] so they
// can be set via environment variables or a configuration file.
func WithViper(prefix string, replacer *strings.Replacer) CommandOption {
	return func(c *Command) {
		c.viperEnabled = true
		c.viperEnvPrefix = prefix
		c.viperEnvKeyReplacer = replacer
	}
}
