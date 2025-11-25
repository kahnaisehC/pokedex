package pokeapiclient

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/kahnaisehC/pokedex/internal/pokecache"
)

type PokeAPIClient struct {
	cache *pokecache.Pokecache
}

func NewClient(interval time.Duration) *PokeAPIClient {
	cache := pokecache.NewPokecache(interval)
	pokeClient := PokeAPIClient{
		cache: cache,
	}
	return &pokeClient
}

type LocationAreaListResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *PokeAPIClient) GetLocationAreasList(url *string) (LocationAreaListResponse, error) {
	locations := LocationAreaListResponse{}
	var err error

	if url == nil {
		u := "https://pokeapi.co/api/v2/location-area/"
		url = &u
	}

	body, ok := c.cache.Get(*url)
	if !ok {
		r, err := http.Get(*url)
		if err != nil {
			return locations, err
		}
		defer r.Body.Close()

		body, err = io.ReadAll(r.Body)
		if err != nil {
			return locations, err
		}
	}

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return locations, err
	}
	c.cache.Add(*url, body)
	return locations, nil
}

type LocationDetailsResponse struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *PokeAPIClient) GetPokemonsInLocation(location string) (LocationDetailsResponse, error) {
	pokemonList := LocationDetailsResponse{}
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	url := baseUrl + location

	_, err := http.Get(url)
	if err != nil {
		return pokemonList, err
	}

	return pokemonList, nil
}
