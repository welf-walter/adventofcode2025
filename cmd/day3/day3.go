package main

import (
	"adventofcode/year2025/cmd/util"
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

func onAllSubbanks(prefix Bank, rest Bank, toadd int, yield func(subbank Bank)) {
	if toadd == 0 {
		yield(prefix)
		return
	}
	if len(rest) == 0 {
		return
	}
	// take it
	onAllSubbanks(append(prefix, rest[0]), rest[1:], toadd-1, yield)
	// or leave it
	onAllSubbanks(prefix, rest[1:], toadd, yield)
}

func calcBank(bank Bank) int {
	val := 0
	for _, joltage := range bank {
		val = val*10 + int(joltage)
	}
	return val
}

func findLargestJoltage(bank Bank, digits int) int {

	largest := 0
	//dumpBank := func(subbank Bank) { fmt.Printf("%v\n", bank) }
	findMaxBank := func(subbank Bank) {
		val := calcBank(subbank)
		if largest < val {
			largest = val
		}
	}
	onAllSubbanks(Bank{}, bank, digits, findMaxBank)
	return largest
}

func sumLargestJoltage(banks []Bank, digits int) int {
	sum := 0
	for _, bank := range banks {
		sum += int(findLargestJoltage(bank, digits))
		fmt.Print(".")
	}
	return sum
}

func main() {
	banks := parseInput(util.LoadInput(3))
	sum := sumLargestJoltage(banks, 2)
	fmt.Println(sum)
	sum = sumLargestJoltage(banks, 4)
	fmt.Println(sum)
}
