package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`

func TestParsing(t *testing.T) {
	assert := assert.New(t)
	p := parseInput(example)

	assert.Equal(6, len(p.shapes))
	assert.Equal([3]bool{false, true, true}, p.shapes[2][0])
	assert.Equal([3]bool{true, true, true}, p.shapes[2][1])
	assert.Equal([3]bool{true, true, false}, p.shapes[2][2])

}
