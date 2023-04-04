package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filename := flag.String("file", "example.txt", "file that will be parsed")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	lines := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !lines[line] {
			lines[line] = true
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning file: %v\n", err)
		os.Exit(1)
	}
	// Get the directory and base name of the input file.
	dir := filepath.Dir(*filename)
	base := filepath.Base(*filename)

	// Append "_new" to the base name to get the output file name.
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext) + "_new" + ext
	outputPath := filepath.Join(dir, name)

	// Create the output file and write the lines to it.
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	fmt.Printf("Duplicate lines removed from %s, new file saved as %s\n", *filename, outputPath)
}
