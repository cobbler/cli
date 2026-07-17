// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"

	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// NewDistroGroupCmd builds the `cobbler distro-group` command tree.
func NewDistroGroupCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "distro-group",
		Short: "Manage distro groups",
		Long:  `Manage Cobbler 4.0.0 DistroGroup items.`,
	}
	cmd.AddCommand(newDistroGroupAddCmd())
	cmd.AddCommand(newDistroGroupCopyCmd())
	cmd.AddCommand(newDistroGroupEditCmd())
	cmd.AddCommand(newDistroGroupFindCmd())
	cmd.AddCommand(newDistroGroupListCmd())
	cmd.AddCommand(newDistroGroupRemoveCmd())
	cmd.AddCommand(newDistroGroupRenameCmd())
	cmd.AddCommand(newDistroGroupReportCmd())
	cmd.AddCommand(newDistroGroupExportCmd())
	return cmd, nil
}

func newDistroGroupAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a distro group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			g := cobbler.NewDistroGroup()
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			g.Name = name
			comment, commentSet, items, itemsSet, err := extractGroupFlags(cmd)
			if err != nil {
				return err
			}
			if commentSet {
				g.Comment = comment
			}
			if itemsSet {
				g.Members = items
			}
			created, err := Client.CreateDistroGroup(g)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Distro group %s created\n", created.Name)
			return nil
		},
	}
	addGroupFlagSet(cmd)
	return cmd
}

func newDistroGroupEditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "edit a distro group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			g, err := Client.GetDistroGroup(name, false, false)
			if err != nil {
				return err
			}
			comment, commentSet, items, itemsSet, err := extractGroupFlags(cmd)
			if err != nil {
				return err
			}
			if commentSet {
				g.Comment = comment
			}
			if itemsSet {
				g.Members = items
			}
			return Client.UpdateDistroGroup(g)
		},
	}
	addGroupFlagSet(cmd)
	return cmd
}

func newDistroGroupCopyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy",
		Short: "copy a distro group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			newName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}
			handle, err := Client.GetDistroGroupHandle(name)
			if err != nil {
				return err
			}
			return Client.CopyDistroGroup(handle, newName)
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringFlags(cmd, copyRenameStringFlagMetadata)
	return cmd
}

func newDistroGroupRenameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename",
		Short: "rename a distro group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			newName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}
			handle, err := Client.GetDistroGroupHandle(name)
			if err != nil {
				return err
			}
			return Client.RenameDistroGroup(handle, newName)
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringFlags(cmd, copyRenameStringFlagMetadata)
	return cmd
}

func newDistroGroupRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove a distro group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			return RemoveItemRecursive(cmd, args, "distro_group")
		},
	}
	cmd.Flags().String("name", "", "the distro group name")
	cmd.Flags().Bool("recursive", false, "also delete child objects")
	return cmd
}

func newDistroGroupFindCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "find distro groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			return FindItemNames(cmd, args, "distro_group")
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringSliceFlags(cmd, groupStringSliceFlagMetadata)
	addPaginationFlags(cmd)
	return cmd
}

func newDistroGroupListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all distro groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			names, err := Client.ListDistroGroupNames()
			if err != nil {
				return err
			}
			listItems(cmd, "distro_groups", names)
			return nil
		},
	}
	return cmd
}

func newDistroGroupReportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "show distro group details",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			names := make([]string, 0)
			if name == "" {
				names, err = Client.ListDistroGroupNames()
				if err != nil {
					return err
				}
			} else {
				names = append(names, name)
			}
			for _, n := range names {
				g, err := Client.GetDistroGroup(n, false, false)
				if err != nil {
					return err
				}
				printStructured(cmd, g)
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")
			}
			return nil
		},
	}
	cmd.Flags().String("name", "", "the distro group name")
	return cmd
}

func newDistroGroupExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "export distro groups",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			if format != "json" && format != "yaml" {
				return fmt.Errorf("format must be json or yaml")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			names := make([]string, 0)
			if name == "" {
				names, err = Client.ListDistroGroupNames()
				if err != nil {
					return err
				}
			} else {
				names = append(names, name)
			}
			for _, n := range names {
				g, err := Client.GetDistroGroup(n, false, false)
				if err != nil {
					return err
				}
				if err := writeExport(cmd, format, g); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().String("name", "", "the distro group name")
	cmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	return cmd
}
