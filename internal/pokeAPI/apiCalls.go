// /internal/pokeAPI/apiCalls.go
package pokeAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"pokedexcli/internal/pokecache"
	"time"
)
var baseURL string = "https://pokeapi.co/api/v2/"
var mapOffset int = 0
var mapLimit int = 20

var cache = pokecache.NewCache(5 * time.Second)

func mapCaller(url string) ([]Results, error) {
	if data, found := cache.Get(url); found {
		var mapCallResults mapCall
		err := json.Unmarshal(data, &mapCallResults)
		if err != nil {
			fmt.Println("Error unmarshaling cached data:", err)
			return []Results{}, err
		}
		return mapCallResults.Results, nil
	}


	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return []Results{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []Results{}, err
	}

	cache.Add(url, body)

	var mapCallResults mapCall
	err = json.Unmarshal(body, &mapCallResults)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return []Results{}, err
	}

	return mapCallResults.Results, nil
}

func CommandMap() (error) {
	url := fmt.Sprintf("%slocation-area/?offset=%d&limit=%d", baseURL, mapOffset, mapLimit)

	results, err := mapCaller(url)
	if err != nil {
		return err
	}
	for _, result := range results {
		fmt.Printf("- %s\n", result.Name)
	}
	mapOffset += len(results)

	return nil
}

func CommandMapb() (error) {
	if mapOffset >= mapLimit * 2 {
		mapOffset -= mapLimit * 2
	} else if mapOffset - mapLimit >= 0 {
		mapOffset -= mapLimit
	} else {
		mapOffset = 0
	}
	url := fmt.Sprintf("%slocation-area/?offset=%d&limit=%d", baseURL, mapOffset, mapLimit)

	results, err := mapCaller(url)
	if err != nil {
		return err
	}
	for _, result := range results {
		fmt.Printf("- %s\n", result.Name)
	}

	mapOffset += mapLimit

	return nil
}