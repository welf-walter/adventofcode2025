package optimize

import (
	"log"
	"math"
)

type Node interface {
	// what nodes can I reach?
	targets() []Path
	// how can I come to this?
	sources() []Path
	// is this a finishing node?
	isFinish() bool
}

type Path interface {
	from() Node
	to() Node
	cost() int
}

type PathCostMap map[Node]int

func calcCostMap(nodes []Node) (costMap PathCostMap) {
	costMap = PathCostMap{}
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
		for _, path := range target.sources() {
			currentCost, found := costMap[path.from()]
			if !found {
				currentCost = math.MaxInt
			}
			log.Printf("    is %v + %v < %v ?", targetCost, path.cost(), currentCost)
			if targetCost+path.cost() < currentCost {
				costMap[path.from()] = targetCost + path.cost()
				todoList = append(todoList, path.from())
			}
		}
	}

	return
}

////////////////////////////////////////////////////////////////////

type SimpleGraph struct {
	nodes      []SimpleNode
	pathes     []SimplePath
	finishNode SimpleNode
}

type SimplePath struct {
	graph     *SimpleGraph
	fromIndex int
	toIndex   int
	costValue int
}

func (path SimplePath) from() Node {
	return path.graph.nodes[path.fromIndex]
}

func (path SimplePath) to() Node {
	return path.graph.nodes[path.toIndex]
}

func (path SimplePath) cost() int {
	return path.costValue
}

type SimpleNode struct {
	graph *SimpleGraph
	index int
}

func (node SimpleNode) targets() []Path {
	t := []Path{}
	for _, path := range node.graph.pathes {
		if path.fromIndex == node.index {
			t = append(t, path)
		}
	}
	return t
}

func (node SimpleNode) sources() []Path {
	t := []Path{}
	for _, path := range node.graph.pathes {
		if path.toIndex == node.index {
			t = append(t, path)
		}
	}
	return t
}

func (node SimpleNode) isFinish() bool {
	return node == node.graph.finishNode
}
