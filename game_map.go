package main

import (
	 geometry "github.com/sidav/golibrl/geometry"
)

var (
	mapW int
	mapH int
)

type gameMap struct {
	tileMap  [][]*tile
	factions []*faction
	pawns    []*pawn
}

func (g *gameMap) areCoordsValid(x, y int) bool {
	return geometry.AreCoordsInRect(x, y, 0, 0, mapW, mapH)
}

func (g *gameMap) addPawn(p *pawn) {
	g.pawns = append(g.pawns, p)
}

func (g *gameMap) addBuilding(b *pawn, asAlreadyConstructed bool) {
	if asAlreadyConstructed {
		b.asBuilding.hasBeenPlaced = true
	}
	g.addPawn(b)
}

func (g *gameMap) putUnitIntoBuilding(unit, building *pawn) {
	g.removePawn(unit)
	building.addPawnToPawnsInside(unit)
}

func (g *gameMap) removePawn(p *pawn) {
	for i := 0; i < len(g.pawns); i++ {
		if p == g.pawns[i] {
			g.pawns = append(g.pawns[:i], g.pawns[i+1:]...) // ow it's fucking... magic!
		}
	}
}

func (g *gameMap) getPawnAtCoordinates(x, y int) *pawn {
	for _, b := range g.pawns {
		if b.isOccupyingCoords(x, y) {
			return b
		}
	}
	return nil
}

func (g *gameMap) placePawnNearPawn(spawnThis, nearThis *pawn) {
	w, h := nearThis.getSize()
	spawnThis.x = nearThis.x + w/2  
	spawnThis.y = nearThis.y + h 
}

// func (g *gameMap) getPawnsInRect(x, y, w, h int) []*pawn {
// 	var arr []*pawn
// 	for _, p := range g.pawns {
// 		cx, cy := p.getCenter()
// 		if p.isBuilding() {
// 			bw, bh := p.getSize()
// 			if geometry.AreTwoCellRectsOverlapping(x, y, w, h, p.x, p.y, bw, bh) {
// 				arr = append(arr, p)
// 			}
// 		} else {
// 			if geometry.AreCoordsInRect(cx, cy, x, y, w, h) {
// 				arr = append(arr, p)
// 			}
// 		}
// 	}
// 	return arr
// }

// func (g *gameMap) getEnemyPawnsInRadiusFromPawn(p *pawn, radius int, f *faction) []*pawn {
// 	var arr []*pawn
// 	for _, p2 := range g.pawns {
// 		if p2.faction != f {
// 			if p.isInDistanceFromPawn(p2, radius) {
// 				arr = append(arr, p2)
// 				continue
// 			}
// 		}
// 	}
// 	return arr
// }

// func (g *gameMap) getBuildingAtCoordinates(x, y int) *pawn {
// 	for _, b := range g.pawns {
// 		if b.isOccupyingCoords(x, y) {
// 			return b
// 		}
// 	}
// 	return nil
// }

// func (g *gameMap) isConstructionSiteBlockedByUnitOrBuilding(x, y, w, h int, tight bool) bool {
// 	for _, p := range g.pawns {
// 		if p.isBuilding() {
// 			si := p.buildingInfo.getBuildingStaticInfo()
// 			if si.allowsTightPlacement && tight {
// 				if geometry.AreTwoCellRectsOverlapping(x, y, w, h, p.x, p.y, si.w, si.h) {
// 					return true
// 				}
// 			} else if geometry.AreTwoCellRectsOverlapping(x-1, y-1, w+2, h+2, p.x, p.y, si.w, si.h) {
// 				// -1s and +2s are to prevent tight placement...
// 				// ..and ensure that there always will be at least 1 cell between buildings.
// 				return true
// 			}
// 		} else {
// 			cx, cy := p.getCenter()
// 			if geometry.AreCoordsInRect(cx, cy, x, y, w, h) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func (g *gameMap) canBuildingBeBuiltAt(b *pawn, cx, cy int) bool {
// 	b_w, b_h := b.getSize()
// 	si := b.buildingInfo.getBuildingStaticInfo()
// 	bx := cx - b_w/2
// 	by := cy - b_h/2
// 	if bx < 0 || by < 0 || bx+b_w >= mapW || by+b_h >= mapH {
// 		return false
// 	}
// 	if si.canBeBuiltOnMetalOnly && g.getNumberOfMetalDepositsInRect(bx, by, b_w, b_h) == 0 {
// 		return false
// 	}
// 	if si.canBeBuiltOnThermalOnly && g.getNumberOfThermalDepositsInRect(bx, by, b_w, b_h) == 0 {
// 		return false
// 	}
// 	for x := bx; x < bx+b_w; x++ {
// 		for y := by; y < by+b_h; y++ {
// 			if !g.tileMap[x][y].isPassable {
// 				return false
// 			}
// 		}
// 	}
// 	if g.isConstructionSiteBlockedByUnitOrBuilding(bx, by, b_w, b_h, si.allowsTightPlacement) {
// 		return false
// 	}
// 	return true
// }
