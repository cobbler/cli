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

// buildisoCmd represents the buildiso command
var buildisoCmd = &cobra.Command{
	Use:   "buildiso",
	Short: "Build an ISO",
	Long:  "Build all profiles into a bootable CD image. All flags are optional.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(buildisoCmd)

	//local flags
	buildisoCmd.Flags().Bool("airgapped", false, "creates a standalone ISO with all distro and repo files for disconnected system installation")
	buildisoCmd.Flags().String("distro", "", "used with --standalone and --airgapped to create a distro-based ISO including all associated profiles/systems")
	buildisoCmd.Flags().Bool("exclude-dns", false, "prevents addition of name server addresses to the kernel boot options")
	buildisoCmd.Flags().String("iso", "", "output ISO to this file")
	buildisoCmd.Flags().String("mkisofs-opts", "", "extra options for mkisofs")
	buildisoCmd.Flags().String("profiles", "", "use these profiles only")
	buildisoCmd.Flags().String("source", "", "used with --standalone to specify a source for the distribution files")
	buildisoCmd.Flags().String("standalone", "", "creates a standalone ISO with all required distro files, but without any added repos")
	buildisoCmd.Flags().String("systems", "", "use these systems only")
	buildisoCmd.Flags().String("tempdir", "", "working directory")
}
