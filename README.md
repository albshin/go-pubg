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
client := pubg.New("YOURAPIKEYHERE")
info := client.GetPlayer("JohnDoe") // Returns JSON unfiltered for player "JohnDoe"
fmt.Printf("%+v\n", info)

steaminfo := client.GetSteamInfo("12345678901234567") // Returns Steam information for user based on SteamId
fmt.Printf("%+v\n", steaminfo)
```

## To-Do
* Documentation
* Tests
* Filtered results
