package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/kahnaisehC/pokedex/internal/pokeapiClient"
)

type commandConfig struct {
	Next   *string
	Prev   *string
	Client *pokeapiclient.PokeAPIClient
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *commandConfig) error
}

func commandExit(cfg *commandConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *commandConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")

	for _, cmd := range commandMap {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandDisplayBMap(cfg *commandConfig) error {
	var err error
	url := cfg.Prev
	client := cfg.Client

	if url == nil {
		return errors.New("already in the first page! try the 'map' command")
	}

	locationAreas, err := client.GetLocationAreasList(url)
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

func commandDisplayMap(cfg *commandConfig) error {
	var err error
	url := cfg.Next
	client := cfg.Client

	locationAreas, err := client.GetLocationAreasList(url)
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
