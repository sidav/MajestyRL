package main

type building struct {
	hasBeenPlaced    bool
	code             string
	beingConstructed *underConstructionData

	currWorkers, currTCs, currGuards, currRoyalGurads int
	pawnsInside                                       []*pawn
	pawnsRegistered                                   []*pawn
}

func (b *building) getAppearance() *buildingAppearance {
	return staticBuildingDataTable[b.code].app
}

func (b *building) isUnderConstruction() bool {
	return b.beingConstructed != nil && !b.beingConstructed.isCompleted()
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
	if b.pawnsRegistered == nil {
		b.pawnsRegistered = make([]*pawn, 0)
	}
	b.pawnsRegistered = append(b.pawnsRegistered, p)
}

func (b *building) AddAndRegisterNewPawn(p *pawn) {
	b.addPawnToPawnsInside(p)
	b.registerPawnHere(p)
}

func (b *building) recalculateCurrValues() {
	b.currWorkers = 0
	for _, p := range b.pawnsRegistered {
		if !p.isBuilding() {
			b.currWorkers++
		}
	}
}

func (b *building) removePawnFromInside(p *pawn) {
	for i, pi := range b.pawnsInside {
		if p == pi {
			b.pawnsInside = append(b.pawnsInside[:i], b.pawnsInside[i+1:]...) // ow it's fucking... magic!
		}
	}
}

func (b *building) removePawnFromRegistered(p *pawn) {
	for i, pi := range b.pawnsRegistered {
		if p == pi {
			b.pawnsRegistered = append(b.pawnsRegistered[:i], b.pawnsRegistered[i+1:]...) // ow it's fucking... magic!
		}
	}
}
