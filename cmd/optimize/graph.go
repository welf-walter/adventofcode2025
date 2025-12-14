package optimize

import (
	"log"
	"math"
	"slices"
)

type Node interface {
	// what nodes can I reach?
	targets() []Edge
	// how can I come to this?
	sources() []Edge
	// is this a finishing node?
	isFinish() bool
}

type Edge interface {
	from() Node
	to() Node
	cost() int
}

type Path []Node

func ForAllPathes(from Node, yield func(path Path)) {
	path := Path{from}
	forAllPathes(path, yield)
}

func forAllPathes(path Path, yield func(path Path)) {

	current := path[len(path)-1]
	log.Printf("current path = %v", current)

	for _, edge := range current.targets() {
		next := edge.to()
		if slices.Contains(path, next) {
			log.Printf("  %v: been there. done that", next)
			continue
		}

		newPath := make([]Node, len(path), len(path)+1)
		copy(newPath, path)
		newPath = append(newPath, next)

		if next.isFinish() {
			log.Printf("  %v: finish", next)
			yield(newPath)
		} else {
			log.Printf("  %v: recurse", next)
			forAllPathes(newPath, yield)
		}
	}
}

////////////////////////////////////////////////////////////////////

type EdgeCostMap map[Node]int

func calcCostMap(nodes []Node) (costMap EdgeCostMap) {
	costMap = EdgeCostMap{}
	todoList := []Node{}

	// init all finish nodes with zero
	for _, n := range nodes {
		if n.isFinish() {
			costMap[n] = 0
			todoList = append(todoList, n)
		}
	}

	for todoIndex := 0; todoIndex < len(todoList); todoIndex++ {
		target := todoList[todoIndex]
		targetCost := costMap[target]
		log.Printf("#%v: Inspect node %v with cost %v", todoIndex, target, targetCost)
		for _, edge := range target.sources() {
			currentCost, found := costMap[edge.from()]
			if !found {
				currentCost = math.MaxInt
			}
			log.Printf("    is %v + %v < %v ?", targetCost, edge.cost(), currentCost)
			if targetCost+edge.cost() < currentCost {
				costMap[edge.from()] = targetCost + edge.cost()
				todoList = append(todoList, edge.from())
			}
		}
	}

	return
}

////////////////////////////////////////////////////////////////////

type SimpleGraph struct {
	nodes      []SimpleNode
	edges      []SimpleEdge
	finishNode SimpleNode
}

type SimpleEdge struct {
	graph     *SimpleGraph
	fromIndex int
	toIndex   int
	costValue int
}

func (edge SimpleEdge) from() Node {
	return edge.graph.nodes[edge.fromIndex]
}

func (edge SimpleEdge) to() Node {
	return edge.graph.nodes[edge.toIndex]
}

func (edge SimpleEdge) cost() int {
	return edge.costValue
}

type SimpleNode struct {
	graph *SimpleGraph
	index int
	name  string
}

func (node SimpleNode) targets() []Edge {
	t := []Edge{}
	for _, edge := range node.graph.edges {
		if edge.fromIndex == node.index {
			t = append(t, edge)
		}
	}
	return t
}

func (node SimpleNode) sources() []Edge {
	t := []Edge{}
	for _, edge := range node.graph.edges {
		if edge.toIndex == node.index {
			t = append(t, edge)
		}
	}
	return t
}

func (node SimpleNode) isFinish() bool {
	return node == node.graph.finishNode
}

func (node SimpleNode) String() string {
	return node.name
}

func (graph *SimpleGraph) addNode(name string) Node {
	index := len(graph.nodes)
	graph.nodes = append(graph.nodes, SimpleNode{graph, index, name})
	return graph.nodes[index]
}

func (graph *SimpleGraph) addEdge(fromIndex, toIndex int) Edge {
	index := len(graph.edges)
	graph.edges = append(graph.edges, SimpleEdge{graph, fromIndex, toIndex, 1})
	return graph.edges[index]
}
