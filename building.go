package main

type building struct {
	hasBeenPlaced    bool
	code             string
	beingConstructed *underConstructionData

	currWorkers, currTCs, currGuards, currRoyalGurads int
	pawnsInside                                       []*pawn
	pawnsRegisteredHere                               []*pawn
}

func (b *building) getAppearance() *buildingAppearance {
	return staticBuildingDataTable[b.code].app
}

func (b *building) getSize() (int, int) {
	w := len(b.getAppearance().chars)
	h := len(b.getAppearance().chars[0])
	return w, h
}

func (b *building) addPawnToPawnsInside(p *pawn) {
	if b.pawnsInside == nil {
		b.pawnsInside = make([]*pawn, 0)
	}
	b.pawnsInside = append(b.pawnsInside, p)
}

func (b *building) registerPawnHere(p *pawn) {
	if b.pawnsRegisteredHere == nil {
		b.pawnsRegisteredHere = make([]*pawn, 0)
	}
	b.pawnsRegisteredHere = append(b.pawnsInside, p)
}

func (b *building) AddAndRegisterNewPawn(p *pawn) {
	b.addPawnToPawnsInside(p)
	b.registerPawnHere(p)
}

func (b *building) recalculateCurrValues() {
	b.currWorkers = 0
	for _, p := range b.pawnsRegisteredHere {
		if !p.isBuilding() {
			b.currWorkers++
		}
	}
}
