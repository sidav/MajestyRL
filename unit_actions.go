package main 

// low level unit actions. 

func (u *pawn) spendTime(time int) {
	u.nextTickToAct = CURRENT_TICK + time 
}

func (u *pawn) doMoveToIntentTarget(desiredAccuracy int) bool { // Returns true if route exists. TODO: rewrite
	order := u.asUnit.intent

	ox, oy := order.x, order.y
	ux, uy := u.getCoords()
	var vx, vy int

	path := CURRENT_MAP.getPathFromTo(ux, uy, ox, oy, desiredAccuracy)
	if path != nil {
		vx, vy = path.GetNextStepVector()
	}

	if true {

		u.x += vx
		u.y += vy
		u.spendTime(TICKS_PER_TURN)
		// TODO: delays 
	}
	return true
}
