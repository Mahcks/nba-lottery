package nba

import (
	"fmt"
	"math/rand"

	"github.com/Mahcks/nba-lottery/internal/structures"
)

func GenerateLottery() *string {
	rand.Seed(structures.RandInt())

	json, err := GetStandings()
	if err != nil {
		fmt.Printf("ERROR HERE: %v", err.Error())
	}

	teams := json.Teams[0:14]

	// Create a map to track the scores of each team
	scores := make(map[string]int)

	// Calculate the total percentage of all teams
	total := float64(0)
	for _, team := range teams {
		scores[team.Slug] = 0
		total += team.FirstPickOdds
	}

	// Repeat the picking process until all teams have been chosen
	var picks []structures.Team = []structures.Team{}
	var pick = 0
	var firstPick *string = nil
	for len(teams) > 0 {
		pick++

		// Generate a random number between 0 and the total percentage
		n := rand.Float64() * total

		// Iterate through the teams and pick the one with the corresponding percentage
		for i, team := range teams {
			if n < team.FirstPickOdds {
				if pick == 1 {
					scores[team.Slug] = scores[team.Slug] + 1
					firstPick = &team.Slug
				}

				// The team is picked, so print its name and remove it from the list
				// fmt.Printf("%v | %v - %v\n", pick, team.Name, team.Percentage)
				picks = append(picks, team)
				teams = append(teams[:i], teams[i+1:]...)
				total -= team.FirstPickOdds
				break
			}
			n -= team.FirstPickOdds
		}
	}

	return firstPick
}
