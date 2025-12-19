package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
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

type pairDistance struct{ i, j, dist int }

// get list of all pairs, sorted by distance ascending
func allPairsDistances(jb []junctionBox) []pairDistance {
	pd := []pairDistance{}
	util.ForAllPairIndices(jb, func(i, j int) {
		dist := calcDistance(
			jb[i].x, jb[i].y, jb[i].z,
			jb[j].x, jb[j].y, jb[j].z)
		pd = append(pd, pairDistance{i, j, dist})
	})
	slices.SortFunc(pd, func(a, b pairDistance) int { return a.dist - b.dist })
	return pd
}

func connect(i, j int, jb []junctionBox) {
	log.Printf("connect %v and %v", i, j)
	newc := min(jb[i].c, jb[j].c)
	oldc := max(jb[i].c, jb[j].c)
	for k := range jb {
		if jb[k].c == oldc {
			jb[k].c = newc
			log.Printf("change index %v from circuit %v to circuit %v", k, oldc, newc)
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

func main() {
	input := util.LoadInput(8)
	jb := parseInput(input)

	pd := allPairsDistances(jb)
	for n := range 1000 {
		log.Printf("Iteration #%v", n)
		i := pd[n].i
		j := pd[n].j
		if jb[i].c != jb[j].c {
			connect(i, j, jb)
		}
		log.Println(determineCircuitSizes(jb))
	}

	circuitSizes := determineCircuitSizes(jb)
	log.Println(circuitSizes)
	fmt.Println(circuitSizes[0] * circuitSizes[1] * circuitSizes[2])

	for n := 1000; ; n++ {
		i := pd[n].i
		j := pd[n].j
		if jb[i].c != jb[j].c {
			connect(i, j, jb)
		}
		circuitSize := determineCircuitSizes(jb)
		if len(circuitSize) == 1 {
			log.Println(n)
			fmt.Println(jb[i].x * jb[j].x)
			return
		}
	}

	//	2025/12/09 23:11:38 Iteration #998
	//
	// 2025/12/09 23:11:38 connect 200 and 885
	// 2025/12/09 23:11:38 change index 200 from circuit 200 to circuit 0
	// 2025/12/09 23:11:38 [1000]
	// 2025/12/09 23:11:38 [1000]
}
