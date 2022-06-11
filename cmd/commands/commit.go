package commands

import (
	"io"
	"os"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
)

var CommitName = "COMMIT"

func NewCommitCommand(globals *config.Data) *CommitCommand {
	var c CommitCommand
	c.Globals = globals
	c.name = CommitName

	return &c
}

type CommitCommand struct {
	config.Base

	name string
}

func (c *CommitCommand) Name() string {
	return c.name
}

func (c *CommitCommand) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	if len(args) > 0 {
		stderr.Write([]byte("no value to set!\n"))
		return nil
	}

	if err := c.Globals.DataStore.Tr.Commit(); err != nil {
		message := []byte(err.Error() + "\n")
		os.Stderr.Write(message)
		return nil
	}

	return nil
}
