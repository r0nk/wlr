package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type fp_pair struct {
	file, placeholder string
}

func usage() {
	fmt.Println("Usage: wlr <input> [<filename>:[placeholder]]...")
	fmt.Println("Examples:")
	fmt.Println("cat wordlist.txt | wlr \"replace: FUZZ\" ")
	fmt.Println("wlr \"FIRST SECOND\" ./test/wordlist.txt:FIRST ./test/wordlist2.txt:SECOND")
}

func clusterbomb(input string, fp_pairs []fp_pair, recursions int) []string {
	var ret []string
	var lines []string

	if len(fp_pairs) <= recursions {
		ret = append(ret, input)
		return ret
	}
	fp := fp_pairs[recursions]
	lines, err := read_lines(fp.file)

	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		modified := strings.Replace(input, fp.placeholder, line, -1)
		//	fmt.Printf("mod,input,line: %s,%s,%s \n", modified, input, line)

		r := clusterbomb(modified, fp_pairs, recursions+1)
		for _, v := range r {
			ret = append(ret, v)
		}
	}
	return ret
}

func pitchfork(input string, fp_pairs []fp_pair) []string {
	var ret []string
	var lines [][]string
	for _, fp := range fp_pairs {
		l, err := read_lines(fp.file)
		fmt.Printf("reading %d lines from file \"%s\"\n", len(l), fp.file)
		lines = append(lines, l)

		if err != nil {
			panic(err)
		}
	}

	for rows, _ := range lines[0] {
		modified := input
		fmt.Printf("rows:%d\n", rows)

		for cols, fp := range fp_pairs {
			fmt.Printf("cols:%d\n", cols)
			if rows >= len(lines[cols]) {
				return ret
			}
			modified = strings.Replace(modified, fp.placeholder, lines[cols][rows], -1)
		}
		ret = append(ret, modified)
	}
	return ret
}

func get_file_replacement_pairs(args []string) []fp_pair {
	var fpp fp_pair
	var ret []fp_pair
	for i, str := range args {
		if i == 0 { // skip the string to replace
			if len(args) == 1 {
				str = "" //read from stdin instead
			} else {
				continue
			}
		}
		parts := strings.Split(str, ":")
		fpp.file = parts[0]
		fpp.placeholder = "FUZZ"
		if len(parts) == 2 {
			fpp.placeholder = parts[1]
		}
		ret = append(ret, fpp)
	}
	return ret
}

func read_lines(filename string) ([]string, error) {
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

func main() {
	var clusterbomb_flag bool
	var pitchfork_flag bool

	flag.BoolVar(&clusterbomb_flag, "clusterbomb", true, "Enable clusterbomb mode (1 1, 1 2, 2 1,2 2)")
	flag.BoolVar(&pitchfork_flag, "pitchfork", false, "Enable pitchfork mode (1 1, 2 2)")

	//TODO add an option to not recursively replace

	flag.Parse()

	args := flag.Args()

	if len(args) <= 0 {
		usage()
		return
	}

	fp_pairs := get_file_replacement_pairs(args)

	input := args[0]

	var r []string

	if clusterbomb_flag {
		r = clusterbomb(input, fp_pairs, 0)
	}

	if pitchfork_flag {
		r = pitchfork(input, fp_pairs)
	}

	for _, out := range r {
		fmt.Printf("%s\n", out)
	}
}
