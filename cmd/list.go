// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List configuration",
	Long: `Lists all configuration which Cobbler can obtain from the saved data. There are also report subcommands for
most of the other Cobbler commands (currently: distro, profile, system, repo, image, mgmtclass, package, file, menu).
Identical to 'cobbler report'`,

	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}

		distroNames, err := Client.ListDistroNames()
		if err != nil {
			return err
		}
		profileNames, err := Client.ListProfileNames()
		if err != nil {
			return err
		}
		systemNames, err := Client.ListSystemNames()
		if err != nil {
			return err
		}
		repoNames, err := Client.ListRepoNames()
		if err != nil {
			return err
		}
		imageNames, err := Client.ListImageNames()
		if err != nil {
			return err
		}
		mgmtClassNames, err := Client.ListMgmtClassNames()
		if err != nil {
			return err
		}
		packageNames, err := Client.ListPackageNames()
		if err != nil {
			return err
		}
		fileNames, err := Client.ListFileNames()
		if err != nil {
			return err
		}
		menuNames, err := Client.ListMenuNames()
		if err != nil {
			return err
		}
		listItems(cmd, "distros", distroNames)
		listItems(cmd, "profiles", profileNames)
		listItems(cmd, "systems", systemNames)
		listItems(cmd, "repos", repoNames)
		listItems(cmd, "images", imageNames)
		listItems(cmd, "mgmtclasses", mgmtClassNames)
		listItems(cmd, "packages", packageNames)
		listItems(cmd, "files", fileNames)
		listItems(cmd, "menus", menuNames)
		return nil
	},
}

func listItems(cmd *cobra.Command, what string, items []string) {
	fmt.Fprintf(cmd.OutOrStdout(), "%s:\n", what)
	sort.Strings(items)
	for _, item := range items {
		fmt.Fprintf(cmd.OutOrStdout(), "   %s\n", item)
	}
	fmt.Fprintln(cmd.OutOrStdout(), "")
}

func init() {
	rootCmd.AddCommand(listCmd)
}
