package cmd

import (
	"github.com/Mahcks/nba-lottery/internal/nba"
	"github.com/spf13/cobra"
)

var updateStandingsCmd = &cobra.Command{
	Use:   "update-standings",
	Short: "Update the local files using real data from ESPN.",
	Run: func(cmd *cobra.Command, args []string) {
		nba.GetStandingsFromESPN()
	},
}

func init() {
	rootCmd.AddCommand(updateStandingsCmd)
}
