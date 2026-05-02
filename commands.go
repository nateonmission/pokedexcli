// commands.go
package main

import (
	"pokedexcli/internal/pokeAPI"
)

type cliCommand struct {
	name	string
	description	string
	callback	func() error
}

var commandMap = getCommand()

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex CLI",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Show this help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Show the next 20 location areas",
			callback: pokeAPI.CommandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Show the previous 20 location areas",
			callback: pokeAPI.CommandMapb,
		},
	}
}