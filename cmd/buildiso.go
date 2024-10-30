// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// buildisoCmd represents the buildiso command
var buildisoCmd = &cobra.Command{
	Use:   "buildiso",
	Short: "Build an ISO",
	Long:  "Build all profiles into a bootable CD image. All flags are optional.",
	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		isoOption, err := cmd.Flags().GetString("iso")
		if err != nil {
			return err
		}
		distroOption, err := cmd.Flags().GetString("distro")
		if err != nil {
			return err
		}
		xorrisofsOption, err := cmd.Flags().GetString("mkisofs-opts")
		if err != nil {
			return err
		}
		profilesOption, err := cmd.Flags().GetStringSlice("profiles")
		if err != nil {
			return err
		}
		sourceOption, err := cmd.Flags().GetString("source")
		if err != nil {
			return err
		}
		systemsOption, err := cmd.Flags().GetStringSlice("systems")
		if err != nil {
			return err
		}
		tempdirOption, err := cmd.Flags().GetString("tempdir")
		if err != nil {
			return err
		}
		standaloneOption, err := cmd.Flags().GetBool("standalone")
		if err != nil {
			return err
		}
		excludeDnsOption, err := cmd.Flags().GetBool("exclude-dns")
		if err != nil {
			return err
		}
		airgappedOption, err := cmd.Flags().GetBool("airgapped")
		if err != nil {
			return err
		}
		buildisoOptions := cobblerclient.BuildisoOptions{
			Iso:           isoOption,
			Profiles:      profilesOption,
			Systems:       systemsOption,
			BuildisoDir:   tempdirOption,
			Distro:        distroOption,
			Standalone:    standaloneOption,
			Airgapped:     airgappedOption,
			Source:        sourceOption,
			ExcludeDns:    excludeDnsOption,
			XorrisofsOpts: xorrisofsOption,
		}
		eventId, err := Client.BackgroundBuildiso(buildisoOptions)
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Event ID: %s\n", eventId)
		return nil
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
	buildisoCmd.Flags().StringSlice("profiles", []string{}, "use these profiles only")
	buildisoCmd.Flags().String("source", "", "used with --standalone to specify a source for the distribution files")
	buildisoCmd.Flags().Bool("standalone", false, "creates a standalone ISO with all required distro files, but without any added repos")
	buildisoCmd.Flags().StringSlice("systems", []string{}, "use these systems only")
	buildisoCmd.Flags().String("tempdir", "", "working directory")
}
