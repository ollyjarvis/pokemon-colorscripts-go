package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/jessevdk/go-flags"
)

var (
	PROGRAM, _       = os.Executable()
	PROGRAM_DIR      = filepath.Dir(PROGRAM)
	COLORSCRIPTS_DIR = fmt.Sprintf("/usr/share/%s/colorscripts", filepath.Base(PROGRAM))
	GENERATIONS      = [][]int{{1, 151}, {152, 251}, {252, 386}, {387, 493}, {494, 649}, {650, 721}, {722, 809}, {810, 898}}
)

const (
	SHINY_RATE = 1 / 128
)

func print_file(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	fmt.Print(string(data))

	return nil
}

func list_pokemon_names() error {
	data, err := os.ReadFile(fmt.Sprintf("%s/pokemon.json", PROGRAM_DIR))
	if err != nil {
		return err
	}

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		name, _, _, _ := jsonparser.Get(value, "name")
		fmt.Println(string(name))
	})

	return nil
}

func show_pokemon_by_name(name string, show_title bool, shiny bool, is_large bool, form string) error {
	dir := COLORSCRIPTS_DIR

	if is_large {
		dir = fmt.Sprintf("%s/large", dir)
	} else {
		dir = fmt.Sprintf("%s/small", dir)
	}
	if shiny {
		dir = fmt.Sprintf("%s/shiny", dir)
	} else {
		dir = fmt.Sprintf("%s/regular", dir)
	}

	if form != "" {
		var alt_forms []string
		found := false

		data, err := os.ReadFile(fmt.Sprintf("%s/pokemon.json", PROGRAM_DIR))
		if err != nil {
			return err
		}

		jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			cur_name, _ := jsonparser.GetString(value, "name")
			if strings.ToLower(cur_name) == strings.ToLower(name) {
				jsonparser.ArrayEach(value, func(formValue []byte, dataType jsonparser.ValueType, offset int, err error) {
					cur_form := string(formValue)
					if cur_form != "regular" {
						if cur_form == form {
							found = true
						}
						alt_forms = append(alt_forms, cur_form)
					}
				}, "forms")
			}
		})

		if found {
			name = fmt.Sprintf("%s-%s", name, form)
		} else {
			fmt.Printf("Invalid form '%s' for pokemon %s\n", form, name)
			if len(alt_forms) == 0 {
				fmt.Printf("No alternative forms available for %s", name)
			} else {
				fmt.Println("Available alternate forms are")
				for _, alt_form := range alt_forms {
					fmt.Printf("- %s", alt_form)
				}
			}
			os.Exit(0)
		}
	}

	file := fmt.Sprintf("%s/%s", dir, name)

	if show_title {
		if shiny {
			name = fmt.Sprintf("%s (shiny)", name)
		}
		fmt.Println(name)
	}

	print_file(file)

	return nil
}

func show_random_pokemon(generations string, show_title bool, shiny bool, is_large bool) error {
	var gen int
	var err error

	if len(strings.Split(generations, ",")) > 1 {
		gens := strings.Split(generations, ",")
		gen_str := gens[rand.IntN(len(gens))]
		gen, err = strconv.Atoi(gen_str)
		if err != nil {
			fmt.Printf("Invalid generation '%s'\n", gen_str)
			return err
		}
	} else if len(strings.Split(generations, "-")) > 1 {
		gens := strings.Split(generations, "-")
		gen_str := gens[rand.IntN(len(gens))]
		gen, err = strconv.Atoi(gen_str)
		if err != nil {
			fmt.Printf("Invalid generation '%s'\n", gen_str)
			return err
		}
	} else if generations != "" {
		gen, err = strconv.Atoi(generations)
		if err != nil {
			fmt.Printf("Invalid generation '%s'\n", generations)
			return err
		}
	} else {
		gen = rand.IntN(7) + 1
	}

	if gen <= 0 || gen > 8 {
		fmt.Printf("Invalid generation '%s'\n", gen)
		return fmt.Errorf("generation out of range")
	}

	pokemon_index := rand.N(GENERATIONS[gen-1][1]-GENERATIONS[gen-1][0]) + GENERATIONS[gen-1][0]

	data, err := os.ReadFile(fmt.Sprintf("%s/pokemon.json", PROGRAM_DIR))
	if err != nil {
		return err
	}

	name, _ := jsonparser.GetString(data, fmt.Sprintf("[%d]", pokemon_index), "name")

	if !shiny {
		shiny = rand.Float32() <= SHINY_RATE
	}

	show_pokemon_by_name(name, show_title, shiny, is_large, "")

	return nil
}

func main() {
	var opts struct {
		Help bool `short:"h" long:"help" description:"Show this help message and exit"`

		All bool `short:"a" long:"all" description:"Print list of all pokemon"`

		Name string `short:"n" long:"name" description:"Select pokemon by name. Generally spelled like in the games.\nA few exceptions are nidoran-f, nidoran-m, mr-mime, farfetchd, flabebe type-null etc.\nPerhaps grep the output of --list if in doubt."`

		Form string `short:"f" long:"form" description:"Show an alternative form of a pokemon"`

		Title bool `long:"no-title" description:"Do not display pokemon name"`

		Shiny bool `short:"s" long:"shiny" description:"Show the shiny version of the pokemon instead"`

		Large bool `short:"l" long:"large" description:"Show a larger version of the sprite"`

		Random string `short:"r" long:"random" default:"1-8" description:"Specify a generations for to be randomly selected from.\nExample usage:\n  Range: 2-6\n  List: 1,3,6\n  Number: 5"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	if opts.All {
		err = list_pokemon_names()
	} else if opts.Name != "" {
		err = show_pokemon_by_name(opts.Name, !opts.Title, opts.Shiny, opts.Large, opts.Form)
	} else {
		err = show_random_pokemon(opts.Random, !opts.Title, opts.Shiny, opts.Large)
	}
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
