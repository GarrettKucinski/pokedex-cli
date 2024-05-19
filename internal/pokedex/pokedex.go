package pokedex

import "fmt"

type Captures struct {
	List map[string]*Pokemon
}

func (c *Captures) Add(pokemon *Pokemon) {
	c.List[pokemon.Name] = pokemon
}

func (c *Captures) Get(pokemonName string) (pokemon *Pokemon, found bool) {
	if pokemon, ok := c.List[pokemonName]; ok {
		return pokemon, ok
	}

	return
}

func (c *Captures) View() {
	fmt.Println("These are the Pokemon you've captured!")
	for _, capture := range c.List {
		fmt.Printf("\t-%s\n", capture.Name)
	}
}
