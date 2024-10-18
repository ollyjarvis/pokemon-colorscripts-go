#!/bin/sh

pkgname=pokemon-colorscripts-go

sudo rm -r "/usr/share/$pkgname"

rm "$(echo $GOPATH)/bin/$pkgname"
