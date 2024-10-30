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

func updateMgmtClassFromFlags(cmd *cobra.Command, mgmtClass *cobbler.MgmtClass) error {
	inPlace, err := cmd.Flags().GetBool("in-place")
	if err != nil {
		return err
	}
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			// If one of the previous flags has had an error just directly return.
			return
		}
		switch flag.Name {
		// The rename & copy operations are special operations as such we cannot blindly set this inside here.
		// Any rename & copy operation must be handled outside of this method.
		case "comment":
			var mgmtClassNewComment string
			mgmtClassNewComment, err = cmd.Flags().GetString("comment")
			if err != nil {
				return
			}
			mgmtClass.Comment = mgmtClassNewComment
		case "owners":
			fallthrough
		case "owners-inherit":
			if cmd.Flags().Lookup("owners-inherit").Changed {
				mgmtClass.Owners.Data = []string{}
				mgmtClass.Owners.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				var mgmtClassNewOwners []string
				mgmtClassNewOwners, err = cmd.Flags().GetStringSlice("owners")
				if err != nil {
					return
				}
				mgmtClass.Owners.IsInherited = false
				mgmtClass.Owners.Data = mgmtClassNewOwners
			}
		case "files":
			var mgmtClassNewFiles []string
			mgmtClassNewFiles, err = cmd.Flags().GetStringSlice("files")
			if err != nil {
				return
			}
			mgmtClass.Files = mgmtClassNewFiles
		case "packages":
			var mgmtClassNewPackages []string
			mgmtClassNewPackages, err = cmd.Flags().GetStringSlice("packages")
			if err != nil {
				return
			}
			mgmtClass.Packages = mgmtClassNewPackages
		case "params":
			var mgmtClassNewParams map[string]string
			mgmtClassNewParams, err = cmd.Flags().GetStringToString("params")
			if err != nil {
				return
			}
			if inPlace {
				err = Client.ModifyItemInPlace(
					"mgmtclass",
					mgmtClass.Name,
					"params",
					convertMapStringToMapInterface(mgmtClassNewParams),
				)
				if err != nil {
					return
				}
			} else {
				mgmtClass.Params = mgmtClassNewParams
			}
		case "class-name":
			var mgmtClassNewClassName string
			mgmtClassNewClassName, err = cmd.Flags().GetString("class-name")
			if err != nil {
				return
			}
			mgmtClass.ClassName = mgmtClassNewClassName
		case "is-definition":
			var mgmtClassNewIsDefinition bool
			mgmtClassNewIsDefinition, err = cmd.Flags().GetBool("is-definition")
			if err != nil {
				return
			}
			mgmtClass.IsDefiniton = mgmtClassNewIsDefinition
		}
	})
	// Don't blindly return nil because maybe one of the flags had an issue retrieving an argument.
	return err
}

// mgmtclassCmd represents the mgmtclass command
var mgmtclassCmd = &cobra.Command{
	Use:   "mgmtclass",
	Short: "Mgmtclass management",
	Long: `Let you manage mgmtclasses.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-mgmtclass for more information.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var mgmtclassAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add mgmtclass",
	Long:  `Adds a mgmtclass.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		newMgmtClass := cobbler.NewMgmtClass()
		var err error

		// Get special name flag
		newMgmtClass.Name, err = cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		// Update with the rest of the flags
		err = updateMgmtClassFromFlags(cmd, &newMgmtClass)
		if err != nil {
			return err
		}
		// Now create the file via XML-RPC
		mgmtClass, err := Client.CreateMgmtClass(newMgmtClass)
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Mgmtclass %s created\n", mgmtClass.Name)
		return nil
	},
}

var mgmtclassCopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy mgmtclass",
	Long:  `Copies a mgmtclass.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Get special name and newname flags
		mgmtClassName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		mgmtClassNewName, err := cmd.Flags().GetString("newname")
		if err != nil {
			return err
		}

		// Get API handle
		mgmtClassHandle, err := Client.GetMgmtClassHandle(mgmtClassName)
		if err != nil {
			return err
		}
		// Copy the mgmtclass server-side
		err = Client.CopyMgmtClass(mgmtClassHandle, mgmtClassNewName)
		if err != nil {
			return err
		}
		// Get the copied mgmtclass
		newMgmtClass, err := Client.GetMgmtClass(mgmtClassNewName, false, false)
		if err != nil {
			return err
		}
		// Update the mgmtclass in-memory
		err = updateMgmtClassFromFlags(cmd, newMgmtClass)
		if err != nil {
			return err
		}
		// Update the mgmtclass via XML-RPC
		return Client.UpdateMgmtClass(newMgmtClass)
	},
}

var mgmtclassEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit mgmtclass",
	Long:  `Edits a mgmtclass.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Collect CLI flags
		mgmtClassName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		// Get mgmtclass from the API
		mgmtClassToEdit, err := Client.GetMgmtClass(mgmtClassName, false, false)
		if err != nil {
			return err
		}
		// Update mgmtclass in-memory
		err = updateMgmtClassFromFlags(cmd, mgmtClassToEdit)
		if err != nil {
			return err
		}
		// Update the mgmtclass via XML-RPC
		return Client.UpdateMgmtClass(mgmtClassToEdit)
	},
}

var mgmtclassFindCmd = &cobra.Command{
	Use:   "find",
	Short: "find mgmtclass",
	Long:  `Finds a given mgmtclass.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		return FindItemNames(cmd, args, "mgmtclass")
	},
}

var mgmtclassListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all mgmtclasses",
	Long:  `Lists all available mgmtclasses.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		mgmtclassNames, err := Client.ListMgmtClassNames()
		if err != nil {
			return err
		}
		listItems(cmd, "mgmtclasses", mgmtclassNames)
		return nil
	},
}

var mgmtclassRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove mgmtclass",
	Long:  `Removes a given mgmtclass.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		return RemoveItemRecursive(cmd, args, "mgmtclass")
	},
}

var mgmtclassRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "rename mgmtclass",
	Long:  `Renames a given mgmtclass.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()

		// Get the special name and newname flags
		mgmtClassName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		mgmtClassNewName, err := cmd.Flags().GetString("newname")
		if err != nil {
			return err
		}

		// Get the mgmtclass handle
		mgmtClassHandle, err := Client.GetMgmtClassHandle(mgmtClassName)
		if err != nil {
			return err
		}
		// Rename the mgmtclass server-side
		err = Client.RenameMgmtClass(mgmtClassHandle, mgmtClassNewName)
		if err != nil {
			return err
		}
		// Get the renamed mgmtclass
		renamedMgmtClass, err := Client.GetMgmtClass(mgmtClassNewName, false, false)
		if err != nil {
			return err
		}
		// Update mgmtclass in-memory
		err = updateMgmtClassFromFlags(cmd, renamedMgmtClass)
		if err != nil {
			return err
		}
		// Update the mgmtclass via XML-RPC
		return Client.UpdateMgmtClass(renamedMgmtClass)
	},
}

func reportMgmtClasses(cmd *cobra.Command, mgmtClassNames []string) error {
	for _, itemName := range mgmtClassNames {
		system, err := Client.GetMgmtClass(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(cmd, system)
		fmt.Fprintln(cmd.OutOrStdout(), "")
	}
	return nil
}

var mgmtclassReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list all mgmtclasses in detail",
	Long:  `Shows detailed information about all mgmtclasses.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		itemNames := make([]string, 0)
		if name == "" {
			itemNames, err = Client.ListMgmtClassNames()
			if err != nil {
				return err
			}
		} else {
			itemNames = append(itemNames, name)
		}
		return reportMgmtClasses(cmd, itemNames)
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
	addCommonArgs(mgmtclassAddCmd)
	addStringFlags(mgmtclassAddCmd, mgmtclassStringFlagMetadata)
	addBoolFlags(mgmtclassAddCmd, mgmtclassBoolFlagMetadata)
	addStringSliceFlags(mgmtclassAddCmd, mgmtclassStringSliceFlagMetadata)
	addMapFlags(mgmtclassAddCmd, mgmtclassStringMapFlagMetadata)

	// local flags for mgmtclass copy
	addCommonArgs(mgmtclassCopyCmd)
	addStringFlags(mgmtclassCopyCmd, mgmtclassStringFlagMetadata)
	addBoolFlags(mgmtclassCopyCmd, mgmtclassBoolFlagMetadata)
	addStringSliceFlags(mgmtclassCopyCmd, mgmtclassStringSliceFlagMetadata)
	addMapFlags(mgmtclassCopyCmd, mgmtclassStringMapFlagMetadata)
	mgmtclassCopyCmd.Flags().String("newname", "", "the new mgmtclass name")
	mgmtclassCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for mgmtclass edit
	addCommonArgs(mgmtclassEditCmd)
	addStringFlags(mgmtclassEditCmd, mgmtclassStringFlagMetadata)
	addBoolFlags(mgmtclassEditCmd, mgmtclassBoolFlagMetadata)
	addStringSliceFlags(mgmtclassEditCmd, mgmtclassStringSliceFlagMetadata)
	addMapFlags(mgmtclassEditCmd, mgmtclassStringMapFlagMetadata)
	mgmtclassEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for mgmtclass find
	addCommonArgs(mgmtclassFindCmd)
	addStringFlags(mgmtclassFindCmd, mgmtclassStringFlagMetadata)
	addBoolFlags(mgmtclassFindCmd, mgmtclassBoolFlagMetadata)
	addStringSliceFlags(mgmtclassFindCmd, mgmtclassStringSliceFlagMetadata)
	addMapFlags(mgmtclassFindCmd, mgmtclassStringMapFlagMetadata)
	addStringFlags(mgmtclassFindCmd, findStringFlagMetadata)
	addIntFlags(mgmtclassFindCmd, findIntFlagMetadata)
	addFloatFlags(mgmtclassFindCmd, findFloatFlagMetadata)

	// local flags for mgmtclass remove
	mgmtclassRemoveCmd.Flags().String("name", "", "the mgmtclass name")
	mgmtclassRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")

	// local flags for mgmtclass rename
	addCommonArgs(mgmtclassRenameCmd)
	addStringFlags(mgmtclassRenameCmd, mgmtclassStringFlagMetadata)
	addBoolFlags(mgmtclassRenameCmd, mgmtclassBoolFlagMetadata)
	addStringSliceFlags(mgmtclassRenameCmd, mgmtclassStringSliceFlagMetadata)
	addMapFlags(mgmtclassRenameCmd, mgmtclassStringMapFlagMetadata)
	mgmtclassRenameCmd.Flags().String("newname", "", "the new mgmtclass name")
	mgmtclassRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")

	// local flags for mgmtclass report
	mgmtclassReportCmd.Flags().String("name", "", "the mgmtclass name")
}
