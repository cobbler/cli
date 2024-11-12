package cmd

import (
	"bytes"
	"fmt"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"io"
	"strings"
	"testing"
)

func Test_ReportCmd(t *testing.T) {
	// Arrange
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "report"})
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
	if !(strings.Contains(stdoutString, "distros:") && strings.Contains(stdoutString, "profiles")) {
		fmt.Println(stdoutString)
		t.Fatal("no heading for distros and profiles present")
	}
}
