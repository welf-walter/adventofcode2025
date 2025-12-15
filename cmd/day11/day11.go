package main

import (
	"adventofcode/year2025/cmd/optimize"
	"adventofcode/year2025/cmd/util"
	"fmt"
	"strings"
)

const startNodeName1 = "you"
const startNodeName2 = "svr"
const finishNodeName = "out"

func parseInput(input string) optimize.SimpleGraph {
	graph := optimize.SimpleGraph{}
	lines := strings.Split(input, "\n")

	graph.AddNode(finishNodeName)
	// add nodes
	for _, line := range lines {
		left_right := strings.Split(line, ": ")
		graph.AddNode(left_right[0])
	}

	// add edges
	for _, line := range lines {
		left_right := strings.Split(line, ": ")
		from := graph.FindNode(left_right[0])
		targets := strings.Split(left_right[1], " ")
		for _, target := range targets {
			graph.AddEdge(from, graph.FindNode(target))
		}
	}

	graph.SetFinish(finishNodeName)
	return graph

}

func maybeProblematic(path optimize.Path) bool {
	dac := false
	fft := false
	for _, node := range path {
		if node.(optimize.SimpleNode).String() == "dac" {
			dac = true
		}
		if node.(optimize.SimpleNode).String() == "fft" {
			fft = true
		}
	}
	return dac && fft
}

func main() {
	graph := parseInput(util.LoadInput(11))

	pathCounter1 := 0
	optimize.ForAllPathes(graph.FindNode(startNodeName1), func(path optimize.Path) {
		pathCounter1++
	})
	fmt.Println(pathCounter1)

	pathCounter2 := 0
	optimize.ForAllPathes(graph.FindNode(startNodeName2), func(path optimize.Path) {
		if maybeProblematic(path) {
			pathCounter2++
		}
	})
	fmt.Println(pathCounter2)

}
