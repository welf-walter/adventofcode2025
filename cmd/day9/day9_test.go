package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func Test1(t *testing.T) {
	assert := assert.New(t)
	tiles := parseInput(example)
	assert.Equal(8, len(tiles))
	assert.Equal(redTile{11, 7}, tiles[2])
	assert.Equal(50, largestRectangle(tiles))
}

func Test2(t *testing.T) {
	assert := assert.New(t)
	tiles := parseInput(example)
	floor := newTileFloor(12)
	connectTiles(&floor, tiles)
	log.Print(floor)

	assert.True(floor.coversRectangle(redTile{7, 3}, redTile{11, 1}))
	assert.False(floor.coversRectangle(redTile{6, 3}, redTile{11, 1}))
	assert.False(floor.coversRectangle(redTile{7, 3}, redTile{11, 0}))
	assert.True(floor.coversRectangle(redTile{9, 7}, redTile{9, 5}))
	assert.True(floor.coversRectangle(redTile{9, 5}, redTile{2, 3}))

	assert.Equal(24, largestRectangleInFloor(floor, tiles))

}
