package cmd

import (
	"fmt"

	"github.com/Mahcks/nba-lottery/internal/nba"
	"github.com/Mahcks/nba-lottery/internal/structures"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var shouldUpgrade bool

var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "View the latest NBA standings.",
	Run: func(cmd *cobra.Command, args []string) {
		json, err := nba.GetStandings(shouldUpgrade)
		if err != nil {
			fmt.Printf("ERROR HERE: %v", err.Error())
			return
		}

		tbl := table.New("PICK", "TEAM", "RECORD", "WIN%", "#1 OVR")

		for i, team := range json.Teams[0:14] {
			// TODO: If two or more teams have the same WIN% than group the rest of the teams under the pick they're fighting for
			tbl.AddRow(i+1, team.DisplayName, fmt.Sprintf("%v-%v", team.Wins, team.Losses), fmt.Sprintf("%.3f", team.WinPercent), fmt.Sprintf("%v %%", structures.ToFixed(team.FirstPickOdds, 3)))
		}
		tbl.Print()

		fmt.Println("================ END OF LOTTERY ================")

		tbl2 := table.New("PICK", "TEAM", "RECORD", "WIN%")
		tbl2.WithHeaderFormatter(func(s string, i ...interface{}) string {
			return ""
		})
		restOfLeague := json.Teams[15:]
		for i, team := range restOfLeague {
			tbl2.AddRow(i+1+15, team.DisplayName, fmt.Sprintf("%v-%v", team.Wins, team.Losses), fmt.Sprintf("%.3f", team.WinPercent))
		}
		tbl2.Print()
	},
}

func init() {
	standingsCmd.Flags().BoolVarP(&shouldUpgrade, "reload", "r", false, "Reload the standings to get more up-to-date information.")

	rootCmd.AddCommand(standingsCmd)
}
