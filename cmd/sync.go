// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync Cobbler",
	Long: `Force a rewrite of all configuration files, distribution files in the TFTP root, and restart managed
services. It is used to repair or rebuild the contents of '/tftpboot' or '/var/www/cobbler' or when something has
changed behind the scenes. It brings the filesystem up to date with the configuration.

See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-sync for more information.`,
	Run: func(cmd *cobra.Command, args []string) {

		// not fully implemented in the cobblerclient library. You cannot use flags at the moment!
		Client.Sync()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	//local flags
	syncCmd.Flags().Bool("dhcp", false, "write DHCP config files and restart service")
	syncCmd.Flags().Bool("dns", false, "write DNS config files and restart service")
	syncCmd.Flags().String("systems", "", "run a sync only on specified systems")
	syncCmd.Flags().Bool("verbose", false, "more verbose output")
}
