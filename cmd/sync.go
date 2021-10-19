/*
Copyright © 2021 Dominik Gedon <dgedon@suse.de>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
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
		// TODO: call cobblerclient
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
