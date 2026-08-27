package cmd

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

func createDistroGroup(client cobbler.Client, name string) (*cobbler.DistroGroup, error) {
	g := cobbler.NewDistroGroup()
	g.Name = name
	return client.CreateDistroGroup(g)
}

func removeDistroGroup(client cobbler.Client, name string) error {
	handle, err := client.GetDistroGroupHandle(name)
	if err != nil {
		return err
	}
	return client.DeleteDistroGroup(handle)
}

func Test_DistroGroupAddCmd(t *testing.T) {
	name := "test-distro-group-add"
	setupClient(t)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "add", "--name", name})
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
	stdoutString := string(stdoutBytes)
	if !strings.Contains(stdoutString, "Distro group "+name+" created") {
		fmt.Println(stdoutString)
		t.Fatal("distro group creation message missing")
	}
}

func Test_DistroGroupEditCmd(t *testing.T) {
	name := "test-distro-group-edit"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "edit", "--name", name, "--comment", "testcomment"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	handle, err := Client.GetDistroGroupHandle(name)
	cobbler.FailOnError(t, err)
	updated, err := Client.GetDistroGroup(handle, false, false)
	cobbler.FailOnError(t, err)
	if updated.Comment != "testcomment" {
		t.Fatal("distro group update wasn't successful")
	}
}

// Test_DistroGroupEditCmd_UID exercises the --uid sibling flag on edit.
func Test_DistroGroupEditCmd_UID(t *testing.T) {
	name := "test-distro-group-edit-uid"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	uid, err := Client.GetDistroGroupHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "edit", "--uid", uid, "--comment", "testcomment-uid"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updated, err := Client.GetDistroGroup(uid, false, false)
	cobbler.FailOnError(t, err)
	if updated.Comment != "testcomment-uid" {
		t.Fatal("distro group update via --uid wasn't successful")
	}
}

func Test_DistroGroupCopyCmd(t *testing.T) {
	name := "test-distro-group-to-copy"
	newName := "test-distro-group-copied"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, newName); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "copy", "--name", name, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetDistroGroupHandle(newName)
	cobbler.FailOnError(t, err)
	_, err = Client.GetDistroGroup(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

func Test_DistroGroupRenameCmd(t *testing.T) {
	name := "test-distro-group-rename"
	newName := "test-distro-group-renamed"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, newName); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "rename", "--name", name, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("distro_group", name)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("distro group not successfully renamed (old name present)")
	}
	resultNewName, err := Client.HasItem("distro_group", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("distro group not successfully renamed (new name not present)")
	}
}

func Test_DistroGroupRemoveCmd(t *testing.T) {
	name := "test-distro-group-remove"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "remove", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("distro_group", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("distro group not successfully removed")
	}
}

// Test_DistroGroupRemoveCmd_UID exercises the --uid sibling flag on remove.
func Test_DistroGroupRemoveCmd_UID(t *testing.T) {
	name := "test-distro-group-remove-uid"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetDistroGroupHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "remove", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("distro_group", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("distro group not successfully removed via --uid")
	}
}

func Test_DistroGroupFindCmd(t *testing.T) {
	name := "test-distro-group-find"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "find", "--name", name})
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
		t.Fatal("distro group not successfully found")
	}
}

func Test_DistroGroupListCmd(t *testing.T) {
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "list"})
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
	if !strings.Contains(string(stdoutBytes), "distro_groups:") {
		t.Fatal("distro group list marker not located in output")
	}
}

func Test_DistroGroupReportCmd(t *testing.T) {
	name := "test-distro-group-report"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "report", "--name", name})
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
		t.Fatal("distro group name missing from report output")
	}
}

// Test_DistroGroupReportCmd_UID exercises the --uid sibling flag on report.
func Test_DistroGroupReportCmd_UID(t *testing.T) {
	name := "test-distro-group-report-uid"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	uid, err := Client.GetDistroGroupHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "report", "--uid", uid})
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
		t.Fatal("distro group name missing from report --uid output")
	}
}

// Test_DistroGroupReportCmd_All exercises the report branch taken when
// neither --name nor --uid is supplied (report every distro group).
func Test_DistroGroupReportCmd_All(t *testing.T) {
	name := "test-distro-group-report-all"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "report"})
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
		t.Fatal("distro group name missing from report --all output")
	}
}

func Test_DistroGroupExportCmd(t *testing.T) {
	name := "test-distro-group-export"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "export", "--name", name, "--format", "json"})
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
		t.Fatal("distro group name missing from json export output")
	}
}

// Test_DistroGroupExportCmd_All exercises the export branch taken when
// neither --name nor --uid is supplied, using the yaml format.
func Test_DistroGroupExportCmd_All(t *testing.T) {
	name := "test-distro-group-export-all"
	setupClient(t)
	_, err := createDistroGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeDistroGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove distro group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "distro-group", "export", "--format", "yaml"})
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
		t.Fatal("distro group name missing from yaml export --all output")
	}
}
