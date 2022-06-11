package commands

import (
	"io"
	"os"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
	"github.com/StuartsHome/key-value-REPL/cmd/errors"
)

var WriteName = "WRITE"

func NewWriteCommand(globals *config.Data) *WriteCommand {
	var c WriteCommand
	c.Globals = globals
	c.name = WriteName

	return &c
}

type WriteCommand struct {
	config.Base

	name string
}

func (c *WriteCommand) Name() string {
	return c.name
}

func (c *WriteCommand) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	if len(args) < 2 || len(args) > 2 {
		err := errors.ErrorIncorrectNumberArguments().Error()
		stderr.Write([]byte(err))
		return nil
	}

	key := args[0]
	val := args[1]

	if err := c.Globals.DataStore.St.Set(key, val, c.Globals.DataStore); err != nil {
		message := []byte(err.Error())
		os.Stderr.Write(message)
		return nil
	}

	return nil
}
