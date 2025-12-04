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

	removePapers(grid, positions)

	positions2 := findForkliftAccessible(grid)
	assert.Equal(12, len(positions2))
	removePapers(grid, positions2)

	positions3 := findForkliftAccessible(grid)
	assert.Equal(7, len(positions3))
	removePapers(grid, positions3)

	positions4 := findForkliftAccessible(grid)
	assert.Equal(5, len(positions4))
	removePapers(grid, positions4)

	positions5 := findForkliftAccessible(grid)
	assert.Equal(2, len(positions5))
	removePapers(grid, positions5)

	positions6 := findForkliftAccessible(grid)
	assert.Equal(1, len(positions6))
	removePapers(grid, positions6)

	positions7 := findForkliftAccessible(grid)
	assert.Equal(1, len(positions7))
	removePapers(grid, positions7)

	positions8 := findForkliftAccessible(grid)
	assert.Equal(1, len(positions8))
	removePapers(grid, positions8)

	positions9 := findForkliftAccessible(grid)
	assert.Equal(1, len(positions9))
	removePapers(grid, positions9)

	assert.Equal(0, countForkliftAccessible(grid))

}
