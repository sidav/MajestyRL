package main

type unit struct {
	intent       *intent
	registeredIn *pawn // building where the unit is registered, its "home"
	code         string
	maxHitpoints int

	carriedResourceAmount int
	carriedResourceType   resourceType
}

func (u *unit) getStaticData() *unitStaticData {
	return getUnitStaticDataFromTable(u.code)
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

func (u *pawn) dropCurrentIntent() {
	if u.asUnit.intent.sourceBid != nil {
		u.asUnit.intent.sourceBid.drop()
	}
	u.asUnit.intent = nil
	if u.asUnit.carriedResourceAmount > 0 {
		u.dropResources()
	}
}

