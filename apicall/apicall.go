package apicall

import (
	"fmt"
    "io"
    "log"
    "net/http"
    "os"
	viper "github.com/spf13/viper"
)

func getEnv() (apiKey string) {
	// Set config file
    viper.SetConfigFile("../.env")
    
	// Try to read config file, return any errors
	if err := viper.ReadInConfig(); err != nil {
        fmt.Println("Error reading config file:", err)
        return
    }

	// Read api_key from config.yaml file
    apiKey = viper.GetString("api_key")
    
	// Return error if key is empty
	if apiKey == "" {
        fmt.Println("API_KEY not found in config file")
        return
    }

	return apiKey
}

func call() string {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}
