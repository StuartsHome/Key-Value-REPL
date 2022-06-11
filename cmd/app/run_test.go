package app

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/StuartsHome/key-value-REPL/cmd/commands"
	"github.com/StuartsHome/key-value-REPL/cmd/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests Run method.
func TestApplicationRun_Success(t *testing.T) {
	// Given
	var stdin bytes.Buffer

	r, w, _ := os.Pipe()
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	opts := NewClientOpts("testApp", &stdin, w, nil)
	var cmds commands.Command

	// When
	err := opts.Run([]commands.Command{cmds})
	require.Nil(t, err)

	w.Close()
	out := <-outC

	// Then
	assert.Equal(t, expectedOutput, out)
}

type QuiterTester struct {
	config.Base

	name string
}

func (c *QuiterTester) Name() string {
	return "testCommand"
}
func (c *QuiterTester) Exec(in io.Reader, out io.Writer, stderr io.Writer, args []string) error {
	return fmt.Errorf("STOP")
}

func TestApplicationRun_Quit_Success(t *testing.T) {
	// Given
	var stdin bytes.Buffer
	stdin.Write([]byte("testCommand\n"))

	r, w, _ := os.Pipe()
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	opts := NewClientOpts("testApp", &stdin, w, w)
	name := []string{"testCommand"}

	testQuiterCommand := &QuiterTester{name: name[0]}
	cmds := []commands.Command{testQuiterCommand}

	// When
	err := opts.Run(cmds)
	require.Nil(t, err)

	w.Close()
	out := <-outC

	// Then
	assert.Equal(t, expectedOutput, out)
}

func TestSelectCommand_Fail(t *testing.T) {
	// Given
	name := []string{"testCommand"}
	cmds := []commands.Command{}

	// When
	cmd, args, found := commands.SelectCommand(name, cmds)
	require.Nil(t, cmd)
	require.Nil(t, args)

	// Then
	assert.False(t, found)
}

func TestStart_Success(t *testing.T) {
	// Given
	name := []string{"testCommand"}
	cmds := []commands.Command{}

	// When
	cmd, args, found := commands.SelectCommand(name, cmds)
	require.Nil(t, cmd)
	require.Nil(t, args)

	// Then
	assert.False(t, found)
}

func TestParseArgs_Success(t *testing.T) {
	// Given
	text := "this is a test"

	// When
	got := parseArgs(text)

	// Then
	expected := []string{"this", "is", "a", "test"}
	assert.Equal(t, expected, got)
}

var expectedOutput = `testApp> Welcome to testApp!
testApp> Type HELP for a list of commands.
testApp> `

/*

func TestApplicationRun_Success(t *testing.T) {
	// Given
	var stdout bytes.Buffer

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	opts := NewClientOpts("testApp", &stdin, w, nil)
	var cmds commands.Command

	// When
	err := opts.Run([]commands.Command{cmds})
	require.Nil(t, err)

	w.Close()
	// os.Stdout = old
	out := <-outC

	// Then
	assert.Equal(t, expectedOutput, out)
}*/
