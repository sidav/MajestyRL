package main

type unitLogic struct{}

// returns true if the unit wants to leave the building
func (ul *unitLogic) wantsToLeaveBuilding(u *pawn) bool {
	return u.asUnit.intent != nil 
}

func (ul *unitLogic) decideNewIntent(p *pawn) {
	if p.asUnit.intent == nil {
		static := staticUnitDataTable[p.asUnit.code]

		if static.canBuild { // try to build and/or repair building
			for _, consideredPawn := range CURRENT_MAP.pawns {
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

}
