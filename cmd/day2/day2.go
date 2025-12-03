package main

import (
	"adventofcode/year2025/cmd/util"
	"fmt"
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

func isInvalid2(id int) bool {
	str := strconv.Itoa(id)
	for l := 1; l <= len(str)/2; l++ {

		//fmt.Printf("%v: l = %v, len(str)%%l = %v\n", id, l, len(str)%l)
		if len(str)%l != 0 {
			continue
		}

		allequal := true
		for r := 1 * l; r+l <= len(str); r += l {
			left := str[0:l]
			right := str[r : r+l]
			if left != right {
				//fmt.Printf("%v: (%v;%v): %v != %v\n", id, l, r, left, right)
				allequal = false
				break
			}
			//fmt.Printf("%v: (%v;%v): %v = %v\n", id, l, r, left, right)
		}
		if allequal {
			//fmt.Printf("invalid %v -------\n", id)
			return true
		}
	}
	return false
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

func main() {
	ranges := parseInput(util.LoadInput(2))
	sum1 := sumInvalidIdsOfRanges(ranges, isInvalid1)
	fmt.Println(sum1)
	sum2 := sumInvalidIdsOfRanges(ranges, isInvalid2)
	fmt.Println(sum2)
}
