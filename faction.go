package main

const (
	TECH_NOTSET uint8 = iota 
	TECH_DENIED
	TECH_ALLOWED 
)

type faction struct {
	cursor                                 *cursor // cursor position
	economy                                *factionEconomy
	factionNumber                          int
	name                                   string
	playerControlled, aiControlled         bool // used as a stub for now
	// aiData                                 *aiData // for AI-controlled factions
	seenTiles, tilesInSight [][] bool

	//tech 
	allowedBuildings map[string]uint8 
}

func createFaction(name string, n int, playerControlled, aiControlled bool) *faction { // temporary
	fctn := &faction{
		playerControlled: playerControlled, aiControlled: aiControlled, name: name, factionNumber: n,
		economy: &factionEconomy{currentGold: 5000, maxGold: 5000}, cursor: &cursor{},
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
	fctn.allowedBuildings = make(map[string]uint8, len(staticBuildingDataTable))
	return fctn
}

func (f *faction) wereCoordsSeen(x, y int) bool {
	return true 
	return f.seenTiles[x][y]
}

func (f *faction) areCoordsInSight(x, y int) bool {
	return true 
	return f.tilesInSight[x][y]
}

func (f *faction) reportToPlayer(text string) {
	plrname := "Unknown One"
	if f != nil {
		plrname = f.name
	}
	log.AppendMessage(plrname + ", " + text)
}

func (f *faction) getFactionColor() int {
	//BLACK        = 0
	//DARK_RED     = 1
	//DARK_GREEN   = 2
	//DARK_YELLOW  = 3
	//DARK_BLUE    = 4
	//DARK_MAGENTA = 5
	//DARK_CYAN    = 6
	//BEIGE        = 7
	//DARK_GRAY    = 8
	//RED          = 9
	//GREEN        = 10
	//YELLOW       = 11
	//BLUE         = 12
	//MAGENTA      = 13
	//CYAN         = 14
	//WHITE        = 15
	if f == nil {
		return 1
	}
	switch f.factionNumber {
	case 0:
		return 14
	case 1:
		return 9
	case 2:
		return 10
	case 3:
		return 11
	}
	return 7
}
