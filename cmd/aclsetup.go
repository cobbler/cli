// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// aclsetupCmd represents the aclsetup command
var aclsetupCmd = &cobra.Command{
	Use:   "aclsetup",
	Short: "Adjust the access control list",
	Long:  "Configures users/groups to run the Cobbler CLI as non-root.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		generateCobblerClient()
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(aclsetupCmd)

	//local flags
	aclsetupCmd.Flags().String("adduser", "", "give acls to this user")
	aclsetupCmd.Flags().String("addgroup", "", "give acls to this group")
	aclsetupCmd.Flags().String("removeuser", "", "remove acls from this user")
	aclsetupCmd.Flags().String("removegroup", "", "remove acls from this user")
}
