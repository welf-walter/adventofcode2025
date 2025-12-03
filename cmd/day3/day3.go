package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
	"strconv"
	"strings"
)

type Joltage int
type Bank []Joltage
type Pair int

func parseLine(line string) Bank {
	jolts := make([]Joltage, len(line))
	for i := 0; i < len(line); i++ {
		jolt, err := strconv.Atoi(line[i : i+1])
		if err != nil {
			panic(err)
		}
		jolts[i] = Joltage(jolt)
	}
	return jolts

}

func parseInput(input string) []Bank {
	banks := []Bank{}
	bankstr_iter := strings.SplitSeq(input, "\n")
	for bank_str := range bankstr_iter {
		banks = append(banks, parseLine(bank_str))
	}
	return banks
}

func findLargestPair(bank Bank) Pair {
	largest := Pair(-1)
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			pair := Pair(int(bank[i])*10 + int(bank[j]))
			if largest < pair {
				largest = pair
			}
		}
	}
	return largest
}

func sumLargestPairs(banks []Bank) int {
	sum := 0
	for _, bank := range banks {
		sum += int(findLargestPair(bank))
	}
	return sum
}

func main() {
	banks := parseInput(util.LoadInput(3))
	sum := sumLargestPairs(banks)
	fmt.Println(sum)
}
