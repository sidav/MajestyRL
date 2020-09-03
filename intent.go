package main

type intentType byte

const (
	INTENT_BUILD intentType = iota
	INTENT_BRING_RESOURCES_TO_CONSTRUCTION
	INTENT_GROW_FOREST
	INTENT_REPAIR
	INTENT_COLLECT_TAXES
	INTENT_RETURN_TAXES
	INTENT_RETURN_HOME
	INTENT_PATROL
	INTENT_MINE
	INTENT_ATTACK

	DROP_INTENT_AFTER byte = 3
)

// Represents anything that the unit is going to do
type intent struct {
	itype      intentType
	x, y       int
	targetPawn *pawn

	sourceBid *bid

	insuccessCount byte
}

func (i *intent) getCoords() (int, int) {
	return i.x, i.y 
}

func (i *intent) isDispatchedFromBid() bool {
	return i.sourceBid != nil 
}

func (i *intent) fulfillBidIfExists() {
	if i.sourceBid != nil {
		i.sourceBid.markFulfilled()
	}
}
