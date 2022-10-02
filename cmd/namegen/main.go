package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ironarachne/namegen"
)

func main() {
	numberOfNames := flag.Int("n", 1, "Number of names to generate")
	gender := flag.String("g", "both", "Gender (male, female, or both)")
	origin := flag.String("o", "english", "Origin of names (e.g., english, spanish, etc.)")
	randomSeed := flag.String("s", "none", "Optional random generator seed (alphanumeric)")

	mode := flag.String("m", "full", "Mode of generation")
	var list = flag.Bool("l", false, "Print available name lists")

	flag.Parse()

	if *list {
		nameLists := []string{
			"anglosaxon",
			"dutch",
			"dwarf",
			"elf",
			"english",
			"estonian",
			"fantasy",
			"finnish",
			"french",
			"german",
			"greek",
			"hindu",
			"indonesian",
			"italian",
			"japanese",
			"korean",
			"mayan",
			"nepalese",
			"norwegian",
			"portuguese",
			"russian",
			"spanish",
			"swedish",
			"thai",
			"ukrainian",
			"somalia",
      "arabic",
			"hawaiian",
		
		}
		fmt.Printf("Available name lists: \n%s\n\n", strings.Join(nameLists, "\n"))
		os.Exit(0)
	}

	if *randomSeed == "none" {
		rand.Seed(time.Now().UnixNano())
	} else {
		namegen.RandomSeedFromString(*randomSeed)
	}

	generator := namegen.NameGeneratorFromType(*origin, *gender)

	output := ""
	var err error
	for i := 0; i < *numberOfNames; i++ {

		switch *mode {
		case "full":
			output, err = generator.CompleteName(*gender)
		case "firstname":
			output, err = generator.FirstName(*gender)
		case "lastname":
			output, err = generator.LastName()
		default:
			output = fmt.Sprintf("Unsupported generation mode %s, supported modes: full, firstname, lastname", *mode)
		}

		if err != nil && errors.Is(err, namegen.ErrorEmptyItems) {
			output = fmt.Sprintf("Unsupported origin %s\nUse \"namegen -l\" to see supported origins", *origin)
		}

		fmt.Println(output)
	}

}
