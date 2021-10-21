// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// reposyncCmd represents the reposync command
var reposyncCmd = &cobra.Command{
	Use:   "reposync",
	Short: "Sync repositories",
	Long: `Update and sync Cobbler repositories. The repositories have to be added beforehand via 'cobbler repo add'.

See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-reposync for more information.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(reposyncCmd)

	//local flags
	reposyncCmd.Flags().Bool("no-fail", false, "do not stop reposyncing if a failure occurs")
	reposyncCmd.Flags().String("only", "", "update only this repository name")
	reposyncCmd.Flags().String("tries", "", "try each repo this many times")
}
