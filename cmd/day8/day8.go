package main

import (
	"log"
	"slices"
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

// square distance, to be precise
func calcDistance(x1, y1, z1, x2, y2, z2 int) int {
	return (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) + (z2-z1)*(z2-z1)
}

// find closest pair which is not yet connected (not in the same circuit)
func findClosestPair(jb []junctionBox) (index1, index2 int) {
	minDist := 99999999
	min1 := -1
	min2 := -2
	for i := 0; i < len(jb); i++ {
		for j := i + 1; j < len(jb); j++ {
			if jb[i].c != jb[j].c {
				dist := calcDistance(
					jb[i].x, jb[i].y, jb[i].z,
					jb[j].x, jb[j].y, jb[j].z)
				if dist < minDist {
					minDist = dist
					min1 = i
					min2 = j
				}
			}
		}
	}
	return min1, min2
}

func connect(i, j int, jb []junctionBox) {
	log.Printf("connect %v and %v", i, j)
	newc := min(jb[i].c, jb[j].c)
	for k := range jb {
		if jb[k].c == jb[i].c || jb[k].c == jb[j].c {
			log.Printf("change index %v from circuit %v to circuit %v", k, jb[k].c, newc)
			jb[k].c = newc
		}
	}
}

// return value is sorted descending
func determineCircuitSizes(jb []junctionBox) []int {
	circuitSize := map[int]int{}
	for i := range jb {
		circuitSize[jb[i].c] += 1
	}
	list := []int{}
	for _, size := range circuitSize {
		list = append(list, size)
	}
	slices.Sort(list)
	slices.Reverse(list)
	return list
}
