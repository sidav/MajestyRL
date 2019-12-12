package main

// bid is something on the map from which a unit creates its intents.
type bid struct {
	intent_type_for_this_bid intentType
	x, y                     int
	targetPawn               *pawn
	factionCreatedBid        *faction
}

func (b *bid) createIntentForThisBid() *intent {
	i := intent{
		itype: b.intent_type_for_this_bid,
		x:     b.x, y: b.y,
		targetPawn: b.targetPawn,
	}
	return &i
}
