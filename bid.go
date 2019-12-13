package main

// bid is something on the map from which a unit creates its intents.
type bid struct {
	intent_type_for_this_bid intentType
	x, y                     int
	targetPawn               *pawn
	factionCreatedBid        *faction
	isFulfilled              bool
}

func (b *bid) createIntentForThisBid() *intent {
	x, y := b.x, b.y
	if x == -1 || y == -1 {
		x, y = b.targetPawn.getCenter()
	}
	i := intent{
		itype: b.intent_type_for_this_bid,
		x: x, y: y, 
		targetPawn: b.targetPawn,
		sourceBid: b,
	}
	return &i
}
