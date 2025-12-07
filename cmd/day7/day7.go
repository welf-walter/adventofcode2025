package main

import "strings"

type splitterRow []int

func parseInput(input string) (start int, splitterRows []splitterRow) {
	lines := strings.Split(input, "\n")
	start = strings.Index(lines[0], "S")
	for y := 1; y < len(lines); y++ {
		row := splitterRow{}
		for x, c := range lines[y] {
			if c == '^' {
				row = append(row, x)
			}
		}
		splitterRows = append(splitterRows, row)
	}
	return
}
