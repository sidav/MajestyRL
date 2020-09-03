package main

import (
	"time"

	cw "github.com/sidav/golibrl/console"
)

type playerController struct {
	exit, playerInControl bool

	last_time               time.Time
	last_time_idle_rendered time.Time

	rerenderNeeded bool
	idleRerenderEachMs int
	endTurnPeriod  int
	curFaction     *faction
}

func (pc *playerController) init() {
	pc.exit = false
	pc.last_time = time.Now()
	pc.rerenderNeeded = true
	pc.endTurnPeriod = 50
	pc.idleRerenderEachMs = 50
}

func (pc *playerController) controlAsFaction(f *faction) {
	pc.curFaction = f
	pc.last_time = time.Now()
	pc.last_time_idle_rendered = time.Now()
	RENDERER.renderScreen(pc.curFaction)
	for !pc.isTimeToAutoEndTurn() || (IS_PAUSED && pc.playerInControl) {
		pc.playerInControl = true
		pc.snapCursorToPawnOrBid()
		pc.mainControlLoop()
		if pc.rerenderNeeded || pc.isTimeToIdleRender() {
			RENDERER.renderScreen(pc.curFaction)
			pc.rerenderNeeded = false
		}
		if pc.exit {
			return
		}
	}
}

func (pc *playerController) doUnconditionalKeyActions(keyPressed string) {
	switch keyPressed {
	case "NOTHING", "NON-KEY":
		if !IS_PAUSED && pc.isTimeToAutoEndTurn() {
			// pc.last_time = time.Now()
			// pc.PLR_LOOP = false // end turn
			return
		} else {
			pc.rerenderNeeded = false
			pc.moveCursorWithMouse()
		}
	case ".": // end turn without unpausing the game
		if IS_PAUSED {
			pc.playerInControl = false
			return
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
			pc.curFaction.reportToPlayer("tactical pause engaged.")
		} else {
			pc.curFaction.reportToPlayer("switched to real-time mode.")
		}
	case "=":
		if pc.endTurnPeriod > 20 {
			pc.endTurnPeriod -= 10
			log.AppendMessagef("Game speed increased to %d", 10-(pc.endTurnPeriod/10))
		} else {
			log.AppendMessage("Can't increase game speed any further.")
		}
	case "-":
		if pc.endTurnPeriod < 2000 {
			pc.endTurnPeriod += 25
			log.AppendMessagef("Game speed decreased to %d", 10-(pc.endTurnPeriod/25))
		} else {
			log.AppendMessage("Can't decrease game speed any further.")
		}

		// case "ENTER", "RETURN":
		// 	u := f.cursor.snappedPawn //m.getUnitAtCoordinates(cx, cy)
		// 	if u == nil {
		// 		return plr_bandboxSelectionWithMouse() // select multiple units
		// 	}
		// 	if u.faction.factionNumber != f.factionNumber {
		// 		LOG.appendMessage("Enemy units can't be selected, Commander.")
		// 		return nil
		// 	}
		// 	return &[]*pawn{f.cursor.snappedPawn}
		// case "TAB":
		// 	trySelectNextIdlePawn()
		// case "C":
		// 	trySnapCursorToCommander()
		// 	return &[]*pawn{f.cursor.snappedPawn}
		// case "?":
		// 	if f.cursor.snappedPawn != nil {
		// 		renderPawnInfo(f.cursor.snappedPawn)
		// 	}
	case "ESCAPE":
		pc.exit = true
		return

	case "DELETE": // test
		for i := 0; i < 3; i++ {
			x, y := rnd.Rand(mapW), rnd.Rand(mapH)
			newbid := &bid{intent_type_for_this_bid: INTENT_BUILD, maxTaken: 2, x: x, y: y, targetPawn: createBuildingAtCoords("GOLDVAULT", false, x, y, pc.curFaction)}
			CURRENT_MAP.addBid(newbid)
			pc.curFaction.reportToPlayer("cheats done")
		}
	case "INSERT": // test
		pc.endTurnPeriod = 1

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
		pc.moveCursorWithMouse()
	}
}

func (pc *playerController) mainControlLoop() *pawn { // returns a pointer to selected pawn.
	pc.curFaction.cursor.currentCursorMode = CURSOR_SELECT
	keyPressed := cw.ReadKeyAsync()
	click := cw.GetMouseClickedButton()
	pc.moveCursorWithMouse()

	if pc.moveCameraIfNeeded() {
		return nil
	}

	u := pc.curFaction.cursor.snappedPawn
	if click == "LEFT" {
		if u != nil && u.faction != nil {
			if u.faction.factionNumber != pc.curFaction.factionNumber {
				pc.curFaction.reportToPlayer("those fools won't obey you!")
				return nil
			}
			return u
		} else { // maybe the resource was clicked on? Create bid if so. 
			pc.createMineBidIfNeeded()
		}
	}
	pc.doUnconditionalKeyActions(keyPressed)
	if keyPressed == "b" {
		pc.constructBuilding()
	}
	return nil
}

func (pc *playerController) constructBuilding() {
	bld_code := pc.selectBuidingToConstruct()
	if bld_code != "" {
		bld := createBuildingAtCoords(bld_code, false, 0, 0, pc.curFaction)
		pc.selectBuildingSiteWithMouse(bld)
	}
}
