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
