// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewMkLoadersCmd builds a new command that represents the mkloaders action
func NewMkLoadersCmd() *cobra.Command {
	mkloadersCmd := &cobra.Command{
		Use:   "mkloaders",
		Short: "Generate GRUB 2 bootloaders",
		Long: `Generate UEFI bootable GRUB 2 bootloaders. If available on the operating system Cobbler is running on,
then this also generates bootloaders for different architectures then the one of the system.
The options are configured in the Cobbler settings file.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			eventId, err := Client.BackgroundMkLoaders()
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Event ID: %s\n", eventId)
			return nil
		},
	}
	return mkloadersCmd
}
