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

func createDistro(client cobbler.Client, name string) (*cobbler.Distro, error) {
	distro := cobbler.NewDistro()
	distro.Name = name
	distro.Kernel = "/extracted_iso_image/install/vmlinuz"
	distro.Initrd = "/extracted_iso_image/install/initrd.gz"
	return client.CreateDistro(distro)
}

func removeDistro(client cobbler.Client, name string) error {
	handle, err := client.GetDistroHandle(name)
	if err != nil {
		return err
	}
	return client.DeleteDistro(handle)
}

func Test_DistroAddCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "distro", "add", "--name", "test-plain", "--kernel", "/extracted_iso_image/install/vmlinuz", "--initrd", "/extracted_iso_image/install/initrd.gz"}},
			want:    "Distro test-plain created",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeDistro(Client, tt.args.command[5])
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

func Test_DistroCopyCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "distro", "copy", "--name", "distro-to-copy", "--newname", "copied-distro"}},
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
				cleanupErr := removeDistro(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
				cleanupErr = removeDistro(Client, tt.args.command[7])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createDistro(Client, tt.args.command[5])
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
			copiedHandle, err := Client.GetDistroHandle(tt.args.command[7])
			cobbler.FailOnError(t, err)
			_, err = Client.GetDistro(copiedHandle, false, false)
			cobbler.FailOnError(t, err)
		})
	}
}

func Test_DistroEditCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "distro", "edit", "--name", "test-distro-edit", "--comment", "testcomment"}},
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
				cleanupErr := removeDistro(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createDistro(Client, tt.args.command[5])
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
			editedHandle, err := Client.GetDistroHandle(tt.args.command[5])
			cobbler.FailOnError(t, err)
			updatedDistro, err := Client.GetDistro(editedHandle, false, false)
			cobbler.FailOnError(t, err)
			if updatedDistro.Comment != "testcomment" {
				t.Fatal("distro update wasn't successful")
			}
		})
	}
}

// Test_DistroCopyCmd_UID exercises the --uid sibling flag on copy.
func Test_DistroCopyCmd_UID(t *testing.T) {
	name := "test-distro-copy-uid"
	newName := "test-distro-copied-uid"
	setupClient(t)
	_, err := createDistro(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistro(Client, name); err != nil {
			t.Errorf("cleanup: remove distro %s: %v", name, err)
		}
	})
	t.Cleanup(func() {
		if err := removeDistro(Client, newName); err != nil {
			t.Errorf("cleanup: remove distro %s: %v", newName, err)
		}
	})
	uid, err := Client.GetDistroHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro", "copy", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetDistroHandle(newName)
	cobbler.FailOnError(t, err)
	_, err = Client.GetDistro(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

// Test_DistroEditCmd_UID exercises the --uid sibling flag on edit.
func Test_DistroEditCmd_UID(t *testing.T) {
	name := "test-distro-edit-uid"
	setupClient(t)
	_, err := createDistro(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistro(Client, name); err != nil {
			t.Errorf("cleanup: remove distro %s: %v", name, err)
		}
	})
	uid, err := Client.GetDistroHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro", "edit", "--uid", uid, "--comment", "testcomment-uid"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updatedDistro, err := Client.GetDistro(uid, false, false)
	cobbler.FailOnError(t, err)
	if updatedDistro.Comment != "testcomment-uid" {
		t.Fatal("distro update via --uid wasn't successful")
	}
}

func Test_DistroEditCmd_SourceTreePath(t *testing.T) {
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
			// The path must be an absolute path that exists on the server; /extracted_iso_image
			// is bind-mounted into the test Cobbler container by testing/compose.yml.
			name:    "plain",
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "distro", "edit", "--name", "test-distro-edit-source-tree-path", "--source-tree-path", "/extracted_iso_image"}},
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
				cleanupErr := removeDistro(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createDistro(Client, tt.args.command[5])
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
			editedHandle, err := Client.GetDistroHandle(tt.args.command[5])
			cobbler.FailOnError(t, err)
			updatedDistro, err := Client.GetDistro(editedHandle, false, false)
			cobbler.FailOnError(t, err)
			if updatedDistro.SourceTreePath != "/extracted_iso_image" {
				t.Fatal("distro source-tree-path update wasn't successful")
			}
		})
	}
}

func Test_DistroFindCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "distro", "find", "--name", "test-distro-find"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			distroName := "test-distro-find"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeDistro(Client, distroName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createDistro(Client, distroName)
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
			if !strings.Contains(stdoutString, distroName) {
				fmt.Println(stdoutString)
				t.Fatal("distro not successfully found")
			}
		})
	}
}

func Test_DistroListCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "distro", "list"}},
			want:    "distros:",
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
				t.Fatal("distro list marker not located in output")
			}
		})
	}
}

func Test_DistroRemoveCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "distro", "remove", "--name", "test-distro-remove"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			setupClient(t)
			_, err := createDistro(Client, tt.args.command[5])
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
			result, err := Client.HasItem("distro", tt.args.command[5])
			cobbler.FailOnError(t, err)
			if result {
				// A missing item means we get "false", as such we error when we find an item.
				t.Fatal("distro not successfully removed")
			}
		})
	}
}

// Test_DistroRemoveCmd_UID exercises the --uid sibling flag added alongside
// --name so a target distro can be identified by its Cobbler UID instead of
// its name.
func Test_DistroRemoveCmd_UID(t *testing.T) {
	distroName := "test-distro-remove-uid"
	t.Run("uid", func(t *testing.T) {
		// Arrange
		setupClient(t)
		_, err := createDistro(Client, distroName)
		cobbler.FailOnError(t, err)
		distroUID, err := Client.GetDistroHandle(distroName)
		cobbler.FailOnError(t, err)
		cobra.OnInitialize(initConfig, setupLogger)
		rootCmd := NewRootCmd()
		rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro", "remove", "--uid", distroUID})
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
		result, err := Client.HasItem("distro", distroName)
		cobbler.FailOnError(t, err)
		if result {
			// A missing item means we get "false", as such we error when we find an item.
			t.Fatal("distro not successfully removed via --uid")
		}
	})
}

// Test_DistroRenameCmd_UID exercises the --uid sibling flag on rename.
func Test_DistroRenameCmd_UID(t *testing.T) {
	name := "test-distro-rename-uid"
	newName := "test-distro-renamed-uid"
	setupClient(t)
	_, err := createDistro(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetDistroHandle(name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistro(Client, newName); err != nil {
			t.Errorf("cleanup: remove distro %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro", "rename", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("distro", name)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("distro not successfully renamed via --uid (old name present)")
	}
	resultNewName, err := Client.HasItem("distro", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("distro not successfully renamed via --uid (new name not present)")
	}
}

func Test_DistroRenameCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "distro", "rename", "--name", "test-distro-rename", "--newname", "test-distro-renamed"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			distroName := "test-distro-rename"
			newDistroName := "test-distro-renamed"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeDistro(Client, newDistroName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createDistro(Client, distroName)
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
			resultOldName, err := Client.HasItem("distro", distroName)
			cobbler.FailOnError(t, err)
			if resultOldName {
				t.Fatal("distro not successfully renamed (old name present)")
			}
			resultNewName, err := Client.HasItem("distro", newDistroName)
			cobbler.FailOnError(t, err)
			if !resultNewName {
				t.Fatal("distro not successfully renamed (new name not present)")
			}
		})
	}
}

func Test_DistroReportCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "distro", "report", "--name", "test-distro-report"}},
			want:    ": test-distro-report",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			distroName := "test-distro-report"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeDistro(Client, distroName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createDistro(Client, distroName)
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

// Test_DistroReportCmd_UID exercises the --uid sibling flag on report.
func Test_DistroReportCmd_UID(t *testing.T) {
	name := "test-distro-report-uid"
	setupClient(t)
	_, err := createDistro(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistro(Client, name); err != nil {
			t.Errorf("cleanup: remove distro %s: %v", name, err)
		}
	})
	uid, err := Client.GetDistroHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro", "report", "--uid", uid})
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
		t.Fatal("distro name missing from report --uid output")
	}
}

// Test_DistroReportCmd_All exercises the report branch taken when neither
// --name nor --uid is supplied (report every distro).
func Test_DistroReportCmd_All(t *testing.T) {
	name := "test-distro-report-all"
	setupClient(t)
	_, err := createDistro(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistro(Client, name); err != nil {
			t.Errorf("cleanup: remove distro %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro", "report"})
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
		t.Fatal("distro name missing from report --all output")
	}
}

// Test_DistroExportCmd exercises the export command's json branch with an
// explicit --name.
func Test_DistroExportCmd(t *testing.T) {
	name := "test-distro-export"
	setupClient(t)
	_, err := createDistro(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistro(Client, name); err != nil {
			t.Errorf("cleanup: remove distro %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro", "export", "--name", name, "--format", "json"})
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
		t.Fatal("distro name missing from json export output")
	}
}

// Test_DistroExportCmd_UID exercises the export command's --uid sibling
// flag.
func Test_DistroExportCmd_UID(t *testing.T) {
	name := "test-distro-export-uid"
	setupClient(t)
	_, err := createDistro(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistro(Client, name); err != nil {
			t.Errorf("cleanup: remove distro %s: %v", name, err)
		}
	})
	uid, err := Client.GetDistroHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro", "export", "--uid", uid, "--format", "json"})
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
		t.Fatal("distro name missing from json export --uid output")
	}
}

// Test_DistroExportCmd_All exercises the export branch taken when neither
// --name nor --uid is supplied, using the yaml format.
func Test_DistroExportCmd_All(t *testing.T) {
	name := "test-distro-export-all"
	setupClient(t)
	_, err := createDistro(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistro(Client, name); err != nil {
			t.Errorf("cleanup: remove distro %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro", "export", "--format", "yaml"})
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
		t.Fatal("distro name missing from yaml export --all output")
	}
}
