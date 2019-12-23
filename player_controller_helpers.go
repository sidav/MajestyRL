package main 

import (
	"time"
	cw "github.com/sidav/golibrl/console"
	cmenu "github.com/sidav/golibrl/console_menu"
)

const (
	PC_CAMERA_MOVE_MARGIN = 2
	PC_CAMERA_MOVE_DELAY = 20 // ms
)

func (pc *playerController) snapCursorToPawn() {
	if !(pc.curFaction.areCoordsInSight(pc.curFaction.cursor.x, pc.curFaction.cursor.y)) {
		return
	}
	b := CURRENT_MAP.getPawnAtCoordinates(pc.curFaction.cursor.x, pc.curFaction.cursor.y)
	if b == nil {
		pc.curFaction.cursor.snappedPawn = nil
	} else {
		pc.curFaction.cursor.x, pc.curFaction.cursor.y = b.getCenter()
		pc.curFaction.cursor.snappedPawn = b
	}
}

func (pc *playerController) moveCursorWithMouse() {
	cx, cy := cw.GetMouseCoords()
	camx, camy := pc.curFaction.cursor.getCameraCoords()

	pc.rerenderNeeded = !(pc.curFaction.cursor.x == camx+cx && pc.curFaction.cursor.y == camy+cy) // rerender is needed if cursor was _actually_ moved

	if CURRENT_MAP.areCoordsValid(camx+cx, camy+cy) {
		pc.curFaction.cursor.x, pc.curFaction.cursor.y = camx+cx, camy+cy
		pc.snapCursorToPawn()
	}
}

func (pc *playerController) moveCameraIfNeeded() bool { // true if camera was moved 
	const scrollSpeed = 2 
	cx, cy := cw.GetMouseCoords()
	crs := pc.curFaction.cursor
	moved := false 
	if cx - PC_CAMERA_MOVE_MARGIN < 0  && crs.cameraX > -VIEWPORT_W / 2 {
		crs.cameraX -= scrollSpeed 
		moved = true 
	}
	if cy - PC_CAMERA_MOVE_MARGIN < 0 && crs.cameraY > -VIEWPORT_H / 2{
		crs.cameraY -= scrollSpeed 
		moved = true 
	}
	if cx + PC_CAMERA_MOVE_MARGIN >= CONSOLE_W && crs.cameraX < mapW - VIEWPORT_W / 2 {
		crs.cameraX += scrollSpeed 
		moved = true 
	}
	if cy + PC_CAMERA_MOVE_MARGIN >= CONSOLE_H && crs.cameraY < mapH - VIEWPORT_H / 2 {
		crs.cameraY += scrollSpeed 
		moved = true 
	}
	if moved {
		time.Sleep(PC_CAMERA_MOVE_DELAY*time.Millisecond)
	}
	pc.rerenderNeeded = moved 
	return moved 
}

func (pc *playerController) isTimeToAutoEndTurn() bool {
	return time.Since(pc.last_time) >= time.Duration(pc.endTurnPeriod)*time.Millisecond
}

func (pc *playerController) selectBuidingToConstruct() string {
	allowedBuildingCodes := make([]string, 0)

	names := make([]string, 0)
	descriptions := make([]string, 0)
	// for _, code := range *allAvailableBuildingCodes {
	// 	if p.faction.tech.areRequirementsSatisfiedForCode(code) {
	// 		allowedBuildingCodes = append(allowedBuildingCodes, code)
	// 		name, desc := getBuildingNameAndDescription(code)
	// 		names = append(names, name)
	// 		descriptions = append(descriptions, desc)
	// 	}
	// }

	index := cmenu.ShowSidebarSingleChoiceMenu("BUILD:", pc.curFaction.getFactionColor(),
		SIDEBAR_X, SIDEBAR_FLOOR_2, SIDEBAR_W, SIDEBAR_H-SIDEBAR_FLOOR_2, names, descriptions)
	if index != -1 {
		return allowedBuildingCodes[index]
	}
	return ""
}

func (pc *playerController) selectBuildingSiteWithMouse(b *pawn, m *gameMap) {
	pc.curFaction.reportToPlayer("Select construction site for " + b.getName())
	pc.rerenderNeeded = true
	for {
		f := pc.curFaction
		cursor := f.cursor
		// cx, cy := cursor.getCoords()
		click := cw.GetMouseClickedButton()
		cursor.currentCursorMode = CURSOR_BUILD

		cursor.buildingToConstruct = b

		cursor.w, cursor.h = b.getSize()
		cursor.w += 2 
		cursor.h += 2 

		// if b.buildingInfo.allowsTightPlacement {
		// 	cursor.w = b.buildingInfo.w
		// 	cursor.h = b.buildingInfo.h
		// } else {
		// 	cursor.w = b.buildingInfo.w + 2
		// 	cursor.h = b.buildingInfo.h + 2
		// }
		

		// cursor.radius = b.getMaxRadiusToFire()

		if pc.rerenderNeeded { // TODO: move all that "if reRenderNeeded" to the renderer itself to keep code more clean.
			RENDERER.renderScreen(f)
		}

		keyPressed := cw.ReadKeyAsync()

		pc.moveCursorWithMouse()

		if pc.moveCameraIfNeeded() {
			continue
		}

		if click == "LEFT" {
			// if m.canBuildingBeBuiltAt(b, cx, cy) {
			// 	b.x = cx - b.buildingInfo.w/2
			// 	b.y = cy - b.buildingInfo.h/2
			// 	p.setOrder(&order{orderType: order_build, x: cx, y: cy, buildingToConstruct: b})
			// 	reRenderNeeded = true
			// 	return
			// } else {
			// 	f.reportToPlayer("This building can't be placed here!")
			// }
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
