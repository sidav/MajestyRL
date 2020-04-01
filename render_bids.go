package main 

import cw "github.com/sidav/golibrl/console"


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
	case INTENT_MINE:
		r.renderMineBid(b, vx, vy)
	}
}

func (r *rendererStruct) renderMineBid(b *bid, vx, vy int) {
	x, y := b.x, b.y 
	tileApp := CURRENT_MAP.tileMap[x][y].getAppearance()
	cw.SetFgColor(cw.BLACK)
	cw.SetBgColor(tileApp.color)
	cw.PutChar(tileApp.char, x-vx, y-vy)
	cw.SetBgColor(cw.BLACK)
}
