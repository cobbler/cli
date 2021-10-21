// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
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

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
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
