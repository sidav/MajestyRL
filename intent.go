package main

const (
	INTENT_BUILD byte = iota
	INTENT_RETURN_HOME
)

// Represents anything that the unit is going to do
type intent struct {
	itype            byte
	targetx, targety int
	targetPawn       *pawn
}
