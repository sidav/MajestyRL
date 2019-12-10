package main

import cw "github.com/sidav/golibrl/console"

func (r *rendererStruct) renderPawnsInViewport(g *gameMap) {
	f := r.currentFactionSeeingTheScreen
	vx, vy := f.cursor.getCameraCoords()
	for _, p := range g.pawns {
		// cx, cy := p.getCenter()
		if p.isBuilding() {
			if p.asBuilding.beingConstructed != nil {
				r.renderBuildingUnderConstruction(f, p, vx, vy, false)
			} else {
				r.renderBuilding(f, p, g, vx, vy, false)
			}
		} else {
			// TODO
		}
	}
}

func (r *rendererStruct) renderBuilding(f *faction, p *pawn, g *gameMap, vx, vy int, inverse bool) {
	b_w, b_h := p.getSize()
	code := p.asBuilding.code
	app := staticBuildingDataTable[code].app
	bx, by := p.getCoords()
	colorToRender := 0
	for x := 0; x < b_w; x++ {
		for y := 0; y < b_h; y++ {
			if true { // p.currentConstructionStatus == nil {
				color := app.colors[x][y]
				if f.areCoordsInSight(bx+x, by+y) {
					if color == -1 {
						colorToRender = p.faction.getFactionColor()
					} else {
						colorToRender = color
					}
				} else {
					colorToRender = cw.DARK_BLUE
				}
			} else { // building is under construction
				colorToRender = cw.DARK_GREEN
				if getCurrentTurn()%2 == 0 {
					colorToRender = cw.GREEN
				}
			}
			if r.areGlobalCoordsOnScreen(bx+x, by+y) && f.wereCoordsSeen(bx+x, by+y) {
				if inverse {
					cw.SetBgColor(colorToRender)
					cw.SetFgColor(cw.BLACK)
				} else {
					cw.SetFgColor(colorToRender)
				}
				cw.PutChar(int32(app.chars[y][x]), bx+x-vx, by+y-vy)
			}
		}
	}
	cw.SetBgColor(cw.BLACK)
}

func (r *rendererStruct) renderBuildingUnderConstruction(f *faction, p *pawn, vx, vy int, inverse bool) {
	b_w, b_h := p.getSize()
	bx, by := p.getCoords()
	constrAmount := p.asBuilding.beingConstructed.currentConstructedAmount
	colorToRender := cw.DARK_YELLOW
	charToRender := '='
	for x := 0; x < b_w; x++ {
		for y := 0; y < b_h; y++ {
			if b_w > 1 && b_h > 1 {
				// the next code is magic
				framex := constrAmount % (2*b_w - 2)
				if framex < b_w && framex == x || framex >= b_w && 2*(b_w-1)-framex == x {
					colorToRender = cw.DARK_YELLOW
					charToRender = '='
				}
				framey := constrAmount % (2*b_h - 2)
				if framey < b_h && framey == y || framey >= b_h && 2*(b_h-1)-framey == y {
					if charToRender == '=' {
						colorToRender = cw.YELLOW
						charToRender = 'X'
					} else {
						colorToRender = cw.DARK_YELLOW
						charToRender = '='
					}
				}
			} else { // another animation for width = 1 (for escaping division by zero above)
				if getCurrentTurn()%2 == 0 {
					charToRender = '='
					colorToRender = cw.DARK_YELLOW
				}
			}
			if r.areGlobalCoordsOnScreen(bx+x, by+y) && f.wereCoordsSeen(bx+x, by+y) {
				if inverse {
					cw.SetBgColor(colorToRender)
					cw.SetFgColor(cw.BLACK)
				} else {
					cw.SetFgColor(colorToRender)
				}
				cw.PutChar(charToRender, bx+x-vx, by+y-vy)
			}
		}
	}
	cw.SetBgColor(cw.BLACK)
}
