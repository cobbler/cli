// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// NewSettingCmd builds a new command that represents the setting action
func NewSettingCmd() *cobra.Command {
	settingCmd := &cobra.Command{
		Use:   "setting",
		Short: "Settings management",
		Long:  `Let you manage settings.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	settingCmd.AddCommand(NewSettingEditCmd())
	settingCmd.AddCommand(NewSettingReportCmd())
	return settingCmd
}

func NewSettingEditCmd() *cobra.Command {
	settingEditCmd := &cobra.Command{
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
				return errors.New("dynamic settings are turned off server-side")
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
	settingEditCmd.Flags().String("name", "", "the settings name to edit (e.g. server)")
	settingEditCmd.Flags().String("value", "", "the new value (e.g. 127.0.0.1)")
	return settingEditCmd
}

func NewSettingReportCmd() *cobra.Command {
	settingReportCmd := &cobra.Command{
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
	settingReportCmd.Flags().String("name", "", "the settings name to show")
	return settingReportCmd
}
