package main

// anything that hax x,y coords on the map and can move
type pawn struct {
	asBuilding    *building
	asUnit        *unit
	hitpoints     int
	x, y          int
	faction       *faction
	nextTickToAct int
}

func (p *pawn) getCoords() (int, int) {
	return p.x, p.y
}

func (p *pawn) isOccupyingCoords(x, y int) bool {
	return p.x == x && p.y == y 
	// TODO: buildings
}
