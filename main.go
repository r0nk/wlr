package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: wlr <input> [filename]")
	fmt.Println("Examples:")
	fmt.Println("cat wordlist.txt | wlr \"replace: FUZZ\" ")
}

func main() {
	var file string

	flag.StringVar(&file, "file", "", "File containing replacement strings")
	//TODO attack modes
	//clusterbomb
	//pitchfork
	//sniper

	flag.Parse()

	args := flag.Args()

	if len(args) <= 0 {
		usage()
		return
	}

	input := args[0]

	for i, str := range args {
		if i == 0 { // skip the string to replace
			if len(args) == 1 {
				str = "" //read from stdin instead
			} else {
				continue
			}
		}
		parts := strings.Split(str, ":")
		file := parts[0]
		placeholder := "FUZZ"
		if len(parts) == 2 {
			placeholder = parts[1]
		}
		var lines []string
		var err error
		lines, err = readLines(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		for _, line := range lines {
			modified := strings.Replace(input, placeholder, line, -1)
			//sniper
			fmt.Println(modified)
		}
	}

}

func readLines(filename string) ([]string, error) {
	var lines []string
	if filename == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return lines, scanner.Err()
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
