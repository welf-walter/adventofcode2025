package day12

import (
	"adventofcode/year2025/cmd/util"
	"strconv"
	"strings"
)

type shape [3][3]bool
type region struct {
	width      int
	height     int
	shapeCount []int
}

type puzzle struct {
	shapes  []shape
	regions []region
}

type regionMap [][]bool

// rotate right
func rotateShape(in shape) (out shape) {
	for y := range 3 {
		for x := range 3 {
			out[x][2-y] = in[y][x]
		}
	}
	return
}

func flipShape(in shape) (out shape) {
	for y := range 3 {
		for x := range 3 {
			out[2-y][x] = in[y][x]
		}
	}
	return
}

func parseShape(input string) shape {
	var s shape
	lines := strings.Split(input, "\n")
	// ignore first line
	if len(lines) != 4 {
		panic(lines)
	}
	if len(lines[1]) != 3 {
		panic(lines[1])
	}
	if len(lines[2]) != 3 {
		panic(lines[2])
	}
	if len(lines[3]) != 3 {
		panic(lines[3])
	}

	for y := range 3 {
		for x := range 3 {
			switch lines[1+y][x] {
			case '.':
				s[y][x] = false
			case '#':
				s[y][x] = true
			default:
				panic(lines[1+y])
			}

		}
	}
	return s
}

func parseRegion(input string) region {
	left_right := strings.Split(input, ": ")
	if len(left_right) != 2 {
		panic(left_right)
	}
	width_height := strings.Split(left_right[0], "x")
	if len(width_height) != 2 {
		panic(width_height)
	}
	width, err := strconv.Atoi(width_height[0])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(width_height[1])
	if err != nil {
		panic(err)
	}

	shapeCount := util.SpaceList2IntSlice(left_right[1])
	return region{width, height, shapeCount}
}

func parseInput(input string) puzzle {
	sections := strings.Split(input, "\n\n")
	shapes := make([]shape, len(sections)-1)
	for i := range shapes {
		shapes[i] = parseShape(sections[i])
	}

	regionLines := strings.Split(sections[len(sections)-1], "\n")
	regions := make([]region, len(regionLines))
	for r := range regionLines {
		regions[r] = parseRegion(regionLines[r])
	}

	return puzzle{shapes, regions}
}
