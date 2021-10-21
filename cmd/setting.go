// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"github.com/spf13/cobra"
)

// settingCmd represents the setting command
var settingCmd = &cobra.Command{
	Use:   "setting",
	Short: "settings management",
	Long:  `Let you manage settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var settingEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit settings",
	Long:  `Edits the settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

var settingReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list settings",
	Long:  `Prints settings to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: call cobblerclient
	},
}

func init() {
	rootCmd.AddCommand(settingCmd)
	settingCmd.AddCommand(settingEditCmd)
	settingCmd.AddCommand(settingReportCmd)

	// local flags for setting edit
	settingEditCmd.Flags().String("name", "", "the settings name to edit (e.g. server)")
	settingEditCmd.Flags().String("value", "", "the new value (e.g. 127.0.0.1)")

	// local flags for setting edit
	settingReportCmd.Flags().String("name", "", "the settings name to show")
}
