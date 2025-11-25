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
