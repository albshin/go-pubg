# go-pubg

API wrapper for Playerunknown's Battlegrounds using the API from http://pubgtracker.com.

## Getting Started

### 1. Install the package

Install the package and import it

```
go get github.com/albshin/go-pubg
```

### 2. Obtain API Key

Create an account at https://pubgtracker.com/site-api and request an API Key.

## Usage

```
client, err := pubg.New("YOURAPIKEYHERE")
if err != nil {
    log.Fatal(err)
}
info, err := client.GetPlayer("JohnDoe") // Returns JSON unfiltered for player "JohnDoe"
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", info)

steaminfo, err := client.GetSteamInfo("76561198396852397") // Returns Steam information for user based on SteamId
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", steaminfo)
```

## To-Do
* Documentation
* Tests
* Filtered results
