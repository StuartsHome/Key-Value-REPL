package commands

import (
	"io"
	"os"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
)

const AbortName = "ABORT"

func NewAbortCommand(globals *config.Data) *AbortCommand {
	var c AbortCommand
	c.Globals = globals
	c.name = AbortName

	return &c
}

type AbortCommand struct {
	config.Base

	name string
}

func (c *AbortCommand) Name() string {
	return c.name
}

func (c *AbortCommand) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	if err := c.Globals.DataStore.Tr.PopTransaction(); err != nil {
		os.Stderr.Write([]byte(err.Error() + "\n"))
		return nil
	}
	return nil
}
