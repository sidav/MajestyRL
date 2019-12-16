package main

const (
	CONSIDER_NONBIDS_EVERY = 1000 // ticks 
)

type unitLogic struct{}

// returns true if the unit wants to leave the building
func (ul *unitLogic) wantsToLeaveBuilding(u *pawn) bool {
	return u.asUnit.intent != nil 
}

func (ul *unitLogic) decideNewIntent(p *pawn) {
	if p.asUnit.intent == nil {
		ul.considerBids(p)
	}
	if p.asUnit.intent == nil && CURRENT_TICK % CONSIDER_NONBIDS_EVERY == 0 {
		ul.considerSituation(p)
	}
}

func (ul *unitLogic) considerBids(p *pawn) {
	startingBid := rnd.Rand(len(CURRENT_MAP.bids))
	static := staticUnitDataTable[p.asUnit.code]
	for i := range CURRENT_MAP.bids {
		consideredBid := CURRENT_MAP.bids[(i+startingBid) % len(CURRENT_MAP.bids)]
		if consideredBid.isFulfilled || !consideredBid.isVacant() {
			continue // skip fulfilled bids (they are cleared automatically)
		}
		switch consideredBid.intent_type_for_this_bid {
		case INTENT_BUILD:
			if static.canBuild {
				p.asUnit.intent = consideredBid.dispatchIntent()
				return 
			}
		}
	}

}

func (ul *unitLogic) considerSituation(p *pawn) {
	static := staticUnitDataTable[p.asUnit.code]
			if static.canBuild { // try to build and/or repair building
			startingPawnIndex := rnd.Rand(len(CURRENT_MAP.pawns))
			for i := range CURRENT_MAP.pawns {
				consideredPawn := CURRENT_MAP.pawns[(i+startingPawnIndex) % len(CURRENT_MAP.pawns)]
				if consideredPawn.isBuilding() && p.faction == consideredPawn.faction {
					// should we build it?
					if consideredPawn.asBuilding.beingConstructed != nil {
						x, y := consideredPawn.getCenter()
						p.asUnit.intent = &intent{itype: INTENT_BUILD, targetPawn: consideredPawn, x: x, y: y}
						return
					}
					// should we repair it?
					// TODO
				}
			}
		}
}
