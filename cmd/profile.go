// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Profile management",
	Long: `Let you manage profiles.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-profile for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add profile",
	Long:  `Adds a profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy profile",
	Long:  `Copies a profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileDumpVarsCmd = &cobra.Command{
	Use:   "dumpvars",
	Short: "dump profile variables",
	Long:  `Prints all profile variables to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit profile",
	Long:  `Edits a profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find profile",
	Long:  `Finds a given profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileGetAutoinstallCmd = &cobra.Command{
	Use:   "get-autoinstall",
	Short: "dump autoinstall XML",
	Long:  `Prints the autoinstall XML file of the given profile to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all profiles",
	Long:  `Lists all available profiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove profile",
	Long:  `Removes a given profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename profile",
	Long:  `Renames a given profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var profileReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all profiles in detail",
	Long:  `Shows detailed information about all profiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.AddCommand(profileAddCmd)
	profileCmd.AddCommand(profileCopyCmd)
	profileCmd.AddCommand(profileDumpVarsCmd)
	profileCmd.AddCommand(profileEditCmd)
	profileCmd.AddCommand(profileFindCmd)
	profileCmd.AddCommand(profileGetAutoinstallCmd)
	profileCmd.AddCommand(profileListCmd)
	profileCmd.AddCommand(profileRemoveCmd)
	profileCmd.AddCommand(profileRenameCmd)
	profileCmd.AddCommand(profileReportCmd)

	// local flags for profile add
	profileAddCmd.Flags().String("name", "", "the profile name")
	profileAddCmd.Flags().String("repos", "", "(repos to auto-assign to this profile")
	profileAddCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	profileAddCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	profileAddCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	profileAddCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	profileAddCmd.Flags().String("distro", "", "the name of a previously defined Cobbler distribution. This value is required")
	profileAddCmd.Flags().String("comment", "", "free form text description")
	profileAddCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	profileAddCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	profileAddCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	profileAddCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	profileAddCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	profileAddCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	profileAddCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	profileAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	profileAddCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	profileAddCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	profileAddCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this profile in the PXE menu?)")
	profileAddCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	profileAddCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	profileAddCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	profileAddCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	profileAddCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	profileAddCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	profileAddCmd.Flags().String("parent", "", "parent profile")
	profileAddCmd.Flags().String("proxy", "", "proxy server URL")
	profileAddCmd.Flags().String("server", "", "server override")
	profileAddCmd.Flags().String("menu", "", "parent boot menu")
	profileAddCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	profileAddCmd.Flags().String("virt-bridge", "", "virt bridge")
	profileAddCmd.Flags().String("virt-cpus", "", "virt CPUs")
	profileAddCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	profileAddCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	profileAddCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	profileAddCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	profileAddCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")

	// local flags for profile copy
	profileCopyCmd.Flags().String("name", "", "the profile name")
	profileCopyCmd.Flags().String("newname", "", "the new profile name")
	profileCopyCmd.Flags().String("repos", "", "(repos to auto-assign to this profile")
	profileCopyCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	profileCopyCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	profileCopyCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	profileCopyCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	profileCopyCmd.Flags().String("distro", "", "the name of a previously defined Cobbler distribution. This value is required")
	profileCopyCmd.Flags().String("comment", "", "free form text description")
	profileCopyCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	profileCopyCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	profileCopyCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	profileCopyCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	profileCopyCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	profileCopyCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	profileCopyCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	profileCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	profileCopyCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	profileCopyCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	profileCopyCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this profile in the PXE menu?)")
	profileCopyCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	profileCopyCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	profileCopyCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	profileCopyCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	profileCopyCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	profileCopyCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	profileCopyCmd.Flags().String("parent", "", "parent profile")
	profileCopyCmd.Flags().String("proxy", "", "proxy server URL")
	profileCopyCmd.Flags().String("server", "", "server override")
	profileCopyCmd.Flags().String("menu", "", "parent boot menu")
	profileCopyCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	profileCopyCmd.Flags().String("virt-bridge", "", "virt bridge")
	profileCopyCmd.Flags().String("virt-cpus", "", "virt CPUs")
	profileCopyCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	profileCopyCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	profileCopyCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	profileCopyCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	profileCopyCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")

	// local flags for profile dumpvars
	profileDumpVarsCmd.Flags().String("name", "", "the profile name")

	// local flags for profile edit
	profileEditCmd.Flags().String("name", "", "the profile name")
	profileEditCmd.Flags().String("repos", "", "(repos to auto-assign to this profile")
	profileEditCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	profileEditCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	profileEditCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	profileEditCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	profileEditCmd.Flags().String("distro", "", "the name of a previously defined Cobbler distribution. This value is required")
	profileEditCmd.Flags().String("comment", "", "free form text description")
	profileEditCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	profileEditCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	profileEditCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	profileEditCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	profileEditCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	profileEditCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	profileEditCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	profileEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	profileEditCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	profileEditCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	profileEditCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this profile in the PXE menu?)")
	profileEditCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	profileEditCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	profileEditCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	profileEditCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	profileEditCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	profileEditCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	profileEditCmd.Flags().String("parent", "", "parent profile")
	profileEditCmd.Flags().String("proxy", "", "proxy server URL")
	profileEditCmd.Flags().String("server", "", "server override")
	profileEditCmd.Flags().String("menu", "", "parent boot menu")
	profileEditCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	profileEditCmd.Flags().String("virt-bridge", "", "virt bridge")
	profileEditCmd.Flags().String("virt-cpus", "", "virt CPUs")
	profileEditCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	profileEditCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	profileEditCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	profileEditCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	profileEditCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")

	// local flags for profile find
	profileFindCmd.Flags().String("name", "", "the profile name")
	profileFindCmd.Flags().String("ctime", "", "")
	profileFindCmd.Flags().String("depth", "", "")
	profileFindCmd.Flags().String("mtime", "", "")
	profileFindCmd.Flags().String("repos", "", "(repos to auto-assign to this profile")
	profileFindCmd.Flags().String("uid", "", "UID")
	profileFindCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	profileFindCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	profileFindCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	profileFindCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	profileFindCmd.Flags().String("distro", "", "the name of a previously defined Cobbler distribution. This value is required")
	profileFindCmd.Flags().String("comment", "", "free form text description")
	profileFindCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	profileFindCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	profileFindCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	profileFindCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	profileFindCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	profileFindCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	profileFindCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	profileFindCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	profileFindCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	profileFindCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this profile in the PXE menu?)")
	profileFindCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	profileFindCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	profileFindCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	profileFindCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	profileFindCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	profileFindCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	profileFindCmd.Flags().String("parent", "", "parent profile")
	profileFindCmd.Flags().String("proxy", "", "proxy server URL")
	profileFindCmd.Flags().String("server", "", "server override")
	profileFindCmd.Flags().String("menu", "", "parent boot menu")
	profileFindCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	profileFindCmd.Flags().String("virt-bridge", "", "virt bridge")
	profileFindCmd.Flags().String("virt-cpus", "", "virt CPUs")
	profileFindCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	profileFindCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	profileFindCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	profileFindCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	profileFindCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")

	// local flags for profile get-autoinstall
	profileGetAutoinstallCmd.Flags().String("name", "", "the profile name")

	// local flags for profile remove
	profileRemoveCmd.Flags().String("name", "", "the profile name")
	profileRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for profile rename
	profileRenameCmd.Flags().String("name", "", "the profile name")
	profileRenameCmd.Flags().String("newname", "", "the new profile name")
	profileRenameCmd.Flags().String("repos", "", "(repos to auto-assign to this profile")
	profileRenameCmd.Flags().String("autoinstall", "", "path to automatic installation template")
	profileRenameCmd.Flags().String("autoinstall-meta", "", "automatic installation metadata")
	profileRenameCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	profileRenameCmd.Flags().String("boot-loaders", "", "boot loader space delimited list (network installation boot loaders). Valid options for list items are: <<inherit>>,grub,pxe,ipxe")
	profileRenameCmd.Flags().String("distro", "", "the name of a previously defined Cobbler distribution. This value is required")
	profileRenameCmd.Flags().String("comment", "", "free form text description")
	profileRenameCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	profileRenameCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	profileRenameCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	profileRenameCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	profileRenameCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	profileRenameCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	profileRenameCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	profileRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	profileRenameCmd.Flags().String("dhcp-tag", "", "DHCP tag (see manpage or leave blank)")
	profileRenameCmd.Flags().Bool("enable-ipxe", false, "enable iPXE? (use iPXE instead of PXELINUX for advanced booting options)")
	profileRenameCmd.Flags().Bool("enable-menu", false, "enable PXE Menu? (show this profile in the PXE menu?)")
	profileRenameCmd.Flags().String("mgmt-parameters", "", "Parameters which will be handed to your management application (must be a valid YAML dictionary))")
	profileRenameCmd.Flags().String("name-servers", "", "name servers (space delimited)")
	profileRenameCmd.Flags().String("name-servers-search", "", "name servers search path (space delimited)")
	profileRenameCmd.Flags().String("next-server-v4", "", "next server (IPv4) override (see manpage or leave blank)")
	profileRenameCmd.Flags().String("next-server-v6", "", "next server (IPv6) override (see manpage or leave blank)")
	profileRenameCmd.Flags().String("filename", "", "DHCP filename override (used to boot non-default bootloaders)")
	profileRenameCmd.Flags().String("parent", "", "parent profile")
	profileRenameCmd.Flags().String("proxy", "", "proxy server URL")
	profileRenameCmd.Flags().String("server", "", "server override")
	profileRenameCmd.Flags().String("menu", "", "parent boot menu")
	profileRenameCmd.Flags().Bool("virt-auto-boot", false, "auto boot this VM?")
	profileRenameCmd.Flags().String("virt-bridge", "", "virt bridge")
	profileRenameCmd.Flags().String("virt-cpus", "", "virt CPUs")
	profileRenameCmd.Flags().String("virt-disk-driver", "", "the on-disk format for the virtualization disk. Valid options: <<inherit>>,raw,qcow2,qed,vdi,vdmk")
	profileRenameCmd.Flags().String("virt-file-size", "", "virt file size in GB")
	profileRenameCmd.Flags().String("virt-path", "", "virt Path (e.g. /directory or VolGroup00)")
	profileRenameCmd.Flags().String("virt-ram", "", "virt RAM size in MB")
	profileRenameCmd.Flags().String("virt-type", "", "virtualization technology to use. Valid options: <<inherit>>,qemu,kvm,xenpv,xenfv,vmware,vmwarew,openvz,auto)")

	// local flags for profile report
	profileReportCmd.Flags().String("name", "", "the profile name")
}
