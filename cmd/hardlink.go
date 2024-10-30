// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// hardlinkCmd represents the hardlink command
var hardlinkCmd = &cobra.Command{
	Use:   "hardlink",
	Short: "Hardlink files",
	Long:  "Hardlink all files where it is possible to improve performance.",
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		eventId, err := Client.BackgroundHardlink()
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Event ID: %s\n", eventId)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hardlinkCmd)
}
