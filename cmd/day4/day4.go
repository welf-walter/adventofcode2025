package main

import (
	"fmt"
	"strings"
)

type Grid [][]bool

func parseInput(input string) Grid {
	grid := Grid{}
	line_iter := strings.SplitSeq(input, "\n")
	for line := range line_iter {
		gridline := []bool{}
		for _, c := range line {
			switch c {
			case '.':
				gridline = append(gridline, false)
			case '@':
				gridline = append(gridline, true)
			default:
				panic(fmt.Sprintf("Unexpected character %v in %v", c, line))
			}
		}
		grid = append(grid, gridline)
	}
	return grid
}

func isPaper(grid Grid, x int, y int) bool {
	if y < 0 || y >= len(grid) {
		return false
	}
	gridline := grid[y]
	if x < 0 || x >= len(gridline) {
		return false
	}
	return gridline[x]
}

func main() {
	fmt.Println("Not yet implemented")
}
