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

func TestWrite_Exec_Success(t *testing.T) {
	// Given
	var stdout bytes.Buffer
	ts := datastore.NewDataStore()

	globals := config.NewData("test", ts)
	testWriteCommand := commands.NewWriteCommand(globals)

	// When
	args := []string{"10", "11"}
	err := testWriteCommand.Exec(nil, &stdout, nil, args)
	require.Nil(t, err)

	// Then - updated store.
	got, err := ts.St.Get("10", ts)
	require.Nil(t, err)

	// Then
	expectedOutput := "11"
	assert.Equal(t, expectedOutput, got)

}

func TestWrite_Exec_Error_TooManyArgs(t *testing.T) {
	// Given
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	ts := datastore.NewDataStore()

	globals := config.NewData("test", ts)
	testWriteCommand := commands.NewWriteCommand(globals)

	// When
	args := []string{"1", "2", "3"}
	err := testWriteCommand.Exec(nil, &stdout, &stderr, args)
	require.Nil(t, err)

	// Then - Error on stderr.
	errOutput := stderr.String()
	expectedStdErr := errors.ErrorIncorrectNumberArguments().Error()
	assert.Equal(t, expectedStdErr, errOutput)

	// Then - Empty Store.
	got, err := ts.St.Get("1", ts)
	require.Empty(t, got)

	expectedError := fmt.Errorf("key 1 not set in global store")
	assert.EqualError(t, err, expectedError.Error())
}

func TestWrite_Exec_Error_TooFewArgs(t *testing.T) {
	// Given
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	ts := datastore.NewDataStore()

	globals := config.NewData("test", ts)
	testWriteCommand := commands.NewWriteCommand(globals)

	// When
	args := []string{"1"}
	err := testWriteCommand.Exec(nil, &stdout, &stderr, args)
	require.Nil(t, err)

	// Then - Error on stderr.
	errOutput := stderr.String()
	expectedStdErr := errors.ErrorIncorrectNumberArguments().Error()
	assert.Equal(t, expectedStdErr, errOutput)

	// Then - Empty Store.
	got, err := ts.St.Get("1", ts)
	require.Empty(t, got)

	expectedError := fmt.Errorf("key 1 not set in global store")
	assert.EqualError(t, err, expectedError.Error())
}

func TestWrite_Name_Success(t *testing.T) {
	// Given
	ts := datastore.NewDataStore()
	globals := config.NewData("test", ts)
	testWriteCommand := commands.NewWriteCommand(globals)

	// When
	name := testWriteCommand.Name()

	// Then
	expectedName := commands.WriteName
	assert.Equal(t, expectedName, name)
}
