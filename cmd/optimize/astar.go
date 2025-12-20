package optimize

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
