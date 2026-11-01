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

func Test_BuildisoStandaloneCmd(t *testing.T) {
	type args struct {
		command []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "plain",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "buildiso", "--standalone"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			name:    "airgapped",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "buildiso", "--airgapped"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
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
			if !strings.Contains(stdoutString, tt.want) {
				fmt.Println(stdoutString)
				t.Fatal("No Event ID present")
			}
		})
	}
}

func Test_BuildisoNetinstallCmd(t *testing.T) {
	type args struct {
		command []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "systems",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "buildiso", "--systems", "test"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			name:    "nodns",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "buildiso", "--systems", "test", "--exclude-dns"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			// Regression test: --exclude-systems used to not exist at all
			// (NewBuildisoCmd never registered it, even though
			// cobblerclient.BuildisoOptions.ExcludeSystems and the Python
			// API's build_iso(exclude_systems=...) already existed
			// end-to-end), so there was no way to narrow the system list
			// on a buildiso without explicitly naming every system to
			// include. This only exercises flag registration/wiring (the
			// RunE invocation would fail outright with "flag accessed but
			// not defined: exclude-systems" before the fix); it does not
			// assert anything about the generated ISO's contents.
			name:    "excludesystems",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "buildiso", "--profiles", "test", "--exclude-systems"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
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
			if !strings.Contains(stdoutString, tt.want) {
				fmt.Println(stdoutString)
				t.Fatal("No Event ID present")
			}
		})
	}
}
