package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	line := flag.String("line", "", "Define line to compute captcha")
	version := flag.Int("v", 1, "Define a version")
	flag.Parse()
	chars := strings.Split(*line, "")
	summ := 0

	switch *version {
	case 1:
		summ = stepCheck(chars, 1)
		break
	case 2:
		summ = stepCheck(chars, len(chars)/2)
		break
	}

	fmt.Printf("V-%d captcha summ: %d", *version, summ)
}

func stepCheck(chars []string, step int) int {
	next := 0
	summ := 0
	for i, c := range chars {
		curr, _ := strconv.Atoi(c)
		ni := (i + step) % len(chars)
		next, _ = strconv.Atoi(chars[ni])
		if curr == next {
			summ += curr
		}
	}
	return summ
}
