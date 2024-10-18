#!/bin/sh

pkgname=pokemon-colorscripts-go

sudo mkdir -p "/usr/share/$pkgname"

sudo cp -rf colorscripts "/usr/share/$pkgname"
sudo cp pokemon.json "/usr/share/$pkgname"

go install -ldflags "-X main.PROGRAM_DIR=/usr/share/$pkgname" . 
