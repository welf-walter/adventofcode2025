package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
	"log"
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

func calcArea(a, b redTile) int {
	return (max(a.x, b.x) - min(a.x, b.x) + 1) *
		(max(a.y, b.y) - min(a.y, b.y) + 1)
}

func largestRectangle(tiles []redTile) (area int) {
	util.ForAllPairs(tiles, func(a, b redTile) { area = max(area, calcArea(a, b)) })
	return
}

func largestRectangleInFloor(floor tileFloor, tiles []redTile) (area int) {
	util.ForAllPairs(tiles, func(a, b redTile) {
		if floor.coversRectangle(a, b) {
			area = max(area, calcArea(a, b))
		}
	})
	return
}

// per line, from where to where are red or green tiles?
type tileFloor struct {
	lines []struct{ from, to int }
}

func newTileFloor(height int) tileFloor {
	return tileFloor{lines: make([]struct {
		from int
		to   int
	}, height)}
}

func (floor *tileFloor) put(x, y int) {
	line := floor.lines[y]
	//if line.from != line.to {
	//	if line.from != x && line.to != x {
	//		panic(fmt.Sprintf("Unexpected: More than two points in one row: %v,%v", x, y))
	//	}
	//}
	if line.from == 0 {
		line.from = x
		line.to = x
	} else {
		line.from = min(line.from, x)
		line.to = max(line.to, x)
	}
	floor.lines[y] = line
}

func (floor *tileFloor) putLine(a redTile, b redTile) {
	log.Printf("Line from (%v,%v) to (%v,%v)", a.x, a.y, b.x, b.y)
	if a.y == b.y {
		floor.put(a.x, a.y)
		floor.put(b.x, b.y)
		return
	}
	if a.x != b.x {
		panic("x != x and y != y")
	}
	miny := min(a.y, b.y)
	maxy := max(a.y, b.y)
	for y := miny; y <= maxy; y++ {
		floor.put(a.x, y)
	}
}

func (floor *tileFloor) coversRectangle(a redTile, b redTile) bool {
	miny := min(a.y, b.y)
	maxy := max(a.y, b.y)
	minx := min(a.x, b.x)
	maxx := max(a.x, b.x)
	for y := miny; y <= maxy; y++ {
		if floor.lines[y].from > minx {
			log.Printf("In line %v, floor is %v..%v but not %v", y, floor.lines[y].from, floor.lines[y].to, minx)
			return false
		}
		if floor.lines[y].to < maxx {
			log.Printf("In line %v, floor is %v..%v but not %v", y, floor.lines[y].from, floor.lines[y].to, maxx)
			return false
		}
	}
	return true
}

func connectTiles(floor *tileFloor, tiles []redTile) {
	for i := range tiles {
		a := tiles[i]
		b := tiles[(i+1)%len(tiles)]
		floor.putLine(a, b)
	}
}

func main() {
	tiles := parseInput(util.LoadInput(9))
	fmt.Println(largestRectangle(tiles))

	floor := newTileFloor(99999)
	connectTiles(&floor, tiles)
	fmt.Println(largestRectangleInFloor(floor, tiles))
}
