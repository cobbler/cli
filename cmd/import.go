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

// TODO: this is my test file at the moment
package cmd

import (
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import operating system distributions",
	Long: `Import operating system distributions into Cobbler. This could be a mounted ISO, network rsync mirror or a tree in the filesystem.
See https://cobbler.readthedocs.io/en/latest/quickstart-guide.html#importing-your-first-distribution for more information.`,

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	//local flags
	importCmd.Flags().String("arch", "", "the architechture of the OS")
	importCmd.Flags().String("autoinstall", "", "assign this autoinstall file")
	importCmd.Flags().String("available-as", "", "do not mirror, the tree is here")
	importCmd.Flags().String("breed", "", "the breed type, e.g. suse, redhat, ubuntu, etc")
	importCmd.Flags().String("name", "", "the name of the imported distro, e.g. openSUSE_Leap_153")
	importCmd.Flags().String("os-version", "", "the version of the OS")
	importCmd.Flags().String("path", "", "local path or rsync location")
	importCmd.Flags().String("rsync-flags", "", "pass additional flags to rsync")
}
