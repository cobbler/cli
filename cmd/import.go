// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import operating system distributions",
	Long: `Import operating system distributions into Cobbler. This could be a mounted ISO, network rsync mirror or a tree in the filesystem.
See https://cobbler.readthedocs.io/en/latest/quickstart-guide.html#importing-your-first-distribution for more information.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}

		archOption, err := cmd.Flags().GetString("arch")
		if err != nil {
			return err
		}
		autoinstallOption, err := cmd.Flags().GetString("autoinstall")
		if err != nil {
			return err
		}
		availableAsOption, err := cmd.Flags().GetString("available-as")
		if err != nil {
			return err
		}
		breedOption, err := cmd.Flags().GetString("breed")
		if err != nil {
			return err
		}
		nameOption, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		osVersionOption, err := cmd.Flags().GetString("os-version")
		if err != nil {
			return err
		}
		pathOption, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}
		rsyncFlagsOption, err := cmd.Flags().GetString("rsync-flags")
		if err != nil {
			return err
		}
		var backgroundOptions = cobblerclient.BackgroundImportOptions{
			Path:            pathOption,
			Name:            nameOption,
			AvailableAs:     availableAsOption,
			AutoinstallFile: autoinstallOption,
			RsyncFlags:      rsyncFlagsOption,
			Arch:            archOption,
			Breed:           breedOption,
			OsVersion:       osVersionOption,
		}
		eventId, err := Client.BackgroundImport(backgroundOptions)
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Event ID: %s\n", eventId)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	//local flags
	importCmd.Flags().String("arch", "", "the architechture of the OS")
	importCmd.Flags().String("autoinstall", "", "assign this autoinstall file")
	importCmd.Flags().String("available-as", "", "do not mirror, the tree is here")
	importCmd.Flags().String("breed", "", "the breed type, e.g. suse, redhat, ubuntu, etc")
	importCmd.Flags().String("name", "", "the name of the imported distro, e.g. openSUSE_Leap_153")
	importCmd.Flags().String("os-version", "", "the version of the OS")
	importCmd.Flags().String("path", "", "local path or rsync location")
	importCmd.Flags().String("rsync-flags", "", "pass additional flags to rsync")
}
