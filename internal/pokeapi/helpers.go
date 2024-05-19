package pokeapi

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetFormattedName(name string) string {
	return cases.Title(language.English).String(strings.ReplaceAll(name, "-", " "))
}
