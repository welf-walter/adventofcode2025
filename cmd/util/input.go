package util

import (
	"fmt"
	"os"
)

func LoadInput(day int) string {
	dat, err := os.ReadFile(fmt.Sprintf("input/day%v.txt", day))
	if err != nil {
		panic(err)
	}
	return string(dat)
}
