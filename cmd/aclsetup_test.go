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

func Test_AclSetupCmd(t *testing.T) {
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
			name:    "adduser",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "aclsetup", "--adduser", "cobbler"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			name:    "addgroup",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "aclsetup", "--addgroup", "cobbler"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			name:    "removeuser",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "aclsetup", "--removeuser", "cobbler"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			name:    "removegroup",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "aclsetup", "--removegroup", "cobbler"}},
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
