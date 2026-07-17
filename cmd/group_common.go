// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

// groupStringSliceFlagMetadata holds the shared --items flag for the three
// 4.0.0 group item types (distro_group, profile_group, system_group).
var groupStringSliceFlagMetadata = map[string]FlagMetadata[[]string]{
	"items": {
		Name:         "items",
		DefaultValue: []string{},
		Usage:        "member item names (comma delimited)",
	},
}

// addGroupFlagSet registers the common flag set used by every group subcommand.
func addGroupFlagSet(cmd *cobra.Command) {
	addCommonArgs(cmd)
	addStringSliceFlags(cmd, groupStringSliceFlagMetadata)
}

// extractGroupFlags reads the --items / --comment flags off cmd and returns
// the values. The bools indicate whether the user explicitly set that flag.
func extractGroupFlags(cmd *cobra.Command) (comment string, commentSet bool, items []string, itemsSet bool, err error) {
	cmd.Flags().Visit(func(flag *pflag.Flag) {
		if err != nil {
			return
		}
		switch flag.Name {
		case "comment":
			comment, err = cmd.Flags().GetString("comment")
			commentSet = true
		case "items":
			items, err = cmd.Flags().GetStringSlice("items")
			itemsSet = true
		}
	})
	return
}

// writeExport marshals v to stdout in the requested format (json or yaml).
func writeExport(cmd *cobra.Command, format string, v interface{}) error {
	switch format {
	case "json":
		out, err := json.Marshal(v)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(out))
	case "yaml":
		out, err := yaml.Marshal(v)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), "---")
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), string(out))
	default:
		return fmt.Errorf("format must be json or yaml")
	}
	return nil
}
