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

func createSystemGroup(client cobbler.Client, name string) (*cobbler.SystemGroup, error) {
	g := cobbler.NewSystemGroup()
	g.Name = name
	return client.CreateSystemGroup(g)
}

func removeSystemGroup(client cobbler.Client, name string) error {
	handle, err := client.GetSystemGroupHandle(name)
	if err != nil {
		return err
	}
	return client.DeleteSystemGroup(handle)
}

func Test_SystemGroupAddCmd(t *testing.T) {
	name := "test-system-group-add"
	setupClient(t)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "add", "--name", name})
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
	if !strings.Contains(stdoutString, "System group "+name+" created") {
		fmt.Println(stdoutString)
		t.Fatal("system group creation message missing")
	}
}

// Test_SystemGroupAddCmd_Items exercises extractGroupFlags' --items branch
// (group_common.go) together with system_group.go's own itemsSet handling.
func Test_SystemGroupAddCmd_Items(t *testing.T) {
	name := "test-system-group-add-items"
	systemName := "test-system-group-add-items-member"
	setupClient(t)
	member, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystem(Client, systemName); err != nil {
			t.Errorf("cleanup: remove system %s: %v", systemName, err)
		}
	})
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "add", "--name", name, "--items", member.Uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	handle, err := Client.GetSystemGroupHandle(name)
	cobbler.FailOnError(t, err)
	created, err := Client.GetSystemGroup(handle, false, false)
	cobbler.FailOnError(t, err)
	if len(created.Members) != 1 || created.Members[0] != member.Uid {
		t.Fatalf("system group members weren't set from --items, got: %v", created.Members)
	}
}

func Test_SystemGroupEditCmd(t *testing.T) {
	name := "test-system-group-edit"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "edit", "--name", name, "--comment", "testcomment"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	handle, err := Client.GetSystemGroupHandle(name)
	cobbler.FailOnError(t, err)
	updated, err := Client.GetSystemGroup(handle, false, false)
	cobbler.FailOnError(t, err)
	if updated.Comment != "testcomment" {
		t.Fatal("system group update wasn't successful")
	}
}

// Test_SystemGroupEditCmd_UID exercises the --uid sibling flag on edit.
func Test_SystemGroupEditCmd_UID(t *testing.T) {
	name := "test-system-group-edit-uid"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	uid, err := Client.GetSystemGroupHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "edit", "--uid", uid, "--comment", "testcomment-uid"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updated, err := Client.GetSystemGroup(uid, false, false)
	cobbler.FailOnError(t, err)
	if updated.Comment != "testcomment-uid" {
		t.Fatal("system group update via --uid wasn't successful")
	}
}

func Test_SystemGroupCopyCmd(t *testing.T) {
	name := "test-system-group-to-copy"
	newName := "test-system-group-copied"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, newName); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "copy", "--name", name, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetSystemGroupHandle(newName)
	cobbler.FailOnError(t, err)
	_, err = Client.GetSystemGroup(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

func Test_SystemGroupRenameCmd(t *testing.T) {
	name := "test-system-group-rename"
	newName := "test-system-group-renamed"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, newName); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "rename", "--name", name, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("system_group", name)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("system group not successfully renamed (old name present)")
	}
	resultNewName, err := Client.HasItem("system_group", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("system group not successfully renamed (new name not present)")
	}
}

func Test_SystemGroupRemoveCmd(t *testing.T) {
	name := "test-system-group-remove"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "remove", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("system_group", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("system group not successfully removed")
	}
}

// Test_SystemGroupRemoveCmd_UID exercises the --uid sibling flag on remove.
func Test_SystemGroupRemoveCmd_UID(t *testing.T) {
	name := "test-system-group-remove-uid"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetSystemGroupHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "remove", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("system_group", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("system group not successfully removed via --uid")
	}
}

func Test_SystemGroupFindCmd(t *testing.T) {
	name := "test-system-group-find"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "find", "--name", name})
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
		t.Fatal("system group not successfully found")
	}
}

func Test_SystemGroupListCmd(t *testing.T) {
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "list"})
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
	if !strings.Contains(string(stdoutBytes), "system_groups:") {
		t.Fatal("system group list marker not located in output")
	}
}

func Test_SystemGroupReportCmd(t *testing.T) {
	name := "test-system-group-report"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "report", "--name", name})
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
		t.Fatal("system group name missing from report output")
	}
}

// Test_SystemGroupReportCmd_UID exercises the --uid sibling flag on report.
func Test_SystemGroupReportCmd_UID(t *testing.T) {
	name := "test-system-group-report-uid"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	uid, err := Client.GetSystemGroupHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "report", "--uid", uid})
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
		t.Fatal("system group name missing from report --uid output")
	}
}

// Test_SystemGroupReportCmd_All exercises the report branch taken when
// neither --name nor --uid is supplied (report every system group).
func Test_SystemGroupReportCmd_All(t *testing.T) {
	name := "test-system-group-report-all"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "report"})
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
		t.Fatal("system group name missing from report --all output")
	}
}

func Test_SystemGroupExportCmd(t *testing.T) {
	name := "test-system-group-export"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "export", "--name", name, "--format", "json"})
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
		t.Fatal("system group name missing from json export output")
	}
}

// Test_SystemGroupExportCmd_All exercises the export branch taken when
// neither --name nor --uid is supplied, using the yaml format.
func Test_SystemGroupExportCmd_All(t *testing.T) {
	name := "test-system-group-export-all"
	setupClient(t)
	_, err := createSystemGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeSystemGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove system group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "system-group", "export", "--format", "yaml"})
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
		t.Fatal("system group name missing from yaml export --all output")
	}
}
