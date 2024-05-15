package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/garrettkucinski/pokedex-cli/internal/pokecache"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (l *Location) GetFormattedName() string {
	return cases.Title(language.English).String(strings.ReplaceAll(l.Name, "-", " "))
}

type Data struct {
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
	Count    int        `json:"count"`
}

func (d *Data) GetLocationData(location string, cache *pokecache.Cache) (responseError error) {
	if val, found := cache.Get(location); found {
		fmt.Println(val, found)
		json.Unmarshal(val, d)
		return
	}

	if res, responseError := http.Get(location); responseError == nil {
		if body, responseError := io.ReadAll(res.Body); responseError == nil {
			json.Unmarshal(body, d)
			cache.Add(location, body)
			res.Body.Close()
		}
	}

	return
}

func (d *Data) PrintMapAreas() (printError error) {
	for _, result := range d.Results {
		fmt.Println(result.GetFormattedName())
	}

	return
}
