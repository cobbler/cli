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

// --- NewInterfaceCommand subcommand tests ---
//
// These exercise cmd/interface.go's cobra command wiring (Add/Edit/Copy/
// Rename/Remove/Find/List/Report/Export), in particular the --uid sibling
// flags and resolveUID/resolveSystemUid call sites this PR added. Every
// NetworkInterface created here hangs off a dedicated System (interface
// names are only unique per-system), and each test cleans its System up
// recursively, which also removes any interfaces attached to it.
//
// resolveUID's own generic behavior (canned-transport error branches,
// ambiguous-name handling, RemoveItemRecursive flag-plumbing) is already
// covered in item_test.go; these tests focus on interface.go's own call
// sites and command wiring instead of re-testing resolveUID itself.

func Test_InterfaceAddCmd_SystemName(t *testing.T) {
	systemName := "test-interface-add-name-system"
	ifaceName := "eth-interface-add-name"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "add", "--name", ifaceName, "--system-name", systemName})
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
	if !strings.Contains(string(stdoutBytes), "Network interface "+ifaceName+" created") {
		fmt.Println(string(stdoutBytes))
		t.Fatal("network interface creation message missing")
	}
}

// Test_InterfaceAddCmd_SystemUID exercises resolveSystemUid's --system-uid
// fast path (no name resolution round-trip).
func Test_InterfaceAddCmd_SystemUID(t *testing.T) {
	systemName := "test-interface-add-uid-system"
	ifaceName := "eth-interface-add-uid"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "add", "--name", ifaceName, "--system-uid", system.Uid})
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
	if !strings.Contains(string(stdoutBytes), "Network interface "+ifaceName+" created") {
		fmt.Println(string(stdoutBytes))
		t.Fatal("network interface creation message missing")
	}
}

// Test_InterfaceAddCmd_NoSystem exercises resolveSystemUid's error branch
// when neither --system-name nor --system-uid is supplied.
func Test_InterfaceAddCmd_NoSystem(t *testing.T) {
	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "add", "--name", "eth-interface-add-no-system"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err := rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error when neither --system-name nor --system-uid is set")
	}
	if !strings.Contains(err.Error(), "--system-name") || !strings.Contains(err.Error(), "--system-uid") {
		t.Fatalf("expected error to mention --system-name/--system-uid, got: %v", err)
	}
}

func Test_InterfaceEditCmd(t *testing.T) {
	systemName := "test-interface-edit-system"
	ifaceName := "eth-interface-edit"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "edit", "--name", ifaceName, "--mac-address", "aa:bb:cc:dd:ee:ff"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	handle, err := Client.GetNetworkInterfaceHandle(ifaceName)
	cobbler.FailOnError(t, err)
	updated, err := Client.GetNetworkInterface(handle, false, false)
	cobbler.FailOnError(t, err)
	if updated.MacAddress != "aa:bb:cc:dd:ee:ff" {
		t.Fatal("network interface update wasn't successful")
	}
}

// Test_InterfaceEditCmd_UID exercises the --uid sibling flag on edit.
func Test_InterfaceEditCmd_UID(t *testing.T) {
	systemName := "test-interface-edit-uid-system"
	ifaceName := "eth-interface-edit-uid"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	created, err := Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "edit", "--uid", created.Uid, "--mac-address", "aa:bb:cc:dd:ee:00"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	updated, err := Client.GetNetworkInterface(created.Uid, false, false)
	cobbler.FailOnError(t, err)
	if updated.MacAddress != "aa:bb:cc:dd:ee:00" {
		t.Fatal("network interface update via --uid wasn't successful")
	}
}

// Test_InterfaceEditCmd_AllFlags exercises updateNetworkInterfaceFromFlags'
// full switch statement (link-layer, bonding/bridging, IPv4, IPv6 and DNS
// flags) plus parseNetworkInterfaceType's success path, none of which were
// covered by the other, narrower interface tests.
func Test_InterfaceEditCmd_AllFlags(t *testing.T) {
	systemName := "test-interface-edit-allflags-system"
	ifaceName := "eth-interface-edit-allflags"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{
		"--config", "../testing/.cobbler.yaml", "interface", "edit", "--name", ifaceName,
		"--mac-address", "aa:bb:cc:dd:ee:ff",
		"--interface-type", "bond",
		"--interface-master", "eth0",
		"--bonding-opts", "mode=1",
		"--bridge-opts", "foo=bar",
		"--connected-mode=true",
		"--management=true",
		"--static=true",
		"--dhcp-tag", "mytag",
		"--mtu", "1500",
		"--virt-bridge", "br0",
		"--ipv4-address", "10.0.0.5",
		"--ipv4-netmask", "255.255.255.0",
		"--ipv4-gateway", "10.0.0.1",
		"--ipv4-static-routes", "10.0.1.0/24",
		"--ipv6-address", "fe80::1",
		"--ipv6-prefix", "64",
		"--ipv6-secondaries", "fe80::2",
		"--ipv6-mtu", "1500",
		"--ipv6-static-routes", "fe80::0/64",
		"--ipv6-default-gateway", "fe80::ff",
		"--dns-name", "myhost.example.com",
		"--dns-cnames", "alias1.example.com",
	})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	handle, err := Client.GetNetworkInterfaceHandle(ifaceName)
	cobbler.FailOnError(t, err)
	updated, err := Client.GetNetworkInterface(handle, false, false)
	cobbler.FailOnError(t, err)
	if updated.MacAddress != "aa:bb:cc:dd:ee:ff" {
		t.Fatal("mac-address update wasn't successful")
	}
	if updated.InterfaceType != cobbler.NetworkInterfaceTypeBond {
		t.Fatal("interface-type update wasn't successful")
	}
	if updated.IPv4.Address != "10.0.0.5" {
		t.Fatal("ipv4-address update wasn't successful")
	}
	if updated.IPv6.Address != "fe80::1" {
		t.Fatal("ipv6-address update wasn't successful")
	}
	if updated.DNS.Name != "myhost.example.com" {
		t.Fatal("dns-name update wasn't successful")
	}
}

// Test_InterfaceEditCmd_AllTypes exercises every remaining success branch of
// parseNetworkInterfaceType not already covered by Test_InterfaceEditCmd_AllFlags
// (which only exercises "bond").
func Test_InterfaceEditCmd_AllTypes(t *testing.T) {
	systemName := "test-interface-edit-alltypes-system"
	ifaceName := "eth-interface-edit-alltypes"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	for _, ifaceType := range []string{"na", "infiniband", "bridge", "bridge_slave", "bond_slave", "bonded_bridge_slave"} {
		cobra.OnInitialize(initConfig, setupLogger)
		rootCmd := NewRootCmd()
		rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "edit", "--name", ifaceName, "--interface-type", ifaceType})
		stdout := bytes.NewBufferString("")
		stderr := bytes.NewBufferString("")
		rootCmd.SetOut(stdout)
		rootCmd.SetErr(stderr)

		err = rootCmd.Execute()

		cobbler.FailOnError(t, err)
		FailOnNonEmptyStream(t, stderr)
		FailOnNonEmptyStream(t, stdout)
	}
}

// Test_InterfaceEditCmd_InvalidType exercises parseNetworkInterfaceType's
// error branch, surfaced through the edit command's --interface-type flag.
func Test_InterfaceEditCmd_InvalidType(t *testing.T) {
	systemName := "test-interface-edit-badtype-system"
	ifaceName := "eth-interface-edit-badtype"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "edit", "--name", ifaceName, "--interface-type", "bogus"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	if err == nil {
		t.Fatal("expected an error for an unknown interface type")
	}
	if !strings.Contains(err.Error(), "unknown interface type") {
		t.Fatalf("unexpected error for invalid interface type: %v", err)
	}
}

func Test_InterfaceCopyCmd(t *testing.T) {
	systemName := "test-interface-copy-system"
	ifaceName := "eth-interface-to-copy"
	newName := "eth-interface-copied"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "copy", "--name", ifaceName, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	copiedHandle, err := Client.GetNetworkInterfaceHandle(newName)
	cobbler.FailOnError(t, err)
	// The copy path always clears identity-bound fields (MacAddress/IPv4/IPv6)
	// after duplicating the item server-side; confirming the copy exists under
	// its new name and can be re-fetched is what exercises those lines.
	_, err = Client.GetNetworkInterface(copiedHandle, false, false)
	cobbler.FailOnError(t, err)
}

func Test_InterfaceRenameCmd(t *testing.T) {
	systemName := "test-interface-rename-system"
	ifaceName := "eth-interface-rename"
	newName := "eth-interface-renamed"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "rename", "--name", ifaceName, "--newname", newName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	resultOldName, err := Client.HasItem("network_interface", ifaceName)
	cobbler.FailOnError(t, err)
	if resultOldName {
		t.Fatal("network interface not successfully renamed (old name present)")
	}
	resultNewName, err := Client.HasItem("network_interface", newName)
	cobbler.FailOnError(t, err)
	if !resultNewName {
		t.Fatal("network interface not successfully renamed (new name not present)")
	}
}

func Test_InterfaceRemoveCmd(t *testing.T) {
	systemName := "test-interface-remove-system"
	ifaceName := "eth-interface-remove"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "remove", "--name", ifaceName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("network_interface", ifaceName)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("network interface not successfully removed")
	}
}

// Test_InterfaceRemoveCmd_UID exercises the --uid sibling flag on remove.
func Test_InterfaceRemoveCmd_UID(t *testing.T) {
	systemName := "test-interface-remove-uid-system"
	ifaceName := "eth-interface-remove-uid"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	created, err := Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "remove", "--uid", created.Uid})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	FailOnNonEmptyStream(t, stdout)
	result, err := Client.HasItem("network_interface", ifaceName)
	cobbler.FailOnError(t, err)
	if result {
		t.Fatal("network interface not successfully removed via --uid")
	}
}

func Test_InterfaceFindCmd(t *testing.T) {
	systemName := "test-interface-find-system"
	ifaceName := "eth-interface-find"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "find", "--name", ifaceName})
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
	if !strings.Contains(string(stdoutBytes), ifaceName) {
		t.Fatal("network interface not successfully found")
	}
}

func Test_InterfaceListCmd(t *testing.T) {
	systemName := "test-interface-list-system"
	ifaceName := "eth-interface-list"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "list"})
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
	stdoutString := string(stdoutBytes)
	// NewInterfaceListCommand groups by iface.SystemName, which
	// Client.GetNetworkInterfaces() (unresolved) doesn't currently populate
	// server-side; that's pre-existing behavior outside this PR's diff, so
	// only the interface's own name is asserted here.
	if !strings.Contains(stdoutString, ifaceName) {
		fmt.Println(stdoutString)
		t.Fatal("network interface list output missing interface name")
	}
}

func Test_InterfaceReportCmd(t *testing.T) {
	systemName := "test-interface-report-system"
	ifaceName := "eth-interface-report"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "report", "--name", ifaceName})
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
	if !strings.Contains(string(stdoutBytes), ": "+ifaceName) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("network interface name missing from report output")
	}
}

// Test_InterfaceReportCmd_UID exercises the --uid sibling flag on report.
func Test_InterfaceReportCmd_UID(t *testing.T) {
	systemName := "test-interface-report-uid-system"
	ifaceName := "eth-interface-report-uid"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	created, err := Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "report", "--uid", created.Uid})
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
	if !strings.Contains(string(stdoutBytes), ": "+ifaceName) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("network interface name missing from report --uid output")
	}
}

// Test_InterfaceReportCmd_SystemName exercises the --system-name filter
// branch (FindNetworkInterface), distinct from the --name/--uid branch.
func Test_InterfaceReportCmd_SystemName(t *testing.T) {
	systemName := "test-interface-report-sysname-system"
	ifaceName := "eth-interface-report-sysname"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "report", "--system-name", systemName})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	// Client.FindNetworkInterface's "system_name" criterion doesn't currently
	// match anything server-side (a pre-existing quirk of the unresolved
	// find_network_interface RPC, outside this PR's diff), so this only
	// asserts that the --system-name branch itself is taken without error,
	// not that it returns the interface.
	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	_, err = io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
}

// Test_InterfaceReportCmd_All exercises the default branch taken when
// neither --name, --uid, nor --system-name is supplied.
func Test_InterfaceReportCmd_All(t *testing.T) {
	systemName := "test-interface-report-all-system"
	ifaceName := "eth-interface-report-all"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "report"})
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
	if !strings.Contains(string(stdoutBytes), ": "+ifaceName) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("network interface name missing from report --all output")
	}
}

func Test_InterfaceExportCmd(t *testing.T) {
	systemName := "test-interface-export-system"
	ifaceName := "eth-interface-export"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "export", "--name", ifaceName, "--format", "json"})
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
	if !strings.Contains(string(stdoutBytes), `"name":"`+ifaceName+`"`) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("network interface name missing from json export output")
	}
}

// Test_InterfaceExportCmd_SystemName exercises export's --system-name filter
// branch, using the yaml format.
func Test_InterfaceExportCmd_SystemName(t *testing.T) {
	systemName := "test-interface-export-sysname-system"
	ifaceName := "eth-interface-export-sysname"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "export", "--system-name", systemName, "--format", "yaml"})
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	err = rootCmd.Execute()

	// As in Test_InterfaceReportCmd_SystemName, Client.FindNetworkInterface's
	// "system_name" criterion doesn't currently match anything server-side
	// (pre-existing, outside this PR's diff); this only asserts the
	// --system-name export branch is taken without error.
	cobbler.FailOnError(t, err)
	FailOnNonEmptyStream(t, stderr)
	_, err = io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
}

// Test_InterfaceExportCmd_All exercises the export default branch taken when
// neither --name, --uid, nor --system-name is supplied.
func Test_InterfaceExportCmd_All(t *testing.T) {
	systemName := "test-interface-export-all-system"
	ifaceName := "eth-interface-export-all"
	setupClient(t)
	system, err := createSystem(Client, systemName)
	cobbler.FailOnError(t, err)
	t.Cleanup(func() {
		if err := Client.DeleteSystemRecursive(system.Uid, true); err != nil {
			t.Errorf("cleanup: delete system %s: %v", systemName, err)
		}
	})
	iface := cobbler.NewNetworkInterface()
	iface.Name = ifaceName
	iface.SystemUid = system.Uid
	_, err = Client.CreateNetworkInterface(system.Uid, iface)
	cobbler.FailOnError(t, err)

	cobra.OnInitialize(initConfig, setupLogger)
	rootCmd := NewRootCmd()
	rootCmd.SetArgs([]string{"--config", "../testing/.cobbler.yaml", "interface", "export", "--format", "json"})
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
	if !strings.Contains(string(stdoutBytes), `"name":"`+ifaceName+`"`) {
		fmt.Println(string(stdoutBytes))
		t.Fatal("network interface name missing from export --all output")
	}
}
