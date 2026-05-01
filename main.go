// main.go
package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {
	
	fmt.Printf("Welcome to the Pokedex!\n")
	scanner := bufio.NewScanner(os.Stdin)
	loop := true
	for loop {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputArr := cleanInput(input)
		if len(inputArr) == 0 {
			continue
		}
		cmd := inputArr[0]
		if command, ok := commandMap[cmd]; ok {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error executing command: %s\n", err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", cmd)
		}
	}
}