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

func updatePackageFromFlags(cmd *cobra.Command, p *cobbler.Package) error {
	// This object type doesn't have the in-place flag
	var err error
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			// If one of the previous flags has had an error just directly return.
			return
		}
		switch flag.Name {
		case "comment":
			var packageNewComment string
			packageNewComment, err = cmd.Flags().GetString("comment")
			if err != nil {
				return
			}
			p.Comment = packageNewComment
		case "owners":
			fallthrough
		case "owners-inherit":
			if cmd.Flags().Lookup("owners-inherit").Changed {
				p.Owners.Data = []string{}
				p.Owners.IsInherited, err = cmd.Flags().GetBool("owners-inherit")
				if err != nil {
					return
				}
			} else {
				var packageNewOwners []string
				packageNewOwners, err = cmd.Flags().GetStringSlice("owners")
				if err != nil {
					return
				}
				p.Owners.IsInherited = false
				p.Owners.Data = packageNewOwners
			}
		case "action":
			var packageNewAction string
			packageNewAction, err = cmd.Flags().GetString("action")
			if err != nil {
				return
			}
			p.Action = packageNewAction
		case "installer":
			var packageNewInstaller string
			packageNewInstaller, err = cmd.Flags().GetString("installer")
			if err != nil {
				return
			}
			p.Action = packageNewInstaller
		case "version":
			var packageNewVersion string
			packageNewVersion, err = cmd.Flags().GetString("version")
			if err != nil {
				return
			}
			p.Action = packageNewVersion
		}
	})
	return err
}

// NewPackageCmd builds a new command that represents the package action
func NewPackageCmd() (*cobra.Command, error) {
	packageCmd := &cobra.Command{
		Use:   "package",
		Short: "Package management",
		Long: `Let you manage packages.
See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-package for more information.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	packageCmd.AddCommand(NewPackageAddCmd())
	packageCmd.AddCommand(NewPackageCopyCmd())
	packageCmd.AddCommand(NewPackageEditCmd())
	packageCmd.AddCommand(NewPackageFindCmd())
	packageCmd.AddCommand(NewPackageListCmd())
	packageCmd.AddCommand(NewPackageRemoveCmd())
	packageCmd.AddCommand(NewPackageRenameCmd())
	packageCmd.AddCommand(NewPackageReportCmd())
	packageCmd.AddCommand(NewPackageExportCmd())
	return packageCmd, nil
}

func NewPackageAddCmd() *cobra.Command {
	packageAddCmd := &cobra.Command{
		Use:   "add",
		Short: "add package",
		Long:  `Adds a package.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			newPackage := cobbler.NewPackage()

			// internal fields (ctime, mtime, depth, uid) cannot be modified
			newPackage.Name, err = cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			// Update package in-memory
			err = updatePackageFromFlags(cmd, &newPackage)
			if err != nil {
				return err
			}
			// Create package via XML-RPC
			linuxpackage, err := Client.CreatePackage(newPackage)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Package %s created\n", linuxpackage.Name)
			return nil
		},
	}
	addCommonArgs(packageAddCmd)
	addStringFlags(packageAddCmd, packageStringFlagMetadata)
	return packageAddCmd
}

func NewPackageCopyCmd() *cobra.Command {
	packageCopyCmd := &cobra.Command{
		Use:   "copy",
		Short: "copy package",
		Long:  `Copies a package.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// Collect CLI flags
			packageName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			packageNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}

			// Get package handle
			packageHandle, err := Client.GetPackageHandle(packageName)
			if err != nil {
				return err
			}
			// Copy the package server-side
			err = Client.CopyPackage(packageHandle, packageNewName)
			if err != nil {
				return err
			}
			// Get the copied package from the API
			newPackage, err := Client.GetPackage(packageNewName, false, false)
			if err != nil {
				return err
			}
			// Update package in-memory
			err = updatePackageFromFlags(cmd, newPackage)
			if err != nil {
				return err
			}
			// Update the package via XML-RPC
			return Client.UpdatePackage(newPackage)
		},
	}
	addCommonArgs(packageCopyCmd)
	addStringFlags(packageCopyCmd, packageStringFlagMetadata)
	packageCopyCmd.Flags().String("newname", "", "the new package name")
	packageCopyCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return packageCopyCmd
}

func NewPackageEditCmd() *cobra.Command {
	packageEditCmd := &cobra.Command{
		Use:   "edit",
		Short: "edit package",
		Long:  `Edits a package.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			packageName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}

			// Get package from the API
			packageToEdit, err := Client.GetPackage(packageName, false, false)
			if err != nil {
				return err
			}
			// Update package in-memory
			err = updatePackageFromFlags(cmd, packageToEdit)
			if err != nil {
				return err
			}
			// Update package via XML-RPC
			return Client.UpdatePackage(packageToEdit)
		},
	}
	addCommonArgs(packageEditCmd)
	addStringFlags(packageEditCmd, packageStringFlagMetadata)
	packageEditCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return packageEditCmd
}

func NewPackageFindCmd() *cobra.Command {
	packageFindCmd := &cobra.Command{
		Use:   "find",
		Short: "find package",
		Long:  `Finds a given package.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return FindItemNames(cmd, args, "package")
		},
	}
	addCommonArgs(packageFindCmd)
	addStringFlags(packageFindCmd, packageStringFlagMetadata)
	addStringFlags(packageFindCmd, findStringFlagMetadata)
	addIntFlags(packageFindCmd, findIntFlagMetadata)
	addFloatFlags(packageFindCmd, findFloatFlagMetadata)
	return packageFindCmd
}

func NewPackageListCmd() *cobra.Command {
	packageListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all packages",
		Long:  `Lists all available packages.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			packageNames, err := Client.ListPackageNames()
			if err != nil {
				return err
			}
			listItems(cmd, "packages", packageNames)
			return nil
		},
	}
	return packageListCmd
}

func NewPackageRemoveCmd() *cobra.Command {
	packageRemoveCmd := &cobra.Command{
		Use:   "remove",
		Short: "remove package",
		Long:  `Removes a given package.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			return RemoveItemRecursive(cmd, args, "package")
		},
	}
	packageRemoveCmd.Flags().String("name", "", "the package name")
	packageRemoveCmd.Flags().Bool("recursive", false, "also delete child objects")
	return packageRemoveCmd
}

func NewPackageRenameCmd() *cobra.Command {
	packageRenameCmd := &cobra.Command{
		Use:   "rename",
		Short: "rename package",
		Long:  `Renames a given package.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// internal fields (ctime, mtime, depth, uid) cannot be modified
			packageName, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			packageNewName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}

			// Get package API handle
			packageHandle, err := Client.GetPackageHandle(packageName)
			if err != nil {
				return err
			}
			// Perform server-side package rename
			err = Client.RenamePackage(packageHandle, packageNewName)
			if err != nil {
				return err
			}
			// Get the renamed package from the API
			newPackage, err := Client.GetPackage(packageNewName, false, false)
			if err != nil {
				return err
			}
			// Update package in-memory
			err = updatePackageFromFlags(cmd, newPackage)
			if err != nil {
				return err
			}
			// Update package via XML-RPC
			return Client.UpdatePackage(newPackage)
		},
	}
	addCommonArgs(packageRenameCmd)
	addStringFlags(packageRenameCmd, packageStringFlagMetadata)
	packageRenameCmd.Flags().String("newname", "", "the new package name")
	packageRenameCmd.Flags().Bool("in-place", false, "edit items in kopts or autoinstall without clearing the other items")
	return packageRenameCmd
}

func reportPackages(cmd *cobra.Command, packageNames []string) error {
	for _, itemName := range packageNames {
		repo, err := Client.GetPackage(itemName, false, false)
		if err != nil {
			return err
		}
		printStructured(cmd, repo)
		fmt.Fprintln(cmd.OutOrStdout(), "")
	}
	return nil
}

func NewPackageReportCmd() *cobra.Command {
	packageReportCmd := &cobra.Command{
		Use:   "report",
		Short: "list all packages in detail",
		Long:  `Shows detailed information about all packages.`,
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
				itemNames, err = Client.ListRepoNames()
				if err != nil {
					return err
				}
			} else {
				itemNames = append(itemNames, name)
			}
			return reportPackages(cmd, itemNames)
		},
	}
	packageReportCmd.Flags().String("name", "", "the package name")
	return packageReportCmd
}

func NewPackageExportCmd() *cobra.Command {
	packageExportCmd := &cobra.Command{
		Use:   "export",
		Short: "export packages",
		Long:  `Export packages.`,
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
				itemNames, err = Client.ListPackageNames()
				if err != nil {
					return err
				}
			} else {
				itemNames = append(itemNames, name)
			}

			for _, itemName := range itemNames {
				linuxpackage, err := Client.GetPackage(itemName, false, false)
				if err != nil {
					return err
				}
				if formatOption == "json" {
					jsonDocument, err := json.Marshal(linuxpackage)
					if err != nil {
						return err
					}
					fmt.Fprintln(cmd.OutOrStdout(), string(jsonDocument))
				}
				if formatOption == "yaml" {
					yamlDocument, err := yaml.Marshal(linuxpackage)
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
	packageExportCmd.Flags().String("name", "", "the package name")
	packageExportCmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	return packageExportCmd
}
