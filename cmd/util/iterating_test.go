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
	ForAllPairs(numbers, do)

	assert.Equal("(1,2)(1,3)(2,3)", s)

	s = ""
	doIndex := func(i, j int) { s = s + fmt.Sprintf("(%v,%v)", i, j) }
	ForAllPairIndices(numbers, doIndex)
	assert.Equal("(0,1)(0,2)(1,2)", s)

}
