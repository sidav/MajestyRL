package main

import "github.com/sidav/golibrl/geometry"

// anything that hax x,y coords on the map and can move
type pawn struct {
	asBuilding    *building
	asUnit        *unit
	hitpoints     int
	x, y          int
	faction       *faction
	nextTickToAct int

	weapon      *item
	currentGold int
}

func (p *pawn) isTimeToAct() bool {
	return CURRENT_TICK >= p.nextTickToAct
}

func (p *pawn) getName() string {
	if p.isBuilding() {
		return getBuildingStaticDataFromTable(p.asBuilding.code).name
	}
	if p.isUnit() {
		return getUnitStaticDataFromTable(p.asUnit.code).name
	}
	return "getName() did some strange garbage"
}

func (p *pawn) setFactionTechAllowance() {
	if p.isBuilding() {
		bsd := getBuildingStaticDataFromTable(p.asBuilding.code)
		for _, allows := range bsd.allowsBuildings {
			if p.faction.allowedBuildings[allows] != TECH_DENIED {
				p.faction.allowedBuildings[allows] = TECH_ALLOWED
			}
		}
		for _, denies := range bsd.deniesBuildings {
			p.faction.allowedBuildings[denies] = TECH_DENIED
		}
	}
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
	if p.isBuilding() {
		bx, by := p.getCoords()
		w, h := p.getSize()
		return geometry.AreCoordsInRect(x, y, bx, by, w, h)
	}
	return p.x == x && p.y == y
}

func (p *pawn) IsCloseupToCoords(x, y int) bool {
	w, h := p.getSize()
	if p.isBuilding() {
		return !geometry.AreCoordsInRect(x, y, p.x, p.y, w, h) &&
			geometry.AreCoordsInRect(x, y, p.x-1, p.y-1, w+2, h+2)
	} else {
		return (x != p.x || y != p.y) && geometry.AreCoordsInRect(x, y, p.x-1, p.y-1, 3, 3)
	}
}

func (p *pawn) getMaxHitpoints() int {
	if p.isBuilding() {
		return getBuildingStaticDataFromTable(p.asBuilding.code).maxHitpoints
	}
	return p.asUnit.maxHitpoints
}

func (p *pawn) isBuilding() bool {
	return p.asBuilding != nil
}

func (p *pawn) isUnit() bool {
	return p.asUnit != nil
}

func (p *pawn) isAlive() bool {
	return p.hitpoints > 0
}
