package namegen

import (
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

var genders []string = []string{"male", "female", "both"}
var origins []string = []string{
	"anglosaxon", "dutch", "dwarf", "elf", "english", "estonian", "fantasy", "finnish",
	"french", "german", "greek", "hindu", "icelandic", "indonesian", "italian", "japanese",
	"korean", "mayan", "nepalese", "norwegian", "portuguese", "russian", "spanish", "swedish",
	"thai",
}

func Test_NameGeneratorFromType(t *testing.T) {
	for _, origin := range origins {
		for _, gender := range genders {
			result := NameGeneratorFromType(origin, gender)
			assert.NotEqual(t, 0, len(result.FemaleFirstNames))
			assert.NotEqual(t, 0, len(result.MaleFirstNames))
			assert.NotEqual(t, 0, len(result.LastNames))
		}
	}
}

func Test_NameGenerator_LastName(t *testing.T) {
	for _, origin := range origins {
		for _, gender := range genders {
			generator := NameGeneratorFromType(origin, gender)
			result, err := generator.LastName()
			assert.Equal(t, nil, err)
			assert.NotEqual(t, 0, len(result))
		}
	}
}

func Test_NameGenerator_FirstName(t *testing.T) {
	for _, origin := range origins {
		for _, gender := range genders {
			generator := NameGeneratorFromType(origin, gender)
			result, err := generator.FirstName(gender)
			assert.Equal(t, nil, err)
			assert.NotEqual(t, 0, len(result))
		}
	}
}

func Test_NameGenerator_CompleteName(t *testing.T) {
	generator := &NameGenerator{
		FemaleFirstNames: germanFemaleFirstNames,
		MaleFirstNames:   germanMaleFirstNames,
	}
	result, err := generator.CompleteName(genders[0])
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(result))

	generator = &NameGenerator{
		LastNames: germanLastNames,
	}
	result, err = generator.CompleteName(genders[0])
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(result))

	for _, origin := range origins {
		for _, gender := range genders {
			generator := NameGeneratorFromType(origin, gender)
			result, err := generator.CompleteName(gender)
			assert.Equal(t, nil, err)
			assert.NotEqual(t, 0, len(result))
			assert.Equal(t, true, strings.Contains(result, " "))
		}
	}
}
