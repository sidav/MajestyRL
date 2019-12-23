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
	u.asUnit.intent.x, u.asUnit.intent.y = tBld.getCenter()
	ux, uy := u.getCoords()
	builderCoeff := 1
	if !tBld.asBuilding.isUnderConstruction() {
		if tBld.asBuilding.beingConstructed != nil {
			tBld.asBuilding.beingConstructed = nil
			u.faction.reportToPlayer("our new building is complete!")
		}
		u.asUnit.intent.fulfillBidIfExists()
		u.asUnit.intent = nil
		return
	}
	if tBld.IsCloseupToCoords(ux, uy) {
		if !tBld.asBuilding.hasBeenPlaced {
			CURRENT_MAP.addBuilding(tBld, true)
		}
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
	if u.asUnit.registeredIn == nil {
		// find new home 
		for _, p := range CURRENT_MAP.pawns {
			if p.isBuilding() && !p.asBuilding.isUnderConstruction() {
				bsd := getBuildingStaticDataFromTable(p.asBuilding.code)
				// TODO: register not only the workers 
				if p.asBuilding.currWorkers < bsd.maxWorkers {
					u.asUnit.registeredIn = p 
					p.registerPawnHere(u)
				}
			}
		}
	}
	tBld := u.asUnit.registeredIn
	if tBld == nil {
		u.asUnit.intent = nil // unit decides to maybe search for other things to do 
		return 
	}
	u.asUnit.intent.targetPawn = tBld
	u.asUnit.intent.x, u.asUnit.intent.y = tBld.getCenter()
	if tBld.IsCloseupToCoords(ux, uy) {
		CURRENT_MAP.putUnitIntoBuilding(u, tBld)
		u.asUnit.intent = nil
	} else {
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST)
	}
}
