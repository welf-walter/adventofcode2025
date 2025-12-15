package optimize

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

func (graph *SimpleGraph) AddNode(name string) Node {
	index := len(graph.nodes)
	graph.nodes = append(graph.nodes, SimpleNode{graph, index, name})
	return graph.nodes[index]
}

func (graph *SimpleGraph) addEdge(fromIndex, toIndex int) Edge {
	index := len(graph.edges)
	graph.edges = append(graph.edges, SimpleEdge{graph, fromIndex, toIndex, 1})
	return graph.edges[index]
}

func (graph *SimpleGraph) AddEdge(from, to SimpleNode) Edge {
	return graph.addEdge(from.index, to.index)
}

func (graph *SimpleGraph) FindNode(name string) SimpleNode {
	for _, n := range graph.nodes {
		if n.name == name {
			return n
		}
	}
	panic(name)
}

func (graph *SimpleGraph) SetFinish(name string) {
	graph.finishNode = graph.FindNode(name)
}

// mainly for tests
func (graph *SimpleGraph) GetTargets(name string) []string {
	targetNames := []string{}
	node := graph.FindNode(name)
	for _, target := range node.targets() {
		targetNames = append(targetNames, target.to().(SimpleNode).name)
	}
	return targetNames
}
