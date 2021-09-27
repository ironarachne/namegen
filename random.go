package namegen

import (
	"crypto/md5"
	"encoding/binary"
	"io"
	"math/rand"
)

// RandomItem returns a random string from an array of strings
func RandomItem(items []string) string {
	return items[rand.Intn(len(items))]
}

// RandomItemFromThresholdMap returns a random weighted string
func RandomItemFromThresholdMap(items map[string]int) string {
	result := ""
	ceiling := 0
	start := 0
	var thresholds = make(map[string]int)

	for item, weight := range items {
		ceiling += weight
		thresholds[item] = start
		start += weight
	}

	randomValue := rand.Intn(ceiling)

	for item, threshold := range thresholds {
		if threshold <= randomValue {
			result = item
		}
	}

	return result
}

// RandomItemInCollection checks to see if a string is in an array of strings
func RandomItemInCollection(item string, collection []string) bool {
	for _, element := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// RandomSeedFromString uses a string to seed the random number generator
func RandomSeedFromString(source string) {
	h := md5.New()
	_, err := io.WriteString(h, source)
	if err != nil {
		panic(err)
	}
	seed := binary.BigEndian.Uint64(h.Sum(nil))
	rand.Seed(int64(seed))
}