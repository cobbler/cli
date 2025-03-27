// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"encoding/json"
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
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

// NewMenuCmd builds a new command that represents the menu action
func NewMenuCmd() (*cobra.Command, error) {
	menuCmd := &cobra.Command{
		Use:   "menu",
		Short: "Menu management",
		Long: `Let you manage menus.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-menu for more information.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	menuAddCmd, err := NewMenuAddCmd()
	if err != nil {
		return nil, err
	}
	menuCmd.AddCommand(menuAddCmd)
	menuCopyCmd, err := NewMenuCopyCmd()
	if err != nil {
		return nil, err
	}
	menuCmd.AddCommand(menuCopyCmd)
	menuCmd.AddCommand(NewMenuEditCmd())
	menuCmd.AddCommand(NewMenuFindCmd())
	menuCmd.AddCommand(NewMenuListCmd())
	menuCmd.AddCommand(NewMenuRemoveCmd())
	menuRenameCmd, err := NewMenuRenameCmd()
	if err != nil {
		return nil, err
	}
	menuCmd.AddCommand(menuRenameCmd)
	menuCmd.AddCommand(NewMenuReportCmd())
	menuCmd.AddCommand(NewMenuExportCmd())
	return menuCmd, nil
}

func NewMenuAddCmd() (*cobra.Command, error) {
	menuAddCmd := &cobra.Command{
		Use:   "add",
		Short: "add menu",
		Long:  `Adds a menu.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			newMenu := cobbler.NewMenu()

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
	addCommonArgs(menuAddCmd)
	addStringFlags(menuAddCmd, menuStringFlagMetadata)
	err := menuAddCmd.MarkFlagRequired("name")
	if err != nil {
		return nil, err
	}
	return menuAddCmd, nil
}

func NewMenuCopyCmd() (*cobra.Command, error) {
	menuCopyCmd := &cobra.Command{
		Use:   "copy",
		Short: "copy menu",
		Long:  `Copies a menu.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
	addCommonArgs(menuCopyCmd)
	addStringFlags(menuCopyCmd, menuStringFlagMetadata)
	addStringFlags(menuCopyCmd, copyRenameStringFlagMetadata)
	menuCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	err := menuCopyCmd.MarkFlagRequired("name")
	if err != nil {
		return nil, err
	}
	err = menuCopyCmd.MarkFlagRequired("newname")
	if err != nil {
		return nil, err
	}
	return menuCopyCmd, nil
}

func NewMenuEditCmd() *cobra.Command {
	menuEditCmd := &cobra.Command{
		Use:   "edit",
		Short: "edit menu",
		Long:  `Edits a menu.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
	addCommonArgs(menuEditCmd)
	addStringFlags(menuEditCmd, menuStringFlagMetadata)
	menuEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return menuEditCmd
}

func NewMenuFindCmd() *cobra.Command {
	menuFindCmd := &cobra.Command{
		Use:   "find",
		Short: "find menu",
		Long:  `Finds a given menu.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return FindItemNames(cmd, args, "menu")
		},
	}
	addCommonArgs(menuFindCmd)
	addStringFlags(menuFindCmd, menuStringFlagMetadata)
	addStringFlags(menuFindCmd, findStringFlagMetadata)
	addIntFlags(menuFindCmd, findIntFlagMetadata)
	addFloatFlags(menuFindCmd, findFloatFlagMetadata)
	return menuFindCmd
}

func NewMenuListCmd() *cobra.Command {
	menuListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all menus",
		Long:  `Lists all available menus.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			menuNames, err := Client.ListMenuNames()
			if err != nil {
				return err
			}
			listItems(cmd, "menus", menuNames)
			return nil
		},
	}
	return menuListCmd
}

func NewMenuRemoveCmd() *cobra.Command {
	menuRemoveCmd := &cobra.Command{
		Use:   "remove",
		Short: "remove menu",
		Long:  `Removes a given menu.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return RemoveItemRecursive(cmd, args, "menu")
		},
	}
	menuRemoveCmd.Flags().String("name", "", "the menu name")
	menuRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")
	return menuRemoveCmd
}

func NewMenuRenameCmd() (*cobra.Command, error) {
	menuRenameCmd := &cobra.Command{
		Use:   "rename",
		Short: "rename menu",
		Long:  `Renames a given menu.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
	addCommonArgs(menuRenameCmd)
	addStringFlags(menuRenameCmd, menuStringFlagMetadata)
	addStringFlags(menuRenameCmd, copyRenameStringFlagMetadata)
	menuRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	err := menuRenameCmd.MarkFlagRequired("name")
	if err != nil {
		return nil, err
	}
	err = menuRenameCmd.MarkFlagRequired("newname")
	if err != nil {
		return nil, err
	}
	return menuRenameCmd, nil
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

func NewMenuReportCmd() *cobra.Command {
	menuReportCmd := &cobra.Command{
		Use:   "report",
		Short: "list all menus in detail",
		Long:  `Shows detailed information about all menus.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

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
	menuReportCmd.Flags().String("name", "", "the menu name")
	return menuReportCmd
}

func NewMenuExportCmd() *cobra.Command {
	menuExportCmd := &cobra.Command{
		Use:   "export",
		Short: "export menus",
		Long:  `Export menus.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			formatOption, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			if formatOption != "json" && formatOption != "yaml" {
				return fmt.Errorf("format must be json or yaml")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			formatOption, err := cmd.Flags().GetString("format")
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

			for _, itemName := range itemNames {
				menu, err := Client.GetMenu(itemName, false, false)
				if err != nil {
					return err
				}
				if formatOption == "json" {
					jsonDocument, err := json.Marshal(menu)
					if err != nil {
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), string(jsonDocument))
				}
				if formatOption == "yaml" {
					yamlDocument, err := yaml.Marshal(menu)
					if err != nil {
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), "---")
					fmt.Fprintln(cmd.OutOrStdout(), string(yamlDocument))
				}
			}
			return nil
		},
	}
	menuExportCmd.Flags().String("name", "", "the menu name")
	menuExportCmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	return menuExportCmd
}
