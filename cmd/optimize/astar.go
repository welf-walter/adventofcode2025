package optimize

import (
	"log"
	"math"
	"slices"
)

// https://de.wikipedia.org/wiki/A*-Algorithmus

type AstarNode interface {
	// muss optimistisch sein!
	// Darf die tatsächlichen Kosten niemals überschätzen!
	estimatedCost() int
	possibleActions() []AstarAction
}

type AstarAction interface {
	targetNode(startNode AstarNode) AstarNode
	actualCost(startNode AstarNode) int
}

type AstarPath []AstarNode

type nodeWithCost struct {
	node AstarAction
	cost int
}
type nodeCost map[AstarNode]int

func dummy() {

}
func findBestNode(costMap nodeCost) (bestNode AstarNode, cost int) {
	mini := math.MaxInt

	for n := range costMap {
		mini = min(mini, costMap[n])
	}

	for n := range costMap {
		if costMap[n] == mini {
			return n, mini
		}
	}
	panic("no best node found")
}

func findMinimalPath(startNode AstarNode) (bestPath AstarPath, cost int) {

	//todoList := []nodeWithCost{startNode,0}
	todoList := map[AstarNode]int{}
	todoList[startNode] = 0

	predecessor := map[AstarNode]AstarNode{}

	doneList := map[AstarNode]bool{}

	for {

		if len(todoList) == 0 {
			panic("No solution found!")
		}

		currentNode, currentCost := findBestNode(todoList)
		log.Printf("Process node %v (cost = %v)", currentNode, currentCost)
		delete(todoList, currentNode)

		if currentNode.estimatedCost() == 0 {
			// build up path
			bestPath = make(AstarPath, 0)
			for n := currentNode; n != startNode; n = predecessor[n] {
				bestPath = append(bestPath, n)
			}
			bestPath = append(bestPath, startNode)
			slices.Reverse(bestPath)
			return bestPath, currentCost

		}
		doneList[currentNode] = true

		actions := currentNode.possibleActions()
		for _, action := range actions {
			log.Printf("  expand action %v", action)
			target := action.targetNode(currentNode)
			if doneList[target] {
				log.Printf("  been there. done that.")
				continue
			}

			actualCost := action.actualCost(currentNode)
			newCost := currentCost + actualCost

			costUpToNow, alternativeWay := todoList[target]
			if alternativeWay {
				if costUpToNow > newCost {
					log.Printf("  found better way to %v. %v + %v < %v", target, currentCost, actualCost, costUpToNow)
					todoList[target] = newCost
					predecessor[target] = currentNode
				} else {
					log.Printf("  no better way to %v. %v + %v >= %v", target, currentCost, actualCost, costUpToNow)
				}
			} else {
				log.Printf("  found way to %v.", target)
				todoList[target] = newCost
				predecessor[target] = currentNode
			}
		}
	}
}
