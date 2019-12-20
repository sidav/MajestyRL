package main

import (
	cw "github.com/sidav/golibrl/console"
	"time"
)

type playerController struct {
	exit, playerInControl bool
	last_time             time.Time
	rerenderNeeded        bool
	endTurnPeriod         int
}

func (pc *playerController) init() {
	pc.exit = false
	pc.last_time = time.Now()
	pc.rerenderNeeded = true
	pc.endTurnPeriod = 50
}

func (pc *playerController) controlAsFaction(f *faction) {
	pc.last_time = time.Now()
	for !pc.isTimeToAutoEndTurn() || IS_PAUSED {
		pc.rerenderNeeded = true
		pc.snapCursorToPawn(f)
		if pc.rerenderNeeded {
			RENDERER.renderScreen(f)
			pc.rerenderNeeded = false
		}
		pc.mainControlLoop(f)
		// keyPressed := cw.ReadKeyAsync()
		// switch keyPressed {
		// // testing
		// case "ENTER":
		// 	for i := 0; i < 10; i++ {
		// 		CURRENT_MAP.addBuilding(createBuildingAtCoords("HUT", false, rnd.Rand(mapW), rnd.Rand(mapH), f), true)
		// 		reportToPlayer("cheats done", f)
		// 	}
		// case " ":
		// 	for i := 0; i < 10; i++ {
		// 		x, y := rnd.Rand(mapW), rnd.Rand(mapH)
		// 		newbid := &bid{intent_type_for_this_bid: INTENT_BUILD, maxTaken: 2, x: x, y: y, targetPawn: createBuildingAtCoords("GOLDVAULT", false, x, y, f)}
		// 		CURRENT_MAP.addBid(newbid)
		// 		reportToPlayer("cheats done", f)
		// 	}
		// default:
		// 	pc.moveCursorWithMouse(f)
		// }
		if pc.exit {
			return 
		}
	}
}

// func (pc *playerController)

func (pc *playerController) mainControlLoop(f *faction) *[]*pawn { // returns a pointer to an array of selected pawns.
	f.cursor.currentCursorMode = CURSOR_SELECT
	keyPressed := cw.ReadKeyAsync()
	click := cw.GetMouseClickedButton()
	pc.moveCursorWithMouse(f)

	if pc.moveCameraIfNeeded(f) {
		return nil
	}

	u := f.cursor.snappedPawn
	if u != nil && click == "LEFT" {
		if u.faction.factionNumber != f.factionNumber {
			reportToPlayer("those fools won't obey you!", f)
			return nil
		}
		return nil
	}
	switch keyPressed {
	case "NOTHING", "NON-KEY":
		if !IS_PAUSED && pc.isTimeToAutoEndTurn() {
			// pc.last_time = time.Now()
			// pc.PLR_LOOP = false // end turn
			return nil
		} else {
			pc.rerenderNeeded = false
		}
	case ".": // end turn without unpausing the game
		if IS_PAUSED {
			pc.playerInControl = false
			return nil
		}
	// case "`":
	// 	mouseEnabled = !mouseEnabled
	// 	if mouseEnabled {
	// 		LOG.appendMessage("Mouse controls enabled.")
	// 	} else {
	// 		LOG.appendMessage("Mouse controls disabled.")
	// 	}
	case "SPACE", " ":
		IS_PAUSED = !IS_PAUSED
		if IS_PAUSED {
			reportToPlayer("tactical pause engaged.", f)
		} else {
			reportToPlayer("switched to real-time mode.", f)
		}
	case "=":
		if pc.endTurnPeriod > 25 {
			pc.endTurnPeriod -= 25
			LOG.AppendMessagef("Game speed increased to %d", 10-(pc.endTurnPeriod/25))
		} else {
			LOG.AppendMessage("Can't increase game speed any further.")
		}
	case "-":
		if pc.endTurnPeriod < 2000 {
			pc.endTurnPeriod += 25
			LOG.AppendMessagef("Game speed decreased to %d", 10-(pc.endTurnPeriod/25))
		} else {
			LOG.AppendMessage("Can't decrease game speed any further.")
		}

	// case "ENTER", "RETURN":
	// 	u := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
	// 	if u == nil {
	// 		return plr_bandboxSelectionWithMouse(f) // select multiple units
	// 	}
	// 	if u.faction.factionNumber != f.factionNumber {
	// 		LOG.appendMessage("Enemy units can't be selected, Commander.")
	// 		return nil
	// 	}
	// 	return &[]*pawn{f.cursor.snappedPawn}
	// case "TAB":
	// 	trySelectNextIdlePawn(f)
	// case "C":
	// 	trySnapCursorToCommander(f)
	// 	return &[]*pawn{f.cursor.snappedPawn}
	// case "?":
	// 	if f.cursor.snappedPawn != nil {
	// 		renderPawnInfo(f.cursor.snappedPawn)
	// 	}
	case "ESCAPE":
		pc.exit = true
		return nil

	// case "DELETE": // cheat
	// 	f.economy.minerals += 10000
	// 	f.economy.vespene += 10000
	// case "E": // test
	// 	CURRENT_MAP.addBuilding(createBuilding("testsmall", f.cursor.x, f.cursor.y, CURRENT_MAP.factions[1]), true)
	// 	LOG.appendMessage("Test enemy building created.")
	// case "B": // test
	// 	CURRENT_MAP.addBuilding(createBuilding("testbig", f.cursor.x, f.cursor.y, CURRENT_MAP.factions[1]), true)
	// 	LOG.appendMessage("LARGE test enemy building created.")
	// case "INSERT": // cheat
	// 	for _, fac := range CURRENT_MAP.factions {
	// 		if fac != f {
	// 			fac.economy.minerals += 500
	// 			fac.economy.vespene += 500
	// 		}
	// 	}
	// 	LOG.appendMessage("Added 500 minerals and gas to the enemies.")
	// case "HOME": // cheat
	// 	// CURRENT_MAP.addBuilding(createBuilding("lturret", f.cursor.x, f.cursor.y, CURRENT_MAP.factions[0]), true)
	// 	endTurnPeriod = 0
	// case "END": // cheat
	// 	CHEAT_IGNORE_FOW = !CHEAT_IGNORE_FOW

	default:
		pc.moveCursorWithMouse(f)
	}
	return nil
}
