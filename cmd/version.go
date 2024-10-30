// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime/debug"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the Cobbler version",
	Long:  `Shows the Cobbler server version.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateCobblerClient()
		version, err := Client.ExtendedVersion()
		if err != nil {
			fmt.Println(err)
		}
		clientVersion, cliVersion, _ := getClientVersion()
		fmt.Printf("Cobbler %s\n", version.Version)
		fmt.Printf("  source: %s, %s\n", version.Gitstamp, version.Gitdate)
		fmt.Printf("  build time: %s\n", version.Builddate)
		fmt.Printf("  cli: %s\n", cliVersion)
		fmt.Printf("  client: %s\n", clientVersion)
		fmt.Fprintf(cmd.OutOrStdout(), "Cobbler %s\n", version.Version)
		fmt.Fprintf(cmd.OutOrStdout(), "  source: %s, %s\n", version.Gitstamp, version.Gitdate)
		fmt.Fprintf(cmd.OutOrStdout(), "  build time: %s\n", version.Builddate)
		fmt.Fprintf(cmd.OutOrStdout(), "  cli: %s\n", cliVersion)
		fmt.Fprintf(cmd.OutOrStdout(), "  client: %s\n", clientVersion)
	},
}

func getClientVersion() (string, string, error) {
	bi, _ := debug.ReadBuildInfo()
	var clientVersion, cliVersion string
	for _, dep := range bi.Deps {
		switch dep.Path {
		case "github.com/cobbler/cli":
			cliVersion = dep.Version
		case "github.com/cobbler/cobblerclient":
			clientVersion = dep.Version
		}
	}
	return clientVersion, cliVersion, nil
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
