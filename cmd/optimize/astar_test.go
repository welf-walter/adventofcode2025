package optimize

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maze() [5]string {
	return [5]string{
		".....",
		"..#..",
		"S.#.Z",
		"..#..",
		"....."}
}

type position struct {
	x, y int
}

type action struct {
	dx, dy int
}

func start() position  { return position{0, 2} }
func target() position { return position{4, 2} }

func (p position) estimatedCost() int {
	return 0 +
		max(p.x-target().x, target().x-p.x) +
		max(p.y-target().y, target().y-p.y)
}

func positionOk(p position) bool {
	if p.y < 0 || p.y >= len(maze()) {
		return false
	}
	if p.x < 0 || p.x >= len(maze()[p.y]) {
		return false
	}
	return maze()[p.y][p.x] != '#'
}

func (p position) possibleActions() (actions /*[]action*/ []AstarAction) {
	if positionOk(position{p.x - 1, p.y}) {
		actions = append(actions, action{-1, 0})
	}
	if positionOk(position{p.x, p.y - 1}) {
		actions = append(actions, action{0, -1})
	}
	if positionOk(position{p.x + 1, p.y}) {
		actions = append(actions, action{+1, 0})
	}
	if positionOk(position{p.x, p.y + 1}) {
		actions = append(actions, action{0, +1})
	}
	return
}

func (a action) targetNode(startNode /*position*/ AstarNode) AstarNode {
	pos := startNode.(position)
	return position{pos.x + a.dx, pos.y + a.dy}
}

func (a action) actualCost(startNode /*position*/ AstarNode) int {
	return 1
}

func concreteActionList(in []AstarAction) (out []action) {
	out = make([]action, len(in))
	for i := range in {
		out[i] = in[i].(action)
	}
	return
}

func TestAstar(t *testing.T) {
	assert := assert.New(t)

	s := start()
	assert.Equal([]action{{0, -1}, {1, 0}, {0, +1}}, concreteActionList(s.possibleActions()))

	assert.Equal(4, s.estimatedCost())

	path, mini := findMinimalPath(s)
	assert.Equal(8, mini)
	/*	assert.Equal(AstarPath{
		position{0, 2},
		position{0, 1},
		position{0, 0},
		position{1, 0},
		position{2, 0},
		position{3, 0},
		position{4, 0},
		position{4, 1},
		position{4, 2},
	}, path)*/
	log.Print(path)

}
