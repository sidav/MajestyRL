package main 

func createBuildingAtCoords(code string, constructed bool, centerX, centerY int, f *faction) *pawn {
	newPawn := &pawn{faction: f}
	newPawn.asBuilding = &building{code: code}
	staticData := getBuildingStaticDataFromTable(code)
	// assign residents array
	newPawn.asBuilding.currentResidents = make(map[string]int)
	for _, code := range staticData.housing_unittypes {
		newPawn.asBuilding.currentResidents[code] = 0
	}
	w, h := newPawn.getSize()
	newPawn.x, newPawn.y = centerX-w/2, centerY-h/2
	if constructed {
		newPawn.hitpoints = staticData.maxHitpoints
	} else {
		newPawn.asBuilding.asBeingConstructed = staticData.underConstructionData.clone()
	}
	return newPawn
}

func createUnitAtCoords(code string, centerX, centerY int, f *faction) *pawn {
	newPawn := &pawn{faction: f}
	newPawn.asUnit = &unit{code: code}
	staticData := getUnitStaticDataFromTable(code)
	newPawn.asUnit.maxHitpoints = rnd.RandInRange(staticData.maxHitpointsMin, staticData.maxHitpointsMax)
	newPawn.hitpoints = newPawn.asUnit.maxHitpoints
	if staticData.defaultWeaponCode != "" {
		newPawn.weapon = createWeaponByCode(staticData.defaultWeaponCode)
	}
	w, h := newPawn.getSize()
	newPawn.x, newPawn.y = centerX-w/2, centerY-h/2
	return newPawn
}
