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

func createRepo(client cobbler.Client, name string) (*cobbler.Repo, error) {
	repo := cobbler.NewRepo()
	repo.Name = name
	return client.CreateRepo(repo)
}

func removeRepo(client cobbler.Client, name string) error {
	handle, err := client.GetRepoHandle(name)
	if err != nil {
		return err
	}
	return client.DeleteRepo(handle)
}

func Test_RepoAddCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "repo", "add", "--name", "test-plain"}},
			want:    "Repo test-plain created",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeRepo(Client, tt.args.command[5])
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

func Test_RepoCopyCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "repo", "copy", "--name", "repo-to-copy", "--newname", "copied-repo"}},
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
				cleanupErr := removeRepo(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
				cleanupErr = removeRepo(Client, tt.args.command[7])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createRepo(Client, tt.args.command[5])
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
			copiedHandle, err := Client.GetRepoHandle(tt.args.command[7])
			cobbler.FailOnError(t, err)
			_, err = Client.GetRepo(copiedHandle, false, false)
			cobbler.FailOnError(t, err)
		})
	}
}

func Test_RepoEditCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "repo", "edit", "--name", "test-repo-edit", "--comment", "testcomment"}},
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
				cleanupErr := removeRepo(Client, tt.args.command[5])
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createRepo(Client, tt.args.command[5])
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
			editedHandle, err := Client.GetRepoHandle(tt.args.command[5])
			cobbler.FailOnError(t, err)
			updatedRepo, err := Client.GetRepo(editedHandle, false, false)
			cobbler.FailOnError(t, err)
			if updatedRepo.Comment != "testcomment" {
				t.Fatal("repo update wasn't successful")
			}
		})
	}
}

// Test_RepoCopyCmd_UID exercises the --uid sibling flag on copy.
func Test_RepoCopyCmd_UID(t *testing.T) {
	name := "test-repo-copy-uid"
	newName := "test-repo-copied-uid"
	setupClient(t)
	_, err := createRepo(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeRepo(Client, name); err != nil {
			t.Errorf("cleanup: remove repo %s: %v", name, err)
		}
	})
	t.Cleanup(func() {
		if err := removeRepo(Client, newName); err != nil {
			t.Errorf("cleanup: remove repo %s: %v", newName, err)
		}
	})
	uid, err := Client.GetRepoHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "repo", "copy", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetRepoHandle(newName)
	cobbler.FailOnError(t, err)
	_, err = Client.GetRepo(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

// Test_RepoEditCmd_UID exercises the --uid sibling flag on edit.
func Test_RepoEditCmd_UID(t *testing.T) {
	name := "test-repo-edit-uid"
	setupClient(t)
	_, err := createRepo(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeRepo(Client, name); err != nil {
			t.Errorf("cleanup: remove repo %s: %v", name, err)
		}
	})
	uid, err := Client.GetRepoHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "repo", "edit", "--uid", uid, "--comment", "testcomment-uid"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updatedRepo, err := Client.GetRepo(uid, false, false)
	cobbler.FailOnError(t, err)
	if updatedRepo.Comment != "testcomment-uid" {
		t.Fatal("repo update via --uid wasn't successful")
	}
}

func Test_RepoFindCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "repo", "find", "--name", "test-repo-find"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			repoName := "test-repo-find"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeRepo(Client, repoName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createRepo(Client, repoName)
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
			if !strings.Contains(stdoutString, repoName) {
				fmt.Println(stdoutString)
				t.Fatal("repo not successfully found")
			}
		})
	}
}

func Test_RepoListCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "repo", "list"}},
			want:    "repos:",
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
				t.Fatal("repo list marker not located in output")
			}
		})
	}
}

func Test_RepoRemoveCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "repo", "remove", "--name", "test-repo-remove"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			setupClient(t)
			_, err := createRepo(Client, tt.args.command[5])
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
			result, err := Client.HasItem("repo", tt.args.command[5])
			cobbler.FailOnError(t, err)
			if result {
				// A missing item means we get "false", as such we error when we find an item.
				t.Fatal("repo not successfully removed")
			}
		})
	}
}

// Test_RepoRemoveCmd_UID exercises the --uid sibling flag on remove.
func Test_RepoRemoveCmd_UID(t *testing.T) {
	name := "test-repo-remove-uid"
	setupClient(t)
	_, err := createRepo(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetRepoHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "repo", "remove", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("repo", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("repo not successfully removed via --uid")
	}
}

func Test_RepoRenameCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "repo", "rename", "--name", "test-repo-rename", "--newname", "test-repo-renamed"}},
			want:    "Event ID:",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			repoName := "test-repo-rename"
			newRepoName := "test-repo-renamed"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeRepo(Client, newRepoName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createRepo(Client, repoName)
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
			resultOldName, err := Client.HasItem("repo", repoName)
			cobbler.FailOnError(t, err)
			if resultOldName {
				t.Fatal("repo not successfully renamed (old name present)")
			}
			resultNewName, err := Client.HasItem("repo", newRepoName)
			cobbler.FailOnError(t, err)
			if !resultNewName {
				t.Fatal("repo not successfully renamed (new name not present)")
			}
		})
	}
}

func Test_RepoReportCmd(t *testing.T) {
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
			args:    args{command: []string{"--config", "../testing/.cobbler.yaml", "repo", "report", "--name", "test-repo-report"}},
			want:    ": test-repo-report",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cleanup
			repoName := "test-repo-report"
			var err error
			defer func() {
				// Client is initialized since this is the cleanup
				cleanupErr := removeRepo(Client, repoName)
				cobbler.FailOnError(t, cleanupErr)
			}()
			// Arrange
			setupClient(t)
			_, err = createRepo(Client, repoName)
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

// Test_RepoRenameCmd_UID exercises the --uid sibling flag on rename.
func Test_RepoRenameCmd_UID(t *testing.T) {
	name := "test-repo-rename-uid"
	newName := "test-repo-renamed-uid"
	setupClient(t)
	_, err := createRepo(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetRepoHandle(name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeRepo(Client, newName); err != nil {
			t.Errorf("cleanup: remove repo %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "repo", "rename", "--uid", uid, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("repo", name)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("repo not successfully renamed via --uid (old name present)")
	}
	resultNewName, err := Client.HasItem("repo", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("repo not successfully renamed via --uid (new name not present)")
	}
}

// Test_RepoReportCmd_UID exercises the --uid sibling flag on report.
func Test_RepoReportCmd_UID(t *testing.T) {
	name := "test-repo-report-uid"
	setupClient(t)
	_, err := createRepo(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeRepo(Client, name); err != nil {
			t.Errorf("cleanup: remove repo %s: %v", name, err)
		}
	})
	uid, err := Client.GetRepoHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "repo", "report", "--uid", uid})
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
		t.Fatal("repo name missing from report --uid output")
	}
}

// Test_RepoReportCmd_All exercises the report branch taken when neither
// --name nor --uid is supplied (report every repo).
func Test_RepoReportCmd_All(t *testing.T) {
	name := "test-repo-report-all"
	setupClient(t)
	_, err := createRepo(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeRepo(Client, name); err != nil {
			t.Errorf("cleanup: remove repo %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "repo", "report"})
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
		t.Fatal("repo name missing from report --all output")
	}
}

// Test_RepoExportCmd exercises the export command's json branch with an
// explicit --name.
func Test_RepoExportCmd(t *testing.T) {
	name := "test-repo-export"
	setupClient(t)
	_, err := createRepo(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeRepo(Client, name); err != nil {
			t.Errorf("cleanup: remove repo %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "repo", "export", "--name", name, "--format", "json"})
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
		t.Fatal("repo name missing from json export output")
	}
}

// Test_RepoExportCmd_UID exercises the export command's --uid sibling flag.
func Test_RepoExportCmd_UID(t *testing.T) {
	name := "test-repo-export-uid"
	setupClient(t)
	_, err := createRepo(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeRepo(Client, name); err != nil {
			t.Errorf("cleanup: remove repo %s: %v", name, err)
		}
	})
	uid, err := Client.GetRepoHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "repo", "export", "--uid", uid, "--format", "json"})
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
		t.Fatal("repo name missing from json export --uid output")
	}
}

// Test_RepoExportCmd_All exercises the export branch taken when neither
// --name nor --uid is supplied, using the yaml format.
func Test_RepoExportCmd_All(t *testing.T) {
	name := "test-repo-export-all"
	setupClient(t)
	_, err := createRepo(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeRepo(Client, name); err != nil {
			t.Errorf("cleanup: remove repo %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "repo", "export", "--format", "yaml"})
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
		t.Fatal("repo name missing from yaml export --all output")
	}
}
