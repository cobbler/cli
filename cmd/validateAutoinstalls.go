// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewValidateAutoinstallsCmd builds a command that represents the validateAutoinstalls action
func NewValidateAutoinstallsCmd() *cobra.Command {
	validateAutoinstallsCmd := &cobra.Command{
		Use:   "validate-autoinstalls",
		Short: "Autoinstall validation",
		Long:  `Validates the autoinstall files.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			eventId, err := Client.BackgroundValidateAutoinstallFiles()
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Event ID: %s\n", eventId)
			return nil
		},
	}
	return validateAutoinstallsCmd
}
