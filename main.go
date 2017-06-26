package pubg

import (
	"net/http"
	"net/url"
)

type API struct {
	APIKey  string
	BaseURL *url.URL
}

type Player struct {
	PlatformID    int      `json:"platformId"`
	AccountID     string   `json:"accountId"`
	Avatar        string   `json:"avatar"`
	Region        string   `json:"region"`
	DefaultSeason string   `json:"defaultSeason"`
	SeasonDisplay string   `json:"seasonDisplay"`
	LastUpdated   string   `json:"LastUpdated"`
	LiveTracking  []Match  `json:"LiveTracking"`
	PlayerName    string   `json:"PlayerName"`
	PUBGTrackerID int      `json:"PubgTrackerId"`
	Stats         []string `json:"Stats"`
}

type Match struct {
	Match        int     `json:"Match"`
	MatchDisplay string  `json:"MatchDisplay"`
	Season       int     `json:"Season"`
	RegionID     int     `json:"RegionId"`
	Region       string  `json:"Region"`
	Date         string  `json:"Date"`
	Delta        float32 `json:"Delta"`
	Value        float32 `json:"Value"`
	Message      string  `json:"message"`
}

type PlayerStats struct {
	Region string       `json:"Region"`
	Season string       `json:"Season"`
	Match  string       `json:"Match"`
	Stats  []MatchStats `json:"Stats"`
}

// Some of these values may return null
type MatchStats struct {
	Partition    string  `json:"partition"`
	Label        string  `json:"label"`
	SubLabel     string  `json:"subLabel"`
	Field        string  `json:"field"`
	Category     string  `json:"category"`
	ValueInt     int     `json:"ValueInt"`
	ValueDec     float32 `json:"ValueDec"`
	Value        string  `json:"value"`
	Rank         int     `json:"rank"`
	Percentile   float32 `json:"percentile"`
	DisplayValue float32 `json:"displayValue"`
}

func Create(key string) *API {
	base, err := url.Parse("https://pubgtracker.com/api/profile/pc/")
	if err != nil {
		panic(err)
	}

	return &API{
		APIKey:  key,
		BaseURL: base,
	}
}

func (a *API) Request(endpoint string) *http.Request {
	end, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	urlStr := a.BaseURL.ResolveReference(end)

	req, err := http.NewRequest("GET", urlStr.String(), nil)
	if err != nil {
		// Handle error
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("trn-api-key", a.APIKey)

	return req
}
