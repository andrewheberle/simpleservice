// The package simplecommand reduces to amount of boilerplate code requried to use [simplecobra]
package simplecommand

import (
	"context"

	"github.com/bep/simplecobra"
)

// Command is the basis for creating your own [simplecobra.Commander] quickly
type Command struct {
	CommandName string
	Short       string
	Long        string
	Deprecated  string

	InitFunction   func(c simplecobra.Commander, cd *simplecobra.Commandeer) error
	PreRunFunction func(c simplecobra.Commander, this, runner *simplecobra.Commandeer) error
	RunFunction    func(c simplecobra.Commander, ctx context.Context, cd *simplecobra.Commandeer, args []string) error

	SubCommands []simplecobra.Commander
}

// New creates a bare minimum *Command with a name and a short description set
func New(name, short string) *Command {
	return &Command{CommandName: name, Short: short}
}

func (c *Command) Name() string {
	return c.CommandName
}

func (c *Command) Commands() []simplecobra.Commander {
	return c.SubCommands
}

func (c *Command) Init(cd *simplecobra.Commandeer) error {
	cmd := cd.CobraCommand
	cmd.Short = c.Short
	cmd.Long = c.Long
	cmd.Deprecated = c.Deprecated

	if c.InitFunction != nil {
		return c.InitFunction(c, cd)
	}

	return nil
}

func (c *Command) PreRun(this, runner *simplecobra.Commandeer) error {
	if c.PreRunFunction != nil {
		return c.PreRunFunction(c, this, runner)
	}

	return nil
}

func (c *Command) Run(ctx context.Context, cd *simplecobra.Commandeer, args []string) error {
	if c.RunFunction != nil {
		return c.RunFunction(c, ctx, cd, args)
	}

	return nil
}
