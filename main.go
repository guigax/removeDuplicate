package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func unique(stringSlice []string, remove bool) []string {
	list := []string{}
	for _, entry := range stringSlice {
		if Index(list, entry) == -1 {
			list = append(list, entry)
		} else if remove {
			list = Remove(list, Index(list, entry))
		}
	}
	return list
}

func Remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func saveLines(filePath string, values []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
	return nil
}

func main() {
	var (
		filename  = flag.String("filename", "", "file that will be parsed")
		removeAll = flag.Bool("removeAll", false, "if true, it will remove all the duplicate occurrences")
	)
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Problem reading %v: %v\n", *filename, err)
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse data from %v: %v\n", *filename, err)
	}

	uniqueLines := unique(lines, *removeAll)

	newFilename := strings.TrimSuffix(*filename, filepath.Ext(*filename)) + "_new" + filepath.Ext(*filename)
	e := saveLines(newFilename, uniqueLines)
	if e != nil {
		log.Fatalf("Problem saving %v: %v\n", *filename, err)
	}

	f.Close()
}
