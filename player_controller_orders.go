package main 

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
