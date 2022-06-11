package commands

import (
	"io"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
	"github.com/StuartsHome/key-value-REPL/cmd/errors"
)

var StartName = "START"

func NewStartCommand(globals *config.Data) *StartCommand {
	var c StartCommand
	c.Globals = globals
	c.name = StartName

	return &c
}

type StartCommand struct {
	config.Base

	name string
}

func (c *StartCommand) Name() string {
	return c.name
}

func (c *StartCommand) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	if len(args) > 0 {
		stderr.Write([]byte(errors.ErrIncorrectNumberArguments.Error()))
	}

	// Start a new Transaction.
	c.Globals.DataStore.Tr.PushTransaction()
	return nil
}
