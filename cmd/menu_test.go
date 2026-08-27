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

func createMenu(client cobbler.Client, name string) (*cobbler.Menu, error) {
	menu := cobbler.NewMenu()
	menu.Name = name
	return client.CreateMenu(menu)
}

func removeMenu(client cobbler.Client, name string) error {
	handle, err := client.GetMenuHandle(name)
	if err != nil {
		return err
	}
	return client.DeleteMenu(handle)
}

func Test_MenuAddCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "menu", "add", "--name", "test-plain"}},
			want:    "Menu test-plain created",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeMenu(Client, tt.args.command[5])
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

func Test_MenuCopyCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "menu", "copy", "--name", "menu-to-copy", "--newname", "copied-menu"}},
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
				cleanupErr := removeMenu(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
				cleanupErr = removeMenu(Client, tt.args.command[7])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createMenu(Client, tt.args.command[5])
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
			copiedHandle, err := Client.GetMenuHandle(tt.args.command[7])
			cobbler.FailOnError(t, err)
			_, err = Client.GetMenu(copiedHandle, false, false)
			cobbler.FailOnError(t, err)
		})
	}
}

func Test_MenuEditCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "menu", "edit", "--name", "test-menu-edit", "--comment", "testcomment"}},
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
				cleanupErr := removeMenu(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createMenu(Client, tt.args.command[5])
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
			editedHandle, err := Client.GetMenuHandle(tt.args.command[5])
			cobbler.FailOnError(t, err)
			updatedMenu, err := Client.GetMenu(editedHandle, false, false)
			cobbler.FailOnError(t, err)
			if updatedMenu.Comment != "testcomment" {
				t.Fatal("menu update wasn't successful")
			}
		})
	}
}

func Test_MenuFindCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "menu", "find", "--name", "test-menu-find"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			menuName := "test-menu-find"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeMenu(Client, menuName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createMenu(Client, menuName)
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
			if !strings.Contains(stdoutString, menuName) {
				fmt.Println(stdoutString)
				t.Fatal("menu not successfully found")
			}
		})
	}
}

func Test_MenuListCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "menu", "list"}},
			want:    "menus:",
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
				t.Fatal("menu list marker not located in output")
			}
		})
	}
}

func Test_MenuRemoveCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "menu", "remove", "--name", "test-menu-remove"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			setupClient(t)
			_, err := createMenu(Client, tt.args.command[5])
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
			result, err := Client.HasItem("menu", tt.args.command[5])
			cobbler.FailOnError(t, err)
			if result {
				// A missing item means we get "false", as such we error when we find an item.
				t.Fatal("menu not successfully removed")
			}
		})
	}
}

func Test_MenuRenameCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "menu", "rename", "--name", "test-menu-rename", "--newname", "test-menu-renamed"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			menuName := "test-menu-rename"
			newMenuName := "test-menu-renamed"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeMenu(Client, newMenuName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createMenu(Client, menuName)
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
			resultOldName, err := Client.HasItem("menu", menuName)
			cobbler.FailOnError(t, err)
			if resultOldName {
				t.Fatal("menu not successfully renamed (old name present)")
			}
			resultNewName, err := Client.HasItem("menu", newMenuName)
			cobbler.FailOnError(t, err)
			if !resultNewName {
				t.Fatal("menu not successfully renamed (new name not present)")
			}
		})
	}
}

// Test_MenuCopyCmd_UID exercises the --uid sibling flag on copy.
func Test_MenuCopyCmd_UID(t *testing.T) {
	name := "test-menu-copy-uid"
	newName := "test-menu-copied-uid"
	setupClient(t)
	_, err := createMenu(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeMenu(Client, name); err != nil {
			t.Errorf("cleanup: remove menu %s: %v", name, err)
		}
	})
	t.Cleanup(func() {
		if err := removeMenu(Client, newName); err != nil {
			t.Errorf("cleanup: remove menu %s: %v", newName, err)
		}
	})
	uid, err := Client.GetMenuHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "menu", "copy", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetMenuHandle(newName)
	cobbler.FailOnError(t, err)
	_, err = Client.GetMenu(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

// Test_MenuEditCmd_UID exercises the --uid sibling flag on edit.
func Test_MenuEditCmd_UID(t *testing.T) {
	name := "test-menu-edit-uid"
	setupClient(t)
	_, err := createMenu(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeMenu(Client, name); err != nil {
			t.Errorf("cleanup: remove menu %s: %v", name, err)
		}
	})
	uid, err := Client.GetMenuHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "menu", "edit", "--uid", uid, "--comment", "testcomment-uid"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updatedMenu, err := Client.GetMenu(uid, false, false)
	cobbler.FailOnError(t, err)
	if updatedMenu.Comment != "testcomment-uid" {
		t.Fatal("menu update via --uid wasn't successful")
	}
}

// Test_MenuRemoveCmd_UID exercises the --uid sibling flag on remove.
func Test_MenuRemoveCmd_UID(t *testing.T) {
	name := "test-menu-remove-uid"
	setupClient(t)
	_, err := createMenu(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetMenuHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "menu", "remove", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("menu", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("menu not successfully removed via --uid")
	}
}

// Test_MenuRenameCmd_UID exercises the --uid sibling flag on rename.
func Test_MenuRenameCmd_UID(t *testing.T) {
	name := "test-menu-rename-uid"
	newName := "test-menu-renamed-uid"
	setupClient(t)
	_, err := createMenu(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetMenuHandle(name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeMenu(Client, newName); err != nil {
			t.Errorf("cleanup: remove menu %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "menu", "rename", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("menu", name)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("menu not successfully renamed via --uid (old name present)")
	}
	resultNewName, err := Client.HasItem("menu", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("menu not successfully renamed via --uid (new name not present)")
	}
}

func Test_MenuReportCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "menu", "report", "--name", "test-menu-report"}},
			want:    ": test-menu-report",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			menuName := "test-menu-report"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeMenu(Client, menuName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createMenu(Client, menuName)
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

// Test_MenuReportCmd_UID exercises the --uid sibling flag on report.
func Test_MenuReportCmd_UID(t *testing.T) {
	name := "test-menu-report-uid"
	setupClient(t)
	_, err := createMenu(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeMenu(Client, name); err != nil {
			t.Errorf("cleanup: remove menu %s: %v", name, err)
		}
	})
	uid, err := Client.GetMenuHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "menu", "report", "--uid", uid})
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
		t.Fatal("menu name missing from report --uid output")
	}
}

// Test_MenuReportCmd_All exercises the report branch taken when neither
// --name nor --uid is supplied (report every menu).
func Test_MenuReportCmd_All(t *testing.T) {
	name := "test-menu-report-all"
	setupClient(t)
	_, err := createMenu(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeMenu(Client, name); err != nil {
			t.Errorf("cleanup: remove menu %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "menu", "report"})
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
		t.Fatal("menu name missing from report --all output")
	}
}

// Test_MenuExportCmd exercises the export command's json branch with an
// explicit --name.
func Test_MenuExportCmd(t *testing.T) {
	name := "test-menu-export"
	setupClient(t)
	_, err := createMenu(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeMenu(Client, name); err != nil {
			t.Errorf("cleanup: remove menu %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "menu", "export", "--name", name, "--format", "json"})
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
		t.Fatal("menu name missing from json export output")
	}
}

// Test_MenuExportCmd_UID exercises the export command's --uid sibling flag.
func Test_MenuExportCmd_UID(t *testing.T) {
	name := "test-menu-export-uid"
	setupClient(t)
	_, err := createMenu(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeMenu(Client, name); err != nil {
			t.Errorf("cleanup: remove menu %s: %v", name, err)
		}
	})
	uid, err := Client.GetMenuHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "menu", "export", "--uid", uid, "--format", "json"})
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
		t.Fatal("menu name missing from json export --uid output")
	}
}

// Test_MenuExportCmd_All exercises the export branch taken when neither
// --name nor --uid is supplied, using the yaml format.
func Test_MenuExportCmd_All(t *testing.T) {
	name := "test-menu-export-all"
	setupClient(t)
	_, err := createMenu(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeMenu(Client, name); err != nil {
			t.Errorf("cleanup: remove menu %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "menu", "export", "--format", "yaml"})
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
		t.Fatal("menu name missing from yaml export --all output")
	}
}
