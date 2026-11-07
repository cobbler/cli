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

// Test_DistroAddCmd_FailedAddDoesNotLeakDanglingName is a regression test for cobbler-cli
// ledger entry 2.8 ("A failed new-mode save_item/add permanently leaks a dangling name
// claim", the client-side fix for cobbler ledger entry 1.4). Before the fix, cobbler-cli
// issued its new_distro/modify_distro/save_distro XML-RPC calls directly with no
// transaction_begin/transaction_commit/transaction_abort wrapping. A failed `distro add`
// (new_distro succeeds server-side and claims the name, but a later modify_distro call fails
// validation before save_distro is ever reached) left the half-created distro's name sitting
// forever in cobblerd's global self.unsaved_items cache (until CACHE_TIMEOUT or a cobblerd
// restart) - with zero CLI-visible trace (`distro list`/`distro report` correctly show no
// such distro). The client-visible symptom is that an immediate retry of `distro add` with
// the same name and otherwise-valid fields fails with "a Distro with the name ... already
// exists", because cobblerd's has_item/get_item_handle falls back to searching
// self.unsaved_items. Wrapping the add in a transaction (see withTransaction in cmd/root.go)
// fixes this: new_distro stages into the server's per-token transaction instead of the
// global unsaved_items cache, and the wrapper's transaction_abort on error discards that
// staged state cleanly, so nothing is left behind for the retry to collide with.
func Test_DistroAddCmd_FailedAddDoesNotLeakDanglingName(t *testing.T) {
	name := "test-distro-add-dangling-name"
	setupClient(t)
	t.Cleanup(func() {
		// Best-effort: only succeeds if the second (valid) add below actually created the
		// distro. If the bug has regressed and only the first (failing) add ran, there is
		// nothing to clean up server-side, so an error here is not itself a test failure.
		_ = removeDistro(Client, name)
	})

	// First attempt: a deliberately invalid --initrd path. Client.CreateDistro's new_distro
	// call succeeds, but the subsequent modify_distro("initrd", ...) call fails server-side
	// validation ("initrd not found"), so CreateDistro returns an error without ever reaching
	// save_distro.
	cobra.OnInitialize(initConfig, setupLogger)
	failCmd := NewRootCmd()
	failCmd.SetArgs([]string{
		"--config", "../testing/.cobbler.yaml",
		"distro", "add", "--name", name,
		"--kernel", "/extracted_iso_image/install/vmlinuz",
		"--initrd", "/extracted_iso_image/install/does-not-exist-initrd.gz",
	})
	failStdout := bytes.NewBufferString("")
	failStderr := bytes.NewBufferString("")
	failCmd.SetOut(failStdout)
	failCmd.SetErr(failStderr)
	if err := failCmd.Execute(); err == nil {
		t.Fatal("expected the first distro add (nonexistent initrd) to fail")
	}

	// Second attempt, immediately afterward: same name, valid fields this time. Before the
	// fix this fails with "a Distro with the name ... already exists" (the dangling name
	// claim left by the first attempt); after the fix it must succeed right away.
	cobra.OnInitialize(initConfig, setupLogger)
	okCmd := NewRootCmd()
	okCmd.SetArgs([]string{
		"--config", "../testing/.cobbler.yaml",
		"distro", "add", "--name", name,
		"--kernel", "/extracted_iso_image/install/vmlinuz",
		"--initrd", "/extracted_iso_image/install/initrd.gz",
	})
	okStdout := bytes.NewBufferString("")
	okStderr := bytes.NewBufferString("")
	okCmd.SetOut(okStdout)
	okCmd.SetErr(okStderr)

	err := okCmd.Execute()
	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, okStderr)

	stdoutBytes, err := io.ReadAll(okStdout)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(stdoutBytes), "Distro "+name+" created") {
		t.Fatalf("expected the immediate retry with the same name to succeed; got stdout: %s", string(stdoutBytes))
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

// Test_UpdateDistroFromFlags_JSONFlags confirms --kernel-options-json, --kernel-options-post-json,
// and --template-files-json each merge on top of their flat stringToString counterpart and can
// carry a value the flat flag's key=value syntax can't represent, mirroring the pre-existing
// --autoinstall-meta-json. Exercises the real flag registration and switch-case code directly and
// does not need a live cobbler server.
func Test_UpdateDistroFromFlags_JSONFlags(t *testing.T) {
	t.Run("kernel-options-json", func(t *testing.T) {
		cmd, err := NewDistroEditCmd()
		cobbler.FailOnError(t, err)
		cobbler.FailOnError(t, cmd.Flags().Set("kernel-options", "foo=bar"))
		cobbler.FailOnError(t, cmd.Flags().Set("kernel-options-json", `{"foo": "overridden", "nested": [1, 2]}`))

		distro := cobbler.NewDistro()
		cobbler.FailOnError(t, updateDistroFromFlags(cmd, &distro))

		if distro.KernelOptions.IsInherited {
			t.Fatal("distro.KernelOptions should not be marked inherited")
		}
		if got := distro.KernelOptions.Data["foo"]; got != "overridden" {
			t.Errorf("expected --kernel-options-json to override the flat \"foo\" value, got %v", got)
		}
		nested, ok := distro.KernelOptions.Data["nested"].([]interface{})
		if !ok || len(nested) != 2 {
			t.Errorf("expected \"nested\" to carry a decoded 2-element list, got %v", distro.KernelOptions.Data["nested"])
		}
	})

	t.Run("kernel-options-post-json", func(t *testing.T) {
		cmd, err := NewDistroEditCmd()
		cobbler.FailOnError(t, err)
		cobbler.FailOnError(t, cmd.Flags().Set("kernel-options-post", "foo=bar"))
		cobbler.FailOnError(t, cmd.Flags().Set("kernel-options-post-json", `{"foo": "overridden", "nested": [1, 2]}`))

		distro := cobbler.NewDistro()
		cobbler.FailOnError(t, updateDistroFromFlags(cmd, &distro))

		if distro.KernelOptionsPost.IsInherited {
			t.Fatal("distro.KernelOptionsPost should not be marked inherited")
		}
		if got := distro.KernelOptionsPost.Data["foo"]; got != "overridden" {
			t.Errorf("expected --kernel-options-post-json to override the flat \"foo\" value, got %v", got)
		}
		nested, ok := distro.KernelOptionsPost.Data["nested"].([]interface{})
		if !ok || len(nested) != 2 {
			t.Errorf("expected \"nested\" to carry a decoded 2-element list, got %v", distro.KernelOptionsPost.Data["nested"])
		}
	})

	t.Run("template-files-json", func(t *testing.T) {
		cmd, err := NewDistroEditCmd()
		cobbler.FailOnError(t, err)
		cobbler.FailOnError(t, cmd.Flags().Set("template-files", "motd=snippets/motd"))
		cobbler.FailOnError(t, cmd.Flags().Set("template-files-json", `{"motd": "snippets/motd-v2", "banner": "snippets/banner"}`))

		distro := cobbler.NewDistro()
		cobbler.FailOnError(t, updateDistroFromFlags(cmd, &distro))

		if got := distro.TemplateFiles["motd"]; got != "snippets/motd-v2" {
			t.Errorf("expected --template-files-json to override the flat \"motd\" value, got %q", got)
		}
		if got := distro.TemplateFiles["banner"]; got != "snippets/banner" {
			t.Errorf("expected --template-files-json to add \"banner\", got %q", got)
		}
	})
}
