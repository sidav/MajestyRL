package main

import "github.com/sidav/golibrl/astar"
import "time"

const (
	PATHFINDING_DEPTH_FASTEST = 5
)

func (g *gameMap) getCostMapForPathfinding() *[][]int {
	width, height := len(g.tileMap), len((g.tileMap)[0])
	if len(g.costMap) == 0 {
		g.costMap = make([][]int, width)
		for j := range g.costMap {
			g.costMap[j] = make([]int, height)
		}
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if !(g.tileMap[i][j].isPassable()) { // || g.getPawnAtCoordinates(i, j) != nil {
				g.costMap[i][j] = -1
			} else {
				g.costMap[i][j] = 0
			}
		}
	}
	for _, p := range g.pawns {
		w, h := p.getSize()
		x, y := p.getCoords()
		if w == 1 && h == 1 {
			g.costMap[x][y] = -1
		} else {
			for i := x; i < x+w; i++ {
				for j := y; j < y+h; j++ {
					if g.areCoordsValid(i, j) {
						g.costMap[i][j] = -1
					}
				}
			}
		}
	}
	return &g.costMap
}

func (g *gameMap) getPathFromTo(fx, fy, tx, ty, accuracy int) *astar.Cell {
	start := time.Now()
	if accuracy <= 0 {
		accuracy = astar.DEFAULT_PATHFINDING_STEPS
	}
	// return astar.FindPath(g.getCostMapForPathfinding(), fx, fy, tx, ty, true, accuracy, true, true)
	res := astar.FindPath(g.getCostMapForPathfinding(), fx, fy, tx, ty, true, accuracy, true, false)
	totalPathfindingTimes += time.Since(start) / time.Nanosecond
	totalPathfindings += 1 
	if totalPathfindings == 10 {
		avgPathfindingTime = totalPathfindingTimes / 10
		totalPathfindingTimes = 0 
		totalPathfindings = 0 
	}
	return res 
}
