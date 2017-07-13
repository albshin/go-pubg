package pubg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type API struct {
	APIKey  string
	BaseURL *url.URL
}

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

type SteamInfo struct {
	AccountID   string `json:"AccountId"`
	Nickname    string `json:"Nickname"`
	AvatarURL   string `json:"AvatarUrl"`
	SteamID     string `json:"SteamId"`
	SteamName   string `json:"SteamName"`
	State       string `json:"State"`
	InviteAllow string `json:"InviteAllow"`
}

func CreateAPI(key string) *API {
	base, err := url.Parse("https://pubgtracker.com/api/")
	if err != nil {
		panic(err)
	}

	return &API{
		APIKey:  key,
		BaseURL: base,
	}
}

func (a *API) NewRequest(endpoint string) *http.Request {
	end, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	urlStr := a.BaseURL.ResolveReference(end)

	req, err := http.NewRequest("GET", urlStr.String(), nil)
	if err != nil {
		// Handle error
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Add("trn-api-key", a.APIKey)

	return req
}

func (a *API) Do(req *http.Request, i interface{}) {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *API) GetPlayer(uname string) *Player {
	endpoint := "profile/pc/" + uname
	req := a.NewRequest(endpoint)

	var player Player
	a.Do(req, &player)

	return &player
}

func (a *API) GetSteamInfo(sid string) *SteamInfo {
	endpoint := "search?steamId=" + sid
	req := a.NewRequest(endpoint)

	var sinfo SteamInfo
	a.Do(req, &sinfo)

	return &sinfo
}
