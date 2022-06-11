package commands

import (
	"fmt"
	"io"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
)

var ReadAllName = "READALL"

func NewReadAllCommand(globals *config.Data) *ReadAllCommand {
	var c ReadAllCommand
	c.Globals = globals
	c.name = ReadAllName

	return &c
}

type ReadAllCommand struct {
	config.Base

	name string
}

func (c *ReadAllCommand) Name() string {
	return c.name
}

func (c *ReadAllCommand) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	if len(args) > 0 {
		stderr.Write([]byte("too many arguments.\n"))
	}

	vals, err := c.Globals.DataStore.St.GetAll(c.Globals.DataStore)
	if err != nil {
		message := fmt.Sprintf("error: %v\n", err)
		stderr.Write([]byte(message))
		return nil
	}

	for _, val := range vals {
		k, v := val[0], val[1]
		out.Write([]byte("Key: " + k + " Value: " + v + "\n"))
	}

	return nil
}
