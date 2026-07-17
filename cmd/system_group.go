// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"

	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// NewSystemGroupCmd builds the `cobbler system-group` command tree.
func NewSystemGroupCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "system-group",
		Short: "Manage system groups",
		Long:  `Manage Cobbler 4.0.0 SystemGroup items.`,
	}
	cmd.AddCommand(newSystemGroupAddCmd())
	cmd.AddCommand(newSystemGroupCopyCmd())
	cmd.AddCommand(newSystemGroupEditCmd())
	cmd.AddCommand(newSystemGroupFindCmd())
	cmd.AddCommand(newSystemGroupListCmd())
	cmd.AddCommand(newSystemGroupRemoveCmd())
	cmd.AddCommand(newSystemGroupRenameCmd())
	cmd.AddCommand(newSystemGroupReportCmd())
	cmd.AddCommand(newSystemGroupExportCmd())
	return cmd, nil
}

func newSystemGroupAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a system group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			g := cobbler.NewSystemGroup()
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
			created, err := Client.CreateSystemGroup(g)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "System group %s created\n", created.Name)
			return nil
		},
	}
	addGroupFlagSet(cmd)
	return cmd
}

func newSystemGroupEditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "edit a system group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			g, err := Client.GetSystemGroup(name, false, false)
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
			return Client.UpdateSystemGroup(g)
		},
	}
	addGroupFlagSet(cmd)
	return cmd
}

func newSystemGroupCopyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy",
		Short: "copy a system group",
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
			handle, err := Client.GetSystemGroupHandle(name)
			if err != nil {
				return err
			}
			return Client.CopySystemGroup(handle, newName)
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringFlags(cmd, copyRenameStringFlagMetadata)
	return cmd
}

func newSystemGroupRenameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename",
		Short: "rename a system group",
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
			handle, err := Client.GetSystemGroupHandle(name)
			if err != nil {
				return err
			}
			return Client.RenameSystemGroup(handle, newName)
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringFlags(cmd, copyRenameStringFlagMetadata)
	return cmd
}

func newSystemGroupRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove a system group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			return RemoveItemRecursive(cmd, args, "system_group")
		},
	}
	cmd.Flags().String("name", "", "the system group name")
	cmd.Flags().Bool("recursive", false, "also delete child objects")
	return cmd
}

func newSystemGroupFindCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "find system groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			return FindItemNames(cmd, args, "system_group")
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringSliceFlags(cmd, groupStringSliceFlagMetadata)
	addPaginationFlags(cmd)
	return cmd
}

func newSystemGroupListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all system groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			names, err := Client.ListSystemGroupNames()
			if err != nil {
				return err
			}
			listItems(cmd, "system_groups", names)
			return nil
		},
	}
	return cmd
}

func newSystemGroupReportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "show system group details",
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
				names, err = Client.ListSystemGroupNames()
				if err != nil {
					return err
				}
			} else {
				names = append(names, name)
			}
			for _, n := range names {
				g, err := Client.GetSystemGroup(n, false, false)
				if err != nil {
					return err
				}
				printStructured(cmd, g)
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")
			}
			return nil
		},
	}
	cmd.Flags().String("name", "", "the system group name")
	return cmd
}

func newSystemGroupExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "export system groups",
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
				names, err = Client.ListSystemGroupNames()
				if err != nil {
					return err
				}
			} else {
				names = append(names, name)
			}
			for _, n := range names {
				g, err := Client.GetSystemGroup(n, false, false)
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
	cmd.Flags().String("name", "", "the system group name")
	cmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	return cmd
}
