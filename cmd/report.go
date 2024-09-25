// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "List configuration in detail",
	Long: `Lists all configuration which Cobbler can obtain from the saved data. There are also report subcommands for
most of the other Cobbler commands (currently: distro, profile, system, repo, image, mgmtclass, package, file, menu).
Identical to 'cobbler list'`,

	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		// Distro
		fmt.Println("distros:")
		fmt.Println("==========")
		distroNames, err := Client.ListDistroNames()
		if err != nil {
			return err
		}
		err = reportDistros(distroNames)
		if err != nil {
			return err
		}
		fmt.Println("")

		// Profile
		fmt.Println("profiles:")
		fmt.Println("==========")
		profileNames, err := Client.ListProfileNames()
		if err != nil {
			return err
		}
		err = reportProfiles(profileNames)
		if err != nil {
			return err
		}
		fmt.Println("")

		// System
		fmt.Println("systems:")
		fmt.Println("==========")
		systemNames, err := Client.ListSystemNames()
		if err != nil {
			return err
		}
		err = reportSystems(systemNames)
		if err != nil {
			return err
		}
		fmt.Println("")

		// Repository
		fmt.Println("repos:")
		fmt.Println("==========")
		repoNames, err := Client.ListRepoNames()
		if err != nil {
			return err
		}
		err = reportRepos(repoNames)
		if err != nil {
			return err
		}
		fmt.Println("")

		// Image
		fmt.Println("images:")
		fmt.Println("==========")
		imageNames, err := Client.ListImageNames()
		if err != nil {
			return err
		}
		err = reportImages(imageNames)
		if err != nil {
			return err
		}
		fmt.Println("")

		// Mgmtclass
		fmt.Println("mgmtclasses:")
		fmt.Println("==========")
		mgmtClassNames, err := Client.ListMgmtClassNames()
		if err != nil {
			return err
		}
		err = reportMgmtClasses(mgmtClassNames)
		if err != nil {
			return err
		}
		fmt.Println("")

		// Package
		fmt.Println("packages:")
		fmt.Println("==========")
		packageNames, err := Client.ListPackageNames()
		if err != nil {
			return err
		}
		err = reportPackages(packageNames)
		if err != nil {
			return err
		}
		fmt.Println("")

		// File
		fmt.Println("files:")
		fmt.Println("==========")
		fileNames, err := Client.ListFileNames()
		if err != nil {
			return err
		}
		err = reportFiles(fileNames)
		if err != nil {
			return err
		}
		fmt.Println("")

		// Menu
		fmt.Println("menus:")
		fmt.Println("==========")
		menuNames, err := Client.ListMenuNames()
		if err != nil {
			return err
		}
		err = reportMenus(menuNames)
		if err != nil {
			return err
		}
		fmt.Println("")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
