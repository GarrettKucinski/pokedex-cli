package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/garrettkucinski/pokedex-cli/internal/pokecache"
	"github.com/garrettkucinski/pokedex-cli/internal/pokedex"
)

type Client struct {
	BaseUrl string           `json:"baseUrl"`
	Cache   *pokecache.Cache `json:"cache"`
}

type Config struct {
	Pokedex  pokedex.Captures
	Client   Client
	Previous *string
	Next     *string
}

func NewConfig() *Config {
	client := &Client{
		Cache:   pokecache.NewCache(1000 * 60 * 10),
		BaseUrl: "https://pokeapi.co/api/v2",
	}

	return &Config{
		Client: *client,
		Pokedex: pokedex.Captures{
			List: make(map[string]pokedex.Pokemon),
		},
	}
}

func (c *Client) GetApiData(url string, data interface{}) (responseError error) {
	if val, found := c.Cache.Get(url); found {
		json.Unmarshal(val, data)

		return
	}

	if res, responseError := http.Get(url); responseError == nil {
		if body, responseError := io.ReadAll(res.Body); responseError == nil {
			json.Unmarshal(body, data)
			c.Cache.Add(url, body)
			res.Body.Close()
		}
	}

	return
}

func (c *Client) CatchPokemon(cfg *Config, pokemonName string) (displayError error) {
	data := new(pokedex.Pokemon)
	pokemonUrl := c.BaseUrl + "/pokemon/" + pokemonName

	if err := c.GetApiData(pokemonUrl, data); err != nil {
		displayError = errors.New("could not throw pokeball")
	}

	captured := data.AttemptCapture()

	if captured {
		fmt.Printf("You caught %s!\n", data.Name)
		cfg.Pokedex.Add(*data)
	} else {
		fmt.Printf("%s escaped!\n", data.Name)
	}

	return
}

func (c *Client) ExploreLocation(locationName string) (displayError error) {
	data := new(Encounters)
	areaUrl := c.BaseUrl + "/location-area/" + locationName

	if err := c.GetApiData(areaUrl, data); err != nil {
		displayError = errors.New("cannot explore area")
	}

	fmt.Println("Exploring", locationName, "...")
	fmt.Println("Found Pokemon:")

	if err := data.PrintPokemonNames(); err != nil {
		displayError = errors.New("could not print pokemon")
	}

	return
}

func (c *Client) DisplayNextLocationList(cfg *Config) (displayError error) {
	locations := new(Locations)
	locationUrl := c.BaseUrl + "/location-area"

	if cfg.Next == nil {
		cfg.Next = &locationUrl
	}

	if err := c.GetApiData(*cfg.Next, locations); err != nil {
		displayError = errors.New("cannot get next location")
	}

	if err := locations.PrintMapAreas(); err != nil {
		displayError = errors.New("could not display locations")
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return
}

func (c *Client) DisplayPrevLocationList(cfg *Config) (displayError error) {
	locations := new(Locations)

	if cfg.Previous != nil {
		if err := c.GetApiData(*cfg.Previous, locations); err != nil {
			displayError = errors.New("cannot get previous location")
		}

		cfg.Next = locations.Next
		cfg.Previous = locations.Previous

	} else {
		displayError = errors.New("cannot go back")
	}

	if err := locations.PrintMapAreas(); err != nil {
		displayError = errors.New("could not display locations")
	}

	return
}
