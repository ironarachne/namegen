package namegen

import (
	"math/rand"
)

func randomItem(items []string) string {
	return items[rand.Intn(len(items))]
}

type NameGenerator struct {
	FirstNames []string
	LastNames  []string
}

func NameGeneratorFromType(origin string) NameGenerator {
	nameGenerators := map[string]NameGenerator{
		"english": {englishFirstNames, englishLastNames},
		"spanish": {spanishFirstNames, spanishLastNames},
		"german":  {germanFirstNames, germanLastNames},
		"thai":    {thaiFirstNames, thaiLastNames},
		"korean":  {koreanFirstNames, koreanLastNames},
    "iceland": {icelandicFirstNames, icelandicLastNames},
    "dutch":   {dutchFirstNames, dutchLastNames},
	}

	return nameGenerators[origin]
}

func (gen NameGenerator) LastName() string {
	return randomItem(gen.LastNames)
}

func (gen NameGenerator) FirstName() string {
	return randomItem(gen.FirstNames)
}

func (gen NameGenerator) CompleteName() string {
	return gen.FirstName() + " " + gen.LastName()
}
