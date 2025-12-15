package main

import (
	"adventofcode/year2025/cmd/optimize"
	"adventofcode/year2025/cmd/util"
	"fmt"
	"strings"
)

const startNodeName = "you"
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

func main() {
	graph := parseInput(util.LoadInput(11))

	pathCounter := 0
	optimize.ForAllPathes(graph.FindNode(startNodeName), func(path optimize.Path) {
		pathCounter++
	})
	fmt.Println(pathCounter)
}
