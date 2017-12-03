package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sumGetter func([]int) int

func main() {
	file := flag.String("i", "", "Define input file to compute checksum")
	delim := flag.String("d", "	", "Define a deliminator for values in matrix")
	version := flag.Int("v", 1, "Define a version")
	flag.Parse()
	matrix := readMatrix(*file, *delim)
	summ := 0

	switch *version {
	case 1:
		summ = getCheckSumForRows(matrix, getMinMaxDiff)
		break
	case 2:
		summ = getCheckSumForRows(matrix, getDivisionCheck)
		break
	}

	fmt.Printf("CheckSum: %d", summ)
}

func readMatrix(input string, delim string) [][]int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matrix [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var arr []int
		for _, s := range strings.Split(line, delim) {
			num, _ := strconv.Atoi(s)
			arr = append(arr, num)
		}
		sort.Ints(arr)
		matrix = append(matrix, arr)
	}

	return matrix
}

func getCheckSumForRows(matrix [][]int, fn sumGetter) int {
	sum := 0
	for _, row := range matrix {
		sum += fn(row)
	}
	return sum
}

func getMinMaxDiff(row []int) int {
	max := row[len(row)-1]
	min := row[0]

	return int(math.Abs(float64(max - min)))
}

func getDivisionCheck(row []int) int {
	for di, divider := range row {
		for i := di + 1; i < len(row); i++ {
			if row[i]%divider == 0 {
				return row[i] / divider
			}
		}
	}
	return 0
}
