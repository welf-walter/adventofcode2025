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
	assert.Equal(ranges[1], Range{95, 115})
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

func TestSumInvalidIds(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(11+22), sumInvalidIds(parseRange("11-22")))
	assert.Equal(int64(99), sumInvalidIds(parseRange("95-115")))
	assert.Equal(int64(1010), sumInvalidIds(parseRange("998-1012")))
	assert.Equal(int64(1188511885), sumInvalidIds(parseRange("1188511880-1188511890")))
	assert.Equal(int64(222222), sumInvalidIds(parseRange("222220-222224")))
	assert.Equal(int64(0), sumInvalidIds(parseRange("1698522-1698528")))
	assert.Equal(int64(446446), sumInvalidIds(parseRange("446443-446449")))
	assert.Equal(int64(38593859), sumInvalidIds(parseRange("38593856-38593862")))
	assert.Equal(int64(0), sumInvalidIds(parseRange("565653-565659")))
	assert.Equal(int64(0), sumInvalidIds(parseRange("824824821-824824827")))
	assert.Equal(int64(0), sumInvalidIds(parseRange("2121212118-2121212124")))

	ranges := parseInput(input1)
	assert.Equal(int64(1227775554), sumInvalidIdsOfRanges(ranges))

}
