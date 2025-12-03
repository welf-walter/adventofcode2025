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

func calcBank(bank Bank) int {
	val := 0
	for _, joltage := range bank {
		val = val*10 + int(joltage)
	}
	return val
}

// is bank1 better than bank2
func isBankBetter(bank1 Bank, bank2 Bank) bool {
	// ein wurstbrot ist besser als nix
	if bank2 == nil {
		return true
	}

	if len(bank1) != len(bank2) {
		panic(fmt.Sprintf("%v -- %v", bank1, bank2))
	}
	for i := 0; i < len(bank1); i++ {
		if bank1[i] > bank2[i] {
			return true
		}
		if bank1[i] < bank2[i] {
			return false
		}
	}
	// "Both are equal"
	return false
}

func drop(bank Bank, index int) Bank {
	// this does not work!! return append(bank[:index], bank[index+1:]...)
	new := Bank{}
	for i := range bank {
		if i != index {
			new = append(new, bank[i])
		}
	}
	return new
}

func findLargestJoltage(bank Bank, digits int) Bank {

	//	fmt.Printf("%v: %v\n", digits, bank)
	// for each Joltage, remove it and check is the bank is then the best
	bestBank := Bank(nil)
	for i := range bank {
		droppedOne := drop(bank, i)

		//		fmt.Printf("    %v: %v\n", i, droppedOne)
		if isBankBetter(droppedOne, bestBank) {
			//fmt.Printf("  better\n")
			bestBank = droppedOne
		}
	}
	//fmt.Printf("%v: %v\n\n", digits, bestBank)

	if digits == len(bestBank) {
		return bestBank
	} else {
		return findLargestJoltage(bestBank, digits)
	}

}

func sumLargestJoltage(banks []Bank, digits int) int {
	sum := 0
	for _, bank := range banks {
		sum += calcBank(findLargestJoltage(bank, digits))
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
