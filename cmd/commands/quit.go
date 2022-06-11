package commands

import (
	"fmt"
	"io"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
	"github.com/StuartsHome/key-value-REPL/cmd/errors"
)

const QuitName = "QUIT"

func NewQuitCommand(globals *config.Data) *QuitCommand {
	var c QuitCommand
	c.Globals = globals
	c.name = QuitName

	return &c
}

type QuitCommand struct {
	config.Base

	name string
}

func (c *QuitCommand) Name() string {
	return c.name
}

func (c *QuitCommand) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	if len(args) > 0 {
		err := errors.ErrIncorrectNumberArguments.Error()
		stderr.Write([]byte(err))
		return nil
	}

	return fmt.Errorf("STOP")
}
