// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// mkgrubCmd represents the mkgrub command
var mkgrubCmd = &cobra.Command{
	Use:   "mkgrub",
	Short: "Generate GRUB bootloaders",
	Long: `Generate UEFI bootable GRUB 2 bootloaders. If available on the operating system Cobbler is running on,
then this also generates bootloaders for different architectures then the one of the system.
The options are configured in the Cobbler settings file.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
		// Check for Cobbler version and decide which command to run
		// Cobbler mkgrub is not yet available with the most recent version, the older version use mkloaders
	},
}

func init() {
	rootCmd.AddCommand(mkgrubCmd)
}
