package main 

func createBuildingAtCoords(code string, constructed bool, centerX, centerY int, f *faction) *pawn {
	newPawn := &pawn{faction: f}
	newPawn.asBuilding = &building{code: code}
	staticData := staticBuildingDataTable[code]
	w, h := newPawn.getSize()
	newPawn.x, newPawn.y = centerX-w/2, centerY-h/2
	if !constructed {
		newPawn.asBuilding.beingConstructed = staticData.underConstructionData.clone()
	}
	return newPawn
}

func createUnitAtCoords(code string, centerX, centerY int, f *faction) *pawn {
	newPawn := &pawn{faction: f}
	newPawn.asUnit = &unit{code: code}
	staticData := staticUnitDataTable[code]
	newPawn.hitpoints = staticData.maxHitpointsMin // TODO: use random here 
	w, h := newPawn.getSize()
	newPawn.x, newPawn.y = centerX-w/2, centerY-h/2
	return newPawn
}
