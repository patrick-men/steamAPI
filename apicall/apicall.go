package apicall

import (
	"fmt"
    "io"
    "log"
    "net/http"
    "os"
	viper "github.com/spf13/viper"
)

func GetEnv() (apiKey, domain string) {
	// Set config file
    viper.SetConfigFile(".env")
	viper.SetConfigType("env")
    
	// Try to read config file, return any errors
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

	return "haha"
}
