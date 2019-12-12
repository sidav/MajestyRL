package main

type unit struct {
	intent       *intent
	registeredIn *pawn // building where the unit is registered, its "home"
	code         string
}
