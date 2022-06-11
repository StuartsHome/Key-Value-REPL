package commands_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/StuartsHome/key-value-REPL/cmd/commands"
	"github.com/StuartsHome/key-value-REPL/cmd/config"
	"github.com/StuartsHome/key-value-REPL/cmd/datastore"
	"github.com/StuartsHome/key-value-REPL/cmd/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuit_Exec_Success(t *testing.T) {

	// Given
	var stdout bytes.Buffer
	ts := datastore.NewDataStore()

	globals := config.NewData("test", ts)
	testQuitCommand := commands.NewQuitCommand(globals)

	// When
	args := []string{}
	err := testQuitCommand.Exec(nil, &stdout, nil, args)

	// Then
	assert.Equal(t, fmt.Errorf("STOP").Error(), err.Error())
}

func TestQuit_Exec_Fail_TooManyArguments(t *testing.T) {
	// Given
	var stderr bytes.Buffer
	ts := datastore.NewDataStore()

	globals := config.NewData("test", ts)
	testQuitCommand := commands.NewQuitCommand(globals)

	// When
	args := []string{"1"}
	err := testQuitCommand.Exec(nil, nil, &stderr, args)

	// Then
	require.Nil(t, err)

	// Then - error on stderr.
	expectedErr := errors.ErrIncorrectNumberArguments.Error()
	assert.Equal(t, expectedErr, stderr.String())
}

func TestQuit_Name_Success(t *testing.T) {
	// Given
	ts := datastore.NewDataStore()
	globals := config.NewData("test", ts)
	testQuitCommand := commands.NewQuitCommand(globals)

	// When
	name := testQuitCommand.Name()

	// Then
	expectedName := commands.QuitName
	assert.Equal(t, expectedName, name)

}
