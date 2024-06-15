// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	cobbler "github.com/cobbler/cobblerclient"
)

var system *cobbler.System //nolint:golint,unused
var systems []*cobbler.System
var iface cobbler.Interface

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "System management",
	Long: `Let you manage systems.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-system for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var systemAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add system",
	Long:  `Adds a system.`,
	Run: func(cmd *cobra.Command, args []string) {

		var newSystem cobbler.System

		// internal fields (ctime, mtime, depth, uid, repos-enabled, ipv6-autoconfiguration) cannot be modified
		newSystem.Autoinstall, _ = cmd.Flags().GetString("autoinstall")
		newSystem.AutoinstallMeta, _ = cmd.Flags().GetStringArray("autoinstall-meta")
		newSystem.BootFiles, _ = cmd.Flags().GetString("bootfiles")
		newSystem.Comment, _ = cmd.Flags().GetString("comment")
		newSystem.EnableGPXE, _ = cmd.Flags().GetBool("enable-ipxe")
		newSystem.FetchableFiles, _ = cmd.Flags().GetStringArray("fetchable-files")
		newSystem.Gateway, _ = cmd.Flags().GetString("gateway")
		newSystem.Hostname, _ = cmd.Flags().GetString("hostname")
		newSystem.Image, _ = cmd.Flags().GetString("image")
		newSystem.IPv6DefaultDevice, _ = cmd.Flags().GetString("ipv6-default-device")
		newSystem.KernelOptions, _ = cmd.Flags().GetStringArray("kernel-options")
		newSystem.KernelOptionsPost, _ = cmd.Flags().GetStringArray("kernel-options-post")
		newSystem.MGMTClasses, _ = cmd.Flags().GetStringArray("mgmt-classes")
		newSystem.MGMTParameters, _ = cmd.Flags().GetString("mgmt-parameters")
		newSystem.Name, _ = cmd.Flags().GetString("name")
		newSystem.NameServers, _ = cmd.Flags().GetStringArray("name-servers")
		newSystem.NameServersSearch, _ = cmd.Flags().GetStringArray("name-servers-search")
		newSystem.NetbootEnabled, _ = cmd.Flags().GetBool("netboot-enabled")
		newSystem.NextServerv4, _ = cmd.Flags().GetString("next-server-v4")
		newSystem.Owners, _ = cmd.Flags().GetStringArray("owners")
		newSystem.PowerAddress, _ = cmd.Flags().GetString("power-address")
		newSystem.PowerID, _ = cmd.Flags().GetString("power-id")
		newSystem.PowerPass, _ = cmd.Flags().GetString("power-pass")
		newSystem.PowerType, _ = cmd.Flags().GetString("power-type")
		newSystem.PowerUser, _ = cmd.Flags().GetString("power-user")
		newSystem.Profile, _ = cmd.Flags().GetString("profile")
		newSystem.Proxy, _ = cmd.Flags().GetString("proxy")
		// newSystem.RedHatManagementKey, _ = cmd.Flags().GetString("redhat-management-key")
		newSystem.Status, _ = cmd.Flags().GetString("status")
		newSystem.TemplateFiles, _ = cmd.Flags().GetStringArray("template-files")
		newSystem.VirtAutoBoot, _ = cmd.Flags().GetString("virt-auto-boot")
		newSystem.VirtCPUs, _ = cmd.Flags().GetString("virt-cpus")
		newSystem.VirtDiskDriver, _ = cmd.Flags().GetString("virt-disk-driver")
		newSystem.VirtFileSize, _ = cmd.Flags().GetString("virt-file-size")
		newSystem.VirtPath, _ = cmd.Flags().GetString("virt-path")
		newSystem.VirtPXEBoot, _ = cmd.Flags().GetInt("virt-pxe-boot")
		newSystem.VirtRAM, _ = cmd.Flags().GetString("virt-ram")
		newSystem.VirtType, _ = cmd.Flags().GetString("virt-type")

		// interface type
		iface.CNAMEs, _ = cmd.Flags().GetStringArray("cnames")
		iface.DHCPTag, _ = cmd.Flags().GetString("dhcp-tag")
		iface.DNSName, _ = cmd.Flags().GetString("dns-name")
		iface.BondingOpts, _ = cmd.Flags().GetString("bonding-opts")
		//iface.BridgeOpts, _ = cmd.Flags().GetString("bridge-opts")
		iface.Gateway, _ = cmd.Flags().GetString("gateway")
		iface.InterfaceType, _ = cmd.Flags().GetString("interface-type")
		iface.InterfaceMaster, _ = cmd.Flags().GetString("interface-master")
		iface.IPAddress, _ = cmd.Flags().GetString("ip-address")
		iface.IPv6Address, _ = cmd.Flags().GetString("ipv6-address")
		iface.IPv6Secondaries, _ = cmd.Flags().GetStringArray("ipv6-secondaries")
		iface.IPv6MTU, _ = cmd.Flags().GetString("ipv6-mtu")
		iface.IPv6StaticRoutes, _ = cmd.Flags().GetStringArray("ipv6-static-routes")
		iface.IPv6DefaultGateway, _ = cmd.Flags().GetString("ipv6-default-gateway")
		iface.MACAddress, _ = cmd.Flags().GetString("mac-address")
		iface.Management, _ = cmd.Flags().GetBool("management")
		iface.Netmask, _ = cmd.Flags().GetString("netmask")
		iface.Static, _ = cmd.Flags().GetBool("static")
		iface.StaticRoutes, _ = cmd.Flags().GetStringArray("static-routes")
		iface.VirtBridge, _ = cmd.Flags().GetString("virt-bridge")

		// TODO: Implementation for more interfaces
		// See https://github.com/cobbler/cli/issues/38
		err = newSystem.CreateInterface("default", iface)
		if checkError(err) != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		system, err = Client.CreateSystem(newSystem)
		if checkError(err) == nil {
			fmt.Printf("System %s created", newSystem.Name)
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

var systemCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy system",
	Long:  `Copies a system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var systemDumpVarsCmd = &cobra.Command{
	Use:   "dumpvars",
	Short: "dump system variables",
	Long:  `Prints all system variables to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var systemEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit system",
	Long:  `Edits a system.`,
	Run: func(cmd *cobra.Command, args []string) {

		// find profile through its name
		pname, _ := cmd.Flags().GetString("name")
		var updateSystem, err = Client.GetSystem(pname)

		if checkError(err) != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		// internal fields (ctime, mtime, depth, uid, repos-enabled, ipv6-autoconfiguration) cannot be modified
		var tmpArgs, _ = cmd.Flags().GetString("autoinstall")
		if tmpArgs != "" {
			updateSystem.Autoinstall, _ = cmd.Flags().GetString("autoinstall")
		}
		var tmpArgsArray, _ = cmd.Flags().GetStringArray("autoinstall-meta")
		if len(tmpArgsArray) > 0 {
			updateSystem.AutoinstallMeta, _ = cmd.Flags().GetStringArray("autoinstall-meta")
		}
		tmpArgs, _ = cmd.Flags().GetString("bootfiles")
		if tmpArgs != "" {
			updateSystem.BootFiles, _ = cmd.Flags().GetString("bootfiles")
		}
		tmpArgs, _ = cmd.Flags().GetString("comment")
		if tmpArgs != "" {
			updateSystem.Comment, _ = cmd.Flags().GetString("comment")
		}
		// TODO
		/*
			var tmpArgsBool, _ = cmd.Flags().GetBool("enable-ipxe")
			if tmpArgsBool != "" {
				updateSystem.EnableGPXE, _ = cmd.Flags().GetBool("enable-ipxe")
			}
		*/
		tmpArgsArray, _ = cmd.Flags().GetStringArray("fetchable-files")
		if len(tmpArgsArray) > 0 {
			updateSystem.FetchableFiles, _ = cmd.Flags().GetStringArray("fetchable-files")
		}
		tmpArgs, _ = cmd.Flags().GetString("gateway")
		if tmpArgs != "" {
			updateSystem.Gateway, _ = cmd.Flags().GetString("gateway")
		}
		tmpArgs, _ = cmd.Flags().GetString("hostname")
		if tmpArgs != "" {
			updateSystem.Hostname, _ = cmd.Flags().GetString("hostname")
		}
		tmpArgs, _ = cmd.Flags().GetString("image")
		if tmpArgs != "" {
			updateSystem.Image, _ = cmd.Flags().GetString("image")
		}
		tmpArgs, _ = cmd.Flags().GetString("ipv6-default-device")
		if tmpArgs != "" {
			updateSystem.IPv6DefaultDevice, _ = cmd.Flags().GetString("ipv6-default-device")
		}
		tmpArgsArray, _ = cmd.Flags().GetStringArray("kernel-options")
		if len(tmpArgsArray) > 0 {
			updateSystem.KernelOptions, _ = cmd.Flags().GetStringArray("kernel-options")
		}
		tmpArgsArray, _ = cmd.Flags().GetStringArray("kernel-options-post")
		if len(tmpArgsArray) > 0 {
			updateSystem.KernelOptionsPost, _ = cmd.Flags().GetStringArray("kernel-options-post")
		}
		tmpArgsArray, _ = cmd.Flags().GetStringArray("mgmt-classes")
		if len(tmpArgsArray) > 0 {
			updateSystem.MGMTClasses, _ = cmd.Flags().GetStringArray("mgmt-classes")
		}
		tmpArgs, _ = cmd.Flags().GetString("mgmt-parameters")
		if tmpArgs != "" {
			updateSystem.MGMTParameters, _ = cmd.Flags().GetString("mgmt-parameters")
		}
		tmpArgs, _ = cmd.Flags().GetString("name")
		if tmpArgs != "" {
			updateSystem.Name, _ = cmd.Flags().GetString("name")
		}
		tmpArgsArray, _ = cmd.Flags().GetStringArray("name-servers")
		if len(tmpArgsArray) > 0 {
			updateSystem.NameServers, _ = cmd.Flags().GetStringArray("name-servers")
		}
		tmpArgsArray, _ = cmd.Flags().GetStringArray("name-servers-search")
		if len(tmpArgsArray) > 0 {
			updateSystem.NameServersSearch, _ = cmd.Flags().GetStringArray("name-servers-search")
		}
		// TODO
		/*
			tmpArgsBool, _ = cmd.Flags().GetBool("netboot-enabled")
			if tmpArgsBool != "" {
				updateSystem.NetbootEnabled, _ = cmd.Flags().GetBool("netboot-enabled")
			}
		*/
		tmpArgs, _ = cmd.Flags().GetString("next-servers")
		if tmpArgs != "" {
			updateSystem.NextServerv4, _ = cmd.Flags().GetString("next-server-v4")
		}
		tmpArgsArray, _ = cmd.Flags().GetStringArray("owners")
		if len(tmpArgsArray) > 0 {
			updateSystem.Owners, _ = cmd.Flags().GetStringArray("owners")
		}
		tmpArgs, _ = cmd.Flags().GetString("power-address")
		if tmpArgs != "" {
			updateSystem.PowerAddress, _ = cmd.Flags().GetString("power-address")
		}
		tmpArgs, _ = cmd.Flags().GetString("power-id")
		if tmpArgs != "" {
			updateSystem.PowerID, _ = cmd.Flags().GetString("power-id")
		}
		tmpArgs, _ = cmd.Flags().GetString("power-pass")
		if tmpArgs != "" {
			updateSystem.PowerPass, _ = cmd.Flags().GetString("power-pass")
		}
		tmpArgs, _ = cmd.Flags().GetString("power-type")
		if tmpArgs != "" {
			updateSystem.PowerType, _ = cmd.Flags().GetString("power-type")
		}
		tmpArgs, _ = cmd.Flags().GetString("power-user")
		if tmpArgs != "" {
			updateSystem.PowerUser, _ = cmd.Flags().GetString("power-user")
		}
		tmpArgs, _ = cmd.Flags().GetString("profile")
		if tmpArgs != "" {
			updateSystem.Profile, _ = cmd.Flags().GetString("profile")
		}
		tmpArgs, _ = cmd.Flags().GetString("proxy")
		if tmpArgs != "" {
			updateSystem.Proxy, _ = cmd.Flags().GetString("proxy")
		}
		/*
			tmpArgs, _ = cmd.Flags().GetString("redhat-management-key")
			if tmpArgs != "" {
				updateSystem.RedHatManagementKey, _ = cmd.Flags().GetString("redhat-management-key")
			}
		*/
		tmpArgs, _ = cmd.Flags().GetString("status")
		if tmpArgs != "" {
			updateSystem.Status, _ = cmd.Flags().GetString("status")
		}
		tmpArgsArray, _ = cmd.Flags().GetStringArray("template-files")
		if len(tmpArgsArray) > 0 {
			updateSystem.TemplateFiles, _ = cmd.Flags().GetStringArray("template-files")
		}
		tmpArgs, _ = cmd.Flags().GetString("virt-auto-boot")
		if tmpArgs != "" {
			updateSystem.VirtAutoBoot, _ = cmd.Flags().GetString("virt-auto-boot")
		}
		tmpArgs, _ = cmd.Flags().GetString("virt-cpus")
		if tmpArgs != "" {
			updateSystem.VirtCPUs, _ = cmd.Flags().GetString("virt-cpus")
		}
		tmpArgs, _ = cmd.Flags().GetString("virt-disk-driver")
		if tmpArgs != "" {
			updateSystem.VirtDiskDriver, _ = cmd.Flags().GetString("virt-disk-driver")
		}
		tmpArgs, _ = cmd.Flags().GetString("virt-file-size")
		if tmpArgs != "" {
			updateSystem.VirtFileSize, _ = cmd.Flags().GetString("virt-file-size")
		}
		tmpArgs, _ = cmd.Flags().GetString("virt-path")
		if tmpArgs != "" {
			updateSystem.VirtPath, _ = cmd.Flags().GetString("virt-path")
		}

		// FIXME: what happens when the int value is acutally 0 instead of 1?
		var tmpArgsInt, _ = cmd.Flags().GetInt("virt-pxe-boot")
		if tmpArgsInt != 0 {
			updateSystem.VirtPXEBoot, _ = cmd.Flags().GetInt("virt-pxe-boot")
		}
		tmpArgs, _ = cmd.Flags().GetString("virt-ram")
		if tmpArgs != "" {
			updateSystem.VirtRAM, _ = cmd.Flags().GetString("virt-ram")
		}
		tmpArgs, _ = cmd.Flags().GetString("virt-type")
		if tmpArgs != "" {
			updateSystem.VirtType, _ = cmd.Flags().GetString("virt-type")
		}

		err = Client.UpdateSystem(updateSystem)

		if checkError(err) != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

var systemFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find system",
	Long:  `Finds a given system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var systemGetAutoinstallCmd = &cobra.Command{
	Use:   "get-autoinstall",
	Short: "dump autoinstall XML",
	Long:  `Prints the autoinstall XML file of the given system to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var systemListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all systems",
	Long:  `Lists all available systems.`,
	Run: func(cmd *cobra.Command, args []string) {

		systems, err = Client.GetSystems()

		if checkError(err) == nil {
			fmt.Println(systems)
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

var systemPowerOffCmd = &cobra.Command{
	Use:   "poweroff",
	Short: "power off system",
	Long:  `Powers off the selected system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var systemPowerOnCmd = &cobra.Command{
	Use:   "poweron",
	Short: "power on system",
	Long:  `Powers on the selected system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var systemPowerStatusCmd = &cobra.Command{
	Use:   "powerstatus",
	Short: "Power status of the system",
	Long:  `Querys the power status of the selected system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var systemRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "reboot system",
	Long:  `Reboots the selected system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var systemRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove system",
	Long:  `Removes a given system.`,
	Run: func(cmd *cobra.Command, args []string) {

		sname, _ := cmd.Flags().GetString("name")
		err := Client.DeleteSystem(sname)
		if checkError(err) != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

var systemRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename system",
	Long:  `Renames a given system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var systemReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all systems in detail",
	Long:  `Shows detailed information about all systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)
	systemCmd.AddCommand(systemAddCmd)
	systemCmd.AddCommand(systemCopyCmd)
	systemCmd.AddCommand(systemDumpVarsCmd)
	systemCmd.AddCommand(systemEditCmd)
	systemCmd.AddCommand(systemFindCmd)
	systemCmd.AddCommand(systemGetAutoinstallCmd)
	systemCmd.AddCommand(systemListCmd)
	systemCmd.AddCommand(systemPowerOffCmd)
	systemCmd.AddCommand(systemPowerOnCmd)
	systemCmd.AddCommand(systemPowerStatusCmd)
	systemCmd.AddCommand(systemRebootCmd)
	systemCmd.AddCommand(systemRemoveCmd)
	systemCmd.AddCommand(systemRenameCmd)
	systemCmd.AddCommand(systemReportCmd)

	// local flags for system add
	systemAddCmd.Flags().String("name", "", "the system name")
	systemAddCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	systemAddCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	systemAddCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	systemAddCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	systemAddCmd.Flags().String("comment", "", "free form text description")
	systemAddCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	systemAddCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	systemAddCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	systemAddCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	systemAddCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	systemAddCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	systemAddCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	systemAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	systemAddCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	systemAddCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	systemAddCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this system in the PXE menu?)")
	systemAddCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	systemAddCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	systemAddCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	systemAddCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	systemAddCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	systemAddCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	systemAddCmd.Flags().String("parent", "", "parent system")
	systemAddCmd.Flags().String("proxy", "", "proxy server URL")
	systemAddCmd.Flags().String("server", "", "server override")
	systemAddCmd.Flags().String("menu", "", "parent boot menu")
	systemAddCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	systemAddCmd.Flags().String("virt-bridge", "", "virt bridge")
	systemAddCmd.Flags().String("virt-cpus", "", "virt CPUs")
	systemAddCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	systemAddCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	systemAddCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	systemAddCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	systemAddCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")
	systemAddCmd.Flags().String("gateway", "", "gateway")
	systemAddCmd.Flags().String("hostname", "", "hostname")
	systemAddCmd.Flags().String("image", "", "parent image (if not a profile)")
	systemAddCmd.Flags().String("ipv6-default-device", "", "IPv6 default device")
	systemAddCmd.Flags().Bool("netboot-enabled", false, "PXE (re)install this machine at next boot?")
	systemAddCmd.Flags().String("power-address", "", "power management address (e.g. power-device.example.org)")
	systemAddCmd.Flags().String("power-id", "", "power management ID (usually a plug number or blade name, if power type requires it)")
	systemAddCmd.Flags().String("power-pass", "", "power management password")
	systemAddCmd.Flags().String("power-type", "", "power management script to use")
	systemAddCmd.Flags().String("power-user", "", "power management username")
	systemAddCmd.Flags().String("power-options", "", "additional options, to be passed to the fencing agent")
	systemAddCmd.Flags().String("power-identity-file", "", "identity file to be passed to the fencing agent (SSH key)")
	systemAddCmd.Flags().String("profile", "", "Parent profile")
	systemAddCmd.Flags().String("status", "", "system status. Valid options: development,testing,acceptance,production")
	systemAddCmd.Flags().Bool("virt-pxe-boot", false, "use PXE to build this VM?")
	systemAddCmd.Flags().String("serial-device", "", "serial device number")
	systemAddCmd.Flags().String("serial-baud-rate", "", "serial Baud Rate. Valid options: 2400,4800,9600,19200,38400,57600,115200")
	systemAddCmd.Flags().String("bonding-opts", "", "bonding opts (should be used with --interface)")
	systemAddCmd.Flags().String("ridge-opts", "", "bridge opts (should be used with --interface")
	systemAddCmd.Flags().String("cnames", "", "Cannonical Name Records, should be used with	--interface (in quotes, space delimited)")
	systemAddCmd.Flags().String("connected-mode", "", "InfiniBand connected mode (should be used with --interface)")
	systemAddCmd.Flags().String("dns-name", "", "DNS name (should be used with --interface)")
	systemAddCmd.Flags().String("if-gateway", "", "per-Interface Gateway (should be used with --interface)")
	systemAddCmd.Flags().String("interface-master", "", "master interface (Should be used with --interface)")
	systemAddCmd.Flags().String("interface-type", "", `interface Type. Valid options: na,bond,bond_slave,bridge,bridge_slave,bonded_bridge_slave,bmc,infiniband.
	(should be used with --interface)`)
	systemAddCmd.Flags().String("ip-address", "", "IPv4 address (should be used with --interface)")
	systemAddCmd.Flags().String("ipv6-address", "", "IPv6 address (should be used with --interface)")
	systemAddCmd.Flags().String("ipv6-default-gateway", "", "IPv6 Default Gateway (should be used with --interface)")
	systemAddCmd.Flags().String("ipv6-mtu", "", "IPv6 MTU")
	systemAddCmd.Flags().String("ipv6-prefix", "", "IPv6 Prefix (should be used with --interface)")
	systemAddCmd.Flags().String("ipv6-secondaries", "", "IPv6 Secondaries (should be used with --interface)")
	systemAddCmd.Flags().String("ipv6-static-routes", "", "IPv6 Static Routes (should be used with --interface)")
	systemAddCmd.Flags().String("mac-address", "", "MAC Address (place 'random' in this field for a	random MAC Address.)")
	systemAddCmd.Flags().Bool("management", false, "declares the interface as management interface (should be used with --interface) ")
	systemAddCmd.Flags().String("mtu", "", "MTU")
	systemAddCmd.Flags().String("netmask", "", "subnet mask (should be used with --interface)")
	systemAddCmd.Flags().Bool("static", false, "Is this interface static? (should be used with --interface)")
	systemAddCmd.Flags().String("static-routes", "", "static routes (should be used with --interface)")
	systemAddCmd.Flags().String("interface", "", "the interface to operate on")
	systemAddCmd.Flags().Bool("delete-interface", false, "delete the given interface (should be used with --interface)")
	systemAddCmd.Flags().String("rename-interface", "", "rename the given interface (should be used with --interface)")

	// local flags for system copy
	systemCopyCmd.Flags().String("name", "", "the system name")
	systemCopyCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	systemCopyCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	systemCopyCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	systemCopyCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	systemCopyCmd.Flags().String("comment", "", "free form text description")
	systemCopyCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	systemCopyCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	systemCopyCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	systemCopyCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	systemCopyCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	systemCopyCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	systemCopyCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	systemCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	systemCopyCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	systemCopyCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	systemCopyCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this system in the PXE menu?)")
	systemCopyCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	systemCopyCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	systemCopyCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	systemCopyCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	systemCopyCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	systemCopyCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	systemCopyCmd.Flags().String("parent", "", "parent system")
	systemCopyCmd.Flags().String("proxy", "", "proxy server URL")
	systemCopyCmd.Flags().String("server", "", "server override")
	systemCopyCmd.Flags().String("menu", "", "parent boot menu")
	systemCopyCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	systemCopyCmd.Flags().String("virt-bridge", "", "virt bridge")
	systemCopyCmd.Flags().String("virt-cpus", "", "virt CPUs")
	systemCopyCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	systemCopyCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	systemCopyCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	systemCopyCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	systemCopyCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")
	systemCopyCmd.Flags().String("gateway", "", "gateway")
	systemCopyCmd.Flags().String("hostname", "", "hostname")
	systemCopyCmd.Flags().String("image", "", "parent image (if not a profile)")
	systemCopyCmd.Flags().String("ipv6-default-device", "", "IPv6 default device")
	systemCopyCmd.Flags().Bool("netboot-enabled", false, "PXE (re)install this machine at next boot?")
	systemCopyCmd.Flags().String("power-address", "", "power management address (e.g. power-device.example.org)")
	systemCopyCmd.Flags().String("power-id", "", "power management ID (usually a plug number or blade name, if power type requires it)")
	systemCopyCmd.Flags().String("power-pass", "", "power management password")
	systemCopyCmd.Flags().String("power-type", "", "power management script to use")
	systemCopyCmd.Flags().String("power-user", "", "power management username")
	systemCopyCmd.Flags().String("power-options", "", "additional options, to be passed to the fencing agent")
	systemCopyCmd.Flags().String("power-identity-file", "", "identity file to be passed to the fencing agent (SSH key)")
	systemCopyCmd.Flags().String("profile", "", "Parent profile")
	systemCopyCmd.Flags().String("status", "", "system status. Valid options: development,testing,acceptance,production")
	systemCopyCmd.Flags().Bool("virt-pxe-boot", false, "use PXE to build this VM?")
	systemCopyCmd.Flags().String("serial-device", "", "serial device number")
	systemCopyCmd.Flags().String("serial-baud-rate", "", "serial Baud Rate. Valid options: 2400,4800,9600,19200,38400,57600,115200")
	systemCopyCmd.Flags().String("bonding-opts", "", "bonding opts (should be used with --interface)")
	systemCopyCmd.Flags().String("ridge-opts", "", "bridge opts (should be used with --interface")
	systemCopyCmd.Flags().String("cnames", "", "Cannonical Name Records, should be used with	--interface (in quotes, space delimited)")
	systemCopyCmd.Flags().String("connected-mode", "", "InfiniBand connected mode (should be used with --interface)")
	systemCopyCmd.Flags().String("dns-name", "", "DNS name (should be used with --interface)")
	systemCopyCmd.Flags().String("if-gateway", "", "per-Interface Gateway (should be used with --interface)")
	systemCopyCmd.Flags().String("interface-master", "", "master interface (Should be used with --interface)")
	systemCopyCmd.Flags().String("interface-type", "", `interface Type. Valid options: na,bond,bond_slave,bridge,bridge_slave,bonded_bridge_slave,bmc,infiniband.
	(should be used with --interface)`)
	systemCopyCmd.Flags().String("ip-address", "", "IPv4 address (should be used with --interface)")
	systemCopyCmd.Flags().String("ipv6-address", "", "IPv6 address (should be used with --interface)")
	systemCopyCmd.Flags().String("ipv6-default-gateway", "", "IPv6 Default Gateway (should be used with --interface)")
	systemCopyCmd.Flags().String("ipv6-mtu", "", "IPv6 MTU")
	systemCopyCmd.Flags().String("ipv6-prefix", "", "IPv6 Prefix (should be used with --interface)")
	systemCopyCmd.Flags().String("ipv6-secondaries", "", "IPv6 Secondaries (should be used with --interface)")
	systemCopyCmd.Flags().String("ipv6-static-routes", "", "IPv6 Static Routes (should be used with --interface)")
	systemCopyCmd.Flags().String("mac-address", "", "MAC Address (place 'random' in this field for a random MAC Address.)")
	systemCopyCmd.Flags().Bool("management", false, "declares the interface as management interface (should be used with --interface) ")
	systemCopyCmd.Flags().String("mtu", "", "MTU")
	systemCopyCmd.Flags().String("netmask", "", "subnet mask (should be used with --interface)")
	systemCopyCmd.Flags().Bool("static", false, "Is this interface static? (should be used with --interface)")
	systemCopyCmd.Flags().String("static-routes", "", "static routes (should be used with --interface)")
	systemCopyCmd.Flags().String("interface", "", "the interface to operate on")
	systemCopyCmd.Flags().Bool("delete-interface", false, "delete the given interface (should be used with --interface)")
	systemCopyCmd.Flags().String("rename-interface", "", "rename the given interface (should be used with --interface)")

	// local flags for system dumpvars
	systemDumpVarsCmd.Flags().String("name", "", "the system name")

	// local flags for system edit
	systemEditCmd.Flags().String("name", "", "the system name")
	systemEditCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	systemEditCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	systemEditCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	systemEditCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	systemEditCmd.Flags().String("comment", "", "free form text description")
	systemEditCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	systemEditCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	systemEditCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	systemEditCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	systemEditCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	systemEditCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	systemEditCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	systemEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	systemEditCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	systemEditCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	systemEditCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this system in the PXE menu?)")
	systemEditCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	systemEditCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	systemEditCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	systemEditCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	systemEditCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	systemEditCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	systemEditCmd.Flags().String("parent", "", "parent system")
	systemEditCmd.Flags().String("proxy", "", "proxy server URL")
	systemEditCmd.Flags().String("server", "", "server override")
	systemEditCmd.Flags().String("menu", "", "parent boot menu")
	systemEditCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	systemEditCmd.Flags().String("virt-bridge", "", "virt bridge")
	systemEditCmd.Flags().String("virt-cpus", "", "virt CPUs")
	systemEditCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	systemEditCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	systemEditCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	systemEditCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	systemEditCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")
	systemEditCmd.Flags().String("gateway", "", "gateway")
	systemEditCmd.Flags().String("hostname", "", "hostname")
	systemEditCmd.Flags().String("image", "", "parent image (if not a profile)")
	systemEditCmd.Flags().String("ipv6-default-device", "", "IPv6 default device")
	systemEditCmd.Flags().Bool("netboot-enabled", false, "PXE (re)install this machine at next boot?")
	systemEditCmd.Flags().String("power-address", "", "power management address (e.g. power-device.example.org)")
	systemEditCmd.Flags().String("power-id", "", "power management ID (usually a plug number or blade name, if power type requires it)")
	systemEditCmd.Flags().String("power-pass", "", "power management password")
	systemEditCmd.Flags().String("power-type", "", "power management script to use")
	systemEditCmd.Flags().String("power-user", "", "power management username")
	systemEditCmd.Flags().String("power-options", "", "additional options, to be passed to the fencing agent")
	systemEditCmd.Flags().String("power-identity-file", "", "identity file to be passed to the fencing agent (SSH key)")
	systemEditCmd.Flags().String("profile", "", "Parent profile")
	systemEditCmd.Flags().String("status", "", "system status. Valid options: development,testing,acceptance,production")
	systemEditCmd.Flags().Bool("virt-pxe-boot", false, "use PXE to build this VM?")
	systemEditCmd.Flags().String("serial-device", "", "serial device number")
	systemEditCmd.Flags().String("serial-baud-rate", "", "serial Baud Rate. Valid options: 2400,4800,9600,19200,38400,57600,115200")
	systemEditCmd.Flags().String("bonding-opts", "", "bonding opts (should be used with --interface)")
	systemEditCmd.Flags().String("ridge-opts", "", "bridge opts (should be used with --interface")
	systemEditCmd.Flags().String("cnames", "", "Cannonical Name Records, should be used with	--interface (in quotes, space delimited)")
	systemEditCmd.Flags().String("connected-mode", "", "InfiniBand connected mode (should be used with --interface)")
	systemEditCmd.Flags().String("dns-name", "", "DNS name (should be used with --interface)")
	systemEditCmd.Flags().String("if-gateway", "", "per-Interface Gateway (should be used with --interface)")
	systemEditCmd.Flags().String("interface-master", "", "master interface (Should be used with --interface)")
	systemEditCmd.Flags().String("interface-type", "", `interface Type. Valid options: na,bond,bond_slave,bridge,bridge_slave,bonded_bridge_slave,bmc,infiniband.
	(should be used with --interface)`)
	systemEditCmd.Flags().String("ip-address", "", "IPv4 address (should be used with --interface)")
	systemEditCmd.Flags().String("ipv6-address", "", "IPv6 address (should be used with --interface)")
	systemEditCmd.Flags().String("ipv6-default-gateway", "", "IPv6 Default Gateway (should be used with --interface)")
	systemEditCmd.Flags().String("ipv6-mtu", "", "IPv6 MTU")
	systemEditCmd.Flags().String("ipv6-prefix", "", "IPv6 Prefix (should be used with --interface)")
	systemEditCmd.Flags().String("ipv6-secondaries", "", "IPv6 Secondaries (should be used with --interface)")
	systemEditCmd.Flags().String("ipv6-static-routes", "", "IPv6 Static Routes (should be used with --interface)")
	systemEditCmd.Flags().String("mac-address", "", "MAC Address (place 'random' in this field for a random MAC Address.)")
	systemEditCmd.Flags().Bool("management", false, "declares the interface as management interface (should be used with --interface) ")
	systemEditCmd.Flags().String("mtu", "", "MTU")
	systemEditCmd.Flags().String("netmask", "", "subnet mask (should be used with --interface)")
	systemEditCmd.Flags().Bool("static", false, "Is this interface static? (should be used with --interface)")
	systemEditCmd.Flags().String("static-routes", "", "static routes (should be used with --interface)")
	systemEditCmd.Flags().String("interface", "", "the interface to operate on")

	// local flags for system find
	systemFindCmd.Flags().String("name", "", "the system name")
	systemFindCmd.Flags().String("ctime", "", "")
	systemFindCmd.Flags().String("depth", "", "")
	systemFindCmd.Flags().String("mtime", "", "")
	systemFindCmd.Flags().String("uid", "", "UID")
	systemFindCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	systemFindCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	systemFindCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	systemFindCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	systemFindCmd.Flags().String("comment", "", "free form text description")
	systemFindCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	systemFindCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	systemFindCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	systemFindCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	systemFindCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	systemFindCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	systemFindCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	systemFindCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	systemFindCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	systemFindCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this system in the PXE menu?)")
	systemFindCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	systemFindCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	systemFindCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	systemFindCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	systemFindCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	systemFindCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	systemFindCmd.Flags().String("parent", "", "parent system")
	systemFindCmd.Flags().String("proxy", "", "proxy server URL")
	systemFindCmd.Flags().String("server", "", "server override")
	systemFindCmd.Flags().String("menu", "", "parent boot menu")
	systemFindCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	systemFindCmd.Flags().String("virt-bridge", "", "virt bridge")
	systemFindCmd.Flags().String("virt-cpus", "", "virt CPUs")
	systemFindCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	systemFindCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	systemFindCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	systemFindCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	systemFindCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")
	systemFindCmd.Flags().Bool("ipv6-autoconfiguration", false, "IPv6 auto configuration")
	systemFindCmd.Flags().Bool("repos-enabled", false, "(re)configure local repos on this machine at next config update?")
	systemFindCmd.Flags().String("gateway", "", "gateway")
	systemFindCmd.Flags().String("hostname", "", "hostname")
	systemFindCmd.Flags().String("image", "", "parent image (if not a profile)")
	systemFindCmd.Flags().String("ipv6-default-device", "", "IPv6 default device")
	systemFindCmd.Flags().Bool("netboot-enabled", false, "PXE (re)install this machine at next boot?")
	systemFindCmd.Flags().String("power-address", "", "power management address (e.g. power-device.example.org)")
	systemFindCmd.Flags().String("power-id", "", "power management ID (usually a plug number or blade name, if power type requires it)")
	systemFindCmd.Flags().String("power-pass", "", "power management password")
	systemFindCmd.Flags().String("power-type", "", "power management script to use")
	systemFindCmd.Flags().String("power-user", "", "power management username")
	systemFindCmd.Flags().String("power-options", "", "additional options, to be passed to the fencing agent")
	systemFindCmd.Flags().String("power-identity-file", "", "identity file to be passed to the fencing agent (SSH key)")
	systemFindCmd.Flags().String("profile", "", "Parent profile")
	systemFindCmd.Flags().String("status", "", "system status. Valid options: development,testing,acceptance,production")
	systemFindCmd.Flags().Bool("virt-pxe-boot", false, "use PXE to build this VM?")
	systemFindCmd.Flags().String("serial-device", "", "serial device number")
	systemFindCmd.Flags().String("serial-baud-rate", "", "serial Baud Rate. Valid options: 2400,4800,9600,19200,38400,57600,115200")
	systemFindCmd.Flags().String("bonding-opts", "", "bonding opts (should be used with --interface)")
	systemFindCmd.Flags().String("ridge-opts", "", "bridge opts (should be used with --interface")
	systemFindCmd.Flags().String("cnames", "", "Cannonical Name Records, should be used with	--interface (in quotes, space delimited)")
	systemFindCmd.Flags().String("connected-mode", "", "InfiniBand connected mode (should be used with --interface)")
	systemFindCmd.Flags().String("dns-name", "", "DNS name (should be used with --interface)")
	systemFindCmd.Flags().String("if-gateway", "", "per-Interface Gateway (should be used with --interface)")
	systemFindCmd.Flags().String("interface-master", "", "master interface (Should be used with --interface)")
	systemFindCmd.Flags().String("interface-type", "", `interface Type. Valid options: na,bond,bond_slave,bridge,bridge_slave,bonded_bridge_slave,bmc,infiniband.
	(should be used with --interface)`)
	systemFindCmd.Flags().String("ip-address", "", "IPv4 address (should be used with --interface)")
	systemFindCmd.Flags().String("ipv6-address", "", "IPv6 address (should be used with --interface)")
	systemFindCmd.Flags().String("ipv6-default-gateway", "", "IPv6 Default Gateway (should be used with --interface)")
	systemFindCmd.Flags().String("ipv6-mtu", "", "IPv6 MTU")
	systemFindCmd.Flags().String("ipv6-prefix", "", "IPv6 Prefix")
	systemFindCmd.Flags().String("ipv6-secondaries", "", "IPv6 Secondaries (should be used with --interface)")
	systemFindCmd.Flags().String("ipv6-static-routes", "", "IPv6 Static Routes (should be used with --interface)")
	systemFindCmd.Flags().String("mac-address", "", "MAC Address (place 'random' in this field for a random MAC Address.)")
	systemFindCmd.Flags().Bool("management", false, "declares the interface as management interface (should be used with --interface) ")
	systemFindCmd.Flags().String("mtu", "", "MTU")
	systemFindCmd.Flags().String("netmask", "", "subnet mask (should be used with --interface)")
	systemFindCmd.Flags().Bool("static", false, "Is this interface static? (should be used with --interface)")
	systemFindCmd.Flags().String("static-routes", "", "static routes (should be used with --interface)")
	systemFindCmd.Flags().String("interface", "", "the interface to operate on")
	systemFindCmd.Flags().Bool("delete-interface", false, "delete the given interface (should be used with --interface)")
	systemFindCmd.Flags().String("rename-interface", "", "rename the given interface (should be used with --interface)")

	// local flags for system get-autoinstall
	systemGetAutoinstallCmd.Flags().String("name", "", "the system name")

	// local flags for system poweroff
	systemPowerOffCmd.Flags().String("name", "", "the system name")

	// local flags for system poweron
	systemPowerOnCmd.Flags().String("name", "", "the system name")

	// local flags for system powerstatus
	systemPowerStatusCmd.Flags().String("name", "", "the system name")

	// local flags for system reboot
	systemRebootCmd.Flags().String("name", "", "the system name")

	// local flags for system remove
	systemRemoveCmd.Flags().String("name", "", "the system name")
	systemRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for system rename
	systemRenameCmd.Flags().String("name", "", "the system name")
	systemRenameCmd.Flags().String("newname", "", "the new system name")
	systemRenameCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	systemRenameCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	systemRenameCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	systemRenameCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	systemRenameCmd.Flags().String("comment", "", "free form text description")
	systemRenameCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	systemRenameCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	systemRenameCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	systemRenameCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	systemRenameCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	systemRenameCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	systemRenameCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	systemRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	systemRenameCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	systemRenameCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	systemRenameCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this system in the PXE menu?)")
	systemRenameCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	systemRenameCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	systemRenameCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	systemRenameCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	systemRenameCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	systemRenameCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	systemRenameCmd.Flags().String("parent", "", "parent system")
	systemRenameCmd.Flags().String("proxy", "", "proxy server URL")
	systemRenameCmd.Flags().String("server", "", "server override")
	systemRenameCmd.Flags().String("menu", "", "parent boot menu")
	systemRenameCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	systemRenameCmd.Flags().String("virt-bridge", "", "virt bridge")
	systemRenameCmd.Flags().String("virt-cpus", "", "virt CPUs")
	systemRenameCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	systemRenameCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	systemRenameCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	systemRenameCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	systemRenameCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")
	systemRenameCmd.Flags().String("gateway", "", "gateway")
	systemRenameCmd.Flags().String("hostname", "", "hostname")
	systemRenameCmd.Flags().String("image", "", "parent image (if not a profile)")
	systemRenameCmd.Flags().String("ipv6-default-device", "", "IPv6 default device")
	systemRenameCmd.Flags().Bool("netboot-enabled", false, "PXE (re)install this machine at next boot?")
	systemRenameCmd.Flags().String("power-address", "", "power management address (e.g. power-device.example.org)")
	systemRenameCmd.Flags().String("power-id", "", "power management ID (usually a plug number or blade name, if power type requires it)")
	systemRenameCmd.Flags().String("power-pass", "", "power management password")
	systemRenameCmd.Flags().String("power-type", "", "power management script to use")
	systemRenameCmd.Flags().String("power-user", "", "power management username")
	systemRenameCmd.Flags().String("power-options", "", "additional options, to be passed to the fencing agent")
	systemRenameCmd.Flags().String("power-identity-file", "", "identity file to be passed to the fencing agent (SSH key)")
	systemRenameCmd.Flags().String("profile", "", "Parent profile")
	systemRenameCmd.Flags().String("status", "", "system status. Valid options: development,testing,acceptance,production")
	systemRenameCmd.Flags().Bool("virt-pxe-boot", false, "use PXE to build this VM?")
	systemRenameCmd.Flags().String("serial-device", "", "serial device number")
	systemRenameCmd.Flags().String("serial-baud-rate", "", "serial Baud Rate. Valid options: 2400,4800,9600,19200,38400,57600,115200")
	systemRenameCmd.Flags().String("bonding-opts", "", "bonding opts (should be used with --interface)")
	systemRenameCmd.Flags().String("ridge-opts", "", "bridge opts (should be used with --interface")
	systemRenameCmd.Flags().String("cnames", "", "Cannonical Name Records, should be used with	--interface (in quotes, space delimited)")
	systemRenameCmd.Flags().String("connected-mode", "", "InfiniBand connected mode (should be used with --interface)")
	systemRenameCmd.Flags().String("dns-name", "", "DNS name (should be used with --interface)")
	systemRenameCmd.Flags().String("if-gateway", "", "per-Interface Gateway (should be used with --interface)")
	systemRenameCmd.Flags().String("interface-master", "", "master interface (Should be used with --interface)")
	systemRenameCmd.Flags().String("interface-type", "", `interface Type. Valid options: na,bond,bond_slave,bridge,bridge_slave,bonded_bridge_slave,bmc,infiniband.
	(should be used with --interface)`)
	systemRenameCmd.Flags().String("ip-address", "", "IPv4 address (should be used with --interface)")
	systemRenameCmd.Flags().String("ipv6-address", "", "IPv6 address (should be used with --interface)")
	systemRenameCmd.Flags().String("ipv6-default-gateway", "", "IPv6 Default Gateway")
	systemRenameCmd.Flags().String("ipv6-mtu", "", "IPv6 MTU")
	systemRenameCmd.Flags().String("ipv6-prefix", "", "IPv6 Prefix")
	systemRenameCmd.Flags().String("ipv6-secondaries", "", "IPv6 Secondaries")
	systemRenameCmd.Flags().String("ipv6-static-routes", "", "IPv6 Static Routes")
	systemRenameCmd.Flags().String("mac-address", "", "MAC Address (place 'random' in this field for a	random MAC Address.)")
	systemRenameCmd.Flags().Bool("management", false, "declares the interface as management interface (should be used with --interface) ")
	systemRenameCmd.Flags().String("mtu", "", "MTU")
	systemRenameCmd.Flags().String("netmask", "", "subnet mask (should be used with --interface)")
	systemRenameCmd.Flags().Bool("static", false, "Is this interface static? (should be used with --interface)")
	systemRenameCmd.Flags().String("static-routes", "", "static routes (should be used with --interface)")
	systemRenameCmd.Flags().String("interface", "", "the interface to operate on")

	// local flags for system report
	systemReportCmd.Flags().String("name", "", "the system name")
}
