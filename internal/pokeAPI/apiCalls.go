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

func CommandMap(args []string) (error) {
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

func CommandMapb(args []string) (error) {
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

func CommandExplore(args []string) (error) {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <location>")
	}
	location := args[0]
	url := fmt.Sprintf("%slocation-area/%s/", baseURL, location)
	var exploreCallResults exploreCall
	if data, found := cache.Get(url); found {
		
		err := json.Unmarshal(data, &exploreCallResults)
		if err != nil {
			fmt.Println("Error unmarshaling cached data:", err)
			return err
		}
	} else {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return err
		}

		cache.Add(url, body)

		err = json.Unmarshal(body, &exploreCallResults)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return err
		}

	}
	fmt.Printf("Exploring %s...\n", location)
	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range exploreCallResults.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func GetPokemonData(pokemonName string) (PokemonData, error) {
	url := fmt.Sprintf("%spokemon/%s/", baseURL, pokemonName)
	
	var pokemonData PokemonData
	if data, found := cache.Get(url); found {
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			fmt.Println("Error unmarshaling cached data:", err)
			return PokemonData{}, err
		}
	} else {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			return PokemonData{}, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return PokemonData{}, err
		}

		cache.Add(url, body)

		err = json.Unmarshal(body, &pokemonData)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return PokemonData{}, err
		}
	}

	return pokemonData, nil
}