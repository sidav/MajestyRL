package main

// cause is something from which a unit creates its intents.
type cause struct {
	intent_type_for_this_cause intentType
	x, y                       int
	targetPawn                 *pawn
	factionCreatedcause        *faction
}
