# removeDuplciate

A CLI to remove duplicates from a file, outputing a new file.

## Requirements

[Go 1.20+](https://go.dev/dl/)

## Build

    go build

## Features

Available flags to use it on this CLI.

### filename

Defines a path to the file that will be parsed. The resulting file will be put on the same directory as the original file with the "_new" suffix.

    -filename example.txt

## Usage

Removing duplicates:

    removeDuplicate -filename example.txt