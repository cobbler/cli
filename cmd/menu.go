// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Menu management",
	Long: `Let you manage menus.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-menu for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var menuAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add menu",
	Long:  `Adds a menu.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var menuCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy menu",
	Long:  `Copies a menu.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var menuEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit menu",
	Long:  `Edits a menu.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var menuFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find menu",
	Long:  `Finds a given menu.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var menuListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all menus",
	Long:  `Lists all available menus.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var menuRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove menu",
	Long:  `Removes a given menu.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var menuRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename menu",
	Long:  `Renames a given menu.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

var menuReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all menus in detail",
	Long:  `Shows detailed information about all menus.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		notImplemented()
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)
	menuCmd.AddCommand(menuAddCmd)
	menuCmd.AddCommand(menuCopyCmd)
	menuCmd.AddCommand(menuEditCmd)
	menuCmd.AddCommand(menuFindCmd)
	menuCmd.AddCommand(menuListCmd)
	menuCmd.AddCommand(menuRemoveCmd)
	menuCmd.AddCommand(menuRenameCmd)
	menuCmd.AddCommand(menuReportCmd)

	// local flags for menu add
	menuAddCmd.Flags().String("name", "", "the menu name")
	menuAddCmd.Flags().String("ctime", "", "")
	menuAddCmd.Flags().String("depth", "", "")
	menuAddCmd.Flags().String("mtime", "", "")
	menuAddCmd.Flags().String("uid", "", "")
	menuAddCmd.Flags().String("comment", "", "free form text description")
	menuAddCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	menuAddCmd.Flags().String("parent", "", "parent menu")
	menuAddCmd.Flags().String("display-name", "", "display name")

	// local flags for menu copy
	menuCopyCmd.Flags().String("name", "", "the menu name")
	menuCopyCmd.Flags().String("ctime", "", "")
	menuCopyCmd.Flags().String("depth", "", "")
	menuCopyCmd.Flags().String("mtime", "", "")
	menuCopyCmd.Flags().String("uid", "", "")
	menuCopyCmd.Flags().String("comment", "", "free form text description")
	menuCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	menuCopyCmd.Flags().String("parent", "", "parent menu")
	menuCopyCmd.Flags().String("display-name", "", "display name")

	// local flags for menu edit
	menuEditCmd.Flags().String("name", "", "the menu name")
	menuEditCmd.Flags().String("ctime", "", "")
	menuEditCmd.Flags().String("depth", "", "")
	menuEditCmd.Flags().String("mtime", "", "")
	menuEditCmd.Flags().String("uid", "", "")
	menuEditCmd.Flags().String("comment", "", "free form text description")
	menuEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	menuEditCmd.Flags().String("parent", "", "parent menu")
	menuEditCmd.Flags().String("display-name", "", "display name")

	// local flags for menu find
	menuFindCmd.Flags().String("name", "", "the menu name")
	menuFindCmd.Flags().String("ctime", "", "")
	menuFindCmd.Flags().String("depth", "", "")
	menuFindCmd.Flags().String("mtime", "", "")
	menuFindCmd.Flags().String("uid", "", "")
	menuFindCmd.Flags().String("comment", "", "free form text description")
	menuFindCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	menuFindCmd.Flags().String("parent", "", "parent menu")
	menuFindCmd.Flags().String("display-name", "", "display name")

	// local flags for menu remove
	menuRemoveCmd.Flags().String("name", "", "the menu name")
	menuRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for menu rename
	menuRenameCmd.Flags().String("name", "", "the menu name")
	menuRenameCmd.Flags().String("ctime", "", "")
	menuRenameCmd.Flags().String("depth", "", "")
	menuRenameCmd.Flags().String("mtime", "", "")
	menuRenameCmd.Flags().String("uid", "", "")
	menuRenameCmd.Flags().String("comment", "", "free form text description")
	menuRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	menuRenameCmd.Flags().String("parent", "", "parent menu")
	menuRenameCmd.Flags().String("display-name", "", "display name")

	// local flags for menu report
	menuReportCmd.Flags().String("name", "", "the menu name")
}
