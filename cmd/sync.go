// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/cobbler/cobblerclient"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		dhcpOption, err := cmd.Flags().GetBool("dhcp")
		if err != nil {
			return err
		}
		dnsOption, err := cmd.Flags().GetBool("dns")
		if err != nil {
			return err
		}
		verboseOption, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return err
		}
		systemsOption, err := cmd.Flags().GetStringSlice("systems")
		if err != nil {
			return err
		}

		var eventId string
		if len(systemsOption) > 0 {
			backgroundSyncSystemsOptions := cobblerclient.BackgroundSyncSystemsOptions{
				Systems: systemsOption,
				Verbose: verboseOption,
			}
			eventId, err = Client.BackgroundSyncSystems(backgroundSyncSystemsOptions)
		} else {
			backgroundSyncOptions := cobblerclient.BackgroundSyncOptions{
				Dhcp:    dhcpOption,
				Dns:     dnsOption,
				Verbose: verboseOption,
			}
			eventId, err = Client.BackgroundSync(backgroundSyncOptions)
		}

		if err != nil {
			return err
		}
		fmt.Printf("Event ID: %s\n", eventId)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	//local flags
	syncCmd.Flags().Bool("dhcp", false, "write DHCP config files and restart service")
	syncCmd.Flags().Bool("dns", false, "write DNS config files and restart service")
	syncCmd.Flags().StringSlice("systems", []string{}, "run a sync only on specified systems")
	syncCmd.Flags().Bool("verbose", false, "more verbose output")
	syncCmd.MarkFlagsMutuallyExclusive("dhcp", "systems")
	syncCmd.MarkFlagsMutuallyExclusive("dns", "systems")
}
