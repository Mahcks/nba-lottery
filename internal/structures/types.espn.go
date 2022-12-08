package structures

type ESPNTeam struct {
	ID               string `json:"id"`
	UID              string `json:"uid"`
	Slug             string `json:"slug"`
	Abbreviation     string `json:"abbreviation"`
	DisplayName      string `json:"displayName"`
	ShortDisplayName string `json:"shortDisplayName"`
	Name             string `json:"name"`
	Nickname         string `json:"nickname"`
	Location         string `json:"location"`
	Color            string `json:"color"`
	AlternateColor   string `json:"alternateColor"`
	IsActive         bool   `json:"isActive"`
	IsAllStar        bool   `json:"isAllStar"`
}

type ESPNTeamStanding struct {
	Team struct {
		ID               string `json:"id"`
		UID              string `json:"uid"`
		Slug             string `json:"slug"`
		Location         string `json:"location"`
		Name             string `json:"name"`
		Abbreviation     string `json:"abbreviation"`
		DisplayName      string `json:"displayName"`
		ShortDisplayName string `json:"shortDisplayName"`
		Color            string `json:"color"`
		AlternateColor   string `json:"alternateColor"`
		IsActive         bool   `json:"isActive"`
		Logos            []struct {
			Href        string   `json:"href"`
			Width       int      `json:"width"`
			Height      int      `json:"height"`
			Alt         string   `json:"alt"`
			Rel         []string `json:"rel"`
			LastUpdated string   `json:"lastUpdated"`
		} `json:"logos"`
		Record struct {
			Items []struct {
				Description string `json:"description"`
				Type        string `json:"type"`
				Summary     string `json:"summary"`
				Stats       []struct {
					Name  string  `json:"name"`
					Value float64 `json:"value"`
				} `json:"stats"`
			} `json:"items"`
		} `json:"record"`
		Groups struct {
			ID     string `json:"id"`
			Parent struct {
				ID string `json:"id"`
			} `json:"parent"`
			IsConference bool `json:"isConference"`
		} `json:"groups"`
		Links []struct {
			Language   string   `json:"language"`
			Rel        []string `json:"rel"`
			Href       string   `json:"href"`
			Text       string   `json:"text"`
			ShortText  string   `json:"shortText"`
			IsExternal bool     `json:"isExternal"`
			IsPremium  bool     `json:"isPremium"`
		} `json:"links"`
		Franchise struct {
			Ref              string `json:"$ref"`
			ID               string `json:"id"`
			UID              string `json:"uid"`
			Slug             string `json:"slug"`
			Location         string `json:"location"`
			Name             string `json:"name"`
			Abbreviation     string `json:"abbreviation"`
			DisplayName      string `json:"displayName"`
			ShortDisplayName string `json:"shortDisplayName"`
			Color            string `json:"color"`
			IsActive         bool   `json:"isActive"`
			Venue            struct {
				Ref       string `json:"$ref"`
				ID        string `json:"id"`
				FullName  string `json:"fullName"`
				ShortName string `json:"shortName"`
				Address   struct {
					City  string `json:"city"`
					State string `json:"state"`
				} `json:"address"`
				Capacity int  `json:"capacity"`
				Grass    bool `json:"grass"`
				Indoor   bool `json:"indoor"`
				Images   []struct {
					Href   string   `json:"href"`
					Width  int      `json:"width"`
					Height int      `json:"height"`
					Alt    string   `json:"alt"`
					Rel    []string `json:"rel"`
				} `json:"images"`
			} `json:"venue"`
			Team struct {
				Ref string `json:"$ref"`
			} `json:"team"`
		} `json:"franchise"`
		NextEvent []struct {
			ID        string `json:"id"`
			Date      string `json:"date"`
			Name      string `json:"name"`
			ShortName string `json:"shortName"`
			Season    struct {
				Year        int    `json:"year"`
				DisplayName string `json:"displayName"`
			} `json:"season"`
			SeasonType struct {
				ID           string `json:"id"`
				Type         int    `json:"type"`
				Name         string `json:"name"`
				Abbreviation string `json:"abbreviation"`
			} `json:"seasonType"`
			TimeValid    bool `json:"timeValid"`
			Competitions []struct {
				ID         string `json:"id"`
				Date       string `json:"date"`
				Attendance int    `json:"attendance"`
				Type       struct {
					ID           string `json:"id"`
					Text         string `json:"text"`
					Abbreviation string `json:"abbreviation"`
					Slug         string `json:"slug"`
					Type         string `json:"type"`
				} `json:"type"`
				TimeValid         bool `json:"timeValid"`
				NeutralSite       bool `json:"neutralSite"`
				BoxscoreAvailable bool `json:"boxscoreAvailable"`
				TicketsAvailable  bool `json:"ticketsAvailable"`
				Venue             struct {
					FullName string `json:"fullName"`
					Address  struct {
						City  string `json:"city"`
						State string `json:"state"`
					} `json:"address"`
				} `json:"venue"`
				Competitors []struct {
					ID       string `json:"id"`
					Type     string `json:"type"`
					Order    int    `json:"order"`
					HomeAway string `json:"homeAway"`
					Team     struct {
						ID               string `json:"id"`
						Location         string `json:"location"`
						Abbreviation     string `json:"abbreviation"`
						DisplayName      string `json:"displayName"`
						ShortDisplayName string `json:"shortDisplayName"`
						Logos            []struct {
							Href        string   `json:"href"`
							Width       int      `json:"width"`
							Height      int      `json:"height"`
							Alt         string   `json:"alt"`
							Rel         []string `json:"rel"`
							LastUpdated string   `json:"lastUpdated"`
						} `json:"logos"`
						Links []struct {
							Rel  []string `json:"rel"`
							Href string   `json:"href"`
							Text string   `json:"text"`
						} `json:"links"`
					} `json:"team"`
				} `json:"competitors"`
				Notes      []interface{} `json:"notes"`
				Broadcasts []interface{} `json:"broadcasts"`
				Tickets    []struct {
					ID              string  `json:"id"`
					Summary         string  `json:"summary"`
					Description     string  `json:"description"`
					MaxPrice        float64 `json:"maxPrice"`
					StartingPrice   float64 `json:"startingPrice"`
					NumberAvailable int     `json:"numberAvailable"`
					TotalPostings   int     `json:"totalPostings"`
					Links           []struct {
						Rel  []string `json:"rel"`
						Href string   `json:"href"`
					} `json:"links"`
				} `json:"tickets"`
				Status struct {
					Clock        float64 `json:"clock"`
					DisplayClock string  `json:"displayClock"`
					Period       int     `json:"period"`
					Type         struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						State       string `json:"state"`
						Completed   bool   `json:"completed"`
						Description string `json:"description"`
						Detail      string `json:"detail"`
						ShortDetail string `json:"shortDetail"`
					} `json:"type"`
				} `json:"status"`
			} `json:"competitions"`
			Links []struct {
				Language   string   `json:"language"`
				Rel        []string `json:"rel"`
				Href       string   `json:"href"`
				Text       string   `json:"text"`
				ShortText  string   `json:"shortText"`
				IsExternal bool     `json:"isExternal"`
				IsPremium  bool     `json:"isPremium"`
			} `json:"links"`
		} `json:"nextEvent"`
		StandingSummary string `json:"standingSummary"`
	} `json:"team"`
}
