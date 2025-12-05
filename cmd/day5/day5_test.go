package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestParsing(t *testing.T) {
	assert := assert.New(t)

	ranges, ingredients := parseInput(example)

	assert.Equal(4, len(ranges))
	assert.Equal(6, len(ingredients))

	assert.Equal(Range{Ingredient(16), Ingredient(20)}, ranges[2])
	assert.Equal(Ingredient(17), ingredients[4])

}
