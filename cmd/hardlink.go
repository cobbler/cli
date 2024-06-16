// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// hardlinkCmd represents the hardlink command
var hardlinkCmd = &cobra.Command{
	Use:   "hardlink",
	Short: "Hardlink files",
	Long:  "Hardlink all files where it is possible to improve performance.",
	Run: func(cmd *cobra.Command, args []string) {
		generateCobblerClient()
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(hardlinkCmd)
}
