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
