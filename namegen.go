package namegen

import (
	"github.com/ironarachne/utility"
)

// NameGenerator is a set of names to use
type NameGenerator struct {
	MaleFirstNames   []string
	FemaleFirstNames []string
	LastNames        []string
}

// NameGeneratorFromType sets up types of names
func NameGeneratorFromType(origin string) NameGenerator {
	nameGenerators := map[string]NameGenerator{
		"dutch":      {dutchMaleFirstNames, dutchFemaleFirstNames, dutchLastNames},
		"english":    {englishMaleFirstNames, englishFemaleFirstNames, englishLastNames},
		"fantasy":    {fantasyMaleFirstNames, fantasyFemaleFirstNames, fantasyLastNames},
		"german":     {germanMaleFirstNames, germanFemaleFirstNames, germanLastNames},
		"greek":      {greekMaleFirstNames, greekFemaleFirstNames, greekLastNames},
		"icelandic":  {icelandicMaleFirstNames, icelandicFemaleFirstNames, icelandicLastNames},
		"indonesian": {indonesianMaleFirstNames, indonesianFemaleFirstNames, indonesianLastNames},
		"korean":     {koreanMaleFirstNames, koreanFemaleFirstNames, koreanLastNames},
		"spanish":    {spanishMaleFirstNames, spanishFemaleFirstNames, spanishLastNames},
	}

	return nameGenerators[origin]
}

// LastName returns a last name
func (gen NameGenerator) LastName() string {
	return utility.RandomItem(gen.LastNames)
}

// FirstName returns a first name
func (gen NameGenerator) FirstName(gender string) string {
	firstNames := gen.MaleFirstNames
	if gender == "female" {
		firstNames = gen.FemaleFirstNames
	} else if gender == "both" {
		firstNames = append(firstNames, gen.FemaleFirstNames...)
	}

	return utility.RandomItem(firstNames)
}

// CompleteName returns a complete name
func (gen NameGenerator) CompleteName(gender string) string {
	return gen.FirstName(gender) + " " + gen.LastName()
}
