package commands_test

import (
	"bytes"
	"testing"

	"github.com/StuartsHome/key-value-REPL/cmd/commands"
	"github.com/StuartsHome/key-value-REPL/cmd/config"
	"github.com/StuartsHome/key-value-REPL/cmd/datastore"

	"github.com/stretchr/testify/require"
)

func TestHelp_Exec_Success(t *testing.T) {

	// Given
	var stdout bytes.Buffer
	ts := datastore.NewDataStore()

	globals := config.NewData("test", ts)
	testHelpCommand := commands.NewHelpCommand(globals)

	// When
	args := []string{}
	err := testHelpCommand.Exec(nil, &stdout, nil, args)
	require.Nil(t, err)

}
