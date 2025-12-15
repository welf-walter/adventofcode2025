package main

import (
	"adventofcode/year2025/cmd/optimize"
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

func pathToStringList(path optimize.Path) []string {
	names := make([]string, len(path))
	for index, node := range path {
		names[index] = node.(optimize.SimpleNode).String()
	}
	return names
}

func Test1(t *testing.T) {
	assert := assert.New(t)
	graph := parseInput(example)

	pathes := []optimize.Path{}
	optimize.ForAllPathes(graph.FindNode(startNodeName), func(path optimize.Path) {
		pathes = append(pathes, path)
	})

	assert.Equal(5, len(pathes))
	assert.Equal([]string{"you", "bbb", "ddd", "ggg", "out"}, pathToStringList(pathes[0]))
	assert.Equal([]string{"you", "bbb", "eee", "out"}, pathToStringList(pathes[1]))
	assert.Equal([]string{"you", "ccc", "ddd", "ggg", "out"}, pathToStringList(pathes[2]))
	assert.Equal([]string{"you", "ccc", "eee", "out"}, pathToStringList(pathes[3]))
	assert.Equal([]string{"you", "ccc", "fff", "out"}, pathToStringList(pathes[4]))
}
