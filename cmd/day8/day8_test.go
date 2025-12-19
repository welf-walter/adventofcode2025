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

	pd := allPairsDistances(jb)
	assert.Equal(20*19/2, len(pd))

	assert.Equal(pairDistance{0, 19, 100427}, pd[0])
	connect(pd[0].i, pd[0].j, jb)
	log.Println(jb)
	log.Println(determineCircuitSizes(jb))

	assert.Equal(pairDistance{0, 7, 103401}, pd[1])
	connect(pd[1].i, pd[1].j, jb)
	log.Println(jb)
	log.Println(determineCircuitSizes(jb))

	connect(pd[2].i, pd[2].j, jb)
	circuitSizes := determineCircuitSizes(jb)
	assert.Equal([]int{3, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, circuitSizes)

	for n := range 10 {
		i := pd[n].i
		j := pd[n].j
		if jb[i].c != jb[j].c {
			connect(i, j, jb)
		}
		log.Println(determineCircuitSizes(jb))
	}

	circuitSizes = determineCircuitSizes(jb)
	assert.Equal([]int{5, 4, 2, 2, 1, 1, 1, 1, 1, 1, 1}, circuitSizes)

	assert.Equal(5*4*2, circuitSizes[0]*circuitSizes[1]*circuitSizes[2])

}

func Test2(t *testing.T) {
	assert := assert.New(t)
	jb := parseInput(example)

	pd := allPairsDistances(jb)

	for n := 0; ; n++ {
		i := pd[n].i
		j := pd[n].j
		if jb[i].c != jb[j].c {
			connect(i, j, jb)
		}
		circuitSize := determineCircuitSizes(jb)
		if len(circuitSize) == 1 {
			assert.Equal(junctionBox{216, 146, 977, 0}, jb[i])
			assert.Equal(junctionBox{117, 168, 530, 0}, jb[j])
			assert.Equal(25272, jb[i].x*jb[j].x)
			return
		}
	}
}
