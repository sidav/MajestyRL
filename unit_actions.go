package main 

func (u *pawn) doMoveToIntentTarget(desiredAccuracy int) bool { // Returns true if route exists. TODO: rewrite
	order := u.asUnit.intent

	ox, oy := order.x, order.y
	ux, uy := u.getCoords()
	var vx, vy int

	path := CURRENT_MAP.getPathFromTo(ux, uy, ox, oy)
	if path != nil {
		vx, vy = path.GetNextStepVector()
	}

	if true {

		u.x += vx
		u.y += vy
		// TODO: delays 
	}
	return true
}
