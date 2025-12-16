package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommaList2IntSlice(t *testing.T) {
	assert := assert.New(t)
	slice := CommaList2IntSlice("1,2,42")
	assert.Equal([]int{1, 2, 42}, slice)
}

func TestSpaceList2IntSlice(t *testing.T) {
	assert := assert.New(t)
	slice := SpaceList2IntSlice("1 2 42")
	assert.Equal([]int{1, 2, 42}, slice)
}
