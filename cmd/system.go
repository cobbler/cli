// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "System management",
	Long: `Let you manage systems.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-system for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add system",
	Long:  `Adds a system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy system",
	Long:  `Copies a system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemDumpVarsCmd = &cobra.Command{
	Use:   "dumpvars",
	Short: "dump system variables",
	Long:  `Prints all system variables to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit system",
	Long:  `Edits a system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find system",
	Long:  `Finds a given system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemGetAutoinstallCmd = &cobra.Command{
	Use:   "get-autoinstall",
	Short: "dump autoinstall XML",
	Long:  `Prints the autoinstall XML file of the given system to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all systems",
	Long:  `Lists all available systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemPowerOffCmd = &cobra.Command{
	Use:   "poweroff",
	Short: "power off system",
	Long:  `Powers off the selected system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemPowerOnCmd = &cobra.Command{
	Use:   "poweron",
	Short: "power on system",
	Long:  `Powers on the selected system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemPowerStatusCmd = &cobra.Command{
	Use:   "powerstatus",
	Short: "Power status of the system",
	Long:  `Querys the power status of the selected system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "reboot system",
	Long:  `Reboots the selected system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove system",
	Long:  `Removes a given system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename system",
	Long:  `Renames a given system.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var systemReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all systems in detail",
	Long:  `Shows detailed information about all systems.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
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
	systemAddCmd.Flags().String("ctime", "", "")
	systemAddCmd.Flags().String("depth", "", "")
	systemAddCmd.Flags().String("mtime", "", "")
	systemAddCmd.Flags().String("uid", "", "UID")
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
	systemAddCmd.Flags().Bool("ipv6-autoconfiguration", false, "IPv6 auto configuration")
	systemAddCmd.Flags().Bool("repos-enabled", false, "(re)configure local repos on this machine at next config update?")
	systemAddCmd.Flags().String("gateway", "", "gateway")
	systemAddCmd.Flags().String("hostname", "", "hostname")
	systemAddCmd.Flags().String("image", "", "parent image (if not a profile)")
	systemAddCmd.Flags().String("ipv6-default-device", "", "IPv6 default device")
	systemAddCmd.Flags().Bool("netboot-enabled", false, "PXE (re)install this machine at next boot?")
	systemAddCmd.Flags().String("power-address", "", "power management address (e.g. power-device.example.org)")
	systemAddCmd.Flags().String("power-id", "", "power management ID (usually a plug number or blade name, if power type requires it)")
	systemAddCmd.Flags().String("power-pass", "", "power management password")
	systemAddCmd.Flags().String("power-type", "", `power management script to use. Valid options: aliyun,alom,amt,apc,apc_snmp,azure_arm,
bladecenter,brocade,cdu,cisco_mds,cisco_ucs,compute,docker,drac5,eaton_snmp,emerson,eps,
evacuate,gce,hds_cb,hpblade,ibmblade,ibmz,idrac,ifmib,ilo,ilo2,ilo3,ilo3_ssh,ilo4,ilo4_ssh,
ilo5,ilo5_ssh,ilo_moonshot,ilo_mp,ilo_ssh,imm,intelmodular,ipdu,ipmilan,ipmilanplus,ironic,
kdump,ldom,lpar,mpath,netio,openstack,powerman,pve,raritan,rcd_serial,redfish,rhevm,rsa,rsb,
sanbox2,sbd,scsi,tripplite_snmp,vbox,virsh,vmware,vmware_rest,wti,xenapi,zvm,zvmip`)
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
	systemAddCmd.Flags().String("management", "", "Defines the management interface (should be used with --interface) ")
	systemAddCmd.Flags().String("mtu", "", "MTU")
	systemAddCmd.Flags().String("netmask", "", "subnet mask (should be used with --interface)")
	systemAddCmd.Flags().Bool("static", false, "Is this interface static? (should be used with --interface)")
	systemAddCmd.Flags().String("static-routes", "", "static routes (should be used with --interface)")
	systemAddCmd.Flags().String("interface", "", "the interface to operate on")
	systemAddCmd.Flags().Bool("delete-interface", false, "delete the given interface (should be used with --interface)")
	systemAddCmd.Flags().String("rename-interface", "", "rename the given interface (should be used with --interface)")

	// local flags for system copy
	systemCopyCmd.Flags().String("name", "", "the system name")
	systemCopyCmd.Flags().String("ctime", "", "")
	systemCopyCmd.Flags().String("depth", "", "")
	systemCopyCmd.Flags().String("mtime", "", "")
	systemCopyCmd.Flags().String("uid", "", "UID")
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
	systemCopyCmd.Flags().Bool("ipv6-autoconfiguration", false, "IPv6 auto configuration")
	systemCopyCmd.Flags().Bool("repos-enabled", false, "(re)configure local repos on this machine at next config update?")
	systemCopyCmd.Flags().String("gateway", "", "gateway")
	systemCopyCmd.Flags().String("hostname", "", "hostname")
	systemCopyCmd.Flags().String("image", "", "parent image (if not a profile)")
	systemCopyCmd.Flags().String("ipv6-default-device", "", "IPv6 default device")
	systemCopyCmd.Flags().Bool("netboot-enabled", false, "PXE (re)install this machine at next boot?")
	systemCopyCmd.Flags().String("power-address", "", "power management address (e.g. power-device.example.org)")
	systemCopyCmd.Flags().String("power-id", "", "power management ID (usually a plug number or blade name, if power type requires it)")
	systemCopyCmd.Flags().String("power-pass", "", "power management password")
	systemCopyCmd.Flags().String("power-type", "", `power management script to use. Valid options: aliyun,alom,amt,apc,apc_snmp,azure_arm,
bladecenter,brocade,cdu,cisco_mds,cisco_ucs,compute,docker,drac5,eaton_snmp,emerson,eps,
evacuate,gce,hds_cb,hpblade,ibmblade,ibmz,idrac,ifmib,ilo,ilo2,ilo3,ilo3_ssh,ilo4,ilo4_ssh,
ilo5,ilo5_ssh,ilo_moonshot,ilo_mp,ilo_ssh,imm,intelmodular,ipdu,ipmilan,ipmilanplus,ironic,
kdump,ldom,lpar,mpath,netio,openstack,powerman,pve,raritan,rcd_serial,redfish,rhevm,rsa,rsb,
sanbox2,sbd,scsi,tripplite_snmp,vbox,virsh,vmware,vmware_rest,wti,xenapi,zvm,zvmip`)
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
	systemCopyCmd.Flags().String("management", "", "Defines the management interface (should be used with --interface) ")
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
	systemEditCmd.Flags().String("ctime", "", "")
	systemEditCmd.Flags().String("depth", "", "")
	systemEditCmd.Flags().String("mtime", "", "")
	systemEditCmd.Flags().String("uid", "", "UID")
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
	systemEditCmd.Flags().Bool("ipv6-autoconfiguration", false, "IPv6 auto configuration")
	systemEditCmd.Flags().Bool("repos-enabled", false, "(re)configure local repos on this machine at next config update?")
	systemEditCmd.Flags().String("gateway", "", "gateway")
	systemEditCmd.Flags().String("hostname", "", "hostname")
	systemEditCmd.Flags().String("image", "", "parent image (if not a profile)")
	systemEditCmd.Flags().String("ipv6-default-device", "", "IPv6 default device")
	systemEditCmd.Flags().Bool("netboot-enabled", false, "PXE (re)install this machine at next boot?")
	systemEditCmd.Flags().String("power-address", "", "power management address (e.g. power-device.example.org)")
	systemEditCmd.Flags().String("power-id", "", "power management ID (usually a plug number or blade name, if power type requires it)")
	systemEditCmd.Flags().String("power-pass", "", "power management password")
	systemEditCmd.Flags().String("power-type", "", `power management script to use. Valid options: aliyun,alom,amt,apc,apc_snmp,azure_arm,
bladecenter,brocade,cdu,cisco_mds,cisco_ucs,compute,docker,drac5,eaton_snmp,emerson,eps,
evacuate,gce,hds_cb,hpblade,ibmblade,ibmz,idrac,ifmib,ilo,ilo2,ilo3,ilo3_ssh,ilo4,ilo4_ssh,
ilo5,ilo5_ssh,ilo_moonshot,ilo_mp,ilo_ssh,imm,intelmodular,ipdu,ipmilan,ipmilanplus,ironic,
kdump,ldom,lpar,mpath,netio,openstack,powerman,pve,raritan,rcd_serial,redfish,rhevm,rsa,rsb,
sanbox2,sbd,scsi,tripplite_snmp,vbox,virsh,vmware,vmware_rest,wti,xenapi,zvm,zvmip`)
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
	systemEditCmd.Flags().String("management", "", "Defines the management interface (should be used with --interface) ")
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
	systemFindCmd.Flags().String("power-type", "", `power management script to use. Valid options: aliyun,alom,amt,apc,apc_snmp,azure_arm,
bladecenter,brocade,cdu,cisco_mds,cisco_ucs,compute,docker,drac5,eaton_snmp,emerson,eps,
evacuate,gce,hds_cb,hpblade,ibmblade,ibmz,idrac,ifmib,ilo,ilo2,ilo3,ilo3_ssh,ilo4,ilo4_ssh,
ilo5,ilo5_ssh,ilo_moonshot,ilo_mp,ilo_ssh,imm,intelmodular,ipdu,ipmilan,ipmilanplus,ironic,
kdump,ldom,lpar,mpath,netio,openstack,powerman,pve,raritan,rcd_serial,redfish,rhevm,rsa,rsb,
sanbox2,sbd,scsi,tripplite_snmp,vbox,virsh,vmware,vmware_rest,wti,xenapi,zvm,zvmip`)
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
	systemFindCmd.Flags().String("management", "", "Defines the management interface (should be used with --interface) ")
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
	systemRenameCmd.Flags().String("ctime", "", "")
	systemRenameCmd.Flags().String("depth", "", "")
	systemRenameCmd.Flags().String("mtime", "", "")
	systemRenameCmd.Flags().String("uid", "", "UID")
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
	systemRenameCmd.Flags().Bool("ipv6-autoconfiguration", false, "IPv6 auto configuration")
	systemRenameCmd.Flags().Bool("repos-enabled", false, "(re)configure local repos on this machine at next config update?")
	systemRenameCmd.Flags().String("gateway", "", "gateway")
	systemRenameCmd.Flags().String("hostname", "", "hostname")
	systemRenameCmd.Flags().String("image", "", "parent image (if not a profile)")
	systemRenameCmd.Flags().String("ipv6-default-device", "", "IPv6 default device")
	systemRenameCmd.Flags().Bool("netboot-enabled", false, "PXE (re)install this machine at next boot?")
	systemRenameCmd.Flags().String("power-address", "", "power management address (e.g. power-device.example.org)")
	systemRenameCmd.Flags().String("power-id", "", "power management ID (usually a plug number or blade name, if power type requires it)")
	systemRenameCmd.Flags().String("power-pass", "", "power management password")
	systemRenameCmd.Flags().String("power-type", "", `power management script to use. Valid options: aliyun,alom,amt,apc,apc_snmp,azure_arm,
bladecenter,brocade,cdu,cisco_mds,cisco_ucs,compute,docker,drac5,eaton_snmp,emerson,eps,
evacuate,gce,hds_cb,hpblade,ibmblade,ibmz,idrac,ifmib,ilo,ilo2,ilo3,ilo3_ssh,ilo4,ilo4_ssh,
ilo5,ilo5_ssh,ilo_moonshot,ilo_mp,ilo_ssh,imm,intelmodular,ipdu,ipmilan,ipmilanplus,ironic,
kdump,ldom,lpar,mpath,netio,openstack,powerman,pve,raritan,rcd_serial,redfish,rhevm,rsa,rsb,
sanbox2,sbd,scsi,tripplite_snmp,vbox,virsh,vmware,vmware_rest,wti,xenapi,zvm,zvmip`)
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
	systemRenameCmd.Flags().String("management", "", "Defines the management interface (should be used with --interface) ")
	systemRenameCmd.Flags().String("mtu", "", "MTU")
	systemRenameCmd.Flags().String("netmask", "", "subnet mask (should be used with --interface)")
	systemRenameCmd.Flags().Bool("static", false, "Is this interface static? (should be used with --interface)")
	systemRenameCmd.Flags().String("static-routes", "", "static routes (should be used with --interface)")
	systemRenameCmd.Flags().String("interface", "", "the interface to operate on")

	// local flags for system report
	systemReportCmd.Flags().String("name", "", "the system name")
}
