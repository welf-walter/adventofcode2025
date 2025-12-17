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

func TestShapes(t *testing.T) {
	assert := assert.New(t)
	s1 := parseShape(`3:
##.
###
##.`)

	s2 := parseShape(`3:
###
###
.#.`)

	s3 := parseShape(`3:
.##
###
.##`)

	s4 := parseShape(`3:
.#.
###
###`)

	assert.Equal(s2, rotateShape(s1))
	assert.Equal(s3, rotateShape(s2))
	assert.Equal(s4, rotateShape(s3))
	assert.Equal(s1, rotateShape(s4))

	assert.Equal(s4, flipShape(s2))

	r := region{5, 5, []int{}}
	m := r.makeMap()

	assert.True(m.canPlace(s1, 0, 0))
	m.doPlace(s1, 0, 0)
	assert.False(m.canPlace(s1, 0, 0))
	assert.False(m.canPlace(s1, 1, 1))
	assert.True(m.canPlace(s1, 2, 2))
	m.doPlace(s1, 2, 2)
	assert.False(m.canPlace(s1, 2, 2))
}
