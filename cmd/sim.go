package cmd

import (
	"fmt"
	"sort"

	"github.com/Mahcks/nba-lottery/internal/nba"
	"github.com/Mahcks/nba-lottery/internal/structures"
	"github.com/spf13/cobra"
)

var times int

var simCmd = &cobra.Command{
	Use:   "sim",
	Short: "Simulate the draft lottery.",
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

		for _, k := range p {
			fmt.Printf("%v. %v | %v/%v\n", k.Value+1, k.Key, k.FirstPicks, times)
		}

		fmt.Println(json.Teams[16:])
	},
}

func init() {
	rootCmd.AddCommand(simCmd)

	simCmd.Flags().IntVarP(&times, "times", "t", 1, "Defines how many times you'd like to simulate the lottery.")
}
