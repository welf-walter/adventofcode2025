package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

func TestParsing(t *testing.T) {
	assert := assert.New(t)
	graph := parseInput(example)

	assert.Equal([]string{"bbb", "ccc"}, graph.GetTargets(startNodeName))
	assert.Equal([]string{"ddd", "eee", "fff"}, graph.GetTargets("ccc"))
	assert.Equal([]string{"out"}, graph.GetTargets("iii"))

}
