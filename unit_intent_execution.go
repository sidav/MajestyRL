package main

// plain "mechanical" execution of intents here.
func (p *pawn) act() {
	if p.asUnit.intent == nil {
		p.asUnit.intent = &intent{itype: INTENT_RETURN_HOME}
	}
	switch p.asUnit.intent.itype {
	case INTENT_BUILD:
		p.executeBuildIntent()
	case INTENT_BRING_RESOURCES_TO_CONSTRUCTION:
		p.executeBringResourcesToConstructionIntent()
	case INTENT_REPAIR:
		p.executeRepairIntent()
	case INTENT_RETURN_HOME:
		p.executeReturnHome()
	case INTENT_COLLECT_TAXES:
		p.executeCollectTaxes()
	case INTENT_RETURN_TAXES:
		p.executeReturnTaxes()
	case INTENT_PATROL:
		p.executePatrolIntent()
	case INTENT_MINE:
		p.executeMineIntent()
	case INTENT_ATTACK:
		p.executeAttackIntent()
	}
}

// HIGHLY EXPERIMENTAL
func (u *pawn) executeBringResourcesToConstructionIntent() {
	tBld := u.asUnit.intent.targetPawn
	u.asUnit.intent.x, u.asUnit.intent.y = tBld.getCenter()
	ux, uy := u.getCoords()
	// bring resources to construction site
	cost := tBld.asBuilding.getStaticData().cost
	brought := tBld.asBuilding.asBeingConstructed.resourcesBroughtToConstruction
	if !tBld.asBuilding.areBroughtResourcesEnoughToStartCostruction() {
		// unit has the resource and it is appropriate for building
		rtype, ramount := u.asUnit.carriedResourceType, u.asUnit.carriedResourceAmount
		if cost.amount[rtype] > brought.amount[rtype] && ramount > 0 {
			if tBld.IsCloseupToCoords(ux, uy) {
				u.giveResourcesToBuilding(tBld)
				return
			} else {
				u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, true)
				return
			}
		} else {
			u.dropResources()
		}
		// 1. Decide what resource to bring
		var decided resourceType
		for rtype := range cost.amount {
			if amount, exists := brought.amount[rtype]; amount < cost.amount[rtype] || !exists {
				decided = rtype
				break
			}
		}
		closestBuilding := CURRENT_MAP.getNearestBuildingWithStorageOfType(ux, uy, decided)
		if closestBuilding.IsCloseupToCoords(ux, uy) {
			closestBuilding.faction.economy.currentResources.amount[decided] -= 5
			u.asUnit.carriedResourceAmount = 5
			u.asUnit.carriedResourceType = decided
		} else {
			cbx, cby := closestBuilding.getCenter()
			u.doMoveToCoords(cbx, cby, PATHFINDING_DEPTH_FASTEST)
		}
	} else {
		u.dropCurrentIntent()
	}
}

func (u *pawn) executeBuildIntent() {
	tBld := u.asUnit.intent.targetPawn
	u.asUnit.intent.x, u.asUnit.intent.y = tBld.getCenter()
	ux, uy := u.getCoords()
	builderCoeff := 1
	if !tBld.asBuilding.isUnderConstruction() {
		if tBld.asBuilding.asBeingConstructed != nil {
			tBld.asBuilding.asBeingConstructed = nil
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
		tBld.asBuilding.asBeingConstructed.currentConstructedAmount += builderCoeff
		// BUG: insufficient HP added for buildings with too large maxHitpoints
		hpToAdd := tBld.getMaxHitpoints() / (tBld.asBuilding.asBeingConstructed.maxConstructedAmount / builderCoeff)
		if hpToAdd == 0 {
			hpToAdd = 1
		}
		// Workaround fix (increasing starting HP) for the bug described above
		if tBld.hitpoints == 0 && tBld.getMaxHitpoints()%(tBld.asBuilding.asBeingConstructed.maxConstructedAmount/builderCoeff) > 0 {
			tBld.hitpoints = tBld.getMaxHitpoints() % (tBld.asBuilding.asBeingConstructed.maxConstructedAmount / builderCoeff)
		}

		tBld.hitpoints += hpToAdd
		if tBld.hitpoints > tBld.getMaxHitpoints() {
			tBld.hitpoints = tBld.getMaxHitpoints()
		}
		u.spendTime(TICKS_PER_TURN)
	} else {
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, true)
	}
}

func (u *pawn) executeRepairIntent() {
	tBld := u.asUnit.intent.targetPawn
	u.asUnit.intent.x, u.asUnit.intent.y = tBld.getCenter()
	ux, uy := u.getCoords()
	builderCoeff := 1
	if !tBld.isDamaged() {
		u.faction.reportToPlayer("our building is repaired!")
		u.asUnit.intent.fulfillBidIfExists()
		u.asUnit.intent = nil
		return
	}
	if tBld.IsCloseupToCoords(ux, uy) {
		hpToAdd := tBld.getMaxHitpoints() / (tBld.getMaxHitpoints() / builderCoeff)
		if hpToAdd == 0 {
			hpToAdd = 1
		}
		tBld.hitpoints += hpToAdd
		if tBld.hitpoints > tBld.getMaxHitpoints() {
			tBld.hitpoints = tBld.getMaxHitpoints()
		}
		u.spendTime(TICKS_PER_TURN)
	} else {
		log.AppendMessage("MOVING TO REPAIR")
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, true)
	}
}

func (u *pawn) executeCollectTaxes() {
	tBld := u.asUnit.intent.targetPawn
	u.asUnit.intent.x, u.asUnit.intent.y = tBld.getCenter()
	ux, uy := u.getCoords()
	if tBld.IsCloseupToCoords(ux, uy) {
		if u.asUnit.carriedResourceType != RESTYPE_GOLD {
			u.asUnit.carriedResourceType = RESTYPE_GOLD
			u.asUnit.carriedResourceAmount = 0
		}
		u.asUnit.carriedResourceAmount += tBld.asBuilding.accumulatedGoldAmount
		tBld.asBuilding.accumulatedGoldAmount = 0
		u.spendTime(TICKS_PER_TURN)
		ULOGIC.reconsiderSituation(u) // decide new intent
	} else {
		log.AppendMessage("MOVING TO COLLECT")
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, true)
	}
}

func (u *pawn) executeReturnTaxes() {
	tBld := u.asUnit.intent.targetPawn
	u.asUnit.intent.x, u.asUnit.intent.y = tBld.getCenter()
	ux, uy := u.getCoords()
	if tBld.IsCloseupToCoords(ux, uy) {
		u.giveResourcesToBuilding(tBld)
		ULOGIC.reconsiderSituation(u) // decide new intent
	} else {
		log.AppendMessage("MOVING TO RETURN")
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, true)
	}
}

func (u *pawn) executeMineIntent() {
	const (
		// TODO: both should be removed 
		TIME_FOR_MINING = 50
		AMOUNT_MINED    = 5
	)
	ux, uy := u.getCoords()
	currIntent := u.asUnit.intent
	ix, iy := currIntent.getCoords()
	if CURRENT_MAP.getResourcesAtCoords(ix, iy) == nil || CURRENT_MAP.getResourcesAtCoords(ix, iy).amount == 0 {
		CURRENT_MAP.removeResourcesAtCoords(ix, iy)
		CURRENT_MAP.removeBid(u.asUnit.intent.sourceBid)
		u.asUnit.intent = nil
		return
	}
	minedType := CURRENT_MAP.getResourcesAtCoords(ix, iy).resType
	// if intent has no target building, select closest TO THE MINING SITE building which can store gold
	if currIntent.targetPawn == nil {
		buildingToReturn := CURRENT_MAP.getNearestBuildingWithStorageOfType(ix, iy, minedType)
		currIntent.targetPawn = buildingToReturn
		if buildingToReturn == nil {
			u.asUnit.intent = nil
		}
	}

	if u.asUnit.carriedResourceAmount == 0 {
		if CURRENT_MAP.getResourcesAtCoords(ix, iy).amount <= 0 || u.faction.economy.currentResources.amount[minedType] >= u.faction.economy.maxResources[minedType] {
			currIntent.sourceBid.drop()
			u.asUnit.intent = nil
			return
		}
		if u.IsCloseupToCoords(ix, iy) {
			u.spendTime(TIME_FOR_MINING)
			u.asUnit.carriedResourceAmount = AMOUNT_MINED
			u.asUnit.carriedResourceType = minedType
			CURRENT_MAP.getResourcesAtCoords(ix, iy).amount -= AMOUNT_MINED
		} else {
			u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, false)
		}
	} else { // return with gold and drop intent
		if currIntent.targetPawn.IsCloseupToCoords(ux, uy) {
			u.faction.economy.currentResources.amount[minedType] += u.asUnit.carriedResourceAmount
			u.asUnit.carriedResourceAmount = 0
			currIntent.sourceBid.drop()
			u.asUnit.intent = nil
			return
		}
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, true)
	}
}

func (u *pawn) executeReturnHome() {
	ux, uy := u.getCenter()
	static := getUnitStaticDataFromTable(u.asUnit.code)
	if u.asUnit.registeredIn == nil {
		// find new home
		for _, p := range CURRENT_MAP.pawns {
			if p.isBuilding() && !p.asBuilding.isUnderConstruction() {
				if p.asBuilding.canAffordNewResident(static.code) {
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
	if tBld.IsCloseupToCoords(ux, uy) {
		CURRENT_MAP.putUnitIntoBuilding(u, tBld)
		u.asUnit.intent = nil
	} else {
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, true)
	}
}

func (u *pawn) executePatrolIntent() {
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
	if currIntent.x == 0 && currIntent.y == 0 && tBld.x != 0 && tBld.y != 0 {
		// set the initial patrol point
		currIntent.x, currIntent.y = tBld.x+w, tBld.y+h
	}

	if ux == currIntent.x && uy == currIntent.y {
		if ULOGIC.checkForEnemiesAndAct(u) {
			return
		}
		// decide next patrol point, moving counter-clockwise
		if ux == tBld.x-1 && uy == tBld.y-1 {
			currIntent.y += h + 1
		}
		if ux == tBld.x-1 && uy == tBld.y+h {
			currIntent.x += w + 1
		}
		if ux == tBld.x+w && uy == tBld.y+h {
			currIntent.y -= h + 1
		}
		if ux == tBld.x+w && uy == tBld.y-1 {
			currIntent.x -= w + 1
		}
	}
	cx, _ := tBld.getCenter()
	if ux == cx-1 && uy == tBld.y+h {
		u.asUnit.intent = nil // finished patrolling
		return
	}
	u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, false)
}

func (u *pawn) executeAttackIntent() {
	ux, uy := u.getCoords()
	// TODO: ranged and magic attacks
	if u.asUnit.intent.targetPawn.IsCloseupToCoords(ux, uy) {
		u.performMeleeAttack(u.asUnit.intent.targetPawn)
		if !u.asUnit.intent.targetPawn.isAlive() {
			enemySet := ULOGIC.checkForEnemiesAndAct(u) // switch to another target
			if !enemySet {
				u.asUnit.intent = nil
			}
		}
	} else {
		u.doMoveToIntentTarget(PATHFINDING_DEPTH_FASTEST, true)
	}
}
