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
	firstNames := []string{}
	lastNames := []string{}

	if origin == "english" {
		firstNames = englishFirstNames
		lastNames = englishLastNames
	} else if origin == "spanish" {
		firstNames = spanishFirstNames
		lastNames = spanishLastNames
	}

	return NameGenerator{
		FirstNames: firstNames,
		LastNames:  lastNames,
	}
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
