package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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

func (d *Data) GetLocationData(location string) (responseError error) {
	if res, responseError := http.Get(location); responseError == nil {
		if body, responseError := io.ReadAll(res.Body); responseError == nil {
			json.Unmarshal(body, d)
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
