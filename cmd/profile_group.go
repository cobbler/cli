// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"

	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// NewProfileGroupCmd builds the `cobbler profile-group` command tree.
func NewProfileGroupCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "profile-group",
		Short: "Manage profile groups",
		Long:  `Manage Cobbler 4.0.0 ProfileGroup items.`,
	}
	cmd.AddCommand(newProfileGroupAddCmd())
	cmd.AddCommand(newProfileGroupCopyCmd())
	cmd.AddCommand(newProfileGroupEditCmd())
	cmd.AddCommand(newProfileGroupFindCmd())
	cmd.AddCommand(newProfileGroupListCmd())
	cmd.AddCommand(newProfileGroupRemoveCmd())
	cmd.AddCommand(newProfileGroupRenameCmd())
	cmd.AddCommand(newProfileGroupReportCmd())
	cmd.AddCommand(newProfileGroupExportCmd())
	return cmd, nil
}

func newProfileGroupAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a profile group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			g := cobbler.NewProfileGroup()
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
			created, err := Client.CreateProfileGroup(g)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Profile group %s created\n", created.Name)
			return nil
		},
	}
	addGroupFlagSet(cmd)
	return cmd
}

func newProfileGroupEditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "edit a profile group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			resolvedUID, err := resolveUID(&Client, "profile_group", name, uid)
			if err != nil {
				return err
			}
			g, err := Client.GetProfileGroup(resolvedUID, false, false)
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
			return Client.UpdateProfileGroup(g)
		},
	}
	addGroupFlagSet(cmd)
	addUIDFlag(cmd, "profile group")
	return cmd
}

func newProfileGroupCopyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "copy",
		Short: "copy a profile group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			newName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}
			handle, err := resolveUID(&Client, "profile_group", name, uid)
			if err != nil {
				return err
			}
			return Client.CopyProfileGroup(handle, newName)
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringFlags(cmd, copyRenameStringFlagMetadata)
	addUIDFlag(cmd, "profile group")
	return cmd
}

func newProfileGroupRenameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rename",
		Short: "rename a profile group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			newName, err := cmd.Flags().GetString("newname")
			if err != nil {
				return err
			}
			handle, err := resolveUID(&Client, "profile_group", name, uid)
			if err != nil {
				return err
			}
			return Client.RenameProfileGroup(handle, newName)
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringFlags(cmd, copyRenameStringFlagMetadata)
	addUIDFlag(cmd, "profile group")
	return cmd
}

func newProfileGroupRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove a profile group",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			return RemoveItemRecursive(cmd, args, "profile_group")
		},
	}
	cmd.Flags().String("name", "", "the profile group name")
	cmd.Flags().Bool("recursive", false, "also delete child objects")
	addUIDFlag(cmd, "profile group")
	return cmd
}

func newProfileGroupFindCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find",
		Short: "find profile groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			return FindItemNames(cmd, args, "profile_group")
		},
	}
	addStringFlags(cmd, commonStringFlagMetadata)
	addStringSliceFlags(cmd, groupStringSliceFlagMetadata)
	addPaginationFlags(cmd)
	return cmd
}

func newProfileGroupListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all profile groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			names, err := Client.ListProfileGroupNames()
			if err != nil {
				return err
			}
			listItems(cmd, "profile_groups", names)
			return nil
		},
	}
	return cmd
}

func newProfileGroupReportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "show profile group details",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := generateCobblerClient(); err != nil {
				return err
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			var groups []*cobbler.ProfileGroup
			if name == "" && uid == "" {
				groups, err = Client.GetProfileGroups()
				if err != nil {
					return err
				}
			} else {
				resolvedUID, err := resolveUID(&Client, "profile_group", name, uid)
				if err != nil {
					return err
				}
				g, err := Client.GetProfileGroup(resolvedUID, false, false)
				if err != nil {
					return err
				}
				groups = []*cobbler.ProfileGroup{g}
			}
			for _, g := range groups {
				printStructured(cmd, g)
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")
			}
			return nil
		},
	}
	cmd.Flags().String("name", "", "the profile group name")
	addUIDFlag(cmd, "profile group")
	return cmd
}

func newProfileGroupExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "export profile groups",
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
			uid, err := cmd.Flags().GetString("uid")
			if err != nil {
				return err
			}
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			var groups []*cobbler.ProfileGroup
			if name == "" && uid == "" {
				groups, err = Client.GetProfileGroups()
				if err != nil {
					return err
				}
			} else {
				resolvedUID, err := resolveUID(&Client, "profile_group", name, uid)
				if err != nil {
					return err
				}
				g, err := Client.GetProfileGroup(resolvedUID, false, false)
				if err != nil {
					return err
				}
				groups = []*cobbler.ProfileGroup{g}
			}
			for _, g := range groups {
				if err := writeExport(cmd, format, g); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().String("name", "", "the profile group name")
	cmd.Flags().String(exportStringMetadata["format"].Name, exportStringMetadata["format"].DefaultValue, exportStringMetadata["format"].Usage)
	addUIDFlag(cmd, "profile group")
	return cmd
}
