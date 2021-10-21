// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// mkloadersCmd represents the mkloaders command
var mkloadersCmd = &cobra.Command{
	Use:   "mkloaders",
	Short: "Generate GRUB 2 bootloaders",
	Long: `Generate UEFI bootable GRUB 2 bootloaders. If available on the operating system Cobbler is running on,
then this also generates bootloaders for different architectures then the one of the system.
The options are configured in the Cobbler settings file.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(mkloadersCmd)
}
