package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
)

var ReadName = "READ"

func NewReadCommand(globals *config.Data) *ReadCommand {
	var c ReadCommand
	c.Globals = globals
	c.name = ReadName

	return &c
}

type ReadCommand struct {
	config.Base

	name string
}

func (c *ReadCommand) Name() string {
	return c.name
}

func (c *ReadCommand) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	if len(args) < 1 || len(args) > 1 {
		stderr.Write([]byte("no value to read!\n"))
		return nil
	}

	val := args[0]
	got, err := c.Globals.DataStore.St.Get(val, c.Globals.DataStore)
	if err != nil {
		errMessage := fmt.Sprintf("error: key not found\n")
		os.Stderr.Write([]byte(errMessage))
		return nil
	}

	out.Write([]byte(got + "\n"))
	return nil
}
