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
		"estonian":   {estonianMaleFirstNames, estonianFemaleFirstNames, estonianLastNames},
		"fantasy":    {fantasyMaleFirstNames, fantasyFemaleFirstNames, fantasyLastNames},
		"finnish":    {finnishMaleFistNames, finnishFemaleFirstNames, finnishLastNames},
		"french":     {frenchMaleFistNames, frenchFemaleFirstNames, frenchLastNames},
		"german":     {germanMaleFirstNames, germanFemaleFirstNames, germanLastNames},
		"greek":      {greekMaleFirstNames, greekFemaleFirstNames, greekLastNames},
		"hindu":      {hinduMaleFirstNames, hinduFemaleFirstNames, hinduLastNames},
		"icelandic":  {getIcelandicFirstNames(), getIcelandicFirstNames(), getIcelandicLastNames(gender)},
		"indonesian": {indonesianMaleFirstNames, indonesianFemaleFirstNames, indonesianLastNames},
		"irish":      {irishMaleFirstNames, irishFemaleFirstNames, irishLastNames},
		"italian":    {italianMaleFirstNames, italianFemaleFirstNames, italianLastNames},
		"japanese":   {japaneseMaleFirstNames, japaneseFemaleFirstNames, japaneseLastNames},
		"korean":     {koreanMaleFirstNames, koreanFemaleFirstNames, koreanLastNames},
		"maori":      {maoriMaleFirstNames, maoriFemaleFirstNames, maoriLastNames},
		"mayan":      {mayanMaleFirstNames, mayanFemaleFirstNames, mayanLastNames},
		"nepalese":   {nepaleseMaleFirstNames, nepaleseFemaleFirstNames, nepaleseLastNames},
		"norwegian":  {norwegianMaleFirstNames, norwegianFemaleFirstNames, norwegianLastNames},
		"portuguese": {portugueseMaleFirstNames, portugueseFemaleFirstNames, portugueseLastNames},
		"russian":    {russianMaleFirstNames, russianFemaleFirstNames, russianLastNames},
		"spanish":    {spanishMaleFirstNames, spanishFemaleFirstNames, spanishLastNames},
		"swedish":    {swedishMaleFirstNames, swedishFemaleFirstNames, swedishLastNames},
		"thai":       {thaiMaleFirstNames, thaiFemaleFirstNames, thaiLastNames},
		"mongolian":  {mongolianMaleFirstNames, mongolianFemaleFirstNames, mongolianLastNames},
		"ukrainian":  {ukrainianMaleFirstNames, ukrainianFemaleFirstNames, ukrainianLastNames},
		"somalia":    {somaliaMaleFirstNames, somaliaFemaleFirstNames, somaliaLastNames},
		"arabic":     {arabicMaleFirstNames, arabicFemaleFirstNames, arabicLastNames},
		"hawaiian":   {hawaiianMaleFirstNames, hawaiianFemaleFirstNames, hawaiianLastNames},
		"turkish":    {turkishMaleFirstNames, turkishFemaleFirstNames, turkishLastNames},
		"serbian":    {serbianMaleFirstNames, serbianFemaleFirstNames, serbianLastNames},
		"nigerian":   {nigerianMaleFirstNames, nigerianFemaleFirstNames, nigerianLastNames},
		"polish":     {polishMaleFirstNames, polishFemaleFirstNames, polishLastNames},
		"chinese":    {chineseMaleFirstNames, chineseFemaleFirstNames, chineseLastNames},
	}

	return nameGenerators[origin]
}

// LastName returns a last name
func (gen NameGenerator) LastName() (string, error) {
	return RandomItem(gen.LastNames)
}

// FirstName returns a first name
func (gen NameGenerator) FirstName(gender string) (string, error) {
	firstNames := gen.MaleFirstNames
	if gender == "female" {
		firstNames = gen.FemaleFirstNames
	} else if gender == "both" {
		firstNames = append(firstNames, gen.FemaleFirstNames...)
	}

	return RandomItem(firstNames)
}

// CompleteName returns a complete name
func (gen NameGenerator) CompleteName(gender string) (string, error) {
	firstName, err := gen.FirstName(gender)
	if err != nil {
		return "", err
	}

	lastName, err := gen.LastName()
	if err != nil {
		return "", err
	}

	fullname := firstName + " " + lastName
	return fullname, nil
}
