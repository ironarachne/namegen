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

	mode := flag.String("m", "full", "Mode of generation")

	flag.Parse()

	if *randomSeed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}

	generator := namegen.NameGeneratorFromType(*origin)

	output := ""
	for i := 0; i < *numberOfNames; i++ {

		switch *mode {
		case "full":
			output = generator.CompleteName()
		case "firstname":
			output = generator.FirstName()
		case "lastname":
			output = generator.LastName()
		default:
			output = fmt.Sprintf("Unsupported generation mode %s, supported modes: full, firstname, lastname", *mode)
		}

		fmt.Println(output)
	}
}
