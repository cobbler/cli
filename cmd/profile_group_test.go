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

func createProfileGroup(client cobbler.Client, name string) (*cobbler.ProfileGroup, error) {
	g := cobbler.NewProfileGroup()
	g.Name = name
	return client.CreateProfileGroup(g)
}

func removeProfileGroup(client cobbler.Client, name string) error {
	handle, err := client.GetProfileGroupHandle(name)
	if err != nil {
		return err
	}
	return client.DeleteProfileGroup(handle)
}

func Test_ProfileGroupAddCmd(t *testing.T) {
	name := "test-profile-group-add"
	setupClient(t)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "add", "--name", name})
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
	if !strings.Contains(stdoutString, "Profile group "+name+" created") {
		fmt.Println(stdoutString)
		t.Fatal("profile group creation message missing")
	}
}

// Test_ProfileGroupAddCmd_Items exercises extractGroupFlags' --items branch
// (group_common.go) together with profile_group.go's own itemsSet handling.
func Test_ProfileGroupAddCmd_Items(t *testing.T) {
	name := "test-profile-group-add-items"
	profileName := "test-profile-group-add-items-member"
	setupClient(t)
	member, err := createProfile(Client, profileName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfile(Client, profileName); err != nil {
			t.Errorf("cleanup: remove profile %s: %v", profileName, err)
		}
	})
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "add", "--name", name, "--items", member.Uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	handle, err := Client.GetProfileGroupHandle(name)
	cobbler.FailOnError(t, err)
	created, err := Client.GetProfileGroup(handle, false, false)
	cobbler.FailOnError(t, err)
	if len(created.Members) != 1 || created.Members[0] != member.Uid {
		t.Fatalf("profile group members weren't set from --items, got: %v", created.Members)
	}
}

func Test_ProfileGroupEditCmd(t *testing.T) {
	name := "test-profile-group-edit"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "edit", "--name", name, "--comment", "testcomment"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	handle, err := Client.GetProfileGroupHandle(name)
	cobbler.FailOnError(t, err)
	updated, err := Client.GetProfileGroup(handle, false, false)
	cobbler.FailOnError(t, err)
	if updated.Comment != "testcomment" {
		t.Fatal("profile group update wasn't successful")
	}
}

// Test_ProfileGroupEditCmd_UID exercises the --uid sibling flag on edit.
func Test_ProfileGroupEditCmd_UID(t *testing.T) {
	name := "test-profile-group-edit-uid"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	uid, err := Client.GetProfileGroupHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "edit", "--uid", uid, "--comment", "testcomment-uid"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updated, err := Client.GetProfileGroup(uid, false, false)
	cobbler.FailOnError(t, err)
	if updated.Comment != "testcomment-uid" {
		t.Fatal("profile group update via --uid wasn't successful")
	}
}

func Test_ProfileGroupCopyCmd(t *testing.T) {
	name := "test-profile-group-to-copy"
	newName := "test-profile-group-copied"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, newName); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "copy", "--name", name, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetProfileGroupHandle(newName)
	cobbler.FailOnError(t, err)
	_, err = Client.GetProfileGroup(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

func Test_ProfileGroupRenameCmd(t *testing.T) {
	name := "test-profile-group-rename"
	newName := "test-profile-group-renamed"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, newName); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "rename", "--name", name, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("profile_group", name)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("profile group not successfully renamed (old name present)")
	}
	resultNewName, err := Client.HasItem("profile_group", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("profile group not successfully renamed (new name not present)")
	}
}

func Test_ProfileGroupRemoveCmd(t *testing.T) {
	name := "test-profile-group-remove"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "remove", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("profile_group", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("profile group not successfully removed")
	}
}

// Test_ProfileGroupRemoveCmd_UID exercises the --uid sibling flag on remove.
func Test_ProfileGroupRemoveCmd_UID(t *testing.T) {
	name := "test-profile-group-remove-uid"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetProfileGroupHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "remove", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("profile_group", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("profile group not successfully removed via --uid")
	}
}

func Test_ProfileGroupFindCmd(t *testing.T) {
	name := "test-profile-group-find"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "find", "--name", name})
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
		t.Fatal("profile group not successfully found")
	}
}

func Test_ProfileGroupListCmd(t *testing.T) {
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "list"})
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
	if !strings.Contains(string(stdoutBytes), "profile_groups:") {
		t.Fatal("profile group list marker not located in output")
	}
}

func Test_ProfileGroupReportCmd(t *testing.T) {
	name := "test-profile-group-report"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "report", "--name", name})
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
		t.Fatal("profile group name missing from report output")
	}
}

// Test_ProfileGroupReportCmd_UID exercises the --uid sibling flag on report.
func Test_ProfileGroupReportCmd_UID(t *testing.T) {
	name := "test-profile-group-report-uid"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	uid, err := Client.GetProfileGroupHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "report", "--uid", uid})
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
		t.Fatal("profile group name missing from report --uid output")
	}
}

// Test_ProfileGroupReportCmd_All exercises the report branch taken when
// neither --name nor --uid is supplied (report every profile group).
func Test_ProfileGroupReportCmd_All(t *testing.T) {
	name := "test-profile-group-report-all"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "report"})
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
		t.Fatal("profile group name missing from report --all output")
	}
}

func Test_ProfileGroupExportCmd(t *testing.T) {
	name := "test-profile-group-export"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "export", "--name", name, "--format", "json"})
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
		t.Fatal("profile group name missing from json export output")
	}
}

// Test_ProfileGroupExportCmd_All exercises the export branch taken when
// neither --name nor --uid is supplied, using the yaml format.
func Test_ProfileGroupExportCmd_All(t *testing.T) {
	name := "test-profile-group-export-all"
	setupClient(t)
	_, err := createProfileGroup(Client, name)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeProfileGroup(Client, name); err != nil {
			t.Errorf("cleanup: remove profile group %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "profile-group", "export", "--format", "yaml"})
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
		t.Fatal("profile group name missing from yaml export --all output")
	}
}
