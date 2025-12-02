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

func right(dial Dial, ticks Dial) (Dial, int) {
	new := dial + ticks
	zerocrosses := 0
	for new >= 100 {
		new = new - 100
		zerocrosses++
	}
	return new, zerocrosses
}

func left(dial Dial, ticks Dial) (Dial, int) {
	new := dial - ticks
	zerocrosses := 0
	for new < 0 {
		new = new + 100
		zerocrosses++
	}
	return new, zerocrosses
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

func execute(operations []Operation) (int, int) {
	zero_counter1 := 0
	zero_counter2 := 0
	dial := START_DIAL
	for _, operation := range operations {
		crossed := 0
		if operation.direction == RIGHT {
			dial, crossed = right(dial, operation.ticks)
			zero_counter2 += crossed
		} else if operation.direction == LEFT {
			was_zero := dial == 0
			dial, crossed = left(dial, operation.ticks)
			zero_counter2 += crossed
			if was_zero {
				zero_counter2--
			}
			if dial == ZERO_DIAL {
				zero_counter2++
			}
		} else {
			panic(fmt.Sprintf("Unexpected: %v", operation.direction))
		}
		if dial == ZERO_DIAL {
			zero_counter1++
		}
		//		fmt.Println(dial)
	}
	return zero_counter1, zero_counter2
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
	zero_counter1, zero_counter2 := execute(operations)
	fmt.Println(zero_counter1)
	fmt.Println(zero_counter2)
}
