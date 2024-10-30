// SPDX-License-Identifier: GPL-2.0-or-later
// SPDX-FileCopyrightText: 2021 Dominik Gedon <dgedon@suse.de>
// SPDX-FileCopyrightText: Copyright SUSE LLC

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
)

// signatureCmd represents the signature command
var signatureCmd = &cobra.Command{
	Use:   "signature",
	Short: "Signature management",
	Long:  `Reloads, reports or updates the signatures of the distinct operating system versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "Please use one of the sub commands!")
		_ = cmd.Help()
	},
}

var signatureReportCmd = &cobra.Command{
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
			fmt.Fprintln(cmd.OutOrStdout(), "Currently loaded signatures")
			breedNameList := make([]string, 0, len(signatures.Breeds))
			for key := range signatures.Breeds {
				breedNameList = append(breedNameList, key)
			}
			sort.Strings(breedNameList)
			for _, breedName := range breedNameList {
				fmt.Fprintln(cmd.OutOrStdout(), breedName)
				totalOsVersions += len(signatures.Breeds[breedName])
				if len(signatures.Breeds[breedName]) > 0 {
					osVersionNameList := make([]string, 0, len(signatures.Breeds[breedName]))
					for key := range signatures.Breeds[breedName] {
						osVersionNameList = append(osVersionNameList, key)
					}
					sort.Strings(osVersionNameList)
					for _, versionName := range osVersionNameList {
						fmt.Fprintf(cmd.OutOrStdout(), "\t%s\n", versionName)
					}
				} else {
					fmt.Fprintln(cmd.OutOrStdout(), "\t(none)")
				}

			}
			fmt.Fprintf(cmd.OutOrStdout(), "\n%d breeds with %d total OS versions loaded\n", len(signatures.Breeds), totalOsVersions)
		} else {
			fmt.Fprintln(cmd.OutOrStdout(), "No  breeds found in the signature, a signature update is recommended")
		}
		return nil
	},
}

var signatureUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the signatures JSON file",
	Long:  `Retrieve an up-to-date "distro_signatures.json" file from the server-side configured webservice.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}

		eventId, _ := Client.BackgroundSignatureUpdate()
		fmt.Fprintf(cmd.OutOrStdout(), "Event ID: %s\n", eventId)
		return nil
	},
}

var signatureReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reloads signatures",
	Long:  `Reloads signatures from the - on the server - local "distro_signatures.json" file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := generateCobblerClient()
		if err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), "This functionality cannot be used in the new CLI until https://github.com/cobbler/cobbler/issues/3791 is implemented!")
		return nil
	},
}

func init() {
	signatureCmd.AddCommand(signatureReloadCmd)
	signatureCmd.AddCommand(signatureReportCmd)
	signatureCmd.AddCommand(signatureUpdateCmd)
	rootCmd.AddCommand(signatureCmd)
}
