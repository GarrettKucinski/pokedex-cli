package pokedex

type Captures struct {
	List map[string]Pokemon
}

func (c *Captures) Add(pokemon Pokemon) {
	c.List[pokemon.Name] = pokemon
}
