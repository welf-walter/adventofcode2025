package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadInput(day int) string {
	dat, err := os.ReadFile(fmt.Sprintf("input/day%v.txt", day))
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func String2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func CommaList2IntSlice(input string) []int {
	ints := []int{}
	numberStrings := strings.Split(input, ",")
	for _, numberString := range numberStrings {
		i := String2Int(numberString)
		ints = append(ints, i)
	}
	return ints
}

func SpaceList2IntSlice(input string) []int {
	ints := []int{}
	numberStrings := strings.Split(input, " ")
	for _, numberString := range numberStrings {
		i := String2Int(numberString)
		ints = append(ints, i)
	}
	return ints
}
