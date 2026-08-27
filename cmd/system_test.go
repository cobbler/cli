package cmd

import (
	"bytes"
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"io"
	"strings"
	"testing"
)

func createSystem(client cobbler.Client, name string) (*cobbler.System, error) {
	// Profile must be the referenced profile's real uid, not its name.
	profileHandle, err := client.GetProfileHandle("Ubuntu-20.04-x86_64")
	if err != nil {
		return nil, err
	}
	profile, err := client.GetProfile(profileHandle, false, false)
	if err != nil {
		return nil, err
	}
	system := cobbler.NewSystem()
	system.Name = name
	system.Profile = profile.Uid
	return client.CreateSystem(system)
}

func removeSystem(client cobbler.Client, name string) error {
	handle, err := client.GetSystemHandle(name)
	if err != nil {
		return err
	}
	return client.DeleteSystem(handle)
}

func Test_SystemAddCmd(t *testing.T) {
	type args struct {
		command []string
	}
	setupClient(t)
	// --profile must be the referenced profile's real uid, not its name.
	profileHandle, err := Client.GetProfileHandle("Ubuntu-20.04-x86_64")
	cobbler.FailOnError(t, err)
	profile, err := Client.GetProfile(profileHandle, false, false)
	cobbler.FailOnError(t, err)
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "plain",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "system", "add", "--name", "test-plain", "--profile", profile.Uid}},
			want:    "System test-plain created",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeSystem(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
			stdout := bytes.NewBufferString("")
			stderr := bytes.NewBufferString("")
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			// Act
			err = rootCmd.Execute()

			// Assert
			cobbler.FailOnError(t, err)
			FailOnNonEmptyStream(t, stderr)
			stdoutBytes, err := io.ReadAll(stdout)
			if err != nil {
				t.Fatal(err)
			}
			stdoutString := string(stdoutBytes)
			if !strings.Contains(stdoutString, tt.want) {
				fmt.Println(stdoutString)
				t.Fatal("Item creation message missing")
			}
		})
	}
}

func Test_SystemCopyCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "system", "copy", "--name", "system-to-copy", "--newname", "copied-system"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeSystem(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
				cleanupErr = removeSystem(Client, tt.args.command[7])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createSystem(Client, tt.args.command[5])
			cobbler.FailOnError(t, err)
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
			stdout := bytes.NewBufferString("")
			stderr := bytes.NewBufferString("")
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			// Act
			err = rootCmd.Execute()

			// Assert
			cobbler.FailOnError(t, err)
			FailOnNonEmptyStream(t, stderr)
			FailOnNonEmptyStream(t, stdout)
			copiedHandle, err := Client.GetSystemHandle(tt.args.command[7])
			cobbler.FailOnError(t, err)
			_, err = Client.GetSystem(copiedHandle, false, false)
			cobbler.FailOnError(t, err)
		})
	}
}

func Test_SystemEditCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "system", "edit", "--name", "test-system-edit", "--comment", "testcomment"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeSystem(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createSystem(Client, tt.args.command[5])
			cobbler.FailOnError(t, err)
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
			stdout := bytes.NewBufferString("")
			stderr := bytes.NewBufferString("")
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			// Act
			err = rootCmd.Execute()

			// Assert
			cobbler.FailOnError(t, err)
			FailOnNonEmptyStream(t, stderr)
			FailOnNonEmptyStream(t, stdout)
			editedHandle, err := Client.GetSystemHandle(tt.args.command[5])
			cobbler.FailOnError(t, err)
			updatedSystem, err := Client.GetSystem(editedHandle, false, false)
			cobbler.FailOnError(t, err)
			if updatedSystem.Comment != "testcomment" {
				t.Fatal("system update wasn't successful")
			}
		})
	}
}

// Test_SystemCopyCmd_UID exercises the --uid sibling flag on copy.
func Test_SystemCopyCmd_UID(t *testing.T) {
	name := "test-system-copy-uid"
	newName := "test-system-copied-uid"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	t.Cleanup(func() {
		if err := removeSystem(Client, newName); err != nil {
			t.Errorf("cleanup: remove system %s: %v", newName, err)
		}
	})
	uid, err := Client.GetSystemHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "copy", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetSystemHandle(newName)
	cobbler.FailOnError(t, err)
	_, err = Client.GetSystem(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

// Test_SystemEditCmd_UID exercises the --uid sibling flag on edit.
func Test_SystemEditCmd_UID(t *testing.T) {
	name := "test-system-edit-uid"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	uid, err := Client.GetSystemHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "edit", "--uid", uid, "--comment", "testcomment-uid"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updatedSystem, err := Client.GetSystem(uid, false, false)
	cobbler.FailOnError(t, err)
	if updatedSystem.Comment != "testcomment-uid" {
		t.Fatal("system update via --uid wasn't successful")
	}
}

// Test_SystemDumpVarsCmd exercises the dumpvars command.
func Test_SystemDumpVarsCmd(t *testing.T) {
	name := "test-system-dumpvars"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "dumpvars", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(stdoutBytes), "name: "+name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("name missing from dumpvars output")
	}
}

// Test_SystemDumpVarsCmd_UID exercises the --uid sibling flag on dumpvars.
func Test_SystemDumpVarsCmd_UID(t *testing.T) {
	name := "test-system-dumpvars-uid"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	uid, err := Client.GetSystemHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "dumpvars", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(stdoutBytes), "name: "+name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("name missing from dumpvars --uid output")
	}
}

// Test_SystemGetAutoinstallCmd exercises the get-autoinstall command, which
// now resolves the system by uid internally (Client.GenerateAutoinstall is
// called with "uid" instead of "name").
func Test_SystemGetAutoinstallCmd(t *testing.T) {
	name := "test-system-get-autoinstall"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "get-autoinstall", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if len(stdoutBytes) == 0 {
		t.Fatal("get-autoinstall produced no output")
	}
}

// Test_SystemGetAutoinstallCmd_UID exercises the --uid sibling flag on
// get-autoinstall.
func Test_SystemGetAutoinstallCmd_UID(t *testing.T) {
	name := "test-system-get-autoinstall-uid"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	uid, err := Client.GetSystemHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "get-autoinstall", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if len(stdoutBytes) == 0 {
		t.Fatal("get-autoinstall --uid produced no output")
	}
}

// Test_SystemGetAutoinstallCmd_NotFound exercises the error path when the
// named system does not exist.
func Test_SystemGetAutoinstallCmd_NotFound(t *testing.T) {
	setupClient(t)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "get-autoinstall", "--name", "does-not-exist-system"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err := rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error for a non-existent system")
	}
}

// Test_SystemPowerOffCmd exercises the poweroff command's resolveUID and
// Client.PowerSystem call. The test system has no power management
// configured, so the RPC call is expected to fail server-side -- the point
// of this test is to exercise the resolution + RPC call path.
func Test_SystemPowerOffCmd(t *testing.T) {
	name := "test-system-poweroff"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "poweroff", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error since the test system has no power management configured")
	}
	if !strings.Contains(err.Error(), "command failed") {
		t.Fatalf("unexpected error from poweroff: %v", err)
	}
}

// Test_SystemPowerOffCmd_UID exercises the --uid sibling flag on poweroff.
func Test_SystemPowerOffCmd_UID(t *testing.T) {
	name := "test-system-poweroff-uid"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	uid, err := Client.GetSystemHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "poweroff", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error since the test system has no power management configured")
	}
	if !strings.Contains(err.Error(), "command failed") {
		t.Fatalf("unexpected error from poweroff --uid: %v", err)
	}
}

// Test_SystemPowerOnCmd exercises the poweron command's resolveUID and
// Client.PowerSystem call.
func Test_SystemPowerOnCmd(t *testing.T) {
	name := "test-system-poweron"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "poweron", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error since the test system has no power management configured")
	}
	if !strings.Contains(err.Error(), "command failed") {
		t.Fatalf("unexpected error from poweron: %v", err)
	}
}

// Test_SystemPowerStatusCmd exercises the powerstatus command's resolveUID
// and Client.PowerSystem call.
func Test_SystemPowerStatusCmd(t *testing.T) {
	name := "test-system-powerstatus"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "powerstatus", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error since the test system has no power management configured")
	}
	if !strings.Contains(err.Error(), "command failed") {
		t.Fatalf("unexpected error from powerstatus: %v", err)
	}
}

// Test_SystemRebootCmd exercises the reboot command's resolveUID and
// Client.PowerSystem call.
func Test_SystemRebootCmd(t *testing.T) {
	name := "test-system-reboot"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "reboot", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error since the test system has no power management configured")
	}
	if !strings.Contains(err.Error(), "command failed") {
		t.Fatalf("unexpected error from reboot: %v", err)
	}
}

func Test_SystemEditCmd_VirtUEFI(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "system", "edit", "--name", "test-system-edit-virt-uefi", "--virt-uefi=true"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeSystem(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createSystem(Client, tt.args.command[5])
			cobbler.FailOnError(t, err)
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
			stdout := bytes.NewBufferString("")
			stderr := bytes.NewBufferString("")
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			// Act
			err = rootCmd.Execute()

			// Assert
			cobbler.FailOnError(t, err)
			FailOnNonEmptyStream(t, stderr)
			FailOnNonEmptyStream(t, stdout)
			editedHandle, err := Client.GetSystemHandle(tt.args.command[5])
			cobbler.FailOnError(t, err)
			updatedSystem, err := Client.GetSystem(editedHandle, false, false)
			cobbler.FailOnError(t, err)
			if !updatedSystem.Virt.UEFI {
				t.Fatal("system virt-uefi update wasn't successful")
			}
		})
	}
}

func Test_SystemFindCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "system", "find", "--name", "test-system-find"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			systemName := "test-system-find"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeSystem(Client, systemName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createSystem(Client, systemName)
			cobbler.FailOnError(t, err)
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
			stdout := bytes.NewBufferString("")
			stderr := bytes.NewBufferString("")
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			// Act
			err = rootCmd.Execute()

			// Assert
			cobbler.FailOnError(t, err)
			FailOnNonEmptyStream(t, stderr)
			stdoutBytes, err := io.ReadAll(stdout)
			if err != nil {
				t.Fatal(err)
			}
			stdoutString := string(stdoutBytes)
			if !strings.Contains(stdoutString, systemName) {
				fmt.Println(stdoutString)
				t.Fatal("system not successfully found")
			}
		})
	}
}

func Test_SystemListCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "system", "list"}},
			want:    "systems:",
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
			cobbler.FailOnError(t, err)
			FailOnNonEmptyStream(t, stderr)
			stdoutBytes, err := io.ReadAll(stdout)
			if err != nil {
				t.Fatal(err)
			}
			stdoutString := string(stdoutBytes)
			if !strings.Contains(stdoutString, tt.want) {
				fmt.Println(stdoutString)
				t.Fatal("system list marker not located in output")
			}
		})
	}
}

func Test_SystemRemoveCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "system", "remove", "--name", "test-system-remove"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			setupClient(t)
			_, err := createSystem(Client, tt.args.command[5])
			cobbler.FailOnError(t, err)
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
			stdout := bytes.NewBufferString("")
			stderr := bytes.NewBufferString("")
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			// Act
			err = rootCmd.Execute()

			// Assert
			cobbler.FailOnError(t, err)
			FailOnNonEmptyStream(t, stderr)
			FailOnNonEmptyStream(t, stdout)
			result, err := Client.HasItem("system", tt.args.command[5])
			cobbler.FailOnError(t, err)
			if result {
				// A missing item means we get "false", as such we error when we find an item.
				t.Fatal("system not successfully removed")
			}
		})
	}
}

// Test_SystemRemoveCmd_UID exercises the --uid sibling flag on remove.
func Test_SystemRemoveCmd_UID(t *testing.T) {
	name := "test-system-remove-uid"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetSystemHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "remove", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("system", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("system not successfully removed via --uid")
	}
}

func Test_SystemRenameCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "system", "rename", "--name", "test-system-rename", "--newname", "test-system-renamed"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			systemName := "test-system-rename"
			newSystemName := "test-system-renamed"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeSystem(Client, newSystemName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createSystem(Client, systemName)
			cobbler.FailOnError(t, err)
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
			stdout := bytes.NewBufferString("")
			stderr := bytes.NewBufferString("")
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			// Act
			err = rootCmd.Execute()

			// Assert
			cobbler.FailOnError(t, err)
			FailOnNonEmptyStream(t, stderr)
			FailOnNonEmptyStream(t, stdout)
			resultOldName, err := Client.HasItem("system", systemName)
			cobbler.FailOnError(t, err)
			if resultOldName {
				t.Fatal("system not successfully renamed (old name present)")
			}
			resultNewName, err := Client.HasItem("system", newSystemName)
			cobbler.FailOnError(t, err)
			if !resultNewName {
				t.Fatal("system not successfully renamed (new name not present)")
			}
		})
	}
}

func Test_SystemReportCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "system", "report", "--name", "test-system-report"}},
			want:    ": test-system-report",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			systemName := "test-system-report"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeSystem(Client, systemName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createSystem(Client, systemName)
			cobbler.FailOnError(t, err)
			cobra.OnInitialize(initConfig, setupLogger)
			rootCmd := NewRootCmd()
			rootCmd.SetArgs(tt.args.command)
			stdout := bytes.NewBufferString("")
			stderr := bytes.NewBufferString("")
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			// Act
			err = rootCmd.Execute()

			// Assert
			cobbler.FailOnError(t, err)
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

// Test_SystemRenameCmd_UID exercises the --uid sibling flag on rename.
func Test_SystemRenameCmd_UID(t *testing.T) {
	name := "test-system-rename-uid"
	newName := "test-system-renamed-uid"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetSystemHandle(name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, newName); err != nil {
			t.Errorf("cleanup: remove system %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "rename", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("system", name)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("system not successfully renamed via --uid (old name present)")
	}
	resultNewName, err := Client.HasItem("system", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("system not successfully renamed via --uid (new name not present)")
	}
}

// Test_SystemReportCmd_UID exercises the --uid sibling flag on report.
func Test_SystemReportCmd_UID(t *testing.T) {
	name := "test-system-report-uid"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	uid, err := Client.GetSystemHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "report", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(stdoutBytes), name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("system name missing from report --uid output")
	}
}

// Test_SystemReportCmd_All exercises the report branch taken when neither
// --name nor --uid is supplied (report every system).
func Test_SystemReportCmd_All(t *testing.T) {
	name := "test-system-report-all"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "report"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(stdoutBytes), name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("system name missing from report --all output")
	}
}

// Test_SystemExportCmd exercises the export command's json branch with an
// explicit --name.
func Test_SystemExportCmd(t *testing.T) {
	name := "test-system-export"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "export", "--name", name, "--format", "json"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(stdoutBytes), `"name":"`+name+`"`) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("system name missing from json export output")
	}
}

// Test_SystemExportCmd_UID exercises the export command's --uid sibling
// flag.
func Test_SystemExportCmd_UID(t *testing.T) {
	name := "test-system-export-uid"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	uid, err := Client.GetSystemHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "export", "--uid", uid, "--format", "json"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(stdoutBytes), `"name":"`+name+`"`) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("system name missing from json export --uid output")
	}
}

// Test_SystemExportCmd_All exercises the export branch taken when neither
// --name nor --uid is supplied, using the yaml format.
func Test_SystemExportCmd_All(t *testing.T) {
	name := "test-system-export-all"
	setupClient(t)
	_, err := createSystem(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, name); err != nil {
			t.Errorf("cleanup: remove system %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system", "export", "--format", "yaml"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	stdoutBytes, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(stdoutBytes), "name: "+name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("system name missing from yaml export --all output")
	}
}
