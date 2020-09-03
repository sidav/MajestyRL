package main

import (
	"fmt"
)

// low level unit actions. 

func (u *pawn) spendTime(time int) {
	u.nextTickToAct = CURRENT_TICK + time
}

func (u *pawn) doMoveToIntentTarget(desiredAccuracy int, moveToTargetPawn bool) bool { // Returns true if route exists. TODO: rewrite
	intent := u.asUnit.intent

	ox, oy := intent.x, intent.y
	if intent.targetPawn != nil && moveToTargetPawn {
		ox, oy = intent.targetPawn.getCenter()
	}
	ux, uy := u.getCoords()
	var vx, vy int

	path := CURRENT_MAP.getPathFromTo(ux, uy, ox, oy, desiredAccuracy)
	if path != nil {
		vx, vy = path.GetNextStepVector()
	}

	if true { // TODO: if canMove() 
		if vx == 0 && vy == 0 {
			u.faction.reportToPlayer("no path to target!")
			u.spendTime(10 * TICKS_PER_TURN)
			u.asUnit.handleIntentUnsuccess()
			return false
		}
		u.x += vx
		u.y += vy
		u.spendTime(TICKS_PER_TURN)
	}
	return true
}

func (u *pawn) doMoveToCoords(ox, oy, desiredAccuracy int) bool { // Returns true if route exists. TODO: rewrite
	ux, uy := u.getCoords()
	var vx, vy int

	path := CURRENT_MAP.getPathFromTo(ux, uy, ox, oy, desiredAccuracy)
	if path != nil {
		vx, vy = path.GetNextStepVector()
	}

	if true { // TODO: if canMove()
		if vx == 0 && vy == 0 {
			u.faction.reportToPlayer("no path to target!")
			u.spendTime(10 * TICKS_PER_TURN)
			return false
		}
		u.x += vx
		u.y += vy
		u.spendTime(TICKS_PER_TURN)
	}
	return true
}

func (u *pawn) performMeleeAttack(target *pawn) {
	damage := u.weapon.weaponData.rollMeleeDamageDice()
	target.hitpoints -= damage
	log.AppendMessage(fmt.Sprintf("%s hits %s for %d damage!", u.getName(), target.getName(), damage))
	u.spendTime(u.weapon.weaponData.attackTime)

	x, y := target.getCoords()
	addBasicDecalToRender(x, y, 2)
}

func (u *pawn) dropResources() {
	rtype := u.asUnit.carriedResourceType
	ramount := u.asUnit.carriedResourceAmount
	u.faction.economy.addResource(ramount, rtype)
	u.asUnit.carriedResourceAmount = 0
}

func (u *pawn) giveResourcesToBuilding(b *pawn) {
	rtype := u.asUnit.carriedResourceType
	ramount := u.asUnit.carriedResourceAmount
	if b.asBuilding.isUnderConstruction() {
		neededAmount := b.asBuilding.getStaticData().cost.amount[rtype]
		currAmount := b.asBuilding.asBeingConstructed.resourcesBroughtToConstruction.amount[rtype]
		if _, exists := b.asBuilding.asBeingConstructed.resourcesBroughtToConstruction.amount[rtype]; exists {
			if currAmount + ramount <= neededAmount {
				b.asBuilding.asBeingConstructed.resourcesBroughtToConstruction.amount[rtype] += ramount
			} else {
				b.asBuilding.asBeingConstructed.resourcesBroughtToConstruction.amount[rtype] = neededAmount
				u.asUnit.carriedResourceAmount = currAmount + ramount - neededAmount
				u.dropResources()
			}
		} else {
			if currAmount + ramount <= neededAmount {
				b.asBuilding.asBeingConstructed.resourcesBroughtToConstruction.amount[rtype] = ramount
			} else {
				b.asBuilding.asBeingConstructed.resourcesBroughtToConstruction.amount[rtype] = neededAmount
				u.asUnit.carriedResourceAmount = currAmount + ramount - neededAmount
				u.dropResources()
			}
		}
	} else {
		b.faction.economy.addResource(ramount, rtype)
	}
	u.asUnit.carriedResourceAmount = 0
	u.spendTime(TICKS_PER_TURN) // TODO: adjust
}

func (u *unit) getCurrentIntentDescription() string {
	if u.intent == nil {
		return "Thinking..."
	}
	switch u.intent.itype {
	case INTENT_BUILD:
		return "Building..."
	case INTENT_COLLECT_TAXES:
		return "Collecting money..."
	case INTENT_RETURN_TAXES:
		return "Returning money..."
	case INTENT_REPAIR:
		return "Repairing..."
	case INTENT_MINE:
		return "Mining..."
	case INTENT_RETURN_HOME:
		return "Going to rest..."
	case INTENT_PATROL:
		return "Patrolling..."
	case INTENT_ATTACK:
		return "Attacking!"
	}
	panic(fmt.Sprintf("NO DESCRIPTION FOR INTENT %v", u.intent.itype))
}
