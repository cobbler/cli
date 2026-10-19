// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCheckCmd builds a new command that represents the check action
func NewCheckCmd() *cobra.Command {
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Check the Cobbler server configuration for common problems.",
		Long: `This command asks the Cobbler server to inspect its own configuration (settings, installed
dependencies, boot-loader files, services, etc.) and reports suggestions for items you may want to fix. It
does not change anything on the server; it only lists potential problems for the administrator to review.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			checks, err := Client.Check()
			if err != nil {
				return err
			}

			if len(*checks) == 0 {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "No configuration problems found.")
				return nil
			}

			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "The following are potential configuration items that you may want to fix:")
			for i, check := range *checks {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%d : %s\n", i+1, check)
			}
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Restart cobblerd and then run 'cobbler sync' to apply changes.")

			return nil
		},
	}
	return checkCmd
}
