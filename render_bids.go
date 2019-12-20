package main 

func (r *rendererStruct) renderBidsInViewport(g *gameMap) {
	f := r.currentFactionSeeingTheScreen
	vx, vy := f.cursor.getCameraCoords()
	for _, b := range g.bids {
		r.renderBid(b, vx, vy)
	}
}

func (r *rendererStruct) renderBid(b *bid, vx, vy int) {
	switch b.intent_type_for_this_bid {
	case INTENT_BUILD:
		r.renderBuilding(b.factionCreatedBid, b.targetPawn, vx, vy, false, true)
	}
}
