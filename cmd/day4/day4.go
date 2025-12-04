package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
	"log"
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

func isForkliftAccessible(grid Grid, x int, y int) bool {
	papers := 0
	if !isPaper(grid, x, y) {
		panic("This function is expected to be run on a paper place")
	}
	if isPaper(grid, x-1, y-1) {
		papers++
	}
	if isPaper(grid, x+0, y-1) {
		papers++
	}
	if isPaper(grid, x+1, y-1) {
		papers++
	}

	if isPaper(grid, x-1, y+0) {
		papers++
	}

	if isPaper(grid, x+1, y+0) {
		papers++
	}

	if isPaper(grid, x-1, y+1) {
		papers++
	}
	if isPaper(grid, x+0, y+1) {
		papers++
	}
	if isPaper(grid, x+1, y+1) {
		papers++
	}
	return papers < 4
}

func countForkliftAccessible(grid Grid) int {
	counter := 0
	for y := range grid {
		gridline := grid[y]
		for x := range gridline {
			if gridline[x] {
				if isForkliftAccessible(grid, x, y) {
					counter++
					log.Printf("Fork lift accessible at %v,%v", x, y)
				}
			}
		}
	}
	return counter
}

func main() {
	grid := parseInput(util.LoadInput(4))
	counter := countForkliftAccessible(grid)
	fmt.Println(counter)
}
