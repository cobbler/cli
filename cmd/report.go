// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewReportCmd builds a new command that represents the report action
func NewReportCmd() *cobra.Command {
	reportCmd := &cobra.Command{
		Use:   "report",
		Short: "List configuration in detail",
		Long: `Lists all configuration which Cobbler can obtain from the saved data. There are also report subcommands for
most of the other Cobbler commands (currently: distro, profile, system, repo, image, mgmtclass, package, file, menu).
Identical to 'cobbler list'`,

		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// Distro
			fmt.Fprintln(cmd.OutOrStdout(), "distros:")
			fmt.Fprintln(cmd.OutOrStdout(), "==========")
			distroNames, err := Client.ListDistroNames()
			if err != nil {
				return err
			}
			err = reportDistros(cmd, distroNames)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")

			// Profile
			fmt.Fprintln(cmd.OutOrStdout(), "profiles:")
			fmt.Fprintln(cmd.OutOrStdout(), "==========")
			profileNames, err := Client.ListProfileNames()
			if err != nil {
				return err
			}
			err = reportProfiles(cmd, profileNames)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")

			// System
			fmt.Fprintln(cmd.OutOrStdout(), "systems:")
			fmt.Fprintln(cmd.OutOrStdout(), "==========")
			systemNames, err := Client.ListSystemNames()
			if err != nil {
				return err
			}
			err = reportSystems(cmd, systemNames)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")

			// Repository
			fmt.Fprintln(cmd.OutOrStdout(), "repos:")
			fmt.Fprintln(cmd.OutOrStdout(), "==========")
			repoNames, err := Client.ListRepoNames()
			if err != nil {
				return err
			}
			err = reportRepos(cmd, repoNames)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")

			// Image
			fmt.Fprintln(cmd.OutOrStdout(), "images:")
			fmt.Fprintln(cmd.OutOrStdout(), "==========")
			imageNames, err := Client.ListImageNames()
			if err != nil {
				return err
			}
			err = reportImages(cmd, imageNames)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")

			// Mgmtclass
			fmt.Fprintln(cmd.OutOrStdout(), "mgmtclasses:")
			fmt.Fprintln(cmd.OutOrStdout(), "==========")
			mgmtClassNames, err := Client.ListMgmtClassNames()
			if err != nil {
				return err
			}
			err = reportMgmtClasses(cmd, mgmtClassNames)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")

			// Package
			fmt.Fprintln(cmd.OutOrStdout(), "packages:")
			fmt.Fprintln(cmd.OutOrStdout(), "==========")
			packageNames, err := Client.ListPackageNames()
			if err != nil {
				return err
			}
			err = reportPackages(cmd, packageNames)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")

			// File
			fmt.Fprintln(cmd.OutOrStdout(), "files:")
			fmt.Fprintln(cmd.OutOrStdout(), "==========")
			fileNames, err := Client.ListFileNames()
			if err != nil {
				return err
			}
			err = reportFiles(cmd, fileNames)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")

			// Menu
			fmt.Fprintln(cmd.OutOrStdout(), "menus:")
			fmt.Fprintln(cmd.OutOrStdout(), "==========")
			menuNames, err := Client.ListMenuNames()
			if err != nil {
				return err
			}
			err = reportMenus(cmd, menuNames)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), "")
			return nil
		},
	}
	return reportCmd
}
