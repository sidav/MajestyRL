package main 

// low level unit actions. 

func (u *pawn) spendTime(time int) {
	u.nextTickToAct = CURRENT_TICK + time 
}

func (u *pawn) doMoveToIntentTarget(desiredAccuracy int) bool { // Returns true if route exists. TODO: rewrite
	intent := u.asUnit.intent

	ox, oy := intent.x, intent.y
	ux, uy := u.getCoords()
	var vx, vy int

	path := CURRENT_MAP.getPathFromTo(ux, uy, ox, oy, desiredAccuracy)
	if path != nil {
		vx, vy = path.GetNextStepVector()
	}

	if true { // TODO: if canMove() 
		if vx == 0 && vy == 0 {
			reportToPlayer("no path to target!", u.faction)
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
