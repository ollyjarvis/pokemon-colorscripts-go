#!/bin/sh

pkgname=pokemon-colorscripts-go

mkdir -p "/usr/share/$pkgname"

cp -rf colorscripts "/usr/share/$pkgname"
cp pokemon.json "/usr/share/$pkgname"

go install -ldflags "-X main.PROGRAM_DIR=/usr/share/$pkgname" . 
