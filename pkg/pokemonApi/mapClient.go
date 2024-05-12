package pokemonApi

import (
	"errors"

	"github.com/garrettkucinski/pokedex-cli/pkg/pokemonApi/internal/locations"
)

type MapClient struct {
	Url  string         `json:"url"`
	Data locations.Data `json:"data"`
}

func (mc *MapClient) DisplayNextLocationList() (displayError error) {
	locationUrl := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

	if mc.Data.Next != nil {
		locationUrl = *mc.Data.Next
	}

	if err := mc.Data.GetLocationData(locationUrl); err != nil {
		displayError = errors.New("cannot get next location")
	}

	if err := mc.Data.PrintMapAreas(); err != nil {
		displayError = errors.New("could not display locations")
	}

	return
}

func (mc *MapClient) DisplayPrevLocationList() (displayError error) {
	if mc.Data.Previous != nil {
		if err := mc.Data.GetLocationData(*mc.Data.Previous); err != nil {
			displayError = errors.New("cannot get previous location")
		}
	} else {
		displayError = errors.New("cannot go back")
	}

	if err := mc.Data.PrintMapAreas(); err != nil {
		displayError = errors.New("could not display locations")
	}

	return
}
