package nba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Mahcks/nba-lottery/internal/structures"
)

// If standings file doesn't exist, create it otherwise read the local data and return it
func GetStandings() (*structures.TeamJSON, error) {
	if _, err := os.Stat("standings.json"); err != nil {
		standings, err := GetStandingsFromESPN()
		if err != nil {
			return nil, err
		}

		return standings, nil
	}

	json, err := GetLocalStandings()
	if err != nil {
		return nil, err
	}

	return json, nil
}

// Grabs standings from local file
func GetLocalStandings() (*structures.TeamJSON, error) {
	file, _ := ioutil.ReadFile("standings.json")

	var standings structures.TeamJSON
	json.Unmarshal(file, &standings)

	return &standings, nil
}

// GetStandingsFromESPN - Grab the latest standings from ESPN which will be saved locally
func GetStandingsFromESPN() (*structures.TeamJSON, error) {
	var jsonObj structures.TeamJSON = structures.TeamJSON{
		LastUpdated: time.Now().Format(time.RFC3339),
	}

	for _, team := range Teams {
		url := fmt.Sprintf("http://site.api.espn.com/apis/site/v2/sports/basketball/nba/teams/%v", team.ID)

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var testObj structures.ESPNTeamStanding
		json.Unmarshal(resBody, &testObj)

		winLoss := strings.Split(testObj.Team.Record.Items[0].Summary, "-")
		wins, err := strconv.Atoi(winLoss[0])
		if err != nil {
			return nil, err
		}

		losses, err := strconv.Atoi(winLoss[1])
		if err != nil {
			return nil, err
		}

		var toStore = structures.Team{
			ID:               testObj.Team.ID,
			Abbreviation:     testObj.Team.Abbreviation,
			Slug:             testObj.Team.Slug,
			DisplayName:      testObj.Team.DisplayName,
			ShortDisplayName: testObj.Team.ShortDisplayName,
			Name:             testObj.Team.Name,
			Location:         testObj.Team.Location,
			Wins:             wins,
			Losses:           losses,
			WinPercent:       structures.ToFixed(testObj.Team.Record.Items[0].Stats[16].Value, 3),
		}
		// roundFloat(testObj.Team.Record.Items[0].Stats[16].Value, 3)

		jsonObj.Teams = append(jsonObj.Teams, toStore)
	}

	// Sort teams by win percentage
	sort.Slice(jsonObj.Teams, func(i, j int) bool {
		return jsonObj.Teams[i].WinPercent < jsonObj.Teams[j].WinPercent
	})

	// Loop through each team and get the odds for a first pick
	for i := 0; i < len(jsonObj.Teams); i++ {
		//fmt.Printf("%v got a %v chance to get first", jsonObj.Teams[i].Name, getPercentage(i))
		jsonObj.Teams[i].FirstPickOdds = getPercentage(i + 1)
	}

	// Prepare and write to standings file
	file, _ := json.MarshalIndent(jsonObj, "", " ")
	ioutil.WriteFile("standings.json", file, 0644)

	return &jsonObj, nil
}

// getPercentage is a helper function that determines the percentage based on their pick
func getPercentage(pick int) float64 {
	switch pick {
	case 1:
		return float64(14.0)
	case 2:
		return float64(14.0)
	case 3:
		return float64(14.0)
	case 4:
		return float64(11.5)
	case 5:
		return float64(11.5)
	case 6:
		return float64(9.0)
	case 7:
		return float64(7.5)
	case 8:
		return float64(6.0)
	case 9:
		return float64(3.2)
	case 10:
		return float64(3.2)
	case 11:
		return float64(3.1)
	case 12:
		return float64(1.5)
	case 13:
		return float64(1.0)
	case 14:
		return float64(0.5)
	default:
		return float64(0)
	}
}
