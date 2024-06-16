// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

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
		generateCobblerClient()
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
