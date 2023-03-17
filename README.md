# removeDuplciate

A CLI to remove duplicates from a file, outputing a new file.

## Requirements

[Go 1.18+](https://go.dev/dl/)

## Build

    go build

## Features

Available flags to use it on this CLI.

### filename

Defines a path to the file that will be parsed. The resulting file will be put on the same directory as the original file with the "_new" suffix.

    -filename="example.txt"

### removeAll

If used, it removes all occurences of the duplicated lines, significantly impacts the performance (default `false`).

## Usage

Removing duplicates:

    removeDuplicate -filename="example.txt"

Removing every occurence of duplicates:

    removeDuplicate -filename="example.txt" - remove