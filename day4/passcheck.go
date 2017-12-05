package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file := flag.String("i", "", "Define input file to check passphrases")
	version := flag.Int("v", 1, "Define a version")
	flag.Parse()

	switch *version {
	case 1:
		fmt.Printf("Valid: %d", countUnique(*file, " "))
		break
	case 2:
		fmt.Printf("Valid: %d", countWithUniqueLetters(*file, " "))
		break
	}
}

func countUnique(input string, delim string) int {
	count := 0
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := make(map[string]bool)
		exist := true
		for _, s := range strings.Split(line, delim) {
			if _, exist = words[s]; exist {
				break
			}
			words[s] = true
		}
		if !exist {
			count++
		}
	}

	return count
}

func countWithUniqueLetters(input string, delim string) int {
	var validLines []int
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for lineNum := 0; scanner.Scan(); lineNum++ {
		line := scanner.Text()
		words := make(map[string]bool)
		exist := true
		for _, s := range strings.Split(line, delim) {
			letters := strings.Split(s, "")
			sort.Strings(letters)
			sorted := strings.Join(letters, "")
			if _, exist = words[sorted]; exist {
				break
			}
			words[sorted] = true
		}
		if !exist {
			validLines = append(validLines, lineNum)
		}
	}

	return len(validLines)
}
