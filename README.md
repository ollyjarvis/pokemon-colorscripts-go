# pokemon-colorscripts-go

A script to print out images of pokemon to the terminal. Inspired by
[DT's colorscripts compilation](https://gitlab.com/dwt1/shell-color-scripts).

Rewritten in Golang so it's a bit faster.

## Description
Prints out colored unicode sprites of pokemon onto your terminal.

Contains almost 900 pokemon from gen 1 to gen 8. Has all the pokemons you could ever want, including shinies and alternate forms (still missing some female forms)

## Visuals
### Demo GIFs

not done lols

### Screenshots

not done lols

## Requirements
The program requires a terminal with true color support,
which most modern terminals have. More on terminals and color support can be found in
[this gist](https://gist.github.com/XVilka/8346728)

## Installation

### Arch Linux
- Clone the repo
```makepkg -si```

### From source
#### Requirements
- Have Go installed and in PATH

#### Install
- Clone the repo
```chmod +x install.sh```
```./install.sh```

#### Uninstall
```chmod +x uninstall.sh```
```./uninstall.sh```

## Usage
You can run the program from the command line to randomly display a pokemon
or by specifying the pokemon name with `-n [POKEMON_NAME]`.

Slightly different than the original pokemon-colorscripts.
- `-r` flag is no longer needed unless specifiying generations to pick from.
- `-a` all flag instead of `-l` list
- `-l` large flag instead of `-b` big

```
Usage:
  pokemon-colorscripts-go [OPTIONS]

Application Options:
  -h, --help      Show this help message and exit
  -a, --all       Print list of all pokemon
  -n, --name=     Select pokemon by name. Generally spelled like in the games.
                  A few exceptions are nidoran-f, nidoran-m, mr-mime, farfetchd, flabebe
                  type-null etc.
                  Perhaps grep the output of --list if in doubt.
  -f, --form=     Show an alternative form of a pokemon
      --no-title  Do not display pokemon name
  -s, --shiny     Show the shiny version of the pokemon instead
  -l, --large     Show a larger version of the sprite
  -r, --random=   Specify a generations for to be randomly selected from.
                  Example usage:
                  Range: 2-6
                  List: 1,3,6
                  Number: 5 (default: 1-8)
```

Example of printing out a specific pokemon
```
pokemon-colorscripts-go -n charizard
```
Example of printing out a specific shiny pokemon
```
pokemon-colorscripts-go -n spheal -s
```
Example of printing out a random pokemon
```
pokemon-colorscripts-go
```
Example of printing out a random pokemon from generation 1
```
pokemon-colorscripts-go -r 1
```
Example of printing out a random pokemon from generations 1-3
```
pokemon-colorscripts-go -r 1-3
```
Example of printing out a random pokemon from generations 1,3 and 6
```
pokemon-colorscripts-go -r 1,3,6
```
Example of printing out a larger sprite
```
pokemon-colorscripts-go -n sphleal -l
```
Example of printing out an alternate form
```
pokemon-colorscripts-go -n deoxys --form defense
```
Some pokemon with spaces or periods or other special characters in their name
might not be spelled as expected, some examples include:
```
farfetch'd -> farfetchd
mr.mime -> mr-mime
```
These are rare exceptions, and if required you can parse the `--all` page to see
the names of all the pokemon.

### Running on terminal startup
#### On Bash and ZSH
You can display a random pokemon whenever a terminal gets launched by adding
the `pokemon-colorscripts` command to your *.bashrc* or *.zshrc*.

#### On Fish
If you have fish as your user shell you can display a random pokemon on terminal
startup by overriding the `fish_greeting` in your `config.fish`
```
function fish_greeting
     pokemon-colorscripts
end
```

## Location of the files

not done lols

## How it works
The program itself is a go binary that prints out text files corresponding
to the relevant pokemon or a randomly selected pokemon.

The sprites are simple text files that use unicode characters and ANSI color codes to display images of pokemon.

The text sprite files can be found in the *colorscripts* folder. The files were generated using sprites taken from [PokéSprite](https://msikma.github.io/pokesprite/) as the
base and converted to unicode sprites using custom scripts that can be found in
[this repo](https://gitlab.com/phoneybadger/pokemon-generator-scripts)

## Similar projects
`pokemon-colorscripts-go` is a rewrite of pokemon-colorscripts and is not the
exclusive nor the first program to print out pokemon sprites to the terminal.
You can check out these really cool projects as well
- [pokemon-colorscripts](https://gitlab.com/phoneybadger/pokemon-colorscripts) (python version)
- [pokeget](https://github.com/talwat/pokeget)
- [pokeshell](https://github.com/acxz/pokeshell)
- [krabby](https://github.com/yannjor/krabby)

An older more minimal shell script version of the project is now being maintained
as

- [pokescript](https://github.com/acxz/pokescript)

a much more comprehensive comparison of the different projects can be found on
the [Pokeshell Readme](https://github.com/acxz/pokeshell#similar-projects)

## Credits
- All the pokemon designs, names, branding etc. are trademarks of [The Pokémon Company](https://pokemon.com)
- The box art sprites were taken from the amazing [PokéSprite](https://msikma.github.io/pokesprite/) database

## Author
ollyjarvis
https://github.com/ollyjarvis

## Credit
Phoney badger:
https://gitlab.com/phoneybadger

## License
The MIT License (MIT)


