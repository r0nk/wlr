package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: wlr [<input>] [-file=<filename>]")
	fmt.Println("Examples:")
	fmt.Println("cat wordlist.txt | wlr \"replace: FUZZ\" ")
}

func main() {
	var input string
	var file string

	flag.StringVar(&file, "file", "", "File containing replacement strings")

	flag.Parse()

	// Get non-flag arguments
	args := flag.Args()

	if len(args) > 0 {
		input = args[0]
	} else {
		usage()
		return
	}

	var lines []string

	// Read lines from file or stdin
	if file != "" {
		var err error
		lines, err = readLines(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	// Replace 'FUZZ' and output the modified strings
	for _, line := range lines {
		modifiedString := strings.Replace(input, "FUZZ", line, -1)
		fmt.Println(modifiedString)
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
