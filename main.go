package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

var commandMap map[string]cliCommand

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	for _, cmd := range commandMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

var pageNumber = 1

type config struct {
	Next string
	Prev string
}

type locationAreas struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func getLocationAreas(url string) (locationAreas, error) {
	locationAreas := locationAreas{}
	r, err := http.Get(url)
	if err != nil {
		return locationAreas, err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return locationAreas, err
	}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return locationAreas, err
	}
	return locationAreas, nil
}

func commandDisplayBMap(cfg *config) error {
	url := cfg.Prev
	if url == "" {
		return errors.New("already in the first page! try the 'map' command")
	}

	locationAreas, err := getLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.Next = locationAreas.Next
	cfg.Prev = locationAreas.Previous
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandDisplayMap(cfg *config) error {
	url := cfg.Next

	locationAreas, err := getLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.Next = locationAreas.Next
	cfg.Prev = locationAreas.Previous
	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

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
		"bmap": {
			name:        "bmap",
			description: "Displays the names of the previous 20 locations in the Pokemon World!",
			callback:    commandDisplayBMap,
		},
	}
}

func main() {
	cfg := config{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Prev: "",
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
