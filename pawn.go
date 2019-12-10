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

func (p *pawn) getSize() (int, int) { 
	if p.isBuilding() {
		return p.asBuilding.getSize()
	}
	return 1, 1 
}

func (p *pawn) getCenter() (int, int) {
	if p.isBuilding() {
		b_w, b_h := p.getSize()
		return p.x + b_w/2, p.y + b_h/2
	} else {
		return p.x, p.y 
	}
}

func (p *pawn) isOccupyingCoords(x, y int) bool {
	return p.x == x && p.y == y 
	// TODO: buildings
}

func (p *pawn) isBuilding() bool {
	return p.asBuilding != nil 
}