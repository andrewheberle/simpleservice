// The package simplecommand reduces to amount of boilerplate code requried to use [simplecobra]
// as it provides a [*Command] type that satisfies the [simplecobra.Commander] that you can embed
// within your own custom type and implement your own [Init], [PreRun] and [Run] methods as required.
package simplecommand

import (
	"context"

	"github.com/bep/simplecobra"
)

// Command is the basis for creating your own [simplecobra.Commander] quickly.
// A [*Command] satisfies the [simplecobra.Commander] interface and is best used
// by embedding it in your own struct.
type Command struct {
	CommandName string
	Short       string
	Long        string
	Deprecated  string

	SubCommands []simplecobra.Commander
}

// New creates a bare minimum *Command with a name and a short description set
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

// Init is where the short and long description of the command are set and also where command line flags can be handled
func (c *Command) Init(cd *simplecobra.Commandeer) error {
	cmd := cd.CobraCommand
	cmd.Short = c.Short
	cmd.Long = c.Long
	cmd.Deprecated = c.Deprecated

	return nil
}

// PreRun is where command line flags have been parsed, so is a place for any initialisation would go for the command
func (c *Command) PreRun(this, runner *simplecobra.Commandeer) error {
	return nil
}

// Run is where the command actually does it's work
func (c *Command) Run(ctx context.Context, cd *simplecobra.Commandeer, args []string) error {
	return nil
}

type CommandOption func(*Command)

func Long(description string) CommandOption {
	return func(c *Command) {
		c.Long = description
	}
}

func Deprecated(reason string) CommandOption {
	return func(c *Command) {
		c.Deprecated = reason
	}
}
