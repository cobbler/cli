package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"

	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// templateContainerName is the name of the running Cobbler test container, matching
// testing/compose.yml's container_name (and used consistently by every project that
// runs this same image for local/e2e testing).
const templateContainerName = "cobbler-dev"

// newTemplateFixtureFile creates a small file inside the Cobbler test container's
// autoinstall_templates_dir (/var/lib/cobbler/templates) and registers a t.Cleanup to
// remove it again.
//
// This is required because Cobbler's Template.content setter unconditionally requires
// uri.path to be non-empty (see cobbler/items/template.py), and cobblerclient's
// CreateTemplate/UpdateTemplate always push every field -- including a zero-value
// Content -- on every create/update. Since the server's default
// autoinstall_templates_allow_new_files setting is false, uri.path must point at a
// file that already exists under autoinstall_templates_dir, or item creation fails
// with "Additional path validation failed!". A dedicated, disposable file per test
// avoids clobbering any of the server's real fixture templates.
func newTemplateFixtureFile(t *testing.T, name string) string {
	t.Helper()
	createCmd := exec.Command("docker", "exec", templateContainerName, "sh", "-c",
		fmt.Sprintf("printf '# fixture content\\n' > /var/lib/cobbler/templates/%s", name))
	if out, err := createCmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to create template fixture file %s: %v: %s", name, err, out)
	}
	t.Cleanup(func() {
		removeCmd := exec.Command("docker", "exec", templateContainerName, "sh", "-c",
			fmt.Sprintf("rm -f /var/lib/cobbler/templates/%s", name))
		if out, err := removeCmd.CombinedOutput(); err != nil {
			t.Errorf("cleanup: remove template fixture file %s: %v: %s", name, err, out)
		}
	})
	return name
}

func createTemplate(client cobbler.Client, name, uriPath string) (*cobbler.Template, error) {
	tpl := cobbler.NewTemplate()
	tpl.Name = name
	tpl.URI.Path = uriPath
	return client.CreateTemplate(tpl)
}

func removeTemplate(client cobbler.Client, name string) error {
	handle, err := client.GetTemplateHandle(name)
	if err != nil {
		return err
	}
	return client.DeleteTemplate(handle)
}

func Test_TemplateAddCmd(t *testing.T) {
	name := "test-template-add"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-add.j2")
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "add", "--name", name, "--uri-path", fixture})
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
	if !strings.Contains(string(stdoutBytes), "Template "+name+" created") {
		fmt.Println(string(stdoutBytes))
		t.Fatal("template creation message missing")
	}
}

func Test_TemplateEditCmd(t *testing.T) {
	name := "test-template-edit"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-edit.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "edit", "--name", name, "--comment", "testcomment"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	handle, err := Client.GetTemplateHandle(name)
	cobbler.FailOnError(t, err)
	updated, err := Client.GetTemplate(handle, false, false)
	cobbler.FailOnError(t, err)
	if updated.Comment != "testcomment" {
		t.Fatal("template update wasn't successful")
	}
}

// Test_TemplateEditCmd_UID exercises the --uid sibling flag on edit.
func Test_TemplateEditCmd_UID(t *testing.T) {
	name := "test-template-edit-uid"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-edit-uid.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	uid, err := Client.GetTemplateHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "edit", "--uid", uid, "--comment", "testcomment-uid"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updated, err := Client.GetTemplate(uid, false, false)
	cobbler.FailOnError(t, err)
	if updated.Comment != "testcomment-uid" {
		t.Fatal("template update via --uid wasn't successful")
	}
}

// Test_TemplateEditCmd_AllFlags exercises updateTemplateFromFlags' template-type,
// uri-schema/uri-path, tags and content-file branches, plus parseTemplateSchema's
// success path -- none of which were covered by the --comment-only edit tests.
func Test_TemplateEditCmd_AllFlags(t *testing.T) {
	name := "test-template-edit-allflags"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-edit-allflags.j2")
	newFixture := newTemplateFixtureFile(t, "test-template-edit-allflags-2.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{
		"--config", "../testing/.cobbler.yaml", "template", "edit", "--name", name,
		"--template-type", "cheetah",
		"--uri-schema", "file",
		"--uri-path", newFixture,
		"--tags", "tag1,tag2",
	})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	handle, err := Client.GetTemplateHandle(name)
	cobbler.FailOnError(t, err)
	updated, err := Client.GetTemplate(handle, false, false)
	cobbler.FailOnError(t, err)
	if updated.TemplateType != "cheetah" {
		t.Fatal("template-type update wasn't successful")
	}
	if updated.URI.Schema != cobbler.TemplateSchemaFile {
		t.Fatal("uri-schema update wasn't successful")
	}
	if updated.URI.Path != newFixture {
		t.Fatal("uri-path update wasn't successful")
	}
	if len(updated.Tags) != 2 || updated.Tags[0] != "tag1" || updated.Tags[1] != "tag2" {
		t.Fatalf("tags update wasn't successful, got: %v", updated.Tags)
	}
}

// Test_TemplateEditCmd_Content exercises updateTemplateFromFlags' inline
// --content branch.
func Test_TemplateEditCmd_Content(t *testing.T) {
	name := "test-template-edit-content"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-edit-content.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "edit", "--name", name, "--content", "hello world"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
}

// Test_TemplateEditCmd_ContentFile exercises updateTemplateFromFlags'
// --content-file branch, which reads the given file off disk.
func Test_TemplateEditCmd_ContentFile(t *testing.T) {
	name := "test-template-edit-content-file"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-edit-content-file.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	contentFile := t.TempDir() + "/content.txt"
	if err := os.WriteFile(contentFile, []byte("content-from-file"), 0o644); err != nil {
		t.Fatalf("failed to write content file: %v", err)
	}
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "edit", "--name", name, "--content-file", contentFile})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
}

// Test_TemplateEditCmd_InvalidSchema exercises parseTemplateSchema's error
// branch, surfaced through the edit command's --uri-schema flag.
func Test_TemplateEditCmd_InvalidSchema(t *testing.T) {
	name := "test-template-edit-badschema"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-edit-badschema.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "edit", "--name", name, "--uri-schema", "bogus"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error for an unknown URI schema")
	}
	if !strings.Contains(err.Error(), "unknown URI schema") {
		t.Fatalf("unexpected error for invalid uri-schema: %v", err)
	}
}

func Test_TemplateCopyCmd(t *testing.T) {
	name := "test-template-to-copy"
	newName := "test-template-copied"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-to-copy.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	t.Cleanup(func() {
		if err := removeTemplate(Client, newName); err != nil {
			t.Errorf("cleanup: remove template %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "copy", "--name", name, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetTemplateHandle(newName)
	cobbler.FailOnError(t, err)
	_, err = Client.GetTemplate(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

func Test_TemplateRenameCmd(t *testing.T) {
	name := "test-template-rename"
	newName := "test-template-renamed"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-rename.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, newName); err != nil {
			t.Errorf("cleanup: remove template %s: %v", newName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "rename", "--name", name, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("template", name)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("template not successfully renamed (old name present)")
	}
	resultNewName, err := Client.HasItem("template", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("template not successfully renamed (new name not present)")
	}
}

func Test_TemplateRemoveCmd(t *testing.T) {
	name := "test-template-remove"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-remove.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "remove", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("template", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("template not successfully removed")
	}
}

// Test_TemplateRemoveCmd_UID exercises the --uid sibling flag on remove.
func Test_TemplateRemoveCmd_UID(t *testing.T) {
	name := "test-template-remove-uid"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-remove-uid.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	uid, err := Client.GetTemplateHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "remove", "--uid", uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("template", name)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("template not successfully removed via --uid")
	}
}

func Test_TemplateFindCmd(t *testing.T) {
	name := "test-template-find"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-find.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "find", "--name", name})
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
		t.Fatal("template not successfully found")
	}
}

func Test_TemplateListCmd(t *testing.T) {
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "list"})
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
	if !strings.Contains(string(stdoutBytes), "templates:") {
		t.Fatal("template list marker not located in output")
	}
}

func Test_TemplateReportCmd(t *testing.T) {
	name := "test-template-report"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-report.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "report", "--name", name})
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
	if !strings.Contains(string(stdoutBytes), ": "+name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("template name missing from report output")
	}
}

// Test_TemplateReportCmd_UID exercises the --uid sibling flag on report.
func Test_TemplateReportCmd_UID(t *testing.T) {
	name := "test-template-report-uid"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-report-uid.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	uid, err := Client.GetTemplateHandle(name)
	cobbler.FailOnError(t, err)
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "report", "--uid", uid})
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
	if !strings.Contains(string(stdoutBytes), ": "+name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("template name missing from report --uid output")
	}
}

// Test_TemplateReportCmd_All exercises the report branch taken when neither
// --name nor --uid is supplied (report every template).
func Test_TemplateReportCmd_All(t *testing.T) {
	name := "test-template-report-all"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-report-all.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "report"})
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
	if !strings.Contains(string(stdoutBytes), ": "+name) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("template name missing from report --all output")
	}
}

func Test_TemplateExportCmd(t *testing.T) {
	name := "test-template-export"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-export.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "export", "--name", name, "--format", "json"})
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
		t.Fatal("template name missing from json export output")
	}
}

// Test_TemplateExportCmd_All exercises the export branch taken when neither
// --name nor --uid is supplied, using the yaml format.
func Test_TemplateExportCmd_All(t *testing.T) {
	name := "test-template-export-all"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-export-all.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "export", "--format", "yaml"})
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
		t.Fatal("template name missing from yaml export --all output")
	}
}

func Test_TemplateContentCmd(t *testing.T) {
	name := "test-template-content"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-content.j2")
	tpl := cobbler.NewTemplate()
	tpl.Name = name
	tpl.URI.Path = fixture
	tpl.Content = "template body for content test"
	_, err := Client.CreateTemplate(tpl)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "content", "--name", name})
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
	if !strings.Contains(string(stdoutBytes), "template body for content test") {
		fmt.Println(string(stdoutBytes))
		t.Fatal("template content command didn't print resolved content")
	}
}

// Test_TemplateContentCmd_UID exercises the --uid sibling flag on content.
func Test_TemplateContentCmd_UID(t *testing.T) {
	name := "test-template-content-uid"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-content-uid.j2")
	tpl := cobbler.NewTemplate()
	tpl.Name = name
	tpl.URI.Path = fixture
	tpl.Content = "template body for content uid test"
	created, err := Client.CreateTemplate(tpl)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "content", "--uid", created.Uid})
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
	if !strings.Contains(string(stdoutBytes), "template body for content uid test") {
		fmt.Println(string(stdoutBytes))
		t.Fatal("template content --uid command didn't print resolved content")
	}
}

func Test_TemplateRefreshCmd(t *testing.T) {
	name := "test-template-refresh"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-refresh.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "refresh", "--name", name})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
}

// Test_TemplateRefreshCmd_All exercises the refresh branch taken when no
// --name is supplied at all (refresh every template).
func Test_TemplateRefreshCmd_All(t *testing.T) {
	name := "test-template-refresh-all"
	setupClient(t)
	fixture := newTemplateFixtureFile(t, "test-template-refresh-all.j2")
	_, err := createTemplate(Client, name, fixture)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := removeTemplate(Client, name); err != nil {
			t.Errorf("cleanup: remove template %s: %v", name, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "template", "refresh"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
}
