package namegen

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"io"
	"math/rand"
)

var (
	ErrorEmptyItems = errors.New("empty items")
)

// RandomItem returns a random string from an array of strings
func RandomItem(items []string) (string, error) {
	itemsLen := len(items)
	if itemsLen <= 0 {
		return "", ErrorEmptyItems
	}

	return items[rand.Intn(itemsLen)], nil
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
