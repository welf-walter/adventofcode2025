package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestParsing(t *testing.T) {
	jb := parseInput(example)
	assert := assert.New(t)
	assert.Equal(20, len(jb))
	assert.Equal(junctionBox{57, 618, 57, 1}, jb[1])
}

func Test1(t *testing.T) {
	jb := parseInput(example)
	assert := assert.New(t)

	i, j := findClosestPair(jb)
	assert.Equal(0, i)
	assert.Equal(19, j)
	jb = connect(i, j, jb)
	fmt.Println(jb)

	i, j = findClosestPair(jb)
	assert.Equal(0, i)
	assert.Equal(7, j)
	jb = connect(i, j, jb)
	fmt.Println(jb)

	circuitSizes := determineCircuitSizes(jb)
	assert.Equal([]int{3, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, circuitSizes)

	for n := 2; n < 10; n++ {
		i, j = findClosestPair(jb)
		jb = connect(i, j, jb)
	}

	circuitSizes = determineCircuitSizes(jb)
	assert.Equal([]int{5, 4, 2, 2, 1, 1, 1, 1, 1, 1, 1}, circuitSizes)

}
