// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// replicateCmd represents the replicate command
var replicateCmd = &cobra.Command{
	Use:   "replicate",
	Short: "Replicate data",
	Long: `Replicate configurations from a master Cobbler server. This feature is intended for load-balancing,
disaster-recovery, backup, or multiple geography support. Each Cobbler server is still expected to have a locally
relevant cobbler.conf and modules.conf, as these files are not synced.

See https://cobbler.readthedocs.io/en/latest/cobbler.html#cobbler-replicate for more information.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		generateCobblerClient()
		distrosOption, err := cmd.Flags().GetString("distros")
		if err != nil {
			return err
		}
		profilesOption, err := cmd.Flags().GetString("profiles")
		if err != nil {
			return err
		}
		systemsOption, err := cmd.Flags().GetString("systems")
		if err != nil {
			return err
		}
		reposOption, err := cmd.Flags().GetString("repos")
		if err != nil {
			return err
		}
		imagesOption, err := cmd.Flags().GetString("image")
		if err != nil {
			return err
		}
		mgmtClassesOption, err := cmd.Flags().GetString("mgmtclasses")
		if err != nil {
			return err
		}
		packagesOption, err := cmd.Flags().GetString("packages")
		if err != nil {
			return err
		}
		filesOption, err := cmd.Flags().GetString("files")
		if err != nil {
			return err
		}
		portOption, err := cmd.Flags().GetString("port")
		if err != nil {
			return err
		}
		masterOption, err := cmd.Flags().GetString("master")
		if err != nil {
			return err
		}
		pruneOption, err := cmd.Flags().GetBool("prune")
		if err != nil {
			return err
		}
		omitDataOption, err := cmd.Flags().GetBool("omit-data")
		if err != nil {
			return err
		}
		syncAllOption, err := cmd.Flags().GetBool("sync-all")
		if err != nil {
			return err
		}
		useSslOption, err := cmd.Flags().GetBool("use-ssl")
		if err != nil {
			return err
		}
		replicateOptions := cobblerclient.ReplicateOptions{
			Master:            masterOption,
			Port:              portOption,
			DistroPatterns:    distrosOption,
			ProfilePatterns:   profilesOption,
			SystemPatterns:    systemsOption,
			RepoPatterns:      reposOption,
			Imagepatterns:     imagesOption,
			MgmtclassPatterns: mgmtClassesOption,
			PackagePatterns:   packagesOption,
			FilePatterns:      filesOption,
			Prune:             pruneOption,
			OmitData:          omitDataOption,
			SyncAll:           syncAllOption,
			UseSsl:            useSslOption,
		}
		eventId, err := Client.BackgroundReplicate(replicateOptions)
		if err != nil {
			return err
		}
		fmt.Printf("EventID: %s\n", eventId)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(replicateCmd)

	//local flags
	replicateCmd.Flags().String("distros", "", "patterns of distros to replicate")
	replicateCmd.Flags().String("files", "", "patterns of files to replicate")
	replicateCmd.Flags().String("image", "", "patterns of images to replicate")
	replicateCmd.Flags().String("master", "", "Cobbler server to replicate from")
	replicateCmd.Flags().String("mgmtclasses", "", "patterns of mgmtclasses to replicate")
	replicateCmd.Flags().Bool("omit-data", false, "do not rsync data")
	replicateCmd.Flags().String("packages", "", "patterns of packages to replicate")
	replicateCmd.Flags().String("port", "", "remote port")
	replicateCmd.Flags().String("profiles", "", "patterns of profiles to replicate")
	replicateCmd.Flags().Bool("prune", false, "remove objects (of all types) not found on the master")
	replicateCmd.Flags().String("repos", "", "patterns of repos to replicate")
	replicateCmd.Flags().Bool("sync-all", false, "sync all data")
	replicateCmd.Flags().String("systems", "", "patterns of systems to replicate")
	replicateCmd.Flags().Bool("use-ssl", false, "use SSL to access the Cobbler master server API")
}
