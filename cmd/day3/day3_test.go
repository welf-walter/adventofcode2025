package main

import (
	"fmt"
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
	assert.Equal(98, calcBank(Bank{Joltage(9), Joltage(8)}))

	assert.Equal(98, findLargestJoltage(banks[0], 2))
	assert.Equal(89, findLargestJoltage(banks[1], 2))
	assert.Equal(78, findLargestJoltage(banks[2], 2))
	assert.Equal(92, findLargestJoltage(banks[3], 2))

	assert.Equal(98+89+78+92, sumLargestJoltage(banks, 2))

}

func TestHelpers(t *testing.T) {
	assert := assert.New(t)

	assert.True(isBankBetter(parseLine("4321"), parseLine("1234")))
	assert.True(isBankBetter(parseLine("4321"), nil))
	assert.True(isBankBetter(parseLine("4321"), parseLine("4320")))
	assert.False(isBankBetter(parseLine("4321"), parseLine("4321")))
	assert.False(isBankBetter(parseLine("4321"), parseLine("4322")))
	bank := parseLine("4321")
	assert.Equal(parseLine("421"), drop(bank, 1))
	assert.Equal(parseLine("321"), drop(bank, 0))
	assert.Equal(parseLine("421"), drop(bank, 1))
	assert.Equal(parseLine("431"), drop(bank, 2))
	assert.Equal(parseLine("432"), drop(bank, 3))

}

func TestOnAllBanks(t *testing.T) {
	assert := assert.New(t)

	vals := []int{}
	collectBank := func(bank Bank) {
		fmt.Printf("%v\n", bank)
		vals = append(vals, calcBank(bank))
	}

	onAllSubbanks(Bank{}, parseLine("1234"), 1, collectBank)
	assert.Equal([]int{1, 2, 3, 4}, vals)

	vals = []int{}
	onAllSubbanks(Bank{}, parseLine("1234"), 2, collectBank)
	assert.Equal([]int{12, 13, 14, 23, 24, 34}, vals)

	vals = []int{}
	onAllSubbanks(Bank{}, parseLine("1234"), 3, collectBank)
	assert.Equal([]int{123, 124, 134, 234}, vals)

	vals = []int{}
	onAllSubbanks(Bank{}, parseLine("1234"), 4, collectBank)
	assert.Equal([]int{1234}, vals)

}

func Test2(t *testing.T) {
	assert := assert.New(t)

	banks := parseInput(input1)

	assert.Equal(987654321111, findLargestJoltage(banks[0], 12))
	assert.Equal(811111111119, findLargestJoltage(banks[1], 12))
	assert.Equal(434234234278, findLargestJoltage(banks[2], 12))
	assert.Equal(888911112111, findLargestJoltage(banks[3], 12))

	assert.Equal(3121910778619, sumLargestJoltage(banks, 12))

}
