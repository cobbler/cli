package cmd

import (
	"bytes"
	"fmt"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_parseSettingValue_IntBounds(t *testing.T) {
	// Arrange
	target := struct {
		I8  int8
		I64 int64
	}{}
	v := reflect.ValueOf(&target).Elem()

	// Act & Assert: value out of range for the field's actual bit size must error, not wrap silently.
	if _, err := parseSettingValue(v.FieldByName("I8"), "1000"); err == nil {
		t.Fatal("expected error for value out of int8 range, got none")
	}

	// A value that fits should still parse correctly, preserving the field's actual type.
	got, err := parseSettingValue(v.FieldByName("I8"), "100")
	cobblerclient.FailOnError(t, err)
	if got != int8(100) {
		t.Fatalf("expected int8(100), got %v (%T)", got, got)
	}

	got, err = parseSettingValue(v.FieldByName("I64"), "9223372036854775807")
	cobblerclient.FailOnError(t, err)
	if got != int64(9223372036854775807) {
		t.Fatalf("expected max int64, got %v (%T)", got, got)
	}
}

func Test_SettingEditCmd(t *testing.T) {
	// Arrange
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "setting", "edit"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	// Act
	err := rootCmd.Execute()

	// Assert
	if err == nil {
		t.Fatal("expected error, got none")
	}
	if err.Error() != "dynamic settings are turned off server-side" {
		t.Fatalf("expected dynamic settings are to be turned off server-side, got %s", err.Error())
	}
}

func Test_SettingReportCmd(t *testing.T) {
	// Arrange
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "setting", "report"})
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
	if !strings.Contains(stdoutString, "scm_track_enabled") {
		fmt.Println(stdoutString)
		t.Fatal("Expected setting couldn't be found")
	}
}

// Test_SettingReportCmd_Name exercises the --name flag, which must filter the
// report down to the single named setting instead of dumping the full list
// (regression test for the --name flag being parsed but silently ignored).
func Test_SettingReportCmd_Name(t *testing.T) {
	// Arrange
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "setting", "report", "--name", "scm_track_enabled"})
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
	if !strings.Contains(stdoutString, "scm_track_enabled") {
		fmt.Println(stdoutString)
		t.Fatal("Expected setting couldn't be found")
	}
	// The whole point of --name is to narrow the output to one setting, not
	// dump everything. Pick an unrelated setting that should NOT appear.
	if strings.Contains(stdoutString, "allow_duplicate_hostnames") {
		fmt.Println(stdoutString)
		t.Fatal("expected --name to filter the report to a single setting, but the full list was printed")
	}
	lines := strings.Split(strings.TrimSpace(stdoutString), "\n")
	if len(lines) != 1 {
		fmt.Println(stdoutString)
		t.Fatalf("expected exactly one line of output for --name, got %d", len(lines))
	}
}

// Test_SettingReportCmd_NameUnknown exercises the --name flag with a setting
// name that doesn't exist, which should error out rather than silently
// printing the full list.
func Test_SettingReportCmd_NameUnknown(t *testing.T) {
	// Arrange
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "setting", "report", "--name", "does_not_exist"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	// Act
	err := rootCmd.Execute()

	// Assert
	if err == nil {
		t.Fatal("expected error for unknown setting name, got none")
	}
}
