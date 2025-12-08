package main

import (
	"log"
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
	assert := assert.New(t)
	jb := parseInput(example)
	log.Println(jb)
	log.Println(determineCircuitSizes(jb))

	i, j := findClosestPair(jb)
	assert.Equal(0, i)
	assert.Equal(19, j)
	connect(i, j, jb)
	log.Println(jb)
	log.Println(determineCircuitSizes(jb))

	i, j = findClosestPair(jb)
	assert.Equal(0, i)
	assert.Equal(7, j)
	connect(i, j, jb)
	log.Println(jb)
	log.Println(determineCircuitSizes(jb))

	circuitSizes := determineCircuitSizes(jb)
	assert.Equal([]int{3, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, circuitSizes)

	for n := 2; n < 10; n++ {
		i, j = findClosestPair(jb)
		connect(i, j, jb)
	}

	circuitSizes = determineCircuitSizes(jb)
	assert.Equal([]int{5, 4, 2, 2, 1, 1, 1, 1, 1, 1, 1}, circuitSizes)

}
