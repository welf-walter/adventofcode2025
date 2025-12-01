package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Dial int8

const START_DIAL = Dial(50)

type Operation struct {
	direction byte
	ticks     Dial
}

const LEFT = 'L'
const RIGHT = 'R'

func right(dial Dial, ticks Dial) Dial {
	new := dial + ticks
	for new >= 100 {
		new = new - 100
	}
	return new
}

func left(dial Dial, ticks Dial) Dial {
	new := dial - ticks
	for new < 0 {
		new = new + 100
	}
	return new
}

func parseInput(input string) []Operation {
	operations := []Operation{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		ticks, _ := strconv.Atoi(line[1:len(line)])
		operation := Operation{
			direction: line[0],
			ticks:     Dial(ticks),
		}
		operations = append(operations, operation)
	}
	return operations
}

func main() {
	dial := Dial(0)
	dial = right(dial, Dial(12))
	fmt.Println(dial)
}
