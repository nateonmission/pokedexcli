// reple.go
package main

import (
	"strings"
	"os"
	"fmt"
)


func cleanInput(text string) []string {
	text = strings.ToLower(text)
	textArr := strings.Fields(text)

	return textArr
}

func commandExit() error {
		fmt.Printf("Closing the Pokedex... Goodbye!\n")
		os.Exit(0)
		return nil
}

func commandHelp() error {
	fmt.Printf("Available commands:\n")
	for _, cmd := range getCommand() {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}

	return nil
}





