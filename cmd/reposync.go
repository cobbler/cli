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

// reposyncCmd represents the reposync command
var reposyncCmd = &cobra.Command{
	Use:   "reposync",
	Short: "Sync repositories",
	Long: `Update and sync Cobbler repositories. The repositories have to be added beforehand via 'cobbler repo add'.

See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-reposync for more information.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(reposyncCmd)

	//local flags
	reposyncCmd.Flags().Bool("no-fail", false, "do not stop reposyncing if a failure occurs")
	reposyncCmd.Flags().String("only", "", "update only this repository name")
	reposyncCmd.Flags().String("tries", "", "try each repo this many times")
}
