package namegen

import (
	"testing"

	"github.com/go-playground/assert"
)

func Test_RandomItem(t *testing.T) {
	result, err := RandomItem([]string{})
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(result))

	generator := NameGeneratorFromType(origins[0], genders[0])
	result, err = RandomItem(generator.FemaleFirstNames)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, len(result))
}

func Test_RandomItemFromThresholdMap(t *testing.T) {
	genderTreshold := map[string]int{}
	for i, gender := range genders {
		genderTreshold[gender] = i
	}
	result := RandomItemFromThresholdMap(genderTreshold)
	assert.NotEqual(t, "", result)
}

func Test_RandomItemInCollection(t *testing.T) {
	for _, origin := range origins {
		ok := RandomItemInCollection(origin, origins)
		assert.Equal(t, true, ok)
	}

	ok := RandomItemInCollection("none", origins)
	assert.Equal(t, false, ok)
}

func Test_RandomSeedFromString(t *testing.T) {
	defer func() {
		err := recover()
		assert.Equal(t, nil, err)
	}()

	RandomSeedFromString("none")
}
