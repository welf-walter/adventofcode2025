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
