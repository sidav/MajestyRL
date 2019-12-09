package main

func getFactionRGB(fn int) (uint8, uint8, uint8) {
	switch fn {
	case 0:
		return 0, 0, 255
	case 1:
		return 255, 0, 0
	case 2:
		return 0, 255, 0
	case 3:
		return 255, 255, 0
	}
	return 32, 32, 32
}

type faction struct {
	cursor                                 *cursor // cursor position
	economy                                *factionEconomy
	factionNumber                          int
	name                                   string
	playerControlled, aiControlled         bool // used as a stub for now
	// aiData                                 *aiData // for AI-controlled factions
	seenTiles, tilesInSight [][] bool
}

func createFaction(name string, n int, playerControlled, aiControlled bool) *faction { // temporary
	fctn := &faction{
		playerControlled: playerControlled, aiControlled: aiControlled, name: name, factionNumber: n,
		economy: &factionEconomy{currentGold: 5000}, cursor: &cursor{},
	}
	if aiControlled {
		// fctn.aiData = ai_createAiData()
	}
	fctn.seenTiles = make([][]bool, mapW)
	for i := range fctn.seenTiles {
		fctn.seenTiles[i] = make([]bool, mapH)
	}
	fctn.tilesInSight = make([][]bool, mapW)
	for i := range fctn.tilesInSight {
		fctn.tilesInSight[i] = make([]bool, mapH)
	}
	return fctn
}

func (f *faction) wereCoordsSeen(x, y int) bool {
	return f.seenTiles[x][y]
}

func (f *faction) areCoordsInSight(x, y int) bool {
	return f.tilesInSight[x][y]
}
