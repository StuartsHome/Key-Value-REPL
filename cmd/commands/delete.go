package commands

import (
	"io"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
	"github.com/StuartsHome/key-value-REPL/cmd/errors"
)

var DeleteName = "DELETE"

func NewDeleteCommand(globals *config.Data) *DeleteCommand {
	var c DeleteCommand
	c.Globals = globals
	c.name = DeleteName

	return &c
}

type DeleteCommand struct {
	config.Base

	name string
}

func (c *DeleteCommand) Name() string {
	return c.name
}

func (c *DeleteCommand) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	if len(args) < 1 || len(args) > 1 {
		stderr.Write([]byte(errors.ErrIncorrectNumberArguments.Error()))
		return nil
	}

	key := args[0]

	if err := c.Globals.DataStore.St.Delete(key, c.Globals.DataStore); err != nil {
		stderr.Write([]byte(errors.ErrorKeyNotFound(key).Error()))
		return nil
	}

	return nil
}
