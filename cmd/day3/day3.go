package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Joltage int
type Bank []Joltage

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

func main() {
	fmt.Println("Not yet implemented")
}
