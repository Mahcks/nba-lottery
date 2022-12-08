package cmd

import (
	"fmt"
	"sort"

	"github.com/Mahcks/nba-lottery/internal/nba"
	"github.com/Mahcks/nba-lottery/internal/structures"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var times int

var simCmd = &cobra.Command{
	Use:     "simulate",
	Aliases: []string{"sim"},
	Short:   "Simulate the draft lottery.",
	Run: func(cmd *cobra.Command, args []string) {
		json, err := nba.GetStandings(false)
		if err != nil {
			fmt.Printf("ERROR HERE: %v", err.Error())
			return
		}

		ts := json.Teams[0:14]

		// Create a map to track the scores of each team
		scores := make(map[string]int)
		for _, team := range ts {
			scores[team.Slug] = 0
		}

		// Repeat the benchmark x amount of times
		for i := 0; i < times; i++ {
			firstPick := nba.GenerateLottery()
			scores[*firstPick] = scores[*firstPick] + 1
		}

		p := make(structures.PairList, len(ts))

		i := 0
		for k, v := range ts {
			p[i] = structures.Pair{Key: v.Name, Value: k, FirstPicks: scores[v.Slug]}
			i++
		}

		sort.Sort(p)

		tbl := table.New("PICK", "TEAM", "TIMES #1")
		for i, k := range p {
			pick := i + 1
			lotPick := k.Value + 1

			if pick == lotPick {
				tbl.AddRow(pick, k.Key, k.FirstPicks)
				continue
			}

			if pick <= lotPick {
				tbl.AddRow(fmt.Sprintf("%v +%v", pick, lotPick-pick), k.Key, k.FirstPicks)
			} else {
				tbl.AddRow(fmt.Sprintf("%v %v", pick, lotPick-pick), k.Key, k.FirstPicks)
			}
		}

		if times == 1 {
			fmt.Println("============ LOTTERY SIMULATION | RAN 1 TIME ============")
		} else {
			fmt.Printf("============ LOTTERY SIMULATION | RAN %v TIMES ============\n", times)
		}
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(simCmd)

	simCmd.Flags().IntVarP(&times, "times", "t", 1, "Defines how many times you'd like to simulate the lottery.")
}
