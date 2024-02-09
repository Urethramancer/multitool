# Multitool
A grab-bag of functionality for the command line.

## What's this for?
I made this to run on routers or other tiny devices with limited CLI tools and options for installing such. Compatible with TinyGo, but access to other packages may be a limitation.

## How to install
First download the package:
```sh
go get -u github.com/Urethramancer/multitool
```

Then run the command to generate its symlinks:
```sh
sudo multitool
```

## How to use
The currently implemented commands are `pwgen`, `rn`, `md5`, `sha1` and `sha512`. Run each to get usage.
