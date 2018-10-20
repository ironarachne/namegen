package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ironarachne/namegen"
)

func main() {
	numberOfNames := flag.Int("n", 1, "Number of names to generate")
	gender := flag.String("g", "both", "Gender (male, female, or both)")
	origin := flag.String("o", "english", "Origin of names (e.g., english, spanish, etc.)")
	randomSeed := flag.Int64("s", 0, "Optional random generator seed")

	mode := flag.String("m", "full", "Mode of generation")
	var list = flag.Bool("l", false, "Print available name lists")

	flag.Parse()

	if *list == true {
		nameLists := []string{
			"anglosaxon",
			"dutch",
			"english",
			"fantasy",
			"german",
			"greek",
			"indonesian",
			"japanese",
			"korean",
			"spanish",
		}
		fmt.Printf("Available name lists: \n%s\n\n", strings.Join(nameLists, "\n"))
	}

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
			output = generator.CompleteName(*gender)
		case "firstname":
			output = generator.FirstName(*gender)
		case "lastname":
			output = generator.LastName()
		default:
			output = fmt.Sprintf("Unsupported generation mode %s, supported modes: full, firstname, lastname", *mode)
		}

		fmt.Println(output)
	}
}
