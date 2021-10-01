package namegen

// NameGenerator is a set of names to use
type NameGenerator struct {
	MaleFirstNames   []string
	FemaleFirstNames []string
	LastNames        []string
}

// NameGeneratorFromType sets up types of names
func NameGeneratorFromType(origin, gender string) NameGenerator {
	nameGenerators := map[string]NameGenerator{
		"anglosaxon": {anglosaxonMaleFirstNames, anglosaxonFemaleFirstNames, anglosaxonLastNames},
		"dutch":      {dutchMaleFirstNames, dutchFemaleFirstNames, dutchLastNames},
		"dwarf":      {dwarfMaleFirstNames, dwarfFemaleFirstNames, getDwarfLastNames(gender)},
		"elf":        {elfMaleFirstNames, elfFemaleFirstNames, elfLastNames},
		"english":    {englishMaleFirstNames, englishFemaleFirstNames, englishLastNames},
		"fantasy":    {fantasyMaleFirstNames, fantasyFemaleFirstNames, fantasyLastNames},
		"german":     {germanMaleFirstNames, germanFemaleFirstNames, germanLastNames},
		"greek":      {greekMaleFirstNames, greekFemaleFirstNames, greekLastNames},
		"icelandic":  {getIcelandicFirstNames(), getIcelandicFirstNames(), getIcelandicLastNames(gender)},
		"indonesian": {indonesianMaleFirstNames, indonesianFemaleFirstNames, indonesianLastNames},
		"japanese":   {japaneseMaleFirstNames, japaneseFemaleFirstNames, japaneseLastNames},
		"korean":     {koreanMaleFirstNames, koreanFemaleFirstNames, koreanLastNames},
		"russian":    {russianMaleFirstNames, russianFemaleFirstNames, russianLastNames},
		"spanish":    {spanishMaleFirstNames, spanishFemaleFirstNames, spanishLastNames},
		"swedish":    {swedishMaleFirstNames, swedishFemaleFirstNames, swedishLastNames},
		"thai":       {thaiMaleFirstNames, thaiFemaleFirstNames, thaiLastNames},
		"portuguese": {portugueseMaleFirstNames, portugueseFemaleFirstNames, portugueseLastNames},
		"hindu":	  {hinduMaleFirstNames, hinduFemaleFirstNames, hinduLastNames},
	}

	return nameGenerators[origin]
}

// LastName returns a last name
func (gen NameGenerator) LastName() string {
	return RandomItem(gen.LastNames)
}

// FirstName returns a first name
func (gen NameGenerator) FirstName(gender string) string {
	firstNames := gen.MaleFirstNames
	if gender == "female" {
		firstNames = gen.FemaleFirstNames
	} else if gender == "both" {
		firstNames = append(firstNames, gen.FemaleFirstNames...)
	}

	return RandomItem(firstNames)
}

// CompleteName returns a complete name
func (gen NameGenerator) CompleteName(gender string) string {
	return gen.FirstName(gender) + " " + gen.LastName()
}
