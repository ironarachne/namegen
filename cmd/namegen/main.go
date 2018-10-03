package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/namegen"
)

func main() {
	numberOfNames := flag.Int("n", 1, "Number of names to generate")
	origin := flag.String("o", "english", "Origin of names (e.g., english, conlang, etc.)")
	randomSeed := flag.Int64("s", 0, "Optional random generator seed")
	flag.Parse()

	name := ""
	nameOrigin := *origin

	if *randomSeed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}

	for i := 0; i < *numberOfNames; i++ {
		if nameOrigin == "conlang" {
			name = naminglanguage.GeneratePersonName()
		} else {
			name = namegen.GetNameByType(nameOrigin)
		}

		fmt.Println(name)
	}
}
