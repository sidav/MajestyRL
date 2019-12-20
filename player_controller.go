package main

import (
	cw "github.com/sidav/golibrl/console"
	"time"
)

type playerController struct {
	exit           bool
	last_time      time.Time
	rerenderNeeded bool
	endTurnPeriod  int
}

func (pc *playerController) init() {
	pc.exit = false
	pc.last_time = time.Now()
	pc.rerenderNeeded = true
	pc.endTurnPeriod = 50
}

func (pc *playerController) controlAsFaction(f *faction) {
	pc.last_time = time.Now()
	pc.rerenderNeeded = true 
	for !pc.isTimeToAutoEndTurn() {
		pc.snapCursorToPawn(f)
		if pc.rerenderNeeded {
			RENDERER.renderScreen(f)
			pc.rerenderNeeded = false
		}
		pc.selectPawnWithMouse(f)
		keyPressed := cw.ReadKeyAsync()
		switch keyPressed {
		case "ESCAPE":
			pc.exit = true

		// testing
		case "ENTER":
			for i := 0; i < 10; i++ {
				CURRENT_MAP.addBuilding(createBuildingAtCoords("HUT", false, rnd.Rand(mapW), rnd.Rand(mapH), f), true)
				reportToPlayer("cheats done", f)
			}
		case " ":
			for i := 0; i < 10; i++ {
				x, y := rnd.Rand(mapW), rnd.Rand(mapH)
				newbid := &bid{intent_type_for_this_bid: INTENT_BUILD, maxTaken: 2, x: x, y: y, targetPawn: createBuildingAtCoords("GOLDVAULT", false, x, y, f)}
				CURRENT_MAP.addBid(newbid)
				reportToPlayer("cheats done", f)
			}
		default:
			pc.moveCursorWithMouse(f)
		}
	}
}


// func (pc *playerController) 

func (pc *playerController) selectPawnWithMouse(f *faction) *[]*pawn { // returns a pointer to an array of selected pawns.
	f.cursor.currentCursorMode = CURSOR_SELECT
	for {
		// keyPressed := cw.ReadKeyAsync()
		// click := cw.GetMouseClickedButton()

		if pc.moveCameraIfNeeded(f) {
			return nil
		}
		return nil 
		// u := f.cursor.snappedPawn
		// if cw.GetMouseHeldButton() == "LEFT" {
		// 	return plr_bandboxSelectionWithMouse(f)
		// }
		// if u != nil && click == "LEFT" {
		// 	if u.faction.factionNumber != f.factionNumber {
		// 		LOG.appendMessage("Enemy units can't be selected, Commander.")
		// 		return nil
		// 	}
		// 	return &[]*pawn{f.cursor.snappedPawn}
		// }

		// switch keyPressed {
		// case "NOTHING", "NON-KEY":
		// 	if !IS_PAUSED && isTimeToAutoEndTurn() {
		// 		last_time = time.Now()
		// 		PLR_LOOP = false // end turn
		// 		return nil
		// 	} else {
		// 		reRenderNeeded = false
		// 	}
		// case ".": // end turn without unpausing the game
		// 	if IS_PAUSED {
		// 		PLR_LOOP = false
		// 		return nil
		// 	}
		// case "`":
		// 	mouseEnabled = !mouseEnabled
		// 	if mouseEnabled {
		// 		LOG.appendMessage("Mouse controls enabled.")
		// 	} else {
		// 		LOG.appendMessage("Mouse controls disabled.")
		// 	}
		// case "SPACE", " ":
		// 	IS_PAUSED = !IS_PAUSED
		// 	if IS_PAUSED {
		// 		LOG.appendMessage("Tactical pause engaged.")
		// 	} else {
		// 		LOG.appendMessage("Switched to real-time mode.")
		// 	}
		// case "=":
		// 	if endTurnPeriod > 100 {
		// 		endTurnPeriod -= 100
		// 		LOG.appendMessagef("Game speed increased to %d", 10-(endTurnPeriod/100))
		// 	} else {
		// 		LOG.appendMessage("Can't increase game speed any further.")
		// 	}
		// case "-":
		// 	if endTurnPeriod < 2000 {
		// 		endTurnPeriod += 100
		// 		LOG.appendMessagef("Game speed decreased to %d", 10-(endTurnPeriod/100))
		// 	} else {
		// 		LOG.appendMessage("Can't decrease game speed any further.")
		// 	}

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
		// case "ESCAPE":
		// 	if cmenu.ShowSimpleYNChoiceModalWindow("Are you sure you want to quit?") {
		// 		GAME_IS_RUNNING = false
		// 		PLR_LOOP = false
		// 		return nil
		// 	}

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

		// default:
		// 	plr_moveCursor(f, keyPressed)
		// }
	}
}
