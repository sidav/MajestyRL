package main 

// plain "mechanical" execution of intents here. 
func (p *pawn) act() {
	if p.asUnit.intent == nil {
		// TODO 
		return 
	}
	switch p.asUnit.intent.itype {
	case INTENT_BUILD:
		p.actByBuildIntent()
	}
}



func (u *pawn) actByBuildIntent() {
	tBld := u.asUnit.intent.targetPawn
	ux, uy := u.getCoords()
	builderCoeff := 1 
	if tBld.IsCloseupToCoords(ux, uy) {
		if !tBld.asBuilding.isUnderConstruction() {
			tBld.asBuilding.beingConstructed = nil
			u.asUnit.intent = nil
			return 
			// u.reportOrderCompletion("Construction completed")
		}
		tBld.asBuilding.beingConstructed.currentConstructedAmount += builderCoeff
		tBld.hitpoints += tBld.getMaxHitpoints() / (tBld.asBuilding.beingConstructed.maxConstructedAmount / builderCoeff)
		if tBld.hitpoints > tBld.getMaxHitpoints() {
			tBld.hitpoints = tBld.getMaxHitpoints()
		}
	} else {
		u.doMoveToIntentTarget(10)
	}
}
