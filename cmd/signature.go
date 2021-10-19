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

// signatureCmd represents the signature command
var signatureCmd = &cobra.Command{
	Use:   "signature",
	Short: "Signature handling",
	Long:  `Reloads, reports or updates the signatures of the distinct operating system versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(signatureCmd)

	//local flags
	signatureCmd.Flags().Bool("reload", false, "reload the signatures file")
	signatureCmd.Flags().Bool("report", false, "list the currently loaded signatures")
	signatureCmd.Flags().Bool("update", false, "update the signatures file")
}
