package main

import (
	"testing"

	"github.com/stretchr/testify/assert" // https://pkg.go.dev/github.com/stretchr/testify/assert
)

const example string = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestParse(t *testing.T) {
	assert := assert.New(t)

	grid := parseInput(example)
	assert.False(isPaper(grid, -1, 0))
	assert.False(isPaper(grid, 0, -1))
	assert.True(isPaper(grid, 0, 1))
	assert.False(isPaper(grid, 1, 0))
}

func Test1(t *testing.T) {
	assert := assert.New(t)

	grid := parseInput(example)
	assert.True(isForkliftAccessible(grid, 0, 1))
	assert.False(isForkliftAccessible(grid, 1, 1))
	assert.Equal(13, countForkliftAccessible(grid))
}

func Test2(t *testing.T) {
	assert := assert.New(t)

	grid := parseInput(example)
	positions := findForkliftAccessible(grid)

	assert.Equal(countForkliftAccessible(grid), len(positions))
	assert.Equal(Position{2, 0}, positions[0])
	assert.Equal(Position{3, 0}, positions[1])
	assert.Equal(Position{5, 0}, positions[2])

}
