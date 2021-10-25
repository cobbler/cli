// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Package management",
	Long: `Let you manage packages.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-package for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var packageAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add package",
	Long:  `Adds a package.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var packageCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy package",
	Long:  `Copies a package.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var packageEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit package",
	Long:  `Edits a package.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var packageFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find package",
	Long:  `Finds a given package.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var packageListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all packages",
	Long:  `Lists all available packages.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var packageRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove package",
	Long:  `Removes a given package.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var packageRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename package",
	Long:  `Renames a given package.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var packageReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all packages in detail",
	Long:  `Shows detailed information about all packages.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

func init() {
	rootCmd.AddCommand(packageCmd)
	packageCmd.AddCommand(packageAddCmd)
	packageCmd.AddCommand(packageCopyCmd)
	packageCmd.AddCommand(packageEditCmd)
	packageCmd.AddCommand(packageFindCmd)
	packageCmd.AddCommand(packageListCmd)
	packageCmd.AddCommand(packageRemoveCmd)
	packageCmd.AddCommand(packageRenameCmd)
	packageCmd.AddCommand(packageReportCmd)

	// local flags for package add
	packageAddCmd.Flags().String("name", "", "the package name")
	packageAddCmd.Flags().String("ctime", "", "")
	packageAddCmd.Flags().String("depth", "", "")
	packageAddCmd.Flags().String("mtime", "", "")
	packageAddCmd.Flags().String("uid", "", "")
	packageAddCmd.Flags().String("comment", "", "free form text description")
	packageAddCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	packageAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	packageAddCmd.Flags().String("action", "", "install or remove package resourc")
	packageAddCmd.Flags().String("installer", "", "package manager")
	packageAddCmd.Flags().String("version", "", "package version")

	// local flags for package copy
	packageCopyCmd.Flags().String("name", "", "the package name")
	packageCopyCmd.Flags().String("newname", "", "the new package name")
	packageCopyCmd.Flags().String("ctime", "", "")
	packageCopyCmd.Flags().String("depth", "", "")
	packageCopyCmd.Flags().String("mtime", "", "")
	packageCopyCmd.Flags().String("uid", "", "")
	packageCopyCmd.Flags().String("comment", "", "free form text description")
	packageCopyCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited)")
	packageCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	packageCopyCmd.Flags().String("action", "", "install or remove package resourc")
	packageCopyCmd.Flags().String("installer", "", "package manager")
	packageCopyCmd.Flags().String("version", "", "package version")

	// local flags for package edit
	packageEditCmd.Flags().String("name", "", "the package name")
	packageEditCmd.Flags().String("ctime", "", "")
	packageEditCmd.Flags().String("depth", "", "")
	packageEditCmd.Flags().String("mtime", "", "")
	packageEditCmd.Flags().String("uid", "", "")
	packageEditCmd.Flags().String("comment", "", "free form text description")
	packageEditCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	packageEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	packageEditCmd.Flags().String("action", "", "install or remove package resourc")
	packageEditCmd.Flags().String("installer", "", "package manager")
	packageEditCmd.Flags().String("version", "", "package version")

	// local flags for package find
	packageFindCmd.Flags().String("name", "", "the package name")
	packageFindCmd.Flags().String("ctime", "", "")
	packageFindCmd.Flags().String("depth", "", "")
	packageFindCmd.Flags().String("mtime", "", "")
	packageFindCmd.Flags().String("uid", "", "")
	packageFindCmd.Flags().String("comment", "", "free form text description")
	packageFindCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	packageFindCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	packageFindCmd.Flags().String("action", "", "install or remove package resourc")
	packageFindCmd.Flags().String("installer", "", "package manager")
	packageFindCmd.Flags().String("version", "", "package version")

	// local flags for package remove
	packageRemoveCmd.Flags().String("name", "", "the package name")
	packageRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for package rename
	packageRenameCmd.Flags().String("name", "", "the package name")
	packageRenameCmd.Flags().String("newname", "", "the new package name")
	packageRenameCmd.Flags().String("ctime", "", "")
	packageRenameCmd.Flags().String("depth", "", "")
	packageRenameCmd.Flags().String("mtime", "", "")
	packageRenameCmd.Flags().String("uid", "", "")
	packageRenameCmd.Flags().String("comment", "", "free form text description")
	packageRenameCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	packageRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	packageRenameCmd.Flags().String("action", "", "install or remove package resourc")
	packageRenameCmd.Flags().String("installer", "", "package manager")
	packageRenameCmd.Flags().String("version", "", "package version")

	// local flags for package report
	packageReportCmd.Flags().String("name", "", "the package name")
}
