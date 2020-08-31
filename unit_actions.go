package main

import "fmt"

// low level unit actions. 

func (u *pawn) spendTime(time int) {
	u.nextTickToAct = CURRENT_TICK + time 
}

func (u *pawn) doMoveToIntentTarget(desiredAccuracy int) bool { // Returns true if route exists. TODO: rewrite
	intent := u.asUnit.intent

	ox, oy := intent.x, intent.y
	if intent.targetPawn != nil {
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
			u.spendTime(10*TICKS_PER_TURN)
			u.asUnit.handleIntentUnsuccess()
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

func (u *unit) getCurrentIntentDescription() string {
	if u.intent == nil {
		return "Thinking..."
	}
	switch u.intent.itype {
	case INTENT_BUILD: 
	return "Building..."
	case INTENT_RETURN_HOME:
		return "Going to rest..."
	case INTENT_PATROL:
		return "Patrolling..."
	case INTENT_ATTACK:
		return "Attacking!"
	}
	return "NO DESCRIPTION FOR INTENT"
}
