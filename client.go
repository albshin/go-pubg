package pubg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// API holds configuration variables for accessing the API.
type API struct {
	APIKey  string
	BaseURL *url.URL
}

// Player holds the information returned from the API.
type Player struct {
	PlatformID     int       `json:"platformId"`
	AccountID      string    `json:"AccountId"`
	Avatar         string    `json:"Avatar"`
	SelectedRegion string    `json:"selectedRegion"`
	DefaultSeason  string    `json:"defaultSeason"`
	SeasonDisplay  string    `json:"seasonDisplay"`
	LastUpdated    time.Time `json:"LastUpdated"`
	LiveTracking   []struct {
		Match        int         `json:"Match"`
		MatchDisplay string      `json:"MatchDisplay"`
		Season       int         `json:"Season"`
		RegionID     int         `json:"RegionId"`
		Region       string      `json:"Region"`
		Date         string      `json:"Date"`
		Delta        float64     `json:"Delta"`
		Value        float64     `json:"Value"`
		Message      interface{} `json:"message"`
	} `json:"LiveTracking"`
	PlayerName    string `json:"PlayerName"`
	PubgTrackerID int    `json:"PubgTrackerId"`
	Stats         []struct {
		Region string `json:"Region"`
		Season string `json:"Season"`
		Match  string `json:"Match"`
		Stats  []struct {
			Partition    interface{} `json:"partition"`
			Label        string      `json:"label"`
			SubLabel     interface{} `json:"subLabel"`
			Field        string      `json:"field"`
			Category     string      `json:"category"`
			ValueInt     interface{} `json:"ValueInt"`
			ValueDec     float64     `json:"ValueDec"`
			Value        string      `json:"value"`
			Rank         interface{} `json:"rank"`
			Percentile   float64     `json:"percentile"`
			DisplayValue string      `json:"displayValue"`
		} `json:"Stats"`
	} `json:"Stats"`
	MatchHistory []struct {
		Updated          string  `json:"Updated"`
		UpdatedJS        string  `json:"UpdatedJS"`
		Season           int     `json:"Season"`
		SeasonDisplay    string  `json:"SeasonDisplay"`
		Match            int     `json:"Match"`
		MatchDisplay     string  `json:"MatchDisplay"`
		Region           int     `json:"Region"`
		RegionDisplay    string  `json:"RegionDisplay"`
		Rounds           int     `json:"Rounds"`
		Wins             int     `json:"Wins"`
		Kills            int     `json:"Kills"`
		Assists          int     `json:"Assists"`
		Top10            int     `json:"Top10"`
		Rating           float64 `json:"Rating"`
		RatingChange     float64 `json:"RatingChange"`
		RatingRank       int     `json:"RatingRank"`
		RatingRankChange int     `json:"RatingRankChange"`
		Kd               float64 `json:"Kd"`
	} `json:"MatchHistory"`
}

// SteamInfo holds information returned from GetSteamInfo.
type SteamInfo struct {
	AccountID   string `json:"AccountId"`
	Nickname    string `json:"Nickname"`
	AvatarURL   string `json:"AvatarUrl"`
	SteamID     string `json:"SteamId"`
	SteamName   string `json:"SteamName"`
	State       string `json:"State"`
	InviteAllow string `json:"InviteAllow"`
}

// New creates a new API client.
func New(key string) (*API, error) {
	base, err := url.Parse("https://pubgtracker.com/api/")
	if err != nil {
		return &API{}, err
	}

	return &API{
		APIKey:  key,
		BaseURL: base,
	}, nil
}

// NewRequest creates the GET request to access the API.
func (a *API) NewRequest(endpoint string) (*http.Request, error) {
	end, err := url.Parse(endpoint)
	if err != nil {
		return &http.Request{}, err
	}
	urlStr := a.BaseURL.ResolveReference(end)

	req, err := http.NewRequest("GET", urlStr.String(), nil)
	if err != nil {
		// Handle error
		return req, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Add("trn-api-key", a.APIKey)

	return req, nil
}

// Do sends out a request to the API and unmarshals the data.
func (a *API) Do(req *http.Request, i interface{}) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &i)
}

// GetPlayer returns a player's stats.
func (a *API) GetPlayer(uname string) (*Player, error) {
	endpoint := "profile/pc/" + uname
	req, err := a.NewRequest(endpoint)

	if err != nil {
		return &Player{}, err
	}

	var player Player
	err = a.Do(req, &player)

	return &player, err
}

// GetSteamInfo retrieves a player's steam information.
func (a *API) GetSteamInfo(sid string) (*SteamInfo, error) {
	endpoint := "search?steamId=" + sid
	req, err := a.NewRequest(endpoint)

	if err != nil {
		return &SteamInfo{}, err
	}

	var sinfo SteamInfo
	err = a.Do(req, &sinfo)

	return &sinfo, err
}
