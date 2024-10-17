# Pokemon colorscripts

A script to print out images of pokemon to the terminal. Inspired by
[DT's colorscripts compilation](https://gitlab.com/dwt1/shell-color-scripts)

## Table of contents
[[_TOC_]]

## Description
Prints out colored unicode sprites of pokemon onto your terminal. Contains
almost 900 pokemon from gen 1 to gen 8. Has all the pokemons you could ever want, including shinies
and alternate forms (okay it's still missing some female forms but cut me some slack).

## Visuals
### Demo GIFs
Demo of the program being used

![demo of program in action](./demo_images/pokemon-colorscripts-demo.gif)

Demo of the program running on terminal startup.

![demo of random pokemons on terminal spawn](./demo_images/terminal-startup-demo.gif)

### Screenshots
![screenshot](./demo_images/pokemon-colorscripts-screenshot-1.png)
![screenshot](./demo_images/pokemon-colorscripts-screenshot-2.png)
![screenshot](./demo_images/pokemon-colorscripts-screenshot-3.png)

## Requirements
The program requires `python3` to run and a terminal with true color support,
which most modern terminals have. More on terminals and color support can be found in
[this gist](https://gist.github.com/XVilka/8346728)

## Installation

### On Arch or Arch based distros
The utility is available as an AUR package
[pokemon-colorscripts-git](https://aur.archlinux.org/packages/pokemon-colorscripts-git).
You can install this manually or using an AUR helper. For manual installation,
download the PKGBUILD file from this repository. Then run
```
makepkg -cf
```
which will create a .pkg.tar.zst file. Then run
```
sudo pacman -U <filename.pkg.tar.zst>
```

Or you can use an AUR Helper such as yay and run
```
yay -S pokemon-colorscripts-git
```

### On other distros and MacOS
Clone or download the repository
```
git clone https://gitlab.com/phoneybadger/pokemon-colorscripts.git
```
`cd` into the directory and run the install script
```
cd pokemon-colorscripts
sudo ./install.sh
```
Now the program should be installed. You can check this by running
```
pokemon-colorscripts
```
Which should print out the help page for the program

### Uninstall
To uninstall the program you can run the uninstall script. Make the script
executable if it already isn't
```
chmod +x ./uninstall.sh
```
and run the script
```
sudo ./uninstall.sh
```

## Usage
You can run the program from the command line to either display a pokemon of your
choice by specifying the pokemon name or make it display a random pokemon.
```
usage: pokemon-colorscripts [OPTION] [POKEMON NAME]

CLI utility to print out unicode image of a pokemon in your shell

optional arguments:
  -h, --help            Show this help message and exit
  -l, --list            Print list of all pokemon
  -n NAME, --name NAME  Select pokemon by name. Generally spelled like in the games. a few exceptions are nidoran-f, nidoran-m, mr-mime, farfetchd,
                        flabebe type-null etc. Perhaps grep the output of --list if in doubt.
  -f FORM, --form FORM  Show an alternate form of a pokemon
  --no-title            Do not display pokemon name
  -s, --shiny           Show the shiny version of the pokemon instead
  -b, --big             Show a larger version of the sprite
  -r [RANDOM], --random [RANDOM]
                        Show a random pokemon. This flag can optionally be followed by a generation number or range (1-8) to show random pokemon
                        from a specific generation or range of generations. The generations can be provided as a continuous range (eg. 1-3) or as a
                        list of generations (1,3,6)
```

Example of printing out a specific pokemon
```
pokemon-colorscripts -n charizard
```
Example of printing out a specific shiny pokemon
```
pokemon-colorscripts -n spheal -s
```
Example of printing out a random pokemon
```
pokemon-colorscripts -r
```
Example of printing out a random pokemon from generation 1
```
pokemon-colorscripts -r 1
```
Example of printing out a random pokemon from generations 1-3
```
pokemon-colorscripts -r 1-3
```
Example of printing out a random pokemon from generations 1,3 and 6
```
pokemon-colorscripts -r 1,3,6
```
Example of printing out a larger sprite
```
pokemon-colorscripts -n sphleal -b
```
Example of printing out an alternate form
```
pokemon-colorscripts -n deoxys --form defense
```
Some pokemon with spaces or periods or other special characters in their name
might not be spelled as expected, some examples include:
```
farfetch'd -> farfetchd
mr.mime -> mr-mime
```
These are rare exceptions, and if required you can parse the `--list` page to see
the names of all the pokemon.

### Running on terminal startup
#### On Bash and ZSH
You can display a random pokemon whenever a terminal gets launched by adding
the `pokemon-colorscripts -r` command to your *.bashrc* or .*zshrc*.

#### On Fish
If you have fish as your user shell you can display a random pokemon on terminal
startup by overriding the `fish_greeting` in your `config.fish`
```
function fish_greeting
     pokemon-colorscripts -r
end
```
A more advaced setup combining multiple colorscripts can be found on pokemon-colorscripts#2

## Location of the files
The program is located at usr/local/opt/pokemon-colorscripts/ with the script being symlinked to /usr/local/bin/

## How it works
The program itself is a simple python script that prints out text files corresponding
to the relevant pokemon or a randomly selected pokemon. The sprites are simple text
files that use unicode characters and ANSI color codes to display images of pokemon.
The text sprite files can be found in the *colorscripts* folder. The files were generated
using sprites taken from [PokéSprite](https://msikma.github.io/pokesprite/) as the
base and converted to unicode sprites using custom scripts that can be found in
[this repo](https://gitlab.com/phoneybadger/pokemon-generator-scripts)

## Similar projects
`pokemon-colorscripts` is not the exclusive nor the first program to print out
pokemon sprites to the terminal. You can check out these really cool projects as
well
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
Phoney badger:
https://gitlab.com/phoneybadger

## License
The MIT License (MIT)


