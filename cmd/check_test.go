package cmd

import (
	"bytes"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"io"
	"strings"
	"testing"
)

func Test_CheckCmd(t *testing.T) {
	// Arrange
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "check"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	// Act
	err := rootCmd.Execute()

	// Assert
	cobblerclient.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	stdoutString := string(stdoutBytes)
	if !strings.Contains(stdoutString, "No configuration problems found.") &&
		!strings.Contains(stdoutString, "The following are potential configuration items") {
		t.Fatalf("expected one of the known check-command headers in output, got: %q", stdoutString)
	}
}
