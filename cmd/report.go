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
most of the other Cobbler commands (currently: distro, profile, system, repo, image, menu).
Identical to 'cobbler list'`,

		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// Distro
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "distros:")
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "==========")
			distros, err := Client.GetDistros()
			if err != nil {
				return err
			}
			err = reportDistros(cmd, distros)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")

			// Profile
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "profiles:")
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "==========")
			profiles, err := Client.GetProfiles()
			if err != nil {
				return err
			}
			err = reportProfiles(cmd, profiles)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")

			// System
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "systems:")
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "==========")
			systems, err := Client.GetSystems()
			if err != nil {
				return err
			}
			err = reportSystems(cmd, systems)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")

			// Repository
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "repos:")
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "==========")
			repos, err := Client.GetRepos()
			if err != nil {
				return err
			}
			err = reportRepos(cmd, repos)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")

			// Image
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "images:")
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "==========")
			images, err := Client.GetImages()
			if err != nil {
				return err
			}
			err = reportImages(cmd, images)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")

			// Menu
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "menus:")
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "==========")
			menus, err := Client.GetMenus()
			if err != nil {
				return err
			}
			err = reportMenus(cmd, menus)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "")
			return nil
		},
	}
	return reportCmd
}
