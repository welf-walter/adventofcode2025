package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
	"log"
	"slices"
	"strings"
)

type splitterRow []int

// how many ways to get there?
type timeLine map[int]int

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

func runRow(tachyons []int, splitters splitterRow) (outTachyons []int, splitCount int) {
	for _, t := range tachyons {
		if slices.Contains(splitters, t) {
			splitCount++
			if !slices.Contains(outTachyons, t-1) {
				outTachyons = append(outTachyons, t-1)
			}
			if !slices.Contains(outTachyons, t+1) {
				outTachyons = append(outTachyons, t+1)
			}
		} else {
			if !slices.Contains(outTachyons, t) {
				outTachyons = append(outTachyons, t)
			}
		}
	}
	return
}

func runRow2(ti timeLine, splitters splitterRow) timeLine {
	out := timeLine{}
	for tachyon, ways := range ti {
		if slices.Contains(splitters, tachyon) {
			out[tachyon-1] += ways
			out[tachyon+1] += ways
		} else {
			out[tachyon] += ways
		}
	}
	return out
}

func runRows(start int, splitterRows []splitterRow) int {
	sumSplits := 0
	t := []int{start}
	var splits int
	for _, row := range splitterRows {
		t, splits = runRow(t, row)
		log.Printf("Tachyons at %v, %v splits\n", t, splits)
		sumSplits += splits
	}
	return sumSplits
}

func runRows2(start int, splitterRows []splitterRow) int {
	t := timeLine{start: 1}
	for _, row := range splitterRows {
		t = runRow2(t, row)
		log.Printf("Timeline = %v\n", t)
	}

	sumSplits := 0
	for _, count := range t {
		sumSplits += count
	}
	return sumSplits
}

func main() {
	input := util.LoadInput(7)
	start, rows := parseInput(input)
	fmt.Println(runRows(start, rows))
	fmt.Println(runRows2(start, rows))
}
