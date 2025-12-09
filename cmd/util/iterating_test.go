package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForAllPairs(t *testing.T) {
	assert := assert.New(t)
	numbers := []int{1, 2, 3}
	s := ""
	do := func(i, j int) { s = s + fmt.Sprintf("(%v,%v)", i, j) }
	forAllPairs(numbers, do)

	assert.Equal("(1,2)(1,3)(2,3)", s)
}
