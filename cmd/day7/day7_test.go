package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func TestParse(t *testing.T) {
	assert := assert.New(t)
	start, rows := parseInput(example)

	assert.Equal(7, start)
	assert.Equal(splitterRow{}, rows[0])
	assert.Equal(splitterRow{7}, rows[1])
	assert.Equal(splitterRow{}, rows[2])
	assert.Equal(splitterRow{6, 8}, rows[3])

}

func Test1(t *testing.T) {
	assert := assert.New(t)
	start, rows := parseInput(example)

	t0 := []int{start}
	t1, splits1 := runRow(t0, rows[0])
	assert.Equal([]int{7}, t1)
	assert.Equal(0, splits1)
	t2, splits2 := runRow(t1, rows[1])
	assert.Equal([]int{6, 8}, t2)
	assert.Equal(1, splits2)

	assert.Equal(21, runRows(start, rows))

}

func Test2(t *testing.T) {
	assert := assert.New(t)
	start, rows := parseInput(example)

	t0 := timeLine{start: 1}
	t1 := runRow2(t0, rows[0])
	assert.Equal(timeLine{7: 1}, t1)
	t2 := runRow2(t1, rows[1])
	assert.Equal(timeLine{6: 1, 8: 1}, t2)
	t3 := runRow2(t2, rows[2])
	assert.Equal(timeLine{6: 1, 8: 1}, t3)
	t4 := runRow2(t3, rows[3])
	assert.Equal(timeLine{5: 1, 7: 2, 9: 1}, t4)

	//	assert.Equal(21, runRows(start, rows))

}
