package util

import (
	"fmt"
	"os"
	"strconv"
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
