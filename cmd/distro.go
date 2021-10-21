// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// distroCmd represents the distro command
var distroCmd = &cobra.Command{
	Use:   "distro",
	Short: "Distribution management",
	Long: `Let you manage distributions.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-distro for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var distroAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add distribution",
	Long:  `Adds a distribution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var distroCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy distribution",
	Long:  `Copies a distribution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var distroEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit distribution",
	Long:  `Edits a distribution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var distroFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find distribution",
	Long:  `Finds a given distribution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var distroListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all distributions",
	Long:  `Lists all available distributions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var distroRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove distribution",
	Long:  `Removes a given distribution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var distroRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename distribution",
	Long:  `Renames a given distribution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var distroReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all distributions in detail",
	Long:  `Shows detailed information about all distributions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(distroCmd)
	distroCmd.AddCommand(distroAddCmd)
	distroCmd.AddCommand(distroCopyCmd)
	distroCmd.AddCommand(distroEditCmd)
	distroCmd.AddCommand(distroFindCmd)
	distroCmd.AddCommand(distroListCmd)
	distroCmd.AddCommand(distroRemoveCmd)
	distroCmd.AddCommand(distroRenameCmd)
	distroCmd.AddCommand(distroReportCmd)

	// local flags for distro add
	distroAddCmd.Flags().String("name", "", "the distro name")
	distroAddCmd.Flags().String("ctime", "", "")
	distroAddCmd.Flags().String("depth", "", "")
	distroAddCmd.Flags().String("mtime", "", "")
	distroAddCmd.Flags().String("source-repos", "", "source repositories")
	distroAddCmd.Flags().String("tree-build-time", "", "tree build time")
	distroAddCmd.Flags().String("uid", "", "UID")
	distroAddCmd.Flags().String("arch", "", "Architecture")
	distroAddCmd.Flags().String("autoinstall-meta", "", "automatic installation template metadata")
	distroAddCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	distroAddCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	distroAddCmd.Flags().String("breed", "", "Breed (what is the type of the distribution?)")
	distroAddCmd.Flags().String("comment", "", "free form text description")
	distroAddCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	distroAddCmd.Flags().String("initrd", "", "initrd (absolute path on filesystem)")
	distroAddCmd.Flags().String("remote-boot-initrd", "", "remote boot initrd (URL the bootloader directly retrieves and boots from)")
	distroAddCmd.Flags().String("kernel", "", "Kernel (absolute path on filesystem)")
	distroAddCmd.Flags().String("remote-boot-kernel", "", "remote boot kernel (URL the bootloader directly retrieves and boots from)")
	distroAddCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	distroAddCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	distroAddCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	distroAddCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	distroAddCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	distroAddCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	distroAddCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	distroAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for distro copy
	distroCopyCmd.Flags().String("name", "", "the distro name")
	distroCopyCmd.Flags().String("newname", "", "the new distro name")
	distroCopyCmd.Flags().String("ctime", "", "")
	distroCopyCmd.Flags().String("depth", "", "")
	distroCopyCmd.Flags().String("mtime", "", "")
	distroCopyCmd.Flags().String("source-repos", "", "source repositories")
	distroCopyCmd.Flags().String("tree-build-time", "", "tree build time")
	distroCopyCmd.Flags().String("uid", "", "UID")
	distroCopyCmd.Flags().String("arch", "", "Architecture")
	distroCopyCmd.Flags().String("autoinstall-meta", "", "automatic installation template metadata")
	distroCopyCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	distroCopyCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	distroCopyCmd.Flags().String("breed", "", "Breed (what is the type of the distribution?)")
	distroCopyCmd.Flags().String("comment", "", "free form text description")
	distroCopyCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	distroCopyCmd.Flags().String("initrd", "", "initrd (absolute path on filesystem)")
	distroCopyCmd.Flags().String("remote-boot-initrd", "", "remote boot initrd (URL the bootloader directly retrieves and boots from)")
	distroCopyCmd.Flags().String("kernel", "", "Kernel (absolute path on filesystem)")
	distroCopyCmd.Flags().String("remote-boot-kernel", "", "remote boot kernel (URL the bootloader directly retrieves and boots from)")
	distroCopyCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	distroCopyCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	distroCopyCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	distroCopyCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	distroCopyCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	distroCopyCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	distroCopyCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	distroCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for distro edit
	distroEditCmd.Flags().String("name", "", "the distro name")
	distroEditCmd.Flags().String("ctime", "", "")
	distroEditCmd.Flags().String("depth", "", "")
	distroEditCmd.Flags().String("mtime", "", "")
	distroEditCmd.Flags().String("source-repos", "", "source repositories")
	distroEditCmd.Flags().String("tree-build-time", "", "tree build time")
	distroEditCmd.Flags().String("uid", "", "UID")
	distroEditCmd.Flags().String("arch", "", "Architecture")
	distroEditCmd.Flags().String("autoinstall-meta", "", "automatic installation template metadata")
	distroEditCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	distroEditCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	distroEditCmd.Flags().String("breed", "", "Breed (what is the type of the distribution?)")
	distroEditCmd.Flags().String("comment", "", "free form text description")
	distroEditCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	distroEditCmd.Flags().String("initrd", "", "initrd (absolute path on filesystem)")
	distroEditCmd.Flags().String("remote-boot-initrd", "", "remote boot initrd (URL the bootloader directly retrieves and boots from)")
	distroEditCmd.Flags().String("kernel", "", "Kernel (absolute path on filesystem)")
	distroEditCmd.Flags().String("remote-boot-kernel", "", "remote boot kernel (URL the bootloader directly retrieves and boots from)")
	distroEditCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	distroEditCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	distroEditCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	distroEditCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	distroEditCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	distroEditCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	distroEditCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	distroEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for distro find
	distroFindCmd.Flags().String("name", "", "the distro name")
	distroFindCmd.Flags().String("ctime", "", "")
	distroFindCmd.Flags().String("depth", "", "")
	distroFindCmd.Flags().String("mtime", "", "")
	distroFindCmd.Flags().String("source-repos", "", "source repositories")
	distroFindCmd.Flags().String("tree-build-time", "", "tree build time")
	distroFindCmd.Flags().String("uid", "", "UID")
	distroFindCmd.Flags().String("arch", "", "Architecture")
	distroFindCmd.Flags().String("autoinstall-meta", "", "automatic installation template metadata")
	distroFindCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	distroFindCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	distroFindCmd.Flags().String("breed", "", "Breed (what is the type of the distribution?)")
	distroFindCmd.Flags().String("comment", "", "free form text description")
	distroFindCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	distroFindCmd.Flags().String("initrd", "", "initrd (absolute path on filesystem)")
	distroFindCmd.Flags().String("remote-boot-initrd", "", "remote boot initrd (URL the bootloader directly retrieves and boots from)")
	distroFindCmd.Flags().String("kernel", "", "Kernel (absolute path on filesystem)")
	distroFindCmd.Flags().String("remote-boot-kernel", "", "remote boot kernel (URL the bootloader directly retrieves and boots from)")
	distroFindCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	distroFindCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	distroFindCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	distroFindCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	distroFindCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	distroFindCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	distroFindCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")

	// local flags for distro remove
	distroRemoveCmd.Flags().String("name", "", "the distro name")
	distroRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for distro rename
	distroRenameCmd.Flags().String("name", "", "the distro name")
	distroRenameCmd.Flags().String("newname", "", "the new distro name")
	distroRenameCmd.Flags().String("ctime", "", "")
	distroRenameCmd.Flags().String("depth", "", "")
	distroRenameCmd.Flags().String("mtime", "", "")
	distroRenameCmd.Flags().String("source-repos", "", "source repositories")
	distroRenameCmd.Flags().String("tree-build-time", "", "tree build time")
	distroRenameCmd.Flags().String("uid", "", "UID")
	distroRenameCmd.Flags().String("arch", "", "Architecture")
	distroRenameCmd.Flags().String("autoinstall-meta", "", "automatic installation template metadata")
	distroRenameCmd.Flags().String("boot-files", "", "TFTP boot files (files copied into tftpboot beyond the kernel/initrd)")
	distroRenameCmd.Flags().String("boot-loaders", "", "boot loaders (network installation boot loaders)")
	distroRenameCmd.Flags().String("breed", "", "Breed (what is the type of the distribution?)")
	distroRenameCmd.Flags().String("comment", "", "free form text description")
	distroRenameCmd.Flags().String("fetchable-files", "", "fetchable files (templates for tftp, wget or curl)")
	distroRenameCmd.Flags().String("initrd", "", "initrd (absolute path on filesystem)")
	distroRenameCmd.Flags().String("remote-boot-initrd", "", "remote boot initrd (URL the bootloader directly retrieves and boots from)")
	distroRenameCmd.Flags().String("kernel", "", "Kernel (absolute path on filesystem)")
	distroRenameCmd.Flags().String("remote-boot-kernel", "", "remote boot kernel (URL the bootloader directly retrieves and boots from)")
	distroRenameCmd.Flags().String("kernel-options", "", "kernel options (e.g. selinux=permissive)")
	distroRenameCmd.Flags().String("kernel-options-post", "", "post install kernel options (e.g. clocksource=pit noapic)")
	distroRenameCmd.Flags().String("mgmt-classes", "", "management classes (for external config management)")
	distroRenameCmd.Flags().String("os-version", "", "OS version (needed for some virtualization optimizations)")
	distroRenameCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	distroRenameCmd.Flags().String("redhat-management-key", "", "RedHat management key (registration key for RHN, Spacewalk, or Satellite)")
	distroRenameCmd.Flags().String("template-files", "", "template files (file mappings for built-in config management)")
	distroRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for distro report
	distroReportCmd.Flags().String("name", "", "the distro name")
}
