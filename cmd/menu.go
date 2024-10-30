// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func updateMenuFromFlags(cmd *cobra.Command, menu *cobbler.Menu) error {
	// This object doesn't have the in-place flag
	var err error
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			// If one of the previous flags has had an error just directly return.
			return
		}
		switch flag.Name {
		case "name":
			return
		case "newname":
			return
		case "comment":
			var menuNewComment string
			menuNewComment, err = cmd.Flags().GetString("comment")
			if err != nil {
				return
			}
			menu.Comment = menuNewComment
		case "parent":
			var menuNewParent string
			menuNewParent, err = cmd.Flags().GetString("parent")
			if err != nil {
				return
			}
			menu.Parent = menuNewParent
		case "display-name":
			var menuNewDisplayName string
			menuNewDisplayName, err = cmd.Flags().GetString("display-name")
			if err != nil {
				return
			}
			menu.DisplayName = menuNewDisplayName
		}
	})
	// Don't blindly return nil because maybe one of the flags had an issue retrieving an argument.
	return err
}

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Menu management",
	Long: `Let you manage menus.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-menu for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var menuAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add menu",
	Long:  `Adds a menu.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		newMenu := cobbler.NewMenu()
		var err error

		// internal fields (ctime, mtime, depth, uid, source-repos, tree-build-time) cannot be modified
		newMenu.Name, err = cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		// Update menu in-memory
		err = updateMenuFromFlags(cmd, &newMenu)
		if err != nil {
			return err
		}
		// Now create the menu via XML-RPC
		menu, err := Client.CreateMenu(newMenu)
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Menu %s created\n", menu.Name)
		return nil
	},
}

var menuCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy menu",
	Long:  `Copies a menu.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		menuName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		menuNewName, err := cmd.Flags().GetString("newname")
		if err != nil {
			return err
		}

		menuHandle, err := Client.GetMenuHandle(menuName)
		if err != nil {
			return err
		}
		err = Client.CopyMenu(menuHandle, menuNewName)
		if err != nil {
			return err
		}
		newMenu, err := Client.GetMenu(menuNewName, false, false)
		if err != nil {
			return err
		}
		err = updateMenuFromFlags(cmd, newMenu)
		if err != nil {
			return err
		}
		return Client.UpdateMenu(newMenu)
	},
}

var menuEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit menu",
	Long:  `Edits a menu.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		menuName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		menuToEdit, err := Client.GetMenu(menuName, false, false)
		if err != nil {
			return err
		}
		err = updateMenuFromFlags(cmd, menuToEdit)
		if err != nil {
			return err
		}
		return Client.UpdateMenu(menuToEdit)
	},
}

var menuFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find menu",
	Long:  `Finds a given menu.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		return FindItemNames(cmd, args, "menu")
	},
}

var menuListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all menus",
	Long:  `Lists all available menus.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		menuNames, err := Client.ListMenuNames()
		if err != nil {
			return err
		}
		listItems(cmd, "menus", menuNames)
		return nil
	},
}

var menuRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove menu",
	Long:  `Removes a given menu.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		return RemoveItemRecursive(cmd, args, "menu")
	},
}

var menuRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename menu",
	Long:  `Renames a given menu.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		menuName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		menuNewName, err := cmd.Flags().GetString("newname")
		if err != nil {
			return err
		}

		menuHandle, err := Client.GetMenuHandle(menuName)
		if err != nil {
			return err
		}
		err = Client.RenameMenu(menuHandle, menuNewName)
		if err != nil {
			return err
		}
		newMenu, err := Client.GetMenu(menuNewName, false, false)
		if err != nil {
			return err
		}
		err = updateMenuFromFlags(cmd, newMenu)
		if err != nil {
			return err
		}
		return Client.UpdateMenu(newMenu)
	},
}

func reportMenus(cmd *cobra.Command, menuNames []string) error {
	for _, itemName := range menuNames {
		menu, err := Client.GetMenu(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(cmd, menu)
		fmt.Fprintln(cmd.OutOrStdout(), "")
	}
	return nil
}

var menuReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all menus in detail",
	Long:  `Shows detailed information about all menus.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		itemNames := make([]string, 0)
		if name == "" {
			itemNames, err = Client.ListMenuNames()
			if err != nil {
				return err
			}
		} else {
			itemNames = append(itemNames, name)
		}
		return reportMenus(cmd, itemNames)
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
	addCommonArgs(menuAddCmd)
	addStringFlags(menuAddCmd, menuStringFlagMetadata)

	// local flags for menu copy
	addCommonArgs(menuCopyCmd)
	addStringFlags(menuCopyCmd, menuStringFlagMetadata)
	menuCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for menu edit
	addCommonArgs(menuEditCmd)
	addStringFlags(menuEditCmd, menuStringFlagMetadata)
	menuEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for menu find
	addCommonArgs(menuFindCmd)
	addStringFlags(menuFindCmd, menuStringFlagMetadata)
	addStringFlags(menuFindCmd, findStringFlagMetadata)
	addIntFlags(menuFindCmd, findIntFlagMetadata)
	addFloatFlags(menuFindCmd, findFloatFlagMetadata)

	// local flags for menu remove
	menuRemoveCmd.Flags().String("name", "", "the menu name")
	menuRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for menu rename
	addCommonArgs(menuRenameCmd)
	addStringFlags(menuRenameCmd, menuStringFlagMetadata)
	menuRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for menu report
	menuReportCmd.Flags().String("name", "", "the menu name")
}
