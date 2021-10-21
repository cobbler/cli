// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// mgmtclassCmd represents the mgmtclass command
var mgmtclassCmd = &cobra.Command{
	Use:   "mgmtclass",
	Short: "mgmtclass management",
	Long: `Let you manage mgmtclasses.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-mgmtclass for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var mgmtclassAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add mgmtclass",
	Long:  `Adds a mgmtclass.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var mgmtclassCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy mgmtclass",
	Long:  `Copies a mgmtclass.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var mgmtclassEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit mgmtclass",
	Long:  `Edits a mgmtclass.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var mgmtclassFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find mgmtclass",
	Long:  `Finds a given mgmtclass.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var mgmtclassListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all mgmtclasses",
	Long:  `Lists all available mgmtclasses.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var mgmtclassRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove mgmtclass",
	Long:  `Removes a given mgmtclass.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var mgmtclassRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename mgmtclass",
	Long:  `Renames a given mgmtclass.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var mgmtclassReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all mgmtclasses in detail",
	Long:  `Shows detailed information about all mgmtclasses.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(mgmtclassCmd)
	mgmtclassCmd.AddCommand(mgmtclassAddCmd)
	mgmtclassCmd.AddCommand(mgmtclassCopyCmd)
	mgmtclassCmd.AddCommand(mgmtclassEditCmd)
	mgmtclassCmd.AddCommand(mgmtclassFindCmd)
	mgmtclassCmd.AddCommand(mgmtclassListCmd)
	mgmtclassCmd.AddCommand(mgmtclassRemoveCmd)
	mgmtclassCmd.AddCommand(mgmtclassRenameCmd)
	mgmtclassCmd.AddCommand(mgmtclassReportCmd)

	// local flags for mgmtclass add
	mgmtclassAddCmd.Flags().String("name", "", "the mgmtclass name")
	mgmtclassAddCmd.Flags().String("ctime", "", "")
	mgmtclassAddCmd.Flags().String("depth", "", "")
	mgmtclassAddCmd.Flags().String("mtime", "", "")
	mgmtclassAddCmd.Flags().String("comment", "", "free form text description")
	mgmtclassAddCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	mgmtclassAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	mgmtclassAddCmd.Flags().String("files", "", "file resources")
	mgmtclassAddCmd.Flags().String("packages", "", "package resources")
	mgmtclassAddCmd.Flags().String("params", "", "list of parameters/variables")
	mgmtclassAddCmd.Flags().String("class-name", "", "actual class name (leave blank to use the	name field)")
	mgmtclassAddCmd.Flags().String("is-definition", "", "is Definition? Treat this class as a definition (puppet only)")
	mgmtclassAddCmd.Flags().String("uid", "", "")

	// local flags for mgmtclass copy
	mgmtclassCopyCmd.Flags().String("name", "", "the mgmtclass name")
	mgmtclassCopyCmd.Flags().String("newname", "", "the new mgmtclass name")
	mgmtclassCopyCmd.Flags().String("ctime", "", "")
	mgmtclassCopyCmd.Flags().String("depth", "", "")
	mgmtclassCopyCmd.Flags().String("mtime", "", "")
	mgmtclassCopyCmd.Flags().String("comment", "", "free form text description")
	mgmtclassCopyCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	mgmtclassCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	mgmtclassCopyCmd.Flags().String("files", "", "file resources")
	mgmtclassCopyCmd.Flags().String("packages", "", "package resources")
	mgmtclassCopyCmd.Flags().String("params", "", "list of parameters/variables")
	mgmtclassCopyCmd.Flags().String("class-name", "", "actual class name (leave blank to use the	name field)")
	mgmtclassCopyCmd.Flags().String("is-definition", "", "is Definition? Treat this class as a definition (puppet only)")
	mgmtclassCopyCmd.Flags().String("uid", "", "")

	// local flags for mgmtclass edit
	mgmtclassEditCmd.Flags().String("name", "", "the mgmtclass name")
	mgmtclassEditCmd.Flags().String("ctime", "", "")
	mgmtclassEditCmd.Flags().String("depth", "", "")
	mgmtclassEditCmd.Flags().String("mtime", "", "")
	mgmtclassEditCmd.Flags().String("comment", "", "free form text description")
	mgmtclassEditCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	mgmtclassEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	mgmtclassEditCmd.Flags().String("files", "", "file resources")
	mgmtclassEditCmd.Flags().String("packages", "", "package resources")
	mgmtclassEditCmd.Flags().String("params", "", "list of parameters/variables")
	mgmtclassEditCmd.Flags().String("class-name", "", "actual class name (leave blank to use the	name field)")
	mgmtclassEditCmd.Flags().String("is-definition", "", "is Definition? Treat this class as a definition (puppet only)")
	mgmtclassEditCmd.Flags().String("uid", "", "")

	// local flags for mgmtclass find
	mgmtclassFindCmd.Flags().String("name", "", "the mgmtclass name")
	mgmtclassFindCmd.Flags().String("ctime", "", "")
	mgmtclassFindCmd.Flags().String("depth", "", "")
	mgmtclassFindCmd.Flags().String("mtime", "", "")
	mgmtclassFindCmd.Flags().String("comment", "", "free form text description")
	mgmtclassFindCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	mgmtclassFindCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	mgmtclassFindCmd.Flags().String("files", "", "file resources")
	mgmtclassFindCmd.Flags().String("packages", "", "package resources")
	mgmtclassFindCmd.Flags().String("params", "", "list of parameters/variables")
	mgmtclassFindCmd.Flags().String("class-name", "", "actual class name (leave blank to use the	name field)")
	mgmtclassFindCmd.Flags().String("is-definition", "", "is Definition? Treat this class as a definition (puppet only)")
	mgmtclassFindCmd.Flags().String("uid", "", "")

	// local flags for mgmtclass remove
	mgmtclassRemoveCmd.Flags().String("name", "", "the mgmtclass name")
	mgmtclassRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for mgmtclass rename
	mgmtclassRenameCmd.Flags().String("name", "", "the mgmtclass name")
	mgmtclassRenameCmd.Flags().String("newname", "", "the new mgmtclass name")
	mgmtclassRenameCmd.Flags().String("ctime", "", "")
	mgmtclassRenameCmd.Flags().String("depth", "", "")
	mgmtclassRenameCmd.Flags().String("mtime", "", "")
	mgmtclassRenameCmd.Flags().String("comment", "", "free form text description")
	mgmtclassRenameCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	mgmtclassRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	mgmtclassRenameCmd.Flags().String("files", "", "file resources")
	mgmtclassRenameCmd.Flags().String("packages", "", "package resources")
	mgmtclassRenameCmd.Flags().String("params", "", "list of parameters/variables")
	mgmtclassRenameCmd.Flags().String("class-name", "", "actual class name (leave blank to use the	name field)")
	mgmtclassRenameCmd.Flags().String("is-definition", "", "is Definition? Treat this class as a definition (puppet only)")
	mgmtclassRenameCmd.Flags().String("uid", "", "")

	// local flags for mgmtclass report
	mgmtclassReportCmd.Flags().String("name", "", "the mgmtclass name")
}
