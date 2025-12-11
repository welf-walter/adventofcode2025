package main

import (
	"adventofcode/year2025/cmd/util"
	"strings"
)

// indices of indicatorLights that this button toggles
type button []int

type machine struct {
	indicatorLights []bool
	buttons         []button
	joltage         []int
}

func parseIndicatorLights(input string) []bool {
	if input[0] != '[' || input[len(input)-1] != ']' {
		panic(input)
	}
	var lights []bool
	for i := 1; i < len(input)-1; i++ {
		switch input[i] {
		case '.':
			lights = append(lights, false)
		case '#':
			lights = append(lights, true)
		default:
			panic(input)
		}
	}
	return lights
}

func parseButton(input string) button {
	if input[0] != '(' || input[len(input)-1] != ')' {
		panic(input)
	}
	return util.CommaList2IntSlice(input[1 : len(input)-1])
}

func parseJoltage(input string) button {
	if input[0] != '{' || input[len(input)-1] != '}' {
		panic(input)
	}
	return util.CommaList2IntSlice(input[1 : len(input)-1])
}

func parseInput(input string) (machines []machine) {
	line_iter := strings.SplitSeq(input, "\n")
	for line := range line_iter {
		var m machine
		token_iter := strings.SplitSeq(line, " ")
		for token := range token_iter {
			switch token[0] {
			case '[':
				m.indicatorLights = parseIndicatorLights(token)
			case '(':
				m.buttons = append(m.buttons, parseButton(token))
			case '{':
				m.joltage = parseJoltage(token)
			default:
				panic(token)
			}

		}
		machines = append(machines, m)
	}
	return
}
