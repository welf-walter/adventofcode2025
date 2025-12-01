package main

import (
	"fmt"
)

type Dial int8

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

func main() {
	dial := Dial(0)
	dial = right(dial, Dial(12))
	fmt.Println(dial)
}
