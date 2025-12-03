package main

import (
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

func findLargestPair(jolts []Joltage) Pair {
	largest := Pair(-1)
	for i := 0; i < len(jolts); i++ {
		for j := i + 1; j < len(jolts); j++ {
			pair := Pair(int(jolts[i])*10 + int(jolts[j]))
			if largest < pair {
				largest = pair
			}
		}
	}
	return largest
}

func main() {
	fmt.Println("Not yet implemented")
}
