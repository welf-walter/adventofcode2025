package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

func TestParsing(t *testing.T) {
	assert := assert.New(t)
	machines := parseInput(example)

	assert.Equal(3, len(machines))
	assert.Equal(indicatorLights{false, true, true, false}, machines[0].lights)
	assert.Equal(button{1, 3}, machines[0].buttons[1])
	assert.Equal([]int{3, 5, 4, 7}, machines[0].joltage)

}

func Test1(t *testing.T) {
	assert := assert.New(t)
	machines := parseInput(example)

	l0, b0 := toggleIndicatorLights(machines[0], 0b000111)
	assert.Equal(machines[0].lights, l0)
	assert.Equal(3, b0)
	l1, b1 := toggleIndicatorLights(machines[0], 0b001010)
	assert.Equal(machines[0].lights, l1)
	assert.Equal(2, b1)
	l2, b2 := toggleIndicatorLights(machines[0], 0b111101)
	assert.Equal(machines[0].lights, l2)
	assert.Equal(5, b2)

	assert.Equal(2, calcMinNumberOfButtons(machines[0]))
	assert.Equal(3, calcMinNumberOfButtons(machines[1]))
	assert.Equal(2, calcMinNumberOfButtons(machines[2]))

}
