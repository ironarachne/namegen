package namegen

import (
	"github.com/ironarachne/utility"
)

// NameGenerator is a set of names to use
type NameGenerator struct {
	FirstNames []string
	LastNames  []string
}

// NameGeneratorFromType sets up types of names
func NameGeneratorFromType(origin string) NameGenerator {
	nameGenerators := map[string]NameGenerator{
		"english":    {englishFirstNames, englishLastNames},
		"spanish":    {spanishFirstNames, spanishLastNames},
		"german":     {germanFirstNames, germanLastNames},
		"thai":       {thaiFirstNames, thaiLastNames},
		"korean":     {koreanFirstNames, koreanLastNames},
		"iceland":    {icelandicFirstNames, icelandicLastNames},
		"dutch":      {dutchFirstNames, dutchLastNames},
		"anglosaxon": {anglosaxonFirstNames, anglosaxonLastNames},
		"greek":      {greekFirstNames, greekLastNames},
		"indonesian": {indonesianFirstNames, indonesianLastNames},
	}

	return nameGenerators[origin]
}

// LastName returns a last name
func (gen NameGenerator) LastName() string {
	return utility.RandomItem(gen.LastNames)
}

// FirstName returns a first name
func (gen NameGenerator) FirstName() string {
	return utility.RandomItem(gen.FirstNames)
}

// CompleteName returns a complete name
func (gen NameGenerator) CompleteName() string {
	return gen.FirstName() + " " + gen.LastName()
}
