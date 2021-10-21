// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// signatureCmd represents the signature command
var signatureCmd = &cobra.Command{
	Use:   "signature",
	Short: "Signature handling",
	Long:  `Reloads, reports or updates the signatures of the distinct operating system versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(signatureCmd)

	//local flags
	signatureCmd.Flags().Bool("reload", false, "reload the signatures file")
	signatureCmd.Flags().Bool("report", false, "list the currently loaded signatures")
	signatureCmd.Flags().Bool("update", false, "update the signatures file")
}
