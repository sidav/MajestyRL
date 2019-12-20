package main 

import (
	"time"
	cw "github.com/sidav/golibrl/console"
)

const (
	PC_CAMERA_MOVE_MARGIN = 2
	PC_CAMERA_MOVE_DELAY = 20 // ms
)

func (pc *playerController) snapCursorToPawn(f *faction) {
	if !(f.areCoordsInSight(f.cursor.x, f.cursor.y)) {
		return
	}
	b := CURRENT_MAP.getPawnAtCoordinates(f.cursor.x, f.cursor.y)
	if b == nil {
		f.cursor.snappedPawn = nil
	} else {
		f.cursor.x, f.cursor.y = b.getCenter()
		f.cursor.snappedPawn = b
	}
}

func (pc *playerController) moveCursorWithMouse(f *faction) {
	cx, cy := cw.GetMouseCoords()
	camx, camy := f.cursor.getCameraCoords()

	pc.rerenderNeeded = !(f.cursor.x == camx+cx && f.cursor.y == camy+cy) // rerender is needed if cursor was _actually_ moved

	if CURRENT_MAP.areCoordsValid(camx+cx, camy+cy) {
		f.cursor.x, f.cursor.y = camx+cx, camy+cy
		pc.snapCursorToPawn(f)
	}
}

func (pc *playerController) moveCameraIfNeeded(f *faction) bool { // true if camera was moved 
	cx, cy := cw.GetMouseCoords()
	crs := f.cursor
	moved := false 
	if cx - PC_CAMERA_MOVE_MARGIN < 0 {
		crs.cameraX -= 1 
		moved = true 
	}
	if cy - PC_CAMERA_MOVE_MARGIN < 0 {
		crs.cameraY -= 1 
		moved = true 
	}
	if cx + PC_CAMERA_MOVE_MARGIN >= CONSOLE_W {
		crs.cameraX += 1 
		moved = true 
	}
	if cy + PC_CAMERA_MOVE_MARGIN >= CONSOLE_H {
		crs.cameraY += 1 
		moved = true 
	}
	if moved {
		time.Sleep(PC_CAMERA_MOVE_DELAY*time.Millisecond)
	}
	return moved 
}

func (pc *playerController) isTimeToAutoEndTurn() bool {
	return time.Since(pc.last_time) >= time.Duration(pc.endTurnPeriod)*time.Millisecond
}
