package main

import (
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

func TestParsing(t *testing.T) {
	assert := assert.New(t)
	tiles := parseInput(example)
	assert.Equal(8, len(tiles))
	assert.Equal(redTile{11, 7}, tiles[2])
	assert.Equal(50, largestRectangle(tiles))
}
