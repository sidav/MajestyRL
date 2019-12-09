package main 

import cw "github.com/sidav/golibrl/console"

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

type rendererStruct struct {}

func (r *rendererStruct) setFgColorByCcell(c *ccell) {
	cw.SetFgColor(c.color)
	// cw.SetFgColorRGB(c.r, c.g, c.b)
}

func (r *rendererStruct) updateBoundsIfNeccessary(force bool) {
	if cw.WasResized() || force {
		CONSOLE_W, CONSOLE_H = cw.GetConsoleSize()
		VIEWPORT_W           = CONSOLE_W / 2
		VIEWPORT_H           = CONSOLE_H - LOG_HEIGHT - 1
		SIDEBAR_X            = VIEWPORT_W + 1
		SIDEBAR_W            = CONSOLE_W - VIEWPORT_W - 1
		SIDEBAR_H            = CONSOLE_H - LOG_HEIGHT
		SIDEBAR_FLOOR_2      = 7  // y-coord right below resources info
		SIDEBAR_FLOOR_3      = 11 // y-coord right below "floor 2"
	}
}

func (r *rendererStruct) renderScreen(f *faction) {
	r.updateBoundsIfNeccessary(false)
	cw.Clear_console()
	r.renderMapInViewport(f, CURRENT_MAP)
	cw.Flush_console()
}

func (r *rendererStruct) renderMapInViewport(f *faction, g *gameMap) {
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