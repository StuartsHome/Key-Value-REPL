package commands

import (
	"io"

	"github.com/StuartsHome/key-value-REPL/cmd/config"
)

type Command interface {
	Name() string
	Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error
}

func SelectCommand(name []string, commands []Command) (Command, []string, bool) {
	for _, command := range commands {
		if len(name) > 0 && command.Name() == name[0] {
			if len(name) > 1 {
				return command, name[1:], true
			}
			return command, nil, true
		}
	}

	return nil, nil, false
}

// DefineCommands specifies all commands to run
// with the application.
func DefineCommands(globals *config.Data) []Command {
	writeCommand := NewWriteCommand(globals)
	readCommand := NewReadCommand(globals)
	deleteCommand := NewDeleteCommand(globals)
	startCommand := NewStartCommand(globals)
	commitCommand := NewCommitCommand(globals)
	abortCommand := NewAbortCommand(globals)
	quitCommand := NewQuitCommand(globals)
	readallcommand := NewReadAllCommand(globals)
	helpCommand := NewHelpCommand(globals)

	return []Command{
		writeCommand,
		readCommand,
		deleteCommand,
		startCommand,
		commitCommand,
		abortCommand,
		quitCommand,
		readallcommand,
		helpCommand,
	}
}
