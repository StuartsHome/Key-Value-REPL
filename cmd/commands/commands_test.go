package commands_test

import (
	"io"
	"testing"

	"github.com/StuartsHome/key-value-REPL/cmd/commands"
	"github.com/StuartsHome/key-value-REPL/cmd/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type QuiterTester struct {
	config.Base

	name string
}

func (c *QuiterTester) Name() string {
	return "testCommand"
}
func (c *QuiterTester) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	stderr.Write([]byte("STOP"))
	return nil
}

func TestSelectCommand_Success(t *testing.T) {
	// Given
	name := []string{"testCommand"}

	testQuiterCommand := &QuiterTester{name: name[0]}
	cmds := []commands.Command{testQuiterCommand}

	// When
	cmd, args, found := commands.SelectCommand(name, cmds)

	// Then
	expectedArgs := []string(nil)
	expectedCmd := testQuiterCommand
	assert.True(t, found)
	assert.Equal(t, expectedArgs, args)
	assert.Equal(t, expectedCmd, cmd)
}

func TestSelectCommand_Fail_EmptyCommands(t *testing.T) {
	// Given
	name := []string{"testCommand"}
	cmds := []commands.Command{}

	// When
	cmd, args, found := commands.SelectCommand(name, cmds)

	// Then
	expectedArgs := []string(nil)
	require.Nil(t, cmd)
	assert.Equal(t, expectedArgs, args)
	assert.False(t, found)
}

type WriterTester struct {
	config.Base

	name string
}

func (c *WriterTester) Name() string {
	return "testCommand"
}
func (c *WriterTester) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	return nil
}

func TestSelectCommand_Success_WithArguments(t *testing.T) {
	// Given
	name := []string{"testCommand", "arg1", "arg2"}

	testWriterCommand := &WriterTester{name: "testCommand"}

	cmds := []commands.Command{testWriterCommand}

	// When
	cmd, args, found := commands.SelectCommand(name, cmds)

	// Then
	expectedArgs := []string{"arg1", "arg2"}
	expectedCmd := testWriterCommand
	assert.True(t, found)
	assert.Equal(t, expectedArgs, args)
	assert.Equal(t, expectedCmd, cmd)
}
