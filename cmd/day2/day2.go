package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	first int
	last  int
}

// parse "11-22"
func parseRange(input string) Range {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		panic(input)
	}
	first, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		panic(err1)
	}

	last, err2 := strconv.Atoi(parts[1])
	if err2 != nil {
		panic(err2)
	}

	return Range{first, last}
}

func parseInput(input string) []Range {
	ranges := []Range{}
	ranges_iter := strings.SplitSeq(input, ",")
	for range_str := range ranges_iter {
		ranges = append(ranges, parseRange(range_str))
	}
	return ranges
}

type InvalidyFunction func(id int) bool

func isInvalid1(id int) bool {
	str := strconv.Itoa(id)
	l := len(str)
	if l%2 != 0 {
		return false
	}
	left := str[0 : l/2]
	right := str[l/2:]
	return left == right
}

func sumInvalidIds(r Range, isInvalid InvalidyFunction) int64 {
	sum := int64(0)
	for id := r.first; id <= r.last; id++ {
		if isInvalid(id) {
			sum += int64(id)
		}
	}
	return sum
}

func sumInvalidIdsOfRanges(ranges []Range, isInvalid InvalidyFunction) int64 {
	sum := int64(0)
	for _, r := range ranges {
		sum += sumInvalidIds(r, isInvalid)
	}
	return sum
}

func loadInput() string {
	dat, err := os.ReadFile("input/day2.txt")
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func main() {
	ranges := parseInput(loadInput())
	sum := sumInvalidIdsOfRanges(ranges, isInvalid1)
	fmt.Println(sum)
}
