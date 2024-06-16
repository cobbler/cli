// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "List configuration in detail",
	Long: `Lists all configuration which Cobbler can obtain from the saved data. There are also report subcommands for
most of the other Cobbler commands (currently: distro, profile, system, repo, image, mgmtclass, package, file, menu).
Identical to 'cobbler list'`,

	Run: func(cmd *cobra.Command, args []string) {
		generateCobblerClient()

		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
