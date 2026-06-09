// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime/debug"
)

// NewVersionCmd builds a new command that represents the version action
func NewVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the Cobbler version",
		Long:  `Shows the Cobbler server version.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}
			version, err := Client.ExtendedVersion()
			if err != nil {
				return err
			}
			clientVersion, cliVersion, _ := getClientVersion()
			fmt.Fprintf(cmd.OutOrStdout(), "Cobbler %s\n", version.Version)
			fmt.Fprintf(cmd.OutOrStdout(), "  source: %s, %s\n", version.Gitstamp, version.Gitdate)
			fmt.Fprintf(cmd.OutOrStdout(), "  build time: %s\n", version.Builddate)
			fmt.Fprintf(cmd.OutOrStdout(), "  cli: %s\n", cliVersion)
			fmt.Fprintf(cmd.OutOrStdout(), "  client: %s\n", clientVersion)
			return nil
		},
	}
	return versionCmd
}

func getClientVersion() (string, string, error) {
	bi, _ := debug.ReadBuildInfo()
	// The CLI is the main module, which never shows up in bi.Deps (that's only for its
	// dependencies) - its own version comes from bi.Main instead.
	cliVersion := bi.Main.Version
	var clientVersion string
	for _, dep := range bi.Deps {
		if dep.Path == "github.com/cobbler/cobblerclient" {
			clientVersion = dep.Version
		}
	}
	return clientVersion, cliVersion, nil
}
