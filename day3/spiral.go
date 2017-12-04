package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	input := flag.Int("i", 1, "Define input to compute steps")
	version := flag.Int("v", 1, "Define a version")
	flag.Parse()

	switch *version {
	case 1:
		x, y := getUlamSpiralCoordinates(*input)
		steps := int(math.Abs(x) + math.Abs(y))
		fmt.Printf("Steps: %d", steps)
		return
	case 2:
		sum := findPointWithGreaterSummValue(*input)
		fmt.Printf("Greater value: %d", sum)
	}

}

// Algorithm was taken from https://math.stackexchange.com/a/1707796
func getUlamSpiralCoordinates(n int) (float64, float64) {
	r := math.Sqrt(float64(n))
	m := math.Mod(r, 1.)
	var p float64
	if math.Mod(r*.5, 1.) > .5 {
		p = 1.
	} else {
		p = -1.
	}

	s := p*1.5 - m*p*2.
	var x float64
	var y float64

	if m < .5 {
		x = r * .5 * p
		y = r*p - r*s
	} else {
		x = r * s
		y = r * .5 * p
	}

	return x, y
}

// probably weird solution with computation and memory overhead
func findPointWithGreaterSummValue(input int) int {
	search := true
	values := make(map[string]int)
	values["0_0"] = 1

	var x int
	var y int

	prevY := 0
	for r := 1; search; r++ {
		prevX := r
		for direction := 0; direction < 4; direction++ {
			x = prevX
			y = prevY
			for ; isInside(x, y, r); x, y = getNextCartesianCoordiantes(x, y, direction) {
				if _, ok := values[fmt.Sprintf("%d_%d", x, y)]; ok {
					continue
				}
				s := getNextSumm(x, y, direction, values)
				if s > input {
					search = false
					return s
				}

				values[fmt.Sprintf("%d_%d", x, y)] = s
				prevX = x
				prevY = y
			}
		}
	}
	return 0
}

func getNextSumm(x int, y int, direction int, grid map[string]int) int {
	sum := 0
	for _, n := range getNeighbors(x, y, direction) {
		if v, ok := grid[fmt.Sprintf("%d_%d", n.x, n.y)]; ok {
			sum += v
		}
	}
	return sum
}

type point struct {
	x int
	y int
}

//directions:
// 0 is for up
// 1 is for right
// 2 is for down
// 3 is for left
func getNeighbors(px int, py int, direction int) []point {
	switch direction {
	case 0:
		return []point{
			point{x: px, y: py - 1},
			point{x: px - 1, y: py},
			point{x: px - 1, y: py - 1},
			point{x: px - 1, y: py + 1},
		}
	case 1:
		return []point{
			point{x: px + 1, y: py},
			point{x: px, y: py - 1},
			point{x: px + 1, y: py - 1},
			point{x: px - 1, y: py - 1},
		}
	case 2:
		return []point{
			point{x: px, y: py + 1},
			point{x: px + 1, y: py},
			point{x: px + 1, y: py + 1},
			point{x: px + 1, y: py - 1},
		}
	case 3:
		return []point{
			point{x: px - 1, y: py},
			point{x: px, y: py + 1},
			point{x: px - 1, y: py + 1},
			point{x: px + 1, y: py + 1},
		}
	default:
		return nil
	}
}

func isInside(x int, y int, r int) bool {
	absX := int(math.Abs(float64(x)))
	absY := int(math.Abs(float64(y)))
	return absX <= r && absY <= r
}

//directions:
// 0 is for up
// 1 is for right
// 2 is for down
// 3 is for left
func getNextCartesianCoordiantes(x int, y int, direction int) (int, int) {
	switch direction {
	case 0:
		return x, y + 1
	case 1:
		return x - 1, y
	case 2:
		return x, y - 1
	case 3:
		return x + 1, y
	default:
		return x, y
	}
}
