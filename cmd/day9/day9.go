package main

import (
	"adventofcode/year2025/cmd/util"
	"strings"
)

type redTile struct {
	x int
	y int
}

func parseInput(input string) []redTile {
	tiles := []redTile{}
	line_iter := strings.SplitSeq(input, "\n")
	for line := range line_iter {
		numbers := strings.Split(line, ",")
		if len(numbers) != 2 {
			panic(numbers)
		}
		x := util.String2Int(numbers[0])
		y := util.String2Int(numbers[1])
		tiles = append(tiles, redTile{x, y})
	}
	return tiles
}
