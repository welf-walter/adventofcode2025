package main

import (
	"testing"

	"github.com/stretchr/testify/assert" // https://pkg.go.dev/github.com/stretchr/testify/assert
)

const input1 = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestParse(t *testing.T) {
	assert := assert.New(t)

	ranges := parseInput(input1)

	assert.Equal(len(ranges), 11)
}

func TestIsInvalid(t *testing.T) {
	assert := assert.New(t)

	assert.True(isInvalid(55))
	assert.False(isInvalid(555))
	assert.True(isInvalid(6464))
	assert.True(isInvalid(123123))
	assert.False(isInvalid(101))
	assert.False(isInvalid(1111111))

}
