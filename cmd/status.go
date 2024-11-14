package cmd

import (
	"fmt"
	cobbler "github.com/cobbler/cobblerclient"
	"github.com/spf13/cobra"
)

func NewStatusCmd() *cobra.Command {
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "View installation status of Cobbler Profiles and Systems.",
		Long: `This command displays the current status of all Cobbler Profiles and Systems. All installations that
run for longer then 100 minutes are considered stalled.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := generateCobblerClient()
			if err != nil {
				return err
			}

			res, err := Client.GetStatus(cobbler.StatusText)
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), res.(string))

			return nil
		},
	}
	return statusCmd
}
