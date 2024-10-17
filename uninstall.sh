#! /bin/sh

# A basic script to uninstall pokemon colorscripts

INSTALL_DIR='/usr/local/opt'
BIN_DIR='/usr/local/bin'

# Remove directories where files have been installed
rm -rf "$INSTALL_DIR/pokemon-colorscripts"
rm -rf "$BIN_DIR/pokemon-colorscripts"
