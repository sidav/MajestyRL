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
	case INTENT_PATROL:
		p.executePatrol()
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
		// BUG: insufficient HP added for buildings with too large maxHitpoints
		hpToAdd := tBld.getMaxHitpoints() / (tBld.asBuilding.beingConstructed.maxConstructedAmount / builderCoeff)
		if hpToAdd == 0 {
			hpToAdd = 1
		}
		tBld.hitpoints += hpToAdd
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
	static := getUnitStaticDataFromTable(u.asUnit.code)
	if u.asUnit.registeredIn == nil {
		// find new home
		for _, p := range CURRENT_MAP.pawns {
			if p.isBuilding() && !p.asBuilding.isUnderConstruction() {
				bsd := getBuildingStaticDataFromTable(p.asBuilding.code)
				// TODO: register not only the workers
				if static.unitType == UTYPE_WORKER {
					if p.asBuilding.currWorkers < bsd.maxWorkers {
						u.asUnit.registeredIn = p
						p.registerPawnHere(u)
					}
				} else if static.unitType == UTYPE_GUARD {
					if p.asBuilding.currGuards < bsd.maxGuards {
						u.asUnit.registeredIn = p
						p.registerPawnHere(u)
					}
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

func (u *pawn) executePatrol() {
	ux, uy := u.getCenter()
	// static := getUnitStaticDataFromTable(u.asUnit.code)
	currIntent := u.asUnit.intent
	tBld := u.asUnit.registeredIn
	if tBld == nil {
		u.faction.reportToPlayer("Nothing to patrol, now procrastinating!")
		u.asUnit.intent = nil // unit decides to maybe search for other things to do
		return
	}

	w, h := tBld.getSize()
	if currIntent.x == 0 && currIntent.y == 0 && tBld.x != 1 && tBld.y != 1 {
		// set the initial patrol point 
		currIntent.x, currIntent.y = tBld.x+w, tBld.y+h  	
	}

	if ux == currIntent.x && uy == currIntent.y {
		// decide next patrol point, moving counter-clockwise 
		if ux == tBld.x-1 && uy == tBld.y-1 {
			currIntent.y += h+1 
		}
		if ux == tBld.x-1 && uy == tBld.y+h {
			currIntent.x += w+1 
		}
		if ux == tBld.x+w && uy == tBld.y+h {
			currIntent.y -= h+1
		}
		if ux == tBld.x+w && uy == tBld.y-1 {
			currIntent.x -= w+1
		}
	}
	cx, _ := tBld.getCenter()
	if ux == cx-1 && uy == tBld.y+h {
		u.asUnit.intent = nil // finished patrolling 
		return 
	}
	u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST)
}
