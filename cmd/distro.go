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

var distro *cobbler.Distro
var distros []*cobbler.Distro

// distroCmd represents the distro command
var distroCmd = &cobra.Command{
	Use:   "distro",
	Short: "Distribution management",
	Long: `Let you manage distributions.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-distro for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		cmd.Help()
	},
}

var distroAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add distribution",
	Long:  `Adds a distribution.`,
	Run: func(cmd *cobra.Command, args []string) {

		var newDistro cobbler.Distro
		// internal fields (ctime, mtime, depth, uid, source-repos, tree-build-time) cannot be modified
		newDistro.Arch, _ = cmd.Flags().GetString("arch")
		newDistro.BootFiles, _ = cmd.Flags().GetString("boot-files")
		newDistro.BootLoader, _ = cmd.Flags().GetString("boot-loaders")
		newDistro.Breed, _ = cmd.Flags().GetString("breed")
		newDistro.Comment, _ = cmd.Flags().GetString("comment")
		newDistro.FetchableFiles, _ = cmd.Flags().GetString("fetchable-files")
		newDistro.Initrd, _ = cmd.Flags().GetString("initrd")
		newDistro.Kernel, _ = cmd.Flags().GetString("kernel")
		newDistro.KernelOptions, _ = cmd.Flags().GetString("kernel-options")
		newDistro.KernelOptionsPost, _ = cmd.Flags().GetString("kernel-options-post")
		newDistro.MGMTClasses, _ = cmd.Flags().GetStringArray("mgmt-classes")
		newDistro.Name, _ = cmd.Flags().GetString("name")
		newDistro.OSVersion, _ = cmd.Flags().GetString("os-version")
		newDistro.Owners, _ = cmd.Flags().GetStringArray("owners")
		newDistro.RedHatManagementKey, _ = cmd.Flags().GetString("redhat-management-key")
		newDistro.TemplateFiles, _ = cmd.Flags().GetString("template-files")

		distro, err = Client.CreateDistro(newDistro)

		if checkError(err) == nil {
			fmt.Printf("Distro %s created", newDistro.Name)
		} else {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

var distroCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy distribution",
	Long:  `Copies a distribution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var distroEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit distribution",
	Long:  `Edits a distribution.`,
	Run: func(cmd *cobra.Command, args []string) {

		// find distro through its name
		dname, _ := cmd.Flags().GetString("name")
		var updateDistro, err = Client.GetDistro(dname)

		if checkError(err) != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// internal fields (ctime, mtime, depth, uid, source-repos, tree-build-time) cannot be modified
		var tmpArgs, _ = cmd.Flags().GetString("arch")

		if tmpArgs != "" {
			updateDistro.Arch, _ = cmd.Flags().GetString("arch")
		}
		tmpArgs, _ = cmd.Flags().GetString("boot-files")
		if tmpArgs != "" {
			updateDistro.BootFiles, _ = cmd.Flags().GetString("boot-files")
		}
		tmpArgs, _ = cmd.Flags().GetString("boot-loaders")
		if tmpArgs != "" {
			updateDistro.BootLoader, _ = cmd.Flags().GetString("boot-loaders")
		}
		tmpArgs, _ = cmd.Flags().GetString("breed")
		if tmpArgs != "" {
			updateDistro.Breed, _ = cmd.Flags().GetString("breed")
		}
		tmpArgs, _ = cmd.Flags().GetString("comment")
		if tmpArgs != "" {
			updateDistro.Comment, _ = cmd.Flags().GetString("comment")
		}
		tmpArgs, _ = cmd.Flags().GetString("fetchable-files")
		if tmpArgs != "" {
			updateDistro.FetchableFiles, _ = cmd.Flags().GetString("fetchable-files")
		}
		tmpArgs, _ = cmd.Flags().GetString("initrd")
		if tmpArgs != "" {
			updateDistro.Initrd, _ = cmd.Flags().GetString("initrd")
		}
		tmpArgs, _ = cmd.Flags().GetString("kernel")
		if tmpArgs != "" {
			updateDistro.Kernel, _ = cmd.Flags().GetString("kernel")
		}
		tmpArgs, _ = cmd.Flags().GetString("kernel-options")
		if tmpArgs != "" {
			updateDistro.KernelOptions, _ = cmd.Flags().GetString("kernel-options")
		}
		tmpArgs, _ = cmd.Flags().GetString("kernel-options-post")
		if tmpArgs != "" {
			updateDistro.KernelOptionsPost, _ = cmd.Flags().GetString("kernel-options-post")
		}
		tmpArgs, _ = cmd.Flags().GetString("mgmt-classes")
		if tmpArgs != "" {
			updateDistro.MGMTClasses, _ = cmd.Flags().GetStringArray("mgmt-classes")
		}
		tmpArgs, _ = cmd.Flags().GetString("name")
		if tmpArgs != "" {
			updateDistro.Name, _ = cmd.Flags().GetString("name")
		}
		tmpArgs, _ = cmd.Flags().GetString("os-version")
		if tmpArgs != "" {
			updateDistro.OSVersion, _ = cmd.Flags().GetString("os-version")
		}
		tmpArgs, _ = cmd.Flags().GetString("owners")
		if tmpArgs != "" {
			updateDistro.Owners, _ = cmd.Flags().GetStringArray("owners")
		}
		tmpArgs, _ = cmd.Flags().GetString("redhat-management-key")
		if tmpArgs != "" {
			updateDistro.RedHatManagementKey, _ = cmd.Flags().GetString("redhat-management-key")
		}
		tmpArgs, _ = cmd.Flags().GetString("template-files")
		if tmpArgs != "" {
			updateDistro.TemplateFiles, _ = cmd.Flags().GetString("template-files")
		}

		err = Client.UpdateDistro(updateDistro)

		if checkError(err) != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

var distroFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find distribution",
	Long:  `Finds a given distribution.`,
	Run: func(cmd *cobra.Command, args []string) {

		/*
			dname, _ := cmd.Flags().GetString("name")
			distro, err = Client.GetDistro(dname)

			if checkError(err) == nil {
			   	str, _ := json.MarshalIndent(distro, "", " ")
			   	fmt.Println(string(str))
			} else {
			   	fmt.Println(err.Error())
			}
		*/
		notImplemented()
	},
}

var distroListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all distributions",
	Long:  `Lists all available distributions.`,
	Run: func(cmd *cobra.Command, args []string) {

		distros, err = Client.GetDistros()

		if checkError(err) == nil {
			fmt.Println(distros)
		} else {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

var distroRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove distribution",
	Long:  `Removes a given distribution.`,
	Run: func(cmd *cobra.Command, args []string) {

		dname, _ := cmd.Flags().GetString("name")
		err := Client.DeleteDistro(dname)
		if checkError(err) != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

var distroRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename distribution",
	Long:  `Renames a given distribution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var distroReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all distributions in detail",
	Long:  `Shows detailed information about all distributions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
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
