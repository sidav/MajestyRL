package main 

import "github.com/sidav/golibrl/astar"

const (
	PATHFINDING_DEPTH_FASTEST = 1
)

func (g *gameMap) createCostMapForPathfinding() *[][]int {
	width, height := len(g.tileMap), len((g.tileMap)[0])

	costmap := make([][]int, width)
	for j := range costmap {
		costmap[j] = make([]int, height)
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// TODO: optimize by iterating through pawns separately
			if !(g.tileMap[i][j].isPassable()) || g.getPawnAtCoordinates(i, j) != nil {
				costmap[i][j] = -1
			}
		}
	}
	return &costmap
}

func (g *gameMap) getPathFromTo(fx, fy, tx, ty, accuracy int) *astar.Cell {
	if accuracy <= 0 {
		accuracy = astar.DEFAULT_PATHFINDING_STEPS
	}
	return astar.FindPath(g.createCostMapForPathfinding(), fx, fy, tx, ty, true, accuracy, true, true)
}
