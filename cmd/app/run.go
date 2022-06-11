package app

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/StuartsHome/key-value-REPL/cmd/commands"
	"github.com/StuartsHome/key-value-REPL/cmd/config"
)

type ClientOpts struct {
	AppName string
	Base    *config.Base

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewClientOpts(
	appName string,
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
) *ClientOpts {
	return &ClientOpts{
		AppName: appName,
		Base: &config.Base{
			Globals: &config.Data{
				AppName: appName,
			},
		},
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
}

// Run with the selected commands.
func (c *ClientOpts) Run(cmds []commands.Command) error {
	out := c.Stdout
	in := c.Stdin
	stderr := c.Stderr

	// Before we begin, print a welcome message.
	c.printWelcomeMessage(out)

	// Create a reader to read from stdin.
	reader := bufio.NewReader(in)

	// Loop.
	for {
		// text := []string{"WRITE", "10", "20"}

		// Parse new bytes on reader.
		text, err := parseInput(reader)
		if err != nil {
			break
		}

		// Find the matching command.
		command, args, found := commands.SelectCommand(text, cmds)
		if found {
			// Run the Exec method on the matching command.
			if err := command.Exec(in, out, stderr, args); err != nil {
				// The QUIT command returns a 'STOP' error upon successful completion.
				if err.Error() == "STOP" {
					return nil
				}
				return err
			}
		} else {
			// If the command was not found, warn the user.
			c.commandNotFound(text[0])
		}

		// Start a new line.
		c.printLineStarter(out)
	}

	return nil
}

// Read buffer and split into slice.
func parseInput(reader *bufio.Reader) ([]string, error) {
	text, err := readBuffer(reader)
	if err != nil {
		return nil, err
	}

	// Split text.
	splitText := parseArgs(text)

	return splitText, nil
}

// readBuffer reads the provided buffer;
// also trims any spaces.
func readBuffer(r *bufio.Reader) (string, error) {
	t, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(t), nil
}

// parseArgs splits the text input into
// a slice of strings.
func parseArgs(text string) []string {
	// Split text into slice.
	splitText := strings.Split(text, " ")
	return splitText
}

func (c *ClientOpts) commandNotFound(text string) {
	t := strings.TrimSuffix(text, "\n")
	if t != "" {
		fmt.Printf("%s> invalid command: %s. Try again!\n", c.AppName, t)
	}
}

// Welcome message.
func (c *ClientOpts) printWelcomeMessage(out io.Writer) {
	out.Write([]byte(fmt.Sprintf("%s> Welcome to %s!\n", c.AppName, c.AppName)))
	out.Write([]byte(fmt.Sprintf("%s> Type HELP for a list of commands.\n", c.AppName)))

	// New line.
	c.printLineStarter(out)
}

// New line.
func (c *ClientOpts) printLineStarter(out io.Writer) {
	message := fmt.Sprintf("%s> ", c.AppName)
	out.Write([]byte(message))
}

// func (c *ClientOpts) help() {
// 	fmt.Printf("%s> Welcome to %s!\n", c.AppName, c.AppName)
// 	fmt.Printf("%s> Available commands:\n", c.AppName)
// 	fmt.Printf("%s> HELP	- Displays Help\n", c.AppName)
// 	fmt.Printf("%s> READ	- Reads value associated with key\n", c.AppName)
// 	fmt.Printf("%s> WRITE	- Stores key with value\n", c.AppName)
// 	fmt.Printf("%s> DELETE	- Deletes key and value associated with key\n", c.AppName)
// 	fmt.Printf("%s> START	- Starts a transaction\n", c.AppName)
// 	fmt.Printf("%s> COMMIT	- Commits a transaction\n", c.AppName)
// 	fmt.Printf("%s> ABORT	- Removes current transaction and all associated data\n", c.AppName)
// 	fmt.Printf("%s> READALL	- Displays all keys and associated values in current transaction\n", c.AppName)
// 	fmt.Printf("%s> QUIT	- Exits %s\n", c.AppName, c.AppName)
// 	// c.printLineStarter()
// }

// // isQuit checks for quit.
// func isQuit(text string) bool {
// 	if strings.EqualFold("QUIT", text) {
// 		return false
// 	}
// 	return true
// }

// // TODO: is this needed?
// func recoverExp(text string) {
// 	if r := recover(); r != nil {
// 		fmt.Println("go-repl> unknown command ", text)
// 	}
// }
