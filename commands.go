// commands.go
package main



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
	}
}