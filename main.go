package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	pokeapiclient "github.com/kahnaisehC/pokedex/internal/pokeapiClient"
)

var commandMap map[string]cliCommand

func init() {
	commandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 locations in the Pokemon World!",
			callback:    commandDisplayMap,
		},
		"bmap": {
			name:        "bmap",
			description: "Displays the names of the previous 20 locations in the Pokemon World!",
			callback:    commandDisplayBMap,
		},
	}
}

func main() {
	cacheCleanUpTime := time.Second * 5
	pokeClient := pokeapiclient.NewClient(cacheCleanUpTime)
	cfg := commandConfig{
		Next:   nil,
		Prev:   nil,
		Client: pokeClient,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		commandString := cleanedInput[0]
		command, ok := commandMap[commandString]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(&cfg)
		if err != nil {
			fmt.Println(err.Error())
		}

	}
}
