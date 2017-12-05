package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file := flag.String("i", "", "Define input file")
	version := flag.Int("v", 1, "Define a version")
	flag.Parse()
	instructions := readInstructions(*file)
	fmt.Printf("Instructions: %d\n", len(instructions))
	switch *version {
	case 1:
		fmt.Printf("Steps: %d", countJumps(instructions, simpleIncrement))
		break
	case 2:
		fmt.Printf("Steps: %d", countJumps(instructions, incrementWithThresholdOf3))
		break
	}
}

type incrementFunc func(int) int

func readInstructions(input string) []int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var jumps []int
	for scanner.Scan() {
		line := scanner.Text()
		j, _ := strconv.Atoi(line)
		jumps = append(jumps, j)
	}

	return jumps
}

func countJumps(instructions []int, fn incrementFunc) int {
	steps := 0
	index := 0

	for index >= 0 && index < len(instructions) {
		jump := instructions[index]
		instructions[index] = fn(instructions[index])
		index += jump
		steps++
	}

	return steps
}

func simpleIncrement(prev int) int {
	return prev + 1
}

func incrementWithThresholdOf3(prev int) int {
	if prev >= 3 {
		return prev - 1
	}
	return prev + 1
}
