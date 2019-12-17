package main 

import (
	cw "github.com/sidav/golibrl/console"
	// geometry "github.com/sidav/golibrl/geometry"
)

func (r *rendererStruct) renderCursor() {
	c := r.currentFactionSeeingTheScreen.cursor
	cx, cy := c.getCoords()
	if !r.areGlobalCoordsOnScreen(cx, cy) {
		return
	}
	switch c.currentCursorMode {
	case CURSOR_SELECT:
		r.renderSelectCursor()
	}
}

func (r *rendererStruct) renderSelectCursor() {
	c := r.currentFactionSeeingTheScreen.cursor
	x, y := c.getOnScreenCoords()
	snap := c.snappedPawn
	// cw.SetFgColorRGB(128, 128, 128)
	if snap == nil {
		cw.SetFgColor(cw.WHITE)
	} else if snap.faction == r.currentFactionSeeingTheScreen {
		cw.SetFgColor(cw.GREEN)
	} else {
		cw.SetFgColor(cw.RED)
	}

	if snap == nil || snap.isUnit() {
		cw.PutChar('[', x-1, y)
		cw.PutChar(']', x+1, y)
	} else {
		w, h := snap.getSize()
		offset := w % 2
		for cy := 0; cy < h; cy++ {
			cw.PutChar('[', x-w/2-1, cy-h/2+y)
			cw.PutChar(']', x+w/2+offset, cy-h/2+y)
		}
	}
	// globx, globy := c.getCoords()
	resInfoString := ""
	// mineralsUnderCursor := CURRENT_MAP.getMineralsAtCoordinates(globx, globy)
	// vespeneUnderCursor := CURRENT_MAP.getVespeneAtCoordinates(globx, globy)
	// if mineralsUnderCursor > 0 {
	// 	resInfoString = fmt.Sprintf(" %dx minerals ", mineralsUnderCursor)
	// }
	// if vespeneUnderCursor > 0 {
	// 	resInfoString = fmt.Sprintf(" %dx vespene ", vespeneUnderCursor)
	// }
	if len(resInfoString) > 0 {
		cw.SetBgColor(cw.DARK_GRAY)
		cw.SetFgColor(cw.WHITE)
		cw.PutString(resInfoString, x+2, y-1)
	}
	cw.SetBgColor(cw.BLACK)

	// outcommented for non-SDL console
	//cw.PutChar(16*13+10, x-1, y-1)
	//cw.PutChar(16*11+15, x+1, y-1)
	//cw.PutChar(16*12, x-1, y+1)
	//cw.PutChar(16*13+9, x+1, y+1)
	// flushView()
}
