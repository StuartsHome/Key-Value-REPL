package commands

import (
	"fmt"
	"io"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
	"github.com/StuartsHome/key-value-REPL/cmd/errors"
)

var HelpName = "HELP"

func NewHelpCommand(globals *config.Data) *HelpCommand {
	var c HelpCommand
	c.Globals = globals
	c.name = HelpName

	return &c
}

type HelpCommand struct {
	config.Base

	name string
}

func (c *HelpCommand) Name() string {
	return c.name
}

func (c *HelpCommand) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	if len(args) > 0 {
		stderr.Write([]byte(errors.ErrIncorrectNumberArguments.Error()))
		return nil
	}

	appName := c.Globals.AppName

	out.Write([]byte(fmt.Sprintf("%s> Available commands:\n", appName)))
	out.Write([]byte(fmt.Sprintf("%s> HELP	- Displays Help\n", appName)))
	out.Write([]byte(fmt.Sprintf("%s> READ	- Reads value associated with key\n", appName)))
	out.Write([]byte(fmt.Sprintf("%s> WRITE	- Stores key with value\n", appName)))
	out.Write([]byte(fmt.Sprintf("%s> DELETE	- Deletes key and value associated with key\n", appName)))
	out.Write([]byte(fmt.Sprintf("%s> START	- Starts a transaction\n", appName)))
	out.Write([]byte(fmt.Sprintf("%s> COMMIT	- Commits a transaction\n", appName)))
	out.Write([]byte(fmt.Sprintf("%s> ABORT	- Removes current transaction and all associated data\n", appName)))
	out.Write([]byte(fmt.Sprintf("%s> READALL- Displays all keys and associated values in current transaction\n", appName)))
	out.Write([]byte(fmt.Sprintf("%s> QUIT	- Exits %s\n", appName, appName)))
	return nil
}
