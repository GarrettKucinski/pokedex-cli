package pokeapi

import "fmt"

type Encounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}

type Encounters struct {
	PokemonEncounters []Encounter `json:"pokemon_encounters"`
}

func (e *Encounters) PrintPokemonNames() (printError error) {
	for _, encounter := range e.PokemonEncounters {
		fmt.Println(GetFormattedName(encounter.Pokemon.Name))
	}

	return
}
