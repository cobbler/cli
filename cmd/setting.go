// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// settingCmd represents the setting command
var settingCmd = &cobra.Command{
	Use:   "setting",
	Short: "Settings management",
	Long:  `Let you manage settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var settingEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit settings",
	Long:  `Edits the settings.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}

		settings, err := Client.GetSettings()
		if err != nil {
			return err
		}
		if !settings.AllowDynamicSettings {
			fmt.Fprintln(cmd.OutOrStdout(), "Dynamic settings are turned off server-side!")
			os.Exit(1)
		}

		settingName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		settingValue, err := cmd.Flags().GetString("value")
		if err != nil {
			return err
		}
		result, err := Client.ModifySetting(settingName, settingValue)
		if err != nil {
			return err
		}
		if result == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "Successfully updated!")
		} else {
			fmt.Fprintln(cmd.OutOrStdout(), "Updating settings failed!")
		}
		return nil
	},
}

var settingReportCmd = &cobra.Command{
	Use:   "report",
	Short: "list settings",
	Long:  `Prints settings to stdout.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}

		settings, err := Client.GetSettings()
		if err != nil {
			return err
		}

		printStructured(cmd, settings)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(settingCmd)
	settingCmd.AddCommand(settingEditCmd)
	settingCmd.AddCommand(settingReportCmd)

	// local flags for setting edit
	settingEditCmd.Flags().String("name", "", "the settings name to edit (e.g. server)")
	settingEditCmd.Flags().String("value", "", "the new value (e.g. 127.0.0.1)")

	// local flags for setting report
	settingReportCmd.Flags().String("name", "", "the settings name to show")
}
