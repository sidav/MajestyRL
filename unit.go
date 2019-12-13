package main

type unit struct {
	intent       *intent
	registeredIn *pawn // building where the unit is registered, its "home"
	code         string
}

func (u *unit) handleIntentUnsuccess() {
	u.intent.insuccessCount++
	if u.intent.insuccessCount > DROP_INTENT_AFTER {
		if u.intent.isDispatchedFromBid() {
			u.intent.sourceBid.drop()
		}
		u.intent = nil
	}
}
