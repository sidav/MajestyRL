package main

import (
	"time"

	cw "github.com/sidav/golibrl/console"
)

const (
	PC_CAMERA_MOVE_MARGIN = 2
	PC_CAMERA_MOVE_DELAY  = 20 // ms
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

	cursorWasMoved := !(pc.curFaction.cursor.x == camx+cx && pc.curFaction.cursor.y == camy+cy)

	if cursorWasMoved && CURRENT_MAP.areCoordsValid(camx+cx, camy+cy) {
		pc.rerenderNeeded = true // rerender is needed if cursor was _actually_ moved
		pc.curFaction.cursor.x, pc.curFaction.cursor.y = camx+cx, camy+cy
		pc.snapCursorToPawn()
	}
}

func (pc *playerController) moveCameraIfNeeded() bool { // true if camera was moved
	const scrollSpeed = 2
	cx, cy := cw.GetMouseCoords()
	crs := pc.curFaction.cursor
	moved := false
	if cx-PC_CAMERA_MOVE_MARGIN < 0 && crs.cameraX > -VIEWPORT_W/2 {
		crs.cameraX -= scrollSpeed
		moved = true
	}
	if cy-PC_CAMERA_MOVE_MARGIN < 0 && crs.cameraY > -VIEWPORT_H/2 {
		crs.cameraY -= scrollSpeed
		moved = true
	}
	if cx+PC_CAMERA_MOVE_MARGIN >= CONSOLE_W && crs.cameraX < mapW-VIEWPORT_W/2 {
		crs.cameraX += scrollSpeed
		moved = true
	}
	if cy+PC_CAMERA_MOVE_MARGIN >= CONSOLE_H && crs.cameraY < mapH-VIEWPORT_H/2 {
		crs.cameraY += scrollSpeed
		moved = true
	}
	if moved {
		time.Sleep(PC_CAMERA_MOVE_DELAY * time.Millisecond)
	}
	pc.rerenderNeeded = moved
	return moved
}

func (pc *playerController) isTimeToAutoEndTurn() bool {
	return time.Since(pc.last_time) >= time.Duration(pc.endTurnPeriod)*time.Millisecond
}

// this is a horrible workaround for my TCell wrapper...
// TODO: get rid of this by cleaning up useless mouse move events in wrapper.
func (pc *playerController) isTimeToIdleRender() bool {
	isTime := time.Since(pc.last_time_idle_rendered) >= time.Duration(pc.idleRerenderEachMs)*time.Millisecond
	if isTime {
		pc.last_time_idle_rendered = time.Now()
	}
	return isTime
}
