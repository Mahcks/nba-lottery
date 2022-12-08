package structures

type TeamOdds struct {
	Name       string
	Percentage float32
}

type TeamJSON struct {
	LastUpdated string `json:"lastUpdated"`
	Teams       []Team `json:"teams"`
}

type Team struct {
	ID               string  `json:"id"`
	Slug             string  `json:"slug"`
	Abbreviation     string  `json:"abbreviation"`
	DisplayName      string  `json:"displayName"`
	ShortDisplayName string  `json:"shortDisplayName"`
	Name             string  `json:"name"`
	Location         string  `json:"location"`
	Wins             int     `json:"wins"`
	Losses           int     `json:"losses"`
	WinPercent       float64 `json:"winPercent"`
	FirstPickOdds    float64 `json:"odds"` // % of getting first pick
}
