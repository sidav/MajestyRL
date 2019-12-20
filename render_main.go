package main

import (
	cw "github.com/sidav/golibrl/console"
	geometry "github.com/sidav/golibrl/geometry"
)

var (
	CONSOLE_W, CONSOLE_H = 80, 25
	VIEWPORT_W           = 40
	VIEWPORT_H           = CONSOLE_H - LOG_HEIGHT
	SIDEBAR_X            = VIEWPORT_W + 1
	SIDEBAR_W            = CONSOLE_W - VIEWPORT_W - 1
	SIDEBAR_H            = CONSOLE_H - LOG_HEIGHT
	SIDEBAR_FLOOR_2      = 7  // y-coord right below resources info
	SIDEBAR_FLOOR_3      = 11 // y-coord right below "floor 2"
)

type rendererStruct struct {
	currentFactionSeeingTheScreen *faction
}

func (r *rendererStruct) renderScreen(f *faction) {
	r.currentFactionSeeingTheScreen = f
	r.updateBoundsIfNeccessary(false)
	cw.Clear_console()
	r.renderMapInViewport(CURRENT_MAP)
	r.renderBidsInViewport(CURRENT_MAP)
	r.renderPawnsInViewport(CURRENT_MAP)
	r.renderUI()
	r.renderCursor()
	log.Render(CONSOLE_H - LOG_HEIGHT)

	if DEBUG_OUTPUT {
		PrintMemUsage()
	}

	cw.Flush_console()
}

func (r *rendererStruct) setFgColorByCcell(c *ccell) {
	cw.SetFgColor(c.color)
	// cw.SetFgColorRGB(c.r, c.g, c.b)
}

func (r *rendererStruct) updateBoundsIfNeccessary(force bool) {
	if cw.WasResized() || force {
		CONSOLE_W, CONSOLE_H = cw.GetConsoleSize()
		VIEWPORT_W = 2 * CONSOLE_W / 3
		VIEWPORT_H = CONSOLE_H - LOG_HEIGHT - 1
		SIDEBAR_X = VIEWPORT_W + 1
		SIDEBAR_W = CONSOLE_W - VIEWPORT_W - 1
		SIDEBAR_H = CONSOLE_H - LOG_HEIGHT
		SIDEBAR_FLOOR_2 = 7  // y-coord right below resources info
		SIDEBAR_FLOOR_3 = 11 // y-coord right below "floor 2"
	}
}

func (r *rendererStruct) renderMapInViewport(g *gameMap) {
	f := r.currentFactionSeeingTheScreen
	r.currentFactionSeeingTheScreen = f
	vx, vy := f.cursor.getCameraCoords()
	for x := vx; x < vx+VIEWPORT_W; x++ {
		for y := vy; y < vy+VIEWPORT_H; y++ {
			if g.areCoordsValid(x, y) {
				if f.wereCoordsSeen(x, y) {
					tileApp := g.tileMap[x][y].getAppearance()
					if f.areCoordsInSight(x, y) {
						r.setFgColorByCcell(tileApp)
					} else {
						cw.SetFgColor(cw.DARK_BLUE)
					}
					cw.PutChar(tileApp.char, x-vx, y-vy)
				} else {
					cw.PutChar(' ', x-vx, y-vy)
				}
			}
		}
	}
}

func (r *rendererStruct) areGlobalCoordsOnScreen(gx, gy int) bool {
	vx, vy := r.currentFactionSeeingTheScreen.cursor.getCameraCoords()
	return geometry.AreCoordsInRect(gx, gy, vx, vy, VIEWPORT_W, VIEWPORT_H)
}

func (r *rendererStruct) areGlobalCoordsOnScreenForFaction(gx, gy int, f *faction) bool {
	vx, vy := f.cursor.getCameraCoords()
	return geometry.AreCoordsInRect(gx, gy, vx, vy, VIEWPORT_W, VIEWPORT_H)
}
