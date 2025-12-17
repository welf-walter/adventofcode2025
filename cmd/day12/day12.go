package day12

import (
	"adventofcode/year2025/cmd/util"
	"log"
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

func (r region) makeMap() regionMap {
	m := make(regionMap, r.height)
	for y := range r.height {
		m[y] = make([]bool, r.width)
	}
	return m
}

func (r regionMap) clone() regionMap {
	m := make(regionMap, len(r))
	for y := range len(r) {
		m[y] = make([]bool, len(r[y]))
	}
	return m
}

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

func (m *regionMap) canPlace(s shape, x, y int) bool {
	for dy := range 3 {
		for dx := range 3 {
			if (*m)[y+dy][x+dx] && s[dy][dx] {
				return false
			}
		}
	}
	return true
}

func (m *regionMap) doPlace(s shape, x, y int) bool {
	for dy := range 3 {
		for dx := range 3 {
			(*m)[y+dy][x+dx] = (*m)[y+dy][x+dx] || s[dy][dx]
		}
	}
	return true
}

func (m regionMap) canIplaceAll(shapes []shape) bool {
	if len(shapes) == 0 {
		return true
	}
	head := shapes[0]
	tail := shapes[1:]

	// todo: skip equal shapes
	variants := []shape{head, flipShape(head)}
	for range 3 {
		head = rotateShape(head)
		variants = append(variants, head)
		variants = append(variants, flipShape(head))
	}
	log.Println(variants)

	// try to put head somewhere
	for y := range len(m) - len(head) + 1 {
		for x := range len(m[y]) - len(head[y]) + 1 {
			for variantIndex := range variants {
				log.Printf("  try to place variant %v at %v, %v.", variantIndex, x, y)
				if m.canPlace(head, x, y) {
					log.Printf("  place variant %v at %v, %v. %v shapes left", variantIndex, x, y, len(tail))
					m2 := m.clone()
					m2.doPlace(head, x, y)
					if m2.canIplaceAll(tail) {
						return true
					}
				}
			}
		}
	}
	return false
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
