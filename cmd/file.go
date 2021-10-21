// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "file management",
	Long: `Let you manage files.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-file for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var fileAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add file",
	Long:  `Adds a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var fileCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy file",
	Long:  `Copies a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var fileEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit file",
	Long:  `Edits a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var fileFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find file",
	Long:  `Finds a given file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var fileListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all files",
	Long:  `Lists all available files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var fileRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove file",
	Long:  `Removes a given file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var fileRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename file",
	Long:  `Renames a given file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var fileReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all files in detail",
	Long:  `Shows detailed information about all files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.AddCommand(fileAddCmd)
	fileCmd.AddCommand(fileCopyCmd)
	fileCmd.AddCommand(fileEditCmd)
	fileCmd.AddCommand(fileFindCmd)
	fileCmd.AddCommand(fileListCmd)
	fileCmd.AddCommand(fileRemoveCmd)
	fileCmd.AddCommand(fileRenameCmd)
	fileCmd.AddCommand(fileReportCmd)

	// local flags for file add
	fileAddCmd.Flags().String("name", "", "the file name")
	fileAddCmd.Flags().String("ctime", "", "")
	fileAddCmd.Flags().String("depth", "", "")
	fileAddCmd.Flags().String("mtime", "", "")
	fileAddCmd.Flags().String("uid", "", "")
	fileAddCmd.Flags().String("comment", "", "free form text description")
	fileAddCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	fileAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	fileAddCmd.Flags().String("action", "", "create or remove file resource")
	fileAddCmd.Flags().String("mode", "", "file modes")
	fileAddCmd.Flags().String("template", "", "the template for the file")
	fileAddCmd.Flags().String("path", "", "the path of the file")
	fileAddCmd.Flags().String("group", "", "file owner group in file system")
	fileAddCmd.Flags().String("owner", "", "file owner user in file system")
	fileAddCmd.Flags().Bool("is-dir", false, "treat file resource as a directory")

	// local flags for file copy
	fileCopyCmd.Flags().String("name", "", "the file name")
	fileCopyCmd.Flags().String("newname", "", "the new file name")
	fileCopyCmd.Flags().String("ctime", "", "")
	fileCopyCmd.Flags().String("depth", "", "")
	fileCopyCmd.Flags().String("mtime", "", "")
	fileCopyCmd.Flags().String("uid", "", "")
	fileCopyCmd.Flags().String("comment", "", "free form text description")
	fileCopyCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	fileCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	fileCopyCmd.Flags().String("action", "", "create or remove file resource")
	fileCopyCmd.Flags().String("mode", "", "file modes")
	fileCopyCmd.Flags().String("template", "", "the template for the file")
	fileCopyCmd.Flags().String("path", "", "the path of the file")
	fileCopyCmd.Flags().String("group", "", "file owner group in file system")
	fileCopyCmd.Flags().String("owner", "", "file owner user in file system")
	fileCopyCmd.Flags().Bool("is-dir", false, "treat file resource as a directory")

	// local flags for file edit
	fileEditCmd.Flags().String("name", "", "the file name")
	fileEditCmd.Flags().String("ctime", "", "")
	fileEditCmd.Flags().String("depth", "", "")
	fileEditCmd.Flags().String("mtime", "", "")
	fileEditCmd.Flags().String("uid", "", "")
	fileEditCmd.Flags().String("comment", "", "free form text description")
	fileEditCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	fileEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	fileEditCmd.Flags().String("action", "", "create or remove file resource")
	fileEditCmd.Flags().String("mode", "", "file modes")
	fileEditCmd.Flags().String("template", "", "the template for the file")
	fileEditCmd.Flags().String("path", "", "the path of the file")
	fileEditCmd.Flags().String("group", "", "file owner group in file system")
	fileEditCmd.Flags().String("owner", "", "file owner user in file system")
	fileEditCmd.Flags().Bool("is-dir", false, "treat file resource as a directory")

	// local flags for file find
	fileFindCmd.Flags().String("name", "", "the file name")
	fileFindCmd.Flags().String("ctime", "", "")
	fileFindCmd.Flags().String("depth", "", "")
	fileFindCmd.Flags().String("mtime", "", "")
	fileFindCmd.Flags().String("uid", "", "")
	fileFindCmd.Flags().String("comment", "", "free form text description")
	fileFindCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	fileFindCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	fileFindCmd.Flags().String("action", "", "create or remove file resource")
	fileFindCmd.Flags().String("mode", "", "file modes")
	fileFindCmd.Flags().String("template", "", "the template for the file")
	fileFindCmd.Flags().String("path", "", "the path of the file")
	fileFindCmd.Flags().String("group", "", "file owner group in file system")
	fileFindCmd.Flags().String("owner", "", "file owner user in file system")
	fileFindCmd.Flags().Bool("is-dir", false, "treat file resource as a directory")

	// local flags for file remove
	fileRemoveCmd.Flags().String("name", "", "the file name")
	fileRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for file rename
	fileRenameCmd.Flags().String("name", "", "the file name")
	fileRenameCmd.Flags().String("newname", "", "the new file name")
	fileRenameCmd.Flags().String("ctime", "", "")
	fileRenameCmd.Flags().String("depth", "", "")
	fileRenameCmd.Flags().String("mtime", "", "")
	fileRenameCmd.Flags().String("uid", "", "")
	fileRenameCmd.Flags().String("comment", "", "free form text description")
	fileRenameCmd.Flags().String("owners", "", "owners list for authz_ownership (space delimited))")
	fileRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	fileRenameCmd.Flags().String("action", "", "create or remove file resource")
	fileRenameCmd.Flags().String("mode", "", "file modes")
	fileRenameCmd.Flags().String("template", "", "the template for the file")
	fileRenameCmd.Flags().String("path", "", "the path of the file")
	fileRenameCmd.Flags().String("group", "", "file owner group in file system")
	fileRenameCmd.Flags().String("owner", "", "file owner user in file system")
	fileRenameCmd.Flags().Bool("is-dir", false, "treat file resource as a directory")

	// local flags for file report
	fileReportCmd.Flags().String("name", "", "the file name")
}
