package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	first int
	last  int
}

func parseInput(input string) []Range {
	ranges := []Range{}
	ranges_str := strings.Split(input, ",")
	for _, range_str := range ranges_str {
		parts := strings.Split(range_str, "-")
		if len(parts) != 2 {
			panic(range_str)
		}
		first, err1 := strconv.Atoi(parts[0])
		if err1 != nil {
			panic(err1)
		}

		last, err2 := strconv.Atoi(parts[0])
		if err2 != nil {
			panic(err2)
		}

		ranges = append(ranges, Range{first, last})
	}
	return ranges
}

func isInvalid(id int) bool {
	str := strconv.Itoa(id)
	l := len(str)
	if l%2 != 0 {
		return false
	}
	left := str[0 : l/2]
	right := str[l/2:]
	return left == right
}

func main() {
	fmt.Println("Hello, World!")
}
