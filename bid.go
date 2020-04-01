package main

// bid is something on the map from which a unit creates its intents.
type bid struct {
	intent_type_for_this_bid intentType
	x, y                     int
	targetPawn               *pawn
	factionCreatedBid        *faction
	_markedAsFulfilled              bool
	currTaken, maxTaken      byte 
}

func (b *bid) isVacant() bool {
	return b.maxTaken == 0 || b.currTaken < b.maxTaken
}

func (b *bid) dispatchIntent() *intent {
	b.take()
	return b._createIntentForThisBid()
}

func (b *bid) take() {
	b.currTaken++
}

func (b *bid) drop() {
	b.currTaken--
}

func (b *bid) markFulfilled() { 
	b._markedAsFulfilled = true 
}

func (b *bid) isFulfilled() bool {
	if b.intent_type_for_this_bid == INTENT_MINE {
		return CURRENT_MAP.getResourcesAtCoords(b.x, b.y).amount <= 0 
	}
	return b._markedAsFulfilled
}

func (b *bid) _createIntentForThisBid() *intent {
	x, y := b.x, b.y
	if x == -1 || y == -1 {
		x, y = b.targetPawn.getCenter()
	}
	i := intent{
		itype: b.intent_type_for_this_bid,
		x:     x, y: y,
		targetPawn: b.targetPawn,
		sourceBid:  b,
	}
	return &i
}
