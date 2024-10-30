// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

// aclsetupCmd represents the aclsetup command
var aclsetupCmd = &cobra.Command{
	Use:   "aclsetup",
	Short: "Adjust the access control list",
	Long:  "Configures users/groups to run the Cobbler CLI as non-root.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}
		addUserOption, err := cmd.Flags().GetString("adduser")
		if err != nil {
			return err
		}
		addGroupOption, err := cmd.Flags().GetString("addgroup")
		if err != nil {
			return err
		}
		removeUserOption, err := cmd.Flags().GetString("removeuser")
		if err != nil {
			return err
		}
		removeGroupOption, err := cmd.Flags().GetString("removegroup")
		if err != nil {
			return err
		}
		aclSetupOptions := cobblerclient.AclSetupOptions{
			AddUser:     addUserOption,
			AddGroup:    addGroupOption,
			RemoveUser:  removeUserOption,
			RemoveGroup: removeGroupOption,
		}
		eventId, err := Client.BackgroundAclSetup(aclSetupOptions)
		if err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), "Event ID: ", eventId)
		return nil
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
