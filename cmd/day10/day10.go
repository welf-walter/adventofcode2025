package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
	"strings"
)

// indices of indicatorLights that this button toggles
type button []int

type indicatorLights []bool

type machine struct {
	lights  indicatorLights
	buttons []button
	joltage []int
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
				m.lights = parseIndicatorLights(token)
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

func toggleIndicatorLights(m machine, buttonBitMask int) (lights indicatorLights, buttonCount int) {
	lights = make(indicatorLights, len(m.lights))
	for index, button := range m.buttons {
		if buttonBitMask&(1<<index) > 0 {
			for _, light := range button {
				lights[light] = !lights[light]
			}
			buttonCount++
		}
	}
	return
}

func indicatorLightsEqual(i, j indicatorLights) bool {
	if len(i) != len(j) {
		panic(i)
	}
	for n := range len(i) {
		if i[n] != j[n] {
			return false
		}
	}
	return true
}

func calcMinNumberOfButtons(m machine) int {
	minButtons := 99999999
	combinations := 1 << len(m.buttons)
	for i := range combinations {
		lights, buttonCount := toggleIndicatorLights(m, i)
		if indicatorLightsEqual(lights, m.lights) {
			minButtons = min(minButtons, buttonCount)
		}
	}
	return minButtons
}

func main() {
	machines := parseInput(util.LoadInput(10))
	sum := 0
	for _, m := range machines {
		sum += calcMinNumberOfButtons(m)
	}
	fmt.Println(sum)
}
