// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2024 Enno Gotthold <egotthold@suse.com>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func NewEventCmd() *cobra.Command {
	eventCmd := &cobra.Command{
		Use:   "event",
		Short: "Show and query events for their status.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	eventCmd.AddCommand(NewEventStatusCmd())
	eventCmd.AddCommand(NewEventListCmd())
	eventCmd.AddCommand(NewEventLogCmd())
	return eventCmd
}

func NewEventStatusCmd() *cobra.Command {
	eventStatusCmd := &cobra.Command{
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
	eventStatusCmd.Flags().String("event-id", "", "the event ID of the background task")
	_ = eventStatusCmd.MarkFlagRequired("event-id")
	return eventStatusCmd
}

func NewEventListCmd() *cobra.Command {
	eventListCmd := &cobra.Command{
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
	eventListCmd.Flags().String("user", "", "giving this parameter will show only events the user hasn't seen yet")
	return eventListCmd
}

func NewEventLogCmd() *cobra.Command {
	eventLogCmd := &cobra.Command{
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
	eventLogCmd.Flags().String("event-id", "", "the event ID of the background task")
	_ = eventLogCmd.MarkFlagRequired("event-id")
	return eventLogCmd
}
