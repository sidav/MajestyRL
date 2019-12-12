package main

// plain "mechanical" execution of intents here.
func (p *pawn) act() {
	if p.asUnit.intent == nil {
		p.asUnit.intent = &intent{itype: INTENT_RETURN_HOME}
	}
	switch p.asUnit.intent.itype {
	case INTENT_BUILD:
		p.executeBuildIntent()
	case INTENT_RETURN_HOME:
		p.executeReturnHome()
	}
}

func (u *pawn) executeBuildIntent() {
	tBld := u.asUnit.intent.targetPawn
	ux, uy := u.getCoords()
	builderCoeff := 10
	if !tBld.asBuilding.isUnderConstruction() {
		if tBld.asBuilding.beingConstructed != nil {
			tBld.asBuilding.beingConstructed = nil
			reportToPlayer("our new building is complete!", u.faction)
		}
		u.asUnit.intent = nil
		return
	}
	if tBld.IsCloseupToCoords(ux, uy) {
		tBld.asBuilding.beingConstructed.currentConstructedAmount += builderCoeff
		tBld.hitpoints += tBld.getMaxHitpoints() / (tBld.asBuilding.beingConstructed.maxConstructedAmount / builderCoeff)
		if tBld.hitpoints > tBld.getMaxHitpoints() {
			tBld.hitpoints = tBld.getMaxHitpoints()
		}
		u.spendTime(TICKS_PER_TURN)
	} else {
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST)
	}
}

func (u *pawn) executeReturnHome() {
	ux, uy := u.getCenter()
	tBld := u.asUnit.registeredIn
	u.asUnit.intent.targetPawn = tBld
	u.asUnit.intent.x, u.asUnit.intent.y = tBld.getCenter()
	if tBld.IsCloseupToCoords(ux, uy) {
		CURRENT_MAP.putUnitIntoBuilding(u, tBld)
		u.asUnit.intent = nil
	} else {
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST)
	}
}
