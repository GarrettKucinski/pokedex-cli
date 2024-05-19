package pokeapi

import (
	"fmt"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Locations struct {
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
	Count    int        `json:"count"`
}

func (l *Locations) PrintMapAreas() (printError error) {
	for _, result := range l.Results {
		fmt.Println(GetFormattedName(result.Name))
	}

	return
}
