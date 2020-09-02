package main

type building struct {
	hasBeenPlaced         bool
	code                  string
	asBeingConstructed    *underConstructionData
	accumulatedGoldAmount int // taxes, treasure, etc

	currentResidents map[string]int
	pawnsInside      []*pawn
	pawnsRegistered  []*pawn
}

func (b *building) getAppearance() *buildingAppearance {
	return getBuildingStaticDataFromTable(b.code).app
}

func (b *building) getStaticData() *buildingStaticData {
	return getBuildingStaticDataFromTable(b.code)
}

func (b *building) isUnderConstruction() bool {
	return b.asBeingConstructed != nil && !b.asBeingConstructed.isCompleted()
}

func (b *building) areBroughtResourcesEnoughToStartCostruction() bool {
	cost := b.getStaticData().cost
	return b.asBeingConstructed.resourcesBroughtToConstruction.canSubstract(cost)
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
	for i := range b.currentResidents {
		b.currentResidents[i] = 0
	}
	for _, p := range b.pawnsRegistered {
		if p.isAlive() {
			residentCode := p.asUnit.getStaticData().code
			b.currentResidents[residentCode] += 1
		} else {
			b.removePawnFromRegistered(p)
		}
	}
}

func (b *building) canAffordNewResident(newResidentCode string) bool {
	// TODO: optimize that bullshit
	for index, code := range b.getStaticData().housing_unittypes {
		if code == newResidentCode && b.currentResidents[code] < b.getStaticData().housing_max_residents[index] {
			return true
		}
	}
	return false
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
