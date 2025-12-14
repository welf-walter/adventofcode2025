package optimize

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func example1() (nodes []Node) {
	graph := SimpleGraph{}
	//  0   ->   1
	//  V        V
	//  2   ->   3
	for i := range 4 {
		graph.nodes = append(graph.nodes, SimpleNode{&graph, i})
		nodes = append(nodes, graph.nodes[i])
	}

	graph.pathes = append(graph.pathes, SimplePath{&graph, 0, 1, 1})
	graph.pathes = append(graph.pathes, SimplePath{&graph, 0, 2, 1})
	graph.pathes = append(graph.pathes, SimplePath{&graph, 1, 3, 1})
	graph.pathes = append(graph.pathes, SimplePath{&graph, 2, 3, 1})

	graph.finishNode = graph.nodes[3]

	return nodes
}

func TestCostMap(t *testing.T) {
	assert := assert.New(t)
	nodes := example1()
	assert.Equal(4, len(nodes))
	assert.Equal(1, len(nodes[1].sources()))
	assert.Equal(1, len(nodes[1].targets()))
	assert.Equal(true, nodes[3].isFinish())

	costMap := calcCostMap(nodes)
	log.Println(costMap)
	assert.Equal(costMap[nodes[0]], 2)
}
