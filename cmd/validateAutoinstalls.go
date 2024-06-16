// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// validateAutoinstallsCmd represents the validateAutoinstalls command
var validateAutoinstallsCmd = &cobra.Command{
	Use:   "validate-autoinstalls",
	Short: "Autoinstall validation",
	Long:  `Validates the autoinstall files.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateCobblerClient()
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(validateAutoinstallsCmd)
}
