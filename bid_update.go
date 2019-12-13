package main 

// func (g *gameMap) refreshBids() {

// 	for _, pwn := range CURRENT_MAP.pawns {
// 		// create build bids 
// 		if pwn.isBuilding() && pwn.asBuilding.isUnderConstruction() {
// 			g.addBid(&bid{intent_type_for_this_bid: INTENT_BUILD, targetPawn: pwn})
// 		}
// 	}
// }

func (g *gameMap) cleanupBids() {
	LOG.AppendMessagef("Cleaninbg bids, %d total", len(g.bids))
	for i := 0; i < len(g.bids); i++ {
		LOG.AppendMessagef("BID: %d total", g.bids[i].currTaken)
		if g.bids[i].isFulfilled {
			g.bids = append(g.bids[:i], g.bids[i+1:]...) // ow it's fucking... magic!
			LOG.AppendMessage("Bid cleaned")
			i-- 
		}
	}
	LOG.AppendMessagef("Bids cleaned, %d remaining", len(g.bids))
}
