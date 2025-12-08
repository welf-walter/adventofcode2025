package main

import (
	"strconv"
	"strings"
)

type junctionBox struct {
	x int
	y int
	z int
	c int // circuit number
}

func parseInput(input string) []junctionBox {
	jb := []junctionBox{}
	line_iter := strings.SplitSeq(input, "\n")
	c := 0
	for line := range line_iter {
		xyz := strings.Split(line, ",")
		x, err := strconv.Atoi(xyz[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(xyz[1])
		if err != nil {
			panic(err)
		}
		z, err := strconv.Atoi(xyz[2])
		if err != nil {
			panic(err)
		}
		jb = append(jb, junctionBox{x, y, z, c})
		c++
	}
	return jb
}
