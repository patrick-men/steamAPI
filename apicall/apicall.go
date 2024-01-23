package apicall

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	viper "github.com/spf13/viper"
)

func getEnv() (apiKey, domain string) {
	// Set config file
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	// Try to read config file, output = Any errors
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Read api_key from config.yaml file
	apiKey = viper.GetString("key")
	domain = viper.GetString("domainName")

	// Return error if key is empty
	if apiKey == "" {
		fmt.Println("key not found in .env file")
		return
	}

	if domain == "" {
		fmt.Println("domainName not found in .env file")
		return
	}

	return apiKey, domain
}

// function thta contains all api calls. requires 4 parameters (if you want to call an api that only uses three, leave the fourth param empty):
// url: options are news, achievementPercentage, summary, friends, achievementsGotten, userStatsPerGame, ownedGames, recentlyPlayed
// apiKey: api key defined in .env; should be made available by call() func
// appID: some api calls require the appID of the game you're looking at
// userID: some api calls require the userID of the user you're looking at
func apiUrl(url, apiKey, appID, userID string) (output string) {

	APIappNews := "http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002/?appid=" + appID + "&count=3&maxlength=300&format=json"
	APIachievementPercentage := "http://api.steampowered.com/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002/?gameid=" + appID + "&format=json"
	APIplayerSummaries := "http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=" + apiKey + "&steamids=" + userID + "&format=json"
	APIfriendsList := "http://api.steampowered.com/ISteamUser/GetFriendList/v0001/?key=" + apiKey + "&steamid=" + userID + "&relationship=friend&format=json"
	APIplayerAchievements := "http://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v0001/?appid=" + appID + "&key=" + apiKey + "&steamid=" + userID + "&format=json"
	APIuserStatsForGame := "http://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?appid=" + appID + "&key=" + apiKey + "&steamid=" + userID + "&format=json"
	APIownedGames := "http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=" + apiKey + "&steamid=" + userID + "&format=json"
	APIrecentlyPlayed := "http://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v0001/?key=" + apiKey + "&steamid=" + userID + "&format=json"

	switch url {
	case "news":
		output = APIappNews
	case "achievementPercentage":
		output = APIachievementPercentage
	case "summary":
		output = APIplayerSummaries
	case "friends":
		output = APIfriendsList
	case "achievementsGotten":
		output = APIplayerAchievements
	case "userStatsPerGame":
		output = APIuserStatsForGame
	case "ownedGames":
		output = APIownedGames
	case "recentlyPlayed":
		output = APIrecentlyPlayed
	}

    return
}

func call(url string) string {

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	return string(responseData)
}


// logic of the api call
func ApiCall() (output string) {


// example call: owned games of random user found in dishonored reviews
    // key, _ := getEnv()

    // output = call(apiUrl("ownedGames", key, "", "76561199227168782"))

    // return
}
