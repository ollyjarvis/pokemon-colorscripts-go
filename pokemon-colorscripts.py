#!/usr/bin/env python3

import argparse
import json
import os
import random
import sys

PROGRAM = os.path.realpath(__file__)
PROGRAM_DIR = os.path.dirname(PROGRAM)
COLORSCRIPTS_DIR = f"{PROGRAM_DIR}/colorscripts"

REGULAR_SUBDIR = "regular"
SHINY_SUBDIR = "shiny"

LARGE_SUBDIR = "large"
SMALL_SUBDIR = "small"

SHINY_RATE = 1 / 128
GENERATIONS = {
    "1": (1, 151),
    "2": (152, 251),
    "3": (252, 386),
    "4": (387, 493),
    "5": (494, 649),
    "6": (650, 721),
    "7": (722, 809),
    "8": (810, 898),
}


def print_file(filepath: str) -> None:
    with open(filepath, "r") as f:
        print(f.read())


def list_pokemon_names() -> None:
    with open(f"{PROGRAM_DIR}/pokemon.json") as file:
        pokemon_json = json.load(file)
        for pokemon in pokemon_json:
            print(pokemon["name"])


def show_pokemon_by_name(
    name: str, show_title: bool, shiny: bool, is_large: bool, form: str = ""
) -> None:
    base_path = COLORSCRIPTS_DIR
    color_subdir = SHINY_SUBDIR if shiny else REGULAR_SUBDIR
    # default to smaller size as this makes sense for most font size + terminal
    # size combinations
    size_subdir = LARGE_SUBDIR if is_large else SMALL_SUBDIR
    with open(f"{PROGRAM_DIR}/pokemon.json") as file:
        pokemon_json = json.load(file)
        pokemon_names = {pokemon["name"] for pokemon in pokemon_json}
        if name not in pokemon_names:
            print(f"Invalid pokemon {name}")
            sys.exit(1)

        if form:
            for pokemon in pokemon_json:
                if pokemon["name"] == name:
                    forms = pokemon["forms"]
                    alternate_forms = [f for f in forms if f != "regular"]
            if form in alternate_forms:
                name += f"-{form}"
            else:
                print(f"Invalid form '{form}' for pokemon {name}")
                if not alternate_forms:
                    print(f"No alternate forms available for {name}")
                else:
                    print(f"Available alternate forms are")
                    for form in alternate_forms:
                        print(f"- {form}")
                sys.exit(1)
    pokemon_file = f"{base_path}/{size_subdir}/{color_subdir}/{name}"
    if show_title:
        if shiny:
            print(f"{name} (shiny)")
        else:
            print(name)
    print_file(pokemon_file)


def show_random_pokemon(
    generations: str, show_title: bool, shiny: bool, is_large: bool
) -> None:
    # Generation list
    if len(generations.split(",")) > 1:
        input_gens = generations.split(",")
        start_gen = random.choice(input_gens)
        end_gen = start_gen
    # Generation range
    elif len(generations.split("-")) > 1:
        start_gen, end_gen = generations.split("-")
    # Single generation
    else:
        start_gen = generations
        end_gen = start_gen

    with open(f"{PROGRAM_DIR}/pokemon.json", "r") as file:
        pokemon = [pokemon["name"] for pokemon in json.load(file)]
    try:
        start_idx = GENERATIONS[start_gen][0]
        end_idx = GENERATIONS[end_gen][1]
        random_idx = random.randint(start_idx, end_idx)
        random_pokemon = pokemon[random_idx - 1]
        # if the shiny flag is not passed, set a small random chance for the
        # pokemon to be shiny. If the flag is set, always show shiny
        if not shiny:
            shiny = random.random() <= SHINY_RATE
        show_pokemon_by_name(random_pokemon, show_title, shiny, is_large)
    except KeyError:
        print(f"Invalid generation '{generations}'")
        sys.exit(1)


def main() -> None:
    parser = argparse.ArgumentParser(
        prog="pokemon-colorscripts",
        description="CLI utility to print out unicode image of a pokemon in your shell",
        usage="pokemon-colorscripts [OPTION] [POKEMON NAME]",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        add_help=False,
    )

    parser.add_argument(
        "-h", "--help", action="help", help="Show this help message and exit"
    )
    parser.add_argument(
        "-l", "--list", help="Print list of all pokemon", action="store_true"
    )
    parser.add_argument(
        "-n",
        "--name",
        type=str,
        help="""Select pokemon by name. Generally spelled like in the games.
                a few exceptions are nidoran-f, nidoran-m, mr-mime, farfetchd, flabebe
                type-null etc. Perhaps grep the output of --list if in
                doubt.""",
    )
    parser.add_argument(
        "-f",
        "--form",
        type=str,
        help="Show an alternate form of a pokemon",
    )
    parser.add_argument(
        "--no-title", action="store_false", help="Do not display pokemon name"
    )
    parser.add_argument(
        "-s",
        "--shiny",
        action="store_true",
        help="Show the shiny version of the pokemon instead",
    )
    # ideally this argument should be --large, but using --big as -l is already
    # taken
    parser.add_argument(
        "-b",
        "--big",
        action="store_true",
        help="Show a larger version of the sprite",
    )
    parser.add_argument(
        "-r",
        "--random",
        type=str,
        const="1-8",
        nargs="?",
        help="""Show a random pokemon. This flag can optionally be
                followed by a generation number or range (1-8) to show random
                pokemon from a specific generation or range of generations.
                The generations can be provided as a continuous range (eg. 1-3)
                or as a list of generations (1,3,6)""",
    )

    args = parser.parse_args()

    if args.list:
        list_pokemon_names()
    elif args.name:
        show_pokemon_by_name(args.name, args.no_title, args.shiny, args.big, args.form)
    elif args.random:
        if args.form:
            print("--form flag unexpected with --random")
            sys.exit(1)
        show_random_pokemon(args.random, args.no_title, args.shiny, args.big)
    else:
        parser.print_help()


if __name__ == "__main__":
    main()
