// reple.go
package main

import (
	"strings"
	"os"
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokeAPI"
)

var pokedex = map[string]pokeAPI.PokemonData{}


func cleanInput(text string) []string {
	text = strings.ToLower(text)
	textArr := strings.Fields(text)

	return textArr
}

func commandExit(args []string) error {
		fmt.Printf("Closing the Pokedex... Goodbye!\n")
		os.Exit(0)
		return nil
}

func commandHelp(args []string) error {
	fmt.Printf("Available commands:\n")
	for _, cmd := range getCommand() {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandCatch(args []string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	pokemon, err := pokeAPI.GetPokemonData(args[0])
	if err != nil {
		return err
	}

	baseExp := pokemon.BaseExperience
	chance := 1.0 - (float64(baseExp) / 400.0)
	if chance < 0.1 {
		chance = 0.1
	}
	if chance > 0.9 {
		chance = 0.9
	}

	randNum := rand.Float64()
	if randNum > chance {
		fmt.Printf("%s escaped the Pokeball!\n", pokemon.Name)
		return nil
	} else {

		fmt.Printf("Caught %s!\n", pokemon.Name)
		pokedex[pokemon.Name] = pokemon
	}


	return nil
}

func commandInventory(args []string) error {
	if len(pokedex) == 0 {
		fmt.Printf("Your Pokedex is empty. Go catch some Pokemon!\n")
		return nil
	}

	fmt.Printf("Your Pokedex:\n")
	for name := range pokedex {
		fmt.Printf("- %s\n", name)
	}

	return nil
}

func commandInspect(args []string) error {
	if len(args) < 1 {
		fmt.Printf("Usage: inspect <pokemon_name>\n")
		return nil
	}

	name := args[0]
	pokemon, found := pokedex[name]	
	if !found {
		fmt.Printf("you have not caught %s yet.\n", name)
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
