package main

import (
	cw "github.com/sidav/golibrl/console"
	cmenu "github.com/sidav/golibrl/console_menu"
)

func (pc *playerController) selectBuidingToConstruct() string {
	allowedBuildingCodes := make([]string, 0)
	for {
		names := make([]string, 0)
		descriptions := make([]string, 0)
		for code, allowed := range pc.curFaction.allowedBuildings {
			if allowed == TECH_ALLOWED {
				allowedBuildingCodes = append(allowedBuildingCodes, code)
				// name, desc := getBuildingNameAndDescription(code)
				names = append(names, getBuildingStaticDataFromTable(code).name)
				descriptions = append(descriptions, "desc")
			}
		}

		index := cmenu.ShowSidebarSingleChoiceMenu("BUILD:", pc.curFaction.getFactionColor(),
			SIDEBAR_X, SIDEBAR_FLOOR_2, SIDEBAR_W, SIDEBAR_H-SIDEBAR_FLOOR_2, names, descriptions)
		if index != -1 {
			code := allowedBuildingCodes[index]
			if getBuildingStaticDataFromTable(code).cost <= pc.curFaction.economy.currentGold {
				return code
			} else {
				pc.curFaction.reportToPlayer("we do not have enough gold!")
				RENDERER.renderScreen(pc.curFaction)
				continue
			}
		}
		return ""
	}
}

func (pc *playerController) selectBuildingSiteWithMouse(b *pawn) {
	pc.curFaction.reportToPlayer("Select construction site for " + b.getName())
	pc.rerenderNeeded = true
	for {
		pc.moveCursorWithMouse()
		pc.moveCameraIfNeeded()
		f := pc.curFaction
		cursor := f.cursor
		// cx, cy := cursor.getCoords()
		click := cw.GetMouseClickedButton()
		cursor.currentCursorMode = CURSOR_BUILD

		cursor.buildingToConstruct = b

		cursor.w, cursor.h = b.getSize()

		if !b.asBuilding.getStaticData().allowsTightPlacement {
			cursor.w += 2
			cursor.h += 2
		}

		// cursor.radius = b.getMaxRadiusToFire()

		keyPressed := cw.ReadKeyAsync()

		pc.moveCursorWithMouse()

		if pc.rerenderNeeded { // TODO: move all that "if reRenderNeeded" to the renderer itself to keep code more clean.
			RENDERER.renderScreen(f)
		}

		if pc.moveCameraIfNeeded() {
			continue
		}

		bw, bh := b.getSize()

		if click == "LEFT" {
			if CURRENT_MAP.canBuildingBeBuiltAt(b, cursor.x, cursor.y) {
				b.x = cursor.x - bw/2
				b.y = cursor.y - bh/2
				newbid := &bid{intent_type_for_this_bid: INTENT_BUILD, maxTaken: 2, x: b.x, y: b.y, targetPawn: b}
				CURRENT_MAP.addBid(newbid)
				pc.curFaction.economy.currentGold -= b.asBuilding.getStaticData().cost
				pc.rerenderNeeded = true
				return
			} else {
				f.reportToPlayer("This building can't be placed here!")
			}
		}
		if click == "RIGHT" {
			pc.rerenderNeeded = true
			f.reportToPlayer("Construction cancelled: " + b.getName())
			return
		}

		switch keyPressed {
		case "ESCAPE":
			pc.rerenderNeeded = true
			f.reportToPlayer("construction cancelled: " + b.getName())
			return
		default:
			pc.rerenderNeeded = false
		}
	}
}

func (pc *playerController) createMineBidIfNeeded() {
	cx, cy := pc.curFaction.cursor.getCoords()
	if CURRENT_MAP.getResourcesAtCoords(cx, cy) != nil {
		for i := 0; i < len(CURRENT_MAP.bids); i++ {
			cbid := CURRENT_MAP.bids[i]
			if cbid.factionCreatedBid == pc.curFaction && cbid.x == cx && cbid.y == cy {
				CURRENT_MAP.removeBid(cbid)
				return 
			}
		}
		newbid := &bid{intent_type_for_this_bid: INTENT_MINE, maxTaken: 3, x: cx, y: cy, factionCreatedBid: pc.curFaction}
		CURRENT_MAP.addBid(newbid)
	}
}

// func (pc *playerController) giveOrderWithMouse(selection *pawn, f *faction) {
// 	selectedPawn := (*selection)[0] //m.getUnitAtCoordinates(cx, cy)
// 	log.appendMessage(selectedPawn.name + " is awaiting orders.")
// 	f.cursor.currentCursorMode = CURSOR_MOVE
// 	reRenderNeeded = true
// 	for {
// 		equivKey := "NONE" // mouse clicked menu result
// 		click := cw.GetMouseClickedButton()
// 		cx, cy := f.cursor.getCoords()
// 		if reRenderNeeded {
// 			r_renderScreenForFaction(f, CURRENT_MAP, selection, false)
// 		}
// 		equivKey = pcm_mouseOrderSelectMenu(selectedPawn)
// 		if reRenderNeeded {
// 			cw.Flush_console()
// 		}

// 		keyPressed := cw.ReadKeyAsync()
// 		if plr_moveCameraOrCursorWithMouseIfNeeded(f) {
// 			continue
// 		}
// 		if click == "LEFT" {
// 			if equivKey == "NONE" && areGlobalCoordsOnScreen(cx, cy) {
// 				reRenderNeeded = true
// 				issueDefaultOrderToUnit(selectedPawn, CURRENT_MAP, cx, cy)
// 				return
// 			} else {
// 				keyPressed = equivKey
// 			}
// 		}
// 		if click == "RIGHT" {
// 			reRenderNeeded = true
// 			return
// 		}

// 		switch keyPressed {
// 		case "a": // attack-move
// 			if selectedPawn.hasWeapons() || selectedPawn.canConstructUnits() {
// 				f.cursor.currentCursorMode = CURSOR_AMOVE
// 				reRenderNeeded = true
// 			}
// 		case "m": // move
// 			f.cursor.currentCursorMode = CURSOR_MOVE
// 			reRenderNeeded = true
// 		case "b": // build
// 			if selectedPawn.canConstructBuildings() {
// 				code := plr_selectBuidingToConstruct(selectedPawn)
// 				if code != "" {
// 					plr_selectBuildingSiteWithMouse(selectedPawn, createBuilding(code, cx, cy, f), CURRENT_MAP)
// 					return
// 				}
// 			}
// 		case "c": // construct units
// 			if selectedPawn.canConstructUnits() {
// 				plr_selectUnitsToConstruct(selectedPawn)
// 				reRenderNeeded = true
// 			}
// 		case "r": // repeat construction queue
// 			if selectedPawn.canConstructUnits() {
// 				selectedPawn.repeatConstructionQueue = !selectedPawn.repeatConstructionQueue
// 				reRenderNeeded = true
// 			}
// 		case "u": // unload units inside
// 			if selectedPawn.canReleaseContainedPawns() {
// 				selectedPawn.setOrder(&order{orderType: order_unload})
// 				return
// 			}
// 		case "ESCAPE":
// 			return
// 		default:
// 			reRenderNeeded = false
// 		}
// 	}
// }
