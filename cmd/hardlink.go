// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewHardlinkCmd builds a new commdand that represents the hardlink action
func NewHardlinkCmd() *cobra.Command {
	hardlinkCmd := &cobra.Command{
		Use:   "hardlink",
		Short: "Hardlink files",
		Long:  "Hardlink all files where it is possible to improve performance.",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			eventId, err := Client.BackgroundHardlink()
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Event ID: %s\n", eventId)
			return nil
		},
	}
	return hardlinkCmd
}
