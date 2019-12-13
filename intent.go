package main

type intentType byte

const (
	INTENT_BUILD intentType = iota
	INTENT_RETURN_HOME

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


func (i *intent) isDispatchedFromBid() bool {
	return i.sourceBid != nil 
}

func (i *intent) fulfillBidIfExists() {
	if i.sourceBid != nil {
		i.sourceBid.isFulfilled = true
	}
}
