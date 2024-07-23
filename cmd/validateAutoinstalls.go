// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// validateAutoinstallsCmd represents the validateAutoinstalls command
var validateAutoinstallsCmd = &cobra.Command{
	Use:   "validate-autoinstalls",
	Short: "Autoinstall validation",
	Long:  `Validates the autoinstall files.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		eventId, err := Client.BackgroundValidateAutoinstallFiles()
		if err != nil {
			return err
		}
		fmt.Printf("Event ID: %s\n", eventId)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateAutoinstallsCmd)
}
