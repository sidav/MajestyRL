package main

const (
	CONSIDER_NONBIDS_EVERY = 10 * TICKS_PER_TURN // ticks
)

type unitLogic struct{}

// returns true if the unit wants to leave the building
func (ul *unitLogic) wantsToLeaveBuilding(u *pawn) bool {
	return u.asUnit.intent != nil
}

func (ul *unitLogic) decideNewIntent(p *pawn) {
	if p.asUnit.intent == nil {
		ul.checkForEnemiesAndAct(p)
		if p.faction == nil {
			// Neutral units behave differently

		} else {
			ul.considerBids(p)
		}
	}

	if p.asUnit.intent == nil && CURRENT_TICK%CONSIDER_NONBIDS_EVERY == 0 {
		ul.considerSituation(p)
	}
}

func (ul *unitLogic) considerBids(p *pawn) {
	startingBid := rnd.Rand(len(CURRENT_MAP.bids))
	static := getUnitStaticDataFromTable(p.asUnit.code)
	for i := range CURRENT_MAP.bids {
		consideredBid := CURRENT_MAP.bids[(i+startingBid)%len(CURRENT_MAP.bids)]
		if consideredBid.isFulfilled() || !consideredBid.isVacant() {
			continue // skip fulfilled bids (they are cleared automatically)
		}
		switch consideredBid.intent_type_for_this_bid {
		case INTENT_BUILD:
			if consideredBid.targetPawn.asBuilding.isUnderConstruction() {
				if static.canBuild {
					if !RESOURCE_HAULING || consideredBid.targetPawn.asBuilding.areBroughtResourcesEnoughToStartCostruction() {
						p.asUnit.intent = consideredBid.dispatchIntent()
					} else {
						p.asUnit.intent = &intent{itype: INTENT_BRING_RESOURCES_TO_CONSTRUCTION, targetPawn: consideredBid.targetPawn}
					}
					return
				}
			}
		case INTENT_MINE:
			rx, ry := consideredBid.x, consideredBid.y
			minedType := CURRENT_MAP.getResourcesAtCoords(rx, ry).resType
			if static.canMine && p.faction.economy.currentResources.amount[minedType] < p.faction.economy.maxResources[minedType] {
				p.asUnit.intent = consideredBid.dispatchIntent()
				return
			}
		}
	}
}

// returns true if intent was changed
func (ul *unitLogic) checkForEnemiesAndAct(p *pawn) bool {
	static := p.asUnit.getStaticData()
	// should we attac something?
	x, y := p.getCenter()
	enemiesInRange := CURRENT_MAP.getAliveEnemyPawnsInRangeFrom(p.faction, static.sightRange, x, y)
	if len(*enemiesInRange) > 0 {
		if p.weapon != nil {
			enemyInRange := (*enemiesInRange)[rnd.Rand(len(*enemiesInRange))]
			log.AppendMessage("Target sighted, should attack now!")
			px, py := enemyInRange.getCenter()
			p.asUnit.intent = &intent{itype: INTENT_ATTACK, targetPawn: enemyInRange, x: px, y: py}
		} else {
			p.asUnit.intent = nil // flee to home
		}
		return true
	}
	return false
}

func (ul *unitLogic) considerSituation(p *pawn) {
	enemiesWereConsidered := ul.checkForEnemiesAndAct(p)
	if enemiesWereConsidered {
		return
	}
	static := getUnitStaticDataFromTable(p.asUnit.code)
	startingPawnIndex := rnd.Rand(len(CURRENT_MAP.pawns))
	// should we return with collected gold?
	if p.asUnit.getStaticData().canCollectTaxes && p.asUnit.carriedResourceAmount >= RETURN_MINIMUM_TAX && p.asUnit.carriedResourceType == RESTYPE_GOLD {
		p.asUnit.intent = &intent{itype: INTENT_RETURN_TAXES, targetPawn: p.asUnit.registeredIn}
	}
	for i := range CURRENT_MAP.pawns {
		consideredPawn := CURRENT_MAP.pawns[(i+startingPawnIndex)%len(CURRENT_MAP.pawns)]
		if consideredPawn.isBuilding() && p.faction == consideredPawn.faction {
			// should we collect taxes from it?
			if p.asUnit.getStaticData().canCollectTaxes && consideredPawn.asBuilding.accumulatedGoldAmount >= COLLECT_THAT_MINIMUM {
				p.asUnit.intent = &intent{itype: INTENT_COLLECT_TAXES, targetPawn: consideredPawn}
			}
			if static.canBuild { // try to build and/or repair building
				// should we build it?
				if consideredPawn.asBuilding.isUnderConstruction() {
					x, y := consideredPawn.getCenter()
					if !RESOURCE_HAULING || consideredPawn.asBuilding.areBroughtResourcesEnoughToStartCostruction() {
						p.asUnit.intent = &intent{itype: INTENT_BUILD, targetPawn: consideredPawn, x: x, y: y}
					} else {
						p.asUnit.intent = &intent{itype: INTENT_BRING_RESOURCES_TO_CONSTRUCTION, targetPawn: consideredPawn, x: x, y: y}
					}
					return
				} else { // should we repair it?
					if consideredPawn.isDamaged() {
						x, y := consideredPawn.getCenter()
						p.asUnit.intent = &intent{itype: INTENT_REPAIR, targetPawn: consideredPawn, x: x, y: y}
						return
					}
				}
			}
		}
	}
	if p.weapon != nil {
		p.asUnit.intent = &intent{itype: INTENT_PATROL}
	}
}

func (ul *unitLogic) reconsiderSituation(p *pawn) {
	p.asUnit.intent = nil
	ul.considerSituation(p)
}
