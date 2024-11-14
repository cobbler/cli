package cmd

import (
	"bytes"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"testing"
)

func Test_StatusCmd(t *testing.T) {
	// Arrange
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "status"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)
	expectedResult := "ip             |target              |start            |state            \n"

	// Act
	err := rootCmd.Execute()

	// Assert
	cobblerclient.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	if stdout.String() != expectedResult {
		t.Errorf(`Expected "%s", got "%s"`, expectedResult, stdout.String())
	}
}
