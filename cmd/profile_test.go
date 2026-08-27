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

func createProfile(client cobbler.Client, name string) (*cobbler.Profile, error) {
	// Distro must be the referenced distro's real uid, not its name.
	distroHandle, err := client.GetDistroHandle("Ubuntu-20.04-x86_64")
	if err != nil {
		return nil, err
	}
	distro, err := client.GetDistro(distroHandle, false, false)
	if err != nil {
		return nil, err
	}
	profile := cobbler.NewProfile()
	profile.Name = name
	profile.Distro = distro.Uid
	return client.CreateProfile(profile)
}

func removeProfile(client cobbler.Client, name string) error {
	handle, err := client.GetProfileHandle(name)
	if err != nil {
		return err
	}
	return client.DeleteProfile(handle)
}

func Test_ProfileAddCmd(t *testing.T) {
	type args struct {
		command []string
	}
	setupClient(t)
	// --distro must be the referenced distro's real uid, not its name.
	distroHandle, err := Client.GetDistroHandle("Ubuntu-20.04-x86_64")
	cobbler.FailOnError(t, err)
	distro, err := Client.GetDistro(distroHandle, false, false)
	cobbler.FailOnError(t, err)
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "plain",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "profile", "add", "--name", "test-plain", "--distro", distro.Uid}},
			want:    "Profile test-plain created",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeProfile(Client, tt.args.command[5])
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

func Test_ProfileCopyCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "profile", "copy", "--name", "profile-to-copy", "--newname", "copied-profile"}},
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
				cleanupErr := removeProfile(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
				cleanupErr = removeProfile(Client, tt.args.command[7])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createProfile(Client, tt.args.command[5])
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
			copiedHandle, err := Client.GetProfileHandle(tt.args.command[7])
			cobbler.FailOnError(t, err)
			_, err = Client.GetProfile(copiedHandle, false, false)
			cobbler.FailOnError(t, err)
		})
	}
}

func Test_ProfileEditCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "profile", "edit", "--name", "test-profile-edit", "--comment", "testcomment"}},
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
				cleanupErr := removeProfile(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createProfile(Client, tt.args.command[5])
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
			editedHandle, err := Client.GetProfileHandle(tt.args.command[5])
			cobbler.FailOnError(t, err)
			updatedProfile, err := Client.GetProfile(editedHandle, false, false)
			cobbler.FailOnError(t, err)
			if updatedProfile.Comment != "testcomment" {
				t.Fatal("profile update wasn't successful")
			}
		})
	}
}

func Test_ProfileEditCmd_VirtUEFI(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "profile", "edit", "--name", "test-profile-edit-virt-uefi", "--virt-uefi=true"}},
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
				cleanupErr := removeProfile(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createProfile(Client, tt.args.command[5])
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
			editedHandle, err := Client.GetProfileHandle(tt.args.command[5])
			cobbler.FailOnError(t, err)
			updatedProfile, err := Client.GetProfile(editedHandle, false, false)
			cobbler.FailOnError(t, err)
			if !updatedProfile.Virt.UEFI {
				t.Fatal("profile virt-uefi update wasn't successful")
			}
		})
	}
}

// Test_ProfileCopyCmd_UID exercises the --uid sibling flag on copy.
func Test_ProfileCopyCmd_UID(t *testing.T) {
	name := "test-profile-copy-uid"
	newName := "test-profile-copied-uid"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, name); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", name, err)
		}
	})
	t.Cleanup(func() {
		if err := removeProfile(Client, newName); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", newName, err)
		}
	})
	uid, err := Client.GetProfileHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "copy", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetProfileHandle(newName)
	cobbler.FailOnError(t, err)
	_, err = Client.GetProfile(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

// Test_ProfileEditCmd_UID exercises the --uid sibling flag on edit.
func Test_ProfileEditCmd_UID(t *testing.T) {
	name := "test-profile-edit-uid"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, name); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", name, err)
		}
	})
	uid, err := Client.GetProfileHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "edit", "--uid", uid, "--comment", "testcomment-uid"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updatedProfile, err := Client.GetProfile(uid, false, false)
	cobbler.FailOnError(t, err)
	if updatedProfile.Comment != "testcomment-uid" {
		t.Fatal("profile update via --uid wasn't successful")
	}
}

// Test_ProfileDumpVarsCmd exercises the dumpvars command.
func Test_ProfileDumpVarsCmd(t *testing.T) {
	name := "test-profile-dumpvars"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, name); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "dumpvars", "--name", name})
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
	if !strings.Contains(string(stdoutBytes), "profile_name: "+name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("profile_name missing from dumpvars output")
	}
}

// Test_ProfileDumpVarsCmd_UID exercises the --uid sibling flag on dumpvars.
func Test_ProfileDumpVarsCmd_UID(t *testing.T) {
	name := "test-profile-dumpvars-uid"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, name); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", name, err)
		}
	})
	uid, err := Client.GetProfileHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "dumpvars", "--uid", uid})
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
	if !strings.Contains(string(stdoutBytes), "profile_name: "+name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("profile_name missing from dumpvars --uid output")
	}
}

// Test_ProfileGetAutoinstallCmd exercises the get-autoinstall command, which
// now resolves the profile by uid internally (Client.GenerateAutoinstall is
// called with "uid" instead of "name").
func Test_ProfileGetAutoinstallCmd(t *testing.T) {
	setupClient(t)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "get-autoinstall", "--name", "Ubuntu-20.04-x86_64"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err := rootCmd.Execute()

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

// Test_ProfileGetAutoinstallCmd_UID exercises the --uid sibling flag on
// get-autoinstall.
func Test_ProfileGetAutoinstallCmd_UID(t *testing.T) {
	setupClient(t)
	uid, err := Client.GetProfileHandle("Ubuntu-20.04-x86_64")
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "get-autoinstall", "--uid", uid})
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

// Test_ProfileGetAutoinstallCmd_NotFound exercises the error path when
// neither --name nor a matching profile exists.
func Test_ProfileGetAutoinstallCmd_NotFound(t *testing.T) {
	setupClient(t)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "get-autoinstall", "--name", "does-not-exist-profile"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err := rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error for a non-existent profile")
	}
}

func Test_ProfileFindCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "profile", "find", "--name", "test-profile-find"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			profileName := "test-profile-find"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeProfile(Client, profileName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createProfile(Client, profileName)
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
			if !strings.Contains(stdoutString, profileName) {
				fmt.Println(stdoutString)
				t.Fatal("profile not successfully found")
			}
		})
	}
}

func Test_ProfileListCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "profile", "list"}},
			want:    "profiles:",
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
				t.Fatal("profile list marker not located in output")
			}
		})
	}
}

func Test_ProfileRemoveCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "profile", "remove", "--name", "test-profile-remove"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			setupClient(t)
			_, err := createProfile(Client, tt.args.command[5])
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
			result, err := Client.HasItem("profile", tt.args.command[5])
			cobbler.FailOnError(t, err)
			if result {
				// A missing item means we get "false", as such we error when we find an item.
				t.Fatal("profile not successfully removed")
			}
		})
	}
}

// Test_ProfileRemoveCmd_UID exercises the --uid sibling flag on remove.
func Test_ProfileRemoveCmd_UID(t *testing.T) {
	name := "test-profile-remove-uid"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetProfileHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "remove", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("profile", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("profile not successfully removed via --uid")
	}
}

func Test_ProfileRenameCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "profile", "rename", "--name", "test-profile-rename", "--newname", "test-profile-renamed"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			profileName := "test-profile-rename"
			newProfileName := "test-profile-renamed"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeProfile(Client, newProfileName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createProfile(Client, profileName)
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
			resultOldName, err := Client.HasItem("profile", profileName)
			cobbler.FailOnError(t, err)
			if resultOldName {
				t.Fatal("profile not successfully renamed (old name present)")
			}
			resultNewName, err := Client.HasItem("profile", newProfileName)
			cobbler.FailOnError(t, err)
			if !resultNewName {
				t.Fatal("profile not successfully renamed (new name not present)")
			}
		})
	}
}

func Test_ProfileReportCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "profile", "report", "--name", "test-profile-report"}},
			want:    ": test-profile-report",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			profileName := "test-profile-report"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeProfile(Client, profileName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createProfile(Client, profileName)
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

// Test_ProfileRenameCmd_UID exercises the --uid sibling flag on rename.
func Test_ProfileRenameCmd_UID(t *testing.T) {
	name := "test-profile-rename-uid"
	newName := "test-profile-renamed-uid"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetProfileHandle(name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, newName); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "rename", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("profile", name)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("profile not successfully renamed via --uid (old name present)")
	}
	resultNewName, err := Client.HasItem("profile", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("profile not successfully renamed via --uid (new name not present)")
	}
}

// Test_ProfileReportCmd_UID exercises the --uid sibling flag on report.
func Test_ProfileReportCmd_UID(t *testing.T) {
	name := "test-profile-report-uid"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, name); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", name, err)
		}
	})
	uid, err := Client.GetProfileHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "report", "--uid", uid})
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
		t.Fatal("profile name missing from report --uid output")
	}
}

// Test_ProfileReportCmd_All exercises the report branch taken when neither
// --name nor --uid is supplied (report every profile).
func Test_ProfileReportCmd_All(t *testing.T) {
	name := "test-profile-report-all"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, name); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "report"})
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
		t.Fatal("profile name missing from report --all output")
	}
}

// Test_ProfileExportCmd exercises the export command's json branch with an
// explicit --name.
func Test_ProfileExportCmd(t *testing.T) {
	name := "test-profile-export"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, name); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "export", "--name", name, "--format", "json"})
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
		t.Fatal("profile name missing from json export output")
	}
}

// Test_ProfileExportCmd_UID exercises the export command's --uid sibling
// flag.
func Test_ProfileExportCmd_UID(t *testing.T) {
	name := "test-profile-export-uid"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, name); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", name, err)
		}
	})
	uid, err := Client.GetProfileHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "export", "--uid", uid, "--format", "json"})
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
		t.Fatal("profile name missing from json export --uid output")
	}
}

// Test_ProfileExportCmd_All exercises the export branch taken when neither
// --name nor --uid is supplied, using the yaml format.
func Test_ProfileExportCmd_All(t *testing.T) {
	name := "test-profile-export-all"
	setupClient(t)
	_, err := createProfile(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, name); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile", "export", "--format", "yaml"})
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
		t.Fatal("profile name missing from yaml export --all output")
	}
}
