package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dial int

const START_DIAL = Dial(50)
const ZERO_DIAL = Dial(0)

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
		ticks, _ := strconv.Atoi(line[1:])
		operation := Operation{
			direction: line[0],
			ticks:     Dial(ticks),
		}
		operations = append(operations, operation)
	}
	return operations
}

func execute(operations []Operation) int {
	zero_counter := 0
	dial := START_DIAL
	for _, operation := range operations {
		if operation.direction == RIGHT {
			dial = right(dial, operation.ticks)
		}
		if operation.direction == LEFT {
			dial = left(dial, operation.ticks)
		}
		if dial == ZERO_DIAL {
			zero_counter++
		}
		//		fmt.Println(dial)
	}
	return zero_counter
}

func loadInput() string {
	dat, err := os.ReadFile("input/day1.txt")
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func main() {
	operations := parseInput(loadInput())
	zero_counter := execute(operations)
	fmt.Println(zero_counter)
}
