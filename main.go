package main

import (
	"flag"
	"log"
	"os"

	"github.com/StuartsHome/key-value-REPL/cmd/app"
	"github.com/StuartsHome/key-value-REPL/cmd/commands"
	"github.com/StuartsHome/key-value-REPL/cmd/config"
	"github.com/StuartsHome/key-value-REPL/cmd/datastore"
)

// Usage example:
// go run main.go --name="web-app"

func main() {
	var name string
	flag.StringVar(&name, "name", "web-app", "The name of the app, if no name is provided it defaults to backend-go")
	flag.Parse()

	// Create datastore.
	datastore := datastore.NewDataStore()

	// Create global data.
	globals := config.NewData(name, datastore)

	// Define Commands.
	cmds := commands.DefineCommands(globals)

	// Define client data.
	opts := app.NewClientOpts(name, os.Stdin, os.Stdout, os.Stderr)

	// Run the app.
	err := opts.Run(cmds)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	log.Println("Finished!")
}
