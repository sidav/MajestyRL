package main

import "fmt"

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

func (i *intent) getDescription() string {
	switch i.itype {
	case INTENT_BUILD:
		return "Building..."
	case INTENT_COLLECT_TAXES:
		return "Collecting money..."
	case INTENT_BRING_RESOURCES_TO_CONSTRUCTION:
		return "Hauling resources..."
	case INTENT_GROW_FOREST:
		return "Growing new forest..."
	case INTENT_RETURN_TAXES:
		return "Returning money..."
	case INTENT_REPAIR:
		return "Repairing..."
	case INTENT_MINE:
		return "Mining..."
	case INTENT_RETURN_HOME:
		return "Going to rest..."
	case INTENT_PATROL:
		return "Patrolling..."
	case INTENT_ATTACK:
		return "Attacking!"
	}
	panic(fmt.Sprintf("NO DESCRIPTION FOR INTENT %v", i.itype))
}
