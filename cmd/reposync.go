// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// reposyncCmd represents the reposync command
var reposyncCmd = &cobra.Command{
	Use:   "reposync",
	Short: "Sync repositories",
	Long: `Update and sync Cobbler repositories. The repositories have to be added beforehand via 'cobbler repo add'.

See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-reposync for more information.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		noFailOption, err := cmd.Flags().GetBool("no-fail")
		if err != nil {
			return err
		}
		onlyOption, err := cmd.Flags().GetString("only")
		if err != nil {
			return err
		}
		triesOption, err := cmd.Flags().GetInt("tries")
		if err != nil {
			return err
		}
		var reposyncOptions = cobblerclient.BackgroundReposyncOptions{
			Repos:  make([]string, 0),
			Only:   onlyOption,
			Nofail: noFailOption,
			Tries:  triesOption,
		}
		eventId, err := Client.BackgroundReposync(reposyncOptions)
		if err != nil {
			return err
		}
		fmt.Printf("Event ID: %s\n", eventId)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(reposyncCmd)

	//local flags
	reposyncCmd.Flags().Bool("no-fail", false, "do not stop reposyncing if a failure occurs")
	reposyncCmd.Flags().String("only", "", "update only this repository name")
	reposyncCmd.Flags().Int("tries", 3, "try each repo this many times")
}
