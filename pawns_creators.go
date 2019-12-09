package main 

func createBuildingAtCoords(code string, centerX, centerY int, f *faction) *pawn {
	newPawn := &pawn{faction: f}
	newPawn.asBuilding = &building{code: code}
	// staticData := staticBuildingDataTable[code]
	w, h := newPawn.getSize()
	newPawn.x, newPawn.y = centerX-w/2, centerY-h/2
	return newPawn
}
