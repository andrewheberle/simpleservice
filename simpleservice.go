// The package simpleservice reduces to amount of boilerplate code requried to use [simplecobra]
package simpleservice

import (
	"context"

	"github.com/bep/simplecobra"
)

// Command is the basis for creating your own [simplecobra.Commander] quickly
type Command struct {
	Short string
	Long  string

	InitFunction   func(c simplecobra.Commander, cd *simplecobra.Commandeer) error
	PreRunFunction func(c simplecobra.Commander, this, runner *simplecobra.Commandeer) error
	RunFunction    func(c simplecobra.Commander, ctx context.Context, cd *simplecobra.Commandeer, args []string) error

	name     string
	commands []simplecobra.Commander
}

func (c *Command) Name() string {
	return c.name
}

func (c *Command) Commands() []simplecobra.Commander {
	return c.commands
}

func (c *Command) Init(cd *simplecobra.Commandeer) error {
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
