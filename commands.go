package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/kahnaisehC/pokedex/internal/pokeapiClient"
)

type commandConfig struct {
	Next            *string
	Prev            *string
	Client          *pokeapiclient.PokeAPIClient
	Arguments       []string
	CatchedPokemons []pokeapiclient.PokemonDetails
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *commandConfig) error
}

func commandExplore(cfg *commandConfig) error {
	if len(cfg.Arguments) == 0 {
		return errors.New("type \"explore <location> \" to explore that location")
	}
	client := cfg.Client
	location := cfg.Arguments[0]

	pokemonList, err := client.GetPokemonsInLocation(location)
	if err != nil {
		return err
	}

	if len(pokemonList.PokemonEncounters) == 0 {
		println("There are no pokemons in " + location + "!!!")
		return nil
	}
	println("in " + location + " you can encounter the following Pokemon:")
	for _, encounter := range pokemonList.PokemonEncounters {
		println(encounter.Pokemon.Name)
	}

	return nil
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

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func commandCatch(cfg *commandConfig) error {
	if len(cfg.Arguments) == 0 {
		return errors.New("type \"catch <pokemons name> \" to try and catch a pokemon")
	}
	const maxExp = 10000
	var err error

	baseUrl := "https://pokeapi.co/api/v2/pokemon/"

	url := baseUrl + cfg.Arguments[0]
	client := cfg.Client
	pokemon, err := client.GetPokemonDetails(url)
	if err != nil {
		return err
	}

	println("Throwing a Pokeball at " + pokemon.Name + "...")
	experience := pokemon.BaseExperience
	chance := maxInt(rand.Intn(maxExp)-experience, 0)
	if chance <= maxExp/2 {
		println(pokemon.Name + " was caught!")
		cfg.CatchedPokemons = append(cfg.CatchedPokemons, pokemon)
	} else {
		println(pokemon.Name + " escaped!")
	}

	return nil
}
