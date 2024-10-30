// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2024 Enno Gotthold <egotthold@suse.com>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Show and query events for their status.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var eventStatusCmd = &cobra.Command{
	Use: "status",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}

		eventId, err := cmd.Flags().GetString("event-id")
		if err != nil {
			return err
		}

		event, err := Client.GetTaskStatus(eventId)
		if err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), event.State)
		return nil
	},
}

var eventListCmd = &cobra.Command{
	Use: "list",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}

		user, err := cmd.Flags().GetString("user")
		if err != nil {
			return err
		}

		events, err := Client.GetEvents(user)
		if err != nil {
			return err
		}
		idWidth := 0
		stateWidth := 10
		stateTimeWidth := 24 // Fixed width
		nameWidth := 0
		for _, event := range events {
			if len(event.ID) > idWidth {
				idWidth = len(event.ID)
			}
			if len(event.Name) > nameWidth {
				nameWidth = len(event.Name)
			}
			if len(event.State) > stateWidth {
				stateWidth = len(event.State)
			}
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%*s | %*s | %*s | %*s | %s \n", idWidth, "ID", nameWidth, "Name", stateWidth, "Task State", stateTimeWidth, "Time (last transitioned)", "Read by Who")
		for _, event := range events {
			stateTimeStruct, err := covertFloatToUtcTime(event.StateTime)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%*s | %*s | %*s | %*s | %s \n", idWidth, event.ID, nameWidth, event.Name, stateWidth, event.State, stateTimeWidth, stateTimeStruct.Format(time.DateTime), event.ReadByWho)
		}
		return nil
	},
}

var eventLogCmd = &cobra.Command{
	Use: "log",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}

		eventId, err := cmd.Flags().GetString("event-id")
		if err != nil {
			return err
		}

		eventLog, err := Client.GetEventLog(eventId)
		if err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), eventLog)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(eventCmd)

	eventCmd.AddCommand(eventStatusCmd)
	eventCmd.AddCommand(eventListCmd)
	eventCmd.AddCommand(eventLogCmd)

	// local flags for status
	eventStatusCmd.Flags().String("event-id", "", "the event ID of the background task")
	_ = eventStatusCmd.MarkFlagRequired("event-id")

	// local flags for list
	eventListCmd.Flags().String("user", "", "giving this parameter will show only events the user hasn't seen yet")

	// local flags for log
	eventLogCmd.Flags().String("event-id", "", "the event ID of the background task")
	_ = eventStatusCmd.MarkFlagRequired("event-id")
}
