package main

import (
	"adventofcode/year2025/cmd/optimize"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example1 = `aaa: you hhh
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
	graph := parseInput(example1)

	assert.Equal([]string{"bbb", "ccc"}, graph.GetTargets(startNodeName1))
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
	graph := parseInput(example1)

	pathes := []optimize.Path{}
	optimize.ForAllPathes(graph.FindNode(startNodeName1), func(path optimize.Path) {
		pathes = append(pathes, path)
	})

	assert.Equal(5, len(pathes))
	assert.Equal([]string{"you", "bbb", "ddd", "ggg", "out"}, pathToStringList(pathes[0]))
	assert.Equal([]string{"you", "bbb", "eee", "out"}, pathToStringList(pathes[1]))
	assert.Equal([]string{"you", "ccc", "ddd", "ggg", "out"}, pathToStringList(pathes[2]))
	assert.Equal([]string{"you", "ccc", "eee", "out"}, pathToStringList(pathes[3]))
	assert.Equal([]string{"you", "ccc", "fff", "out"}, pathToStringList(pathes[4]))
}

const example2 = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`

func Test2(t *testing.T) {
	assert := assert.New(t)
	graph := parseInput(example2)

	allPathCount := 0
	pathes := []optimize.Path{}
	optimize.ForAllPathes(graph.FindNode(startNodeName2), func(path optimize.Path) {
		allPathCount++
		if maybeProblematic(path) {
			pathes = append(pathes, path)
		}
	})

	assert.Equal(8, allPathCount)

	assert.Equal(2, len(pathes))
	assert.Equal([]string{"svr", "aaa", "fft", "ccc", "eee", "dac", "fff", "ggg", "out"}, pathToStringList(pathes[0]))
	assert.Equal([]string{"svr", "aaa", "fft", "ccc", "eee", "dac", "fff", "hhh", "out"}, pathToStringList(pathes[1]))

}
