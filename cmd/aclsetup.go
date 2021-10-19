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

// aclsetupCmd represents the aclsetup command
var aclsetupCmd = &cobra.Command{
	Use:   "aclsetup",
	Short: "Adjust the access control list",
	Long:  "Configures users/groups to run the Cobbler CLI as non-root.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(aclsetupCmd)

	//local flags
	aclsetupCmd.Flags().String("adduser", "", "give acls to this user")
	aclsetupCmd.Flags().String("addgroup", "", "give acls to this group")
	aclsetupCmd.Flags().String("removeuser", "", "remove acls from this user")
	aclsetupCmd.Flags().String("removegroup", "", "remove acls from this user")
}
