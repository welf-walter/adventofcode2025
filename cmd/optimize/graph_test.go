package optimize

import (
	"fmt"
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
		name := fmt.Sprintf("Node%v", i)
		node := graph.AddNode(name)
		nodes = append(nodes, node)
	}

	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 3)

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

func TestForAllPathes(t *testing.T) {
	assert := assert.New(t)
	pathes := []Path{}

	nodes := example1()
	ForAllPathes(nodes[0], func(path Path) {
		pathes = append(pathes, path)
	})
	assert.Equal(Path{nodes[0], nodes[1], nodes[3]}, pathes[0])
	assert.Equal(Path{nodes[0], nodes[2], nodes[3]}, pathes[1])

}
