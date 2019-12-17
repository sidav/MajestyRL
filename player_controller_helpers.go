package main 

import (
	"time"
	cw "github.com/sidav/golibrl/console"
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

func (pc *playerController) isTimeToAutoEndTurn() bool {
	return time.Since(pc.last_time) >= time.Duration(pc.endTurnPeriod)*time.Millisecond
}
