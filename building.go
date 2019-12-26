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
	return getBuildingStaticDataFromTable(b.code).app
}

func (b *building) isUnderConstruction() bool {
	return b.beingConstructed != nil && !b.beingConstructed.isCompleted()
}

func (b *building) getSize() (int, int) {
	h := len(b.getAppearance().chars)
	w := len(b.getAppearance().chars[0])
	return w, h
}

func (bld *pawn) addPawnToPawnsInside(p *pawn) {
	b := bld.asBuilding
	if b.pawnsInside == nil {
		b.pawnsInside = make([]*pawn, 0)
	}
	b.pawnsInside = append(b.pawnsInside, p)
}

func (bld *pawn) registerPawnHere(p *pawn) {
	b := bld.asBuilding
	if b.pawnsRegistered == nil {
		b.pawnsRegistered = make([]*pawn, 0)
	}
	b.pawnsRegistered = append(b.pawnsRegistered, p)
	p.asUnit.registeredIn = bld 
	bld.asBuilding.recalculateCurrResidents()
}

func (bld *pawn) AddAndRegisterNewPawn(p *pawn) {
	bld.addPawnToPawnsInside(p)
	bld.registerPawnHere(p)
}

func (b *building) recalculateCurrResidents() {
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
			return 
		}
	}
}

func (b *building) addPawnToInside(p *pawn) {
	b.pawnsInside = append(b.pawnsInside, p)
}

func (b *building) removePawnFromRegistered(p *pawn) {
	for i, pi := range b.pawnsRegistered {
		if p == pi {
			b.pawnsRegistered = append(b.pawnsRegistered[:i], b.pawnsRegistered[i+1:]...) // ow it's fucking... magic!
			return 
		}
	}
}
