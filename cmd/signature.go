// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
)

// NewSignatureCmd builds a new command that represents the signature action
func NewSignatureCmd() *cobra.Command {
	signatureCmd := &cobra.Command{
		Use:   "signature",
		Short: "Signature management",
		Long:  `Reloads, reports or updates the signatures of the distinct operating system versions.`,
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Please use one of the sub commands!")
			_ = cmd.Help()
		},
	}
	signatureCmd.AddCommand(NewSignatureReloadCmd())
	signatureCmd.AddCommand(NewSignatureReportCmd())
	signatureCmd.AddCommand(NewSignatureUpdateCmd())
	return signatureCmd
}

func NewSignatureReportCmd() *cobra.Command {
	signatureReportCmd := &cobra.Command{
		Use:   "report",
		Short: "Report the loaded signatures",
		Long:  `Report the loaded signatures`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			// Get signatures
			signatures, err := Client.GetSignatures()
			if err != nil {
				return err
			}

			if len(signatures.Breeds) > 0 {
				// Counters
				var totalOsVersions int

				// Print signatures
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Currently loaded signatures")
				breedNameList := make([]string, 0, len(signatures.Breeds))
				for key := range signatures.Breeds {
					breedNameList = append(breedNameList, key)
				}
				sort.Strings(breedNameList)
				for _, breedName := range breedNameList {
					_, _ = fmt.Fprintln(cmd.OutOrStdout(), breedName)
					totalOsVersions += len(signatures.Breeds[breedName])
					if len(signatures.Breeds[breedName]) > 0 {
						osVersionNameList := make([]string, 0, len(signatures.Breeds[breedName]))
						for key := range signatures.Breeds[breedName] {
							osVersionNameList = append(osVersionNameList, key)
						}
						sort.Strings(osVersionNameList)
						for _, versionName := range osVersionNameList {
							_, _ = fmt.Fprintf(cmd.OutOrStdout(), "\t%s\n", versionName)
						}
					} else {
						_, _ = fmt.Fprintln(cmd.OutOrStdout(), "\t(none)")
					}

				}
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "\n%d breeds with %d total OS versions loaded\n", len(signatures.Breeds), totalOsVersions)
			} else {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "No  breeds found in the signature, a signature update is recommended")
			}
			return nil
		},
	}
	return signatureReportCmd
}

func NewSignatureUpdateCmd() *cobra.Command {
	signatureUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update the signatures JSON file",
		Long:  `Retrieve an up-to-date "distro_signatures.json" file from the server-side configured webservice.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			eventId, _ := Client.BackgroundSignatureUpdate()
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Event ID: %s\n", eventId)
			return nil
		},
	}
	return signatureUpdateCmd
}

func NewSignatureReloadCmd() *cobra.Command {
	signatureReloadCmd := &cobra.Command{
		Use:   "reload",
		Short: "Reloads signatures",
		Long:  `Reloads signatures from the - on the server - local "distro_signatures.json" file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			eventId, _ := Client.BackgroundSignatureReload()
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Event ID: %s\n", eventId)
			return nil
		},
	}
	return signatureReloadCmd
}
