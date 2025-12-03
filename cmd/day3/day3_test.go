package main

import (
	"testing"

	"github.com/stretchr/testify/assert" // https://pkg.go.dev/github.com/stretchr/testify/assert
)

const input1 = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestParseLine(t *testing.T) {
	assert := assert.New(t)

	jolts := parseLine("987654321111111")
	assert.Equal(15, len(jolts))
	assert.Equal(Joltage(9), jolts[0])
	assert.Equal(Joltage(1), jolts[14])
}

func Test1(t *testing.T) {
	assert := assert.New(t)

	banks := parseInput(input1)
	assert.Equal(4, len(banks))
	assert.Equal(Joltage(8), banks[3][2])

}
