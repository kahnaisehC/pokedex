package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...any) error
}

var commandMap map[string]cliCommand

func commandExit(...any) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(...any) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	for _, cmd := range commandMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

var pageNumber = 1

func commandDisplayMap(...any) error {
	// api endpoint: https://pokeapi.co/api/v2/location-area/{id or name}/

	/*{
	RAW JSON
	  "id": 1,
	  "name": "canalave-city-area",
	  "game_index": 1,
	  "encounter_method_rates": [
	    {
	      "encounter_method": {
	        "name": "old-rod",
	        "url": "https://pokeapi.co/api/v2/encounter-method/2/"
	      },
	      "version_details": [
	        {
	          "rate": 25,
	          "version": {
	            "name": "platinum",
	            "url": "https://pokeapi.co/api/v2/version/14/"
	          }
	        }
	      ]
	    }
	  ],
	  "location": {
	    "name": "canalave-city",
	    "url": "https://pokeapi.co/api/v2/location/1/"
	  },
	  "names": [
	    {
	      "name": "",
	      "language": {
	        "name": "en",
	        "url": "https://pokeapi.co/api/v2/language/9/"
	      }
	    }
	  ],
	  "pokemon_encounters": [
	    {
	      "pokemon": {
	        "name": "tentacool",
	        "url": "https://pokeapi.co/api/v2/pokemon/72/"
	      },
	      "version_details": [
	        {
	          "version": {
	            "name": "diamond",
	            "url": "https://pokeapi.co/api/v2/version/12/"
	          },
	          "max_chance": 60,
	          "encounter_details": [
	            {
	              "min_level": 20,
	              "max_level": 30,
	              "condition_values": [],
	              "chance": 60,
	              "method": {
	                "name": "surf",
	                "url": "https://pokeapi.co/api/v2/encounter-method/5/"
	              }
	            }
	          ]
	        }
	      ]
	    }
	  ]
	}
	*/

	return nil
}

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
	}
}

func main() {
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
		err := command.callback()
		if err != nil {
			fmt.Println(err.Error())
		}

	}
}
