package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`

func TestParsing(t *testing.T) {
	assert := assert.New(t)
	problems := parseInput(example)

	assert.Equal(4, len(problems))
	assert.Equal(problem{operands: []int{328, 64, 98}, operator: ADD}, problems[1])

}

func Test1(t *testing.T) {
	assert := assert.New(t)
	problems := parseInput(example)

	assert.Equal(33210, solveProblem(problems[0]))
	assert.Equal(490, solveProblem(problems[1]))
	assert.Equal(4243455, solveProblem(problems[2]))
	assert.Equal(401, solveProblem(problems[3]))
}
