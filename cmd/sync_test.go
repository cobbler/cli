package cmd

import (
	"bytes"
	"fmt"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"io"
	"strings"
	"testing"
	"time"
)

func Test_SyncCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "sync"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			name:    "dns",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "sync", "--dns"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			name:    "dhcp",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "sync", "--dhcp"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			name:    "dhcpdns",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "sync", "--dns", "--dhcp"}},
			want:    "Event ID:",
			wantErr: false,
		},
		{
			name:    "systems",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "sync", "--systems", "a.b.c,a.d.c"}},
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

			// Cleanup - Sleep after each test to let dhcpd restart properly
			time.Sleep(1 * time.Second)
		})
	}
}
