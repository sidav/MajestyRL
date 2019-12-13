package main 

import (
	cw "github.com/sidav/golibrl/console"
	"fmt"
)

func (r *rendererStruct) renderUI() {
	r.renderUIOutline()
	r.renderFactionStats()
}

func (r *rendererStruct) renderUIOutline() {
	// if IS_PAUSED {
	// 	cw.SetBgColor(f.getFactionColor())
	// } else {
	// 	cw.SetFgColor(f.getFactionColor())
	// }
	cw.SetFgColor(r.currentFactionSeeingTheScreen.getFactionColor())
	for y := 0; y < VIEWPORT_H; y++ {
		cw.PutChar('|', VIEWPORT_W, y)
	}
	for x := 0; x < CONSOLE_W; x++ {
		cw.PutChar('-', x, VIEWPORT_H)
	}
	cw.PutChar('+', VIEWPORT_W, VIEWPORT_H)
	cw.SetBgColor(cw.BLACK)
}

func (r *rendererStruct) renderFactionStats() {
	f := r.currentFactionSeeingTheScreen
	eco := f.economy
	statsx := VIEWPORT_W + 1

	// fr, fg, fb := getFactionRGB(f.factionNumber)
	// cw.SetFgColorRGB(fr, fg, fb)
	if IS_PAUSED {
		cw.SetFgColor(f.getFactionColor())
		cw.PutString(f.name+", ", statsx, 0)
		cw.SetFgColor(cw.YELLOW)
		cw.PutString(fmt.Sprintf("turn %d (PAUSED)", getCurrentTurn()), statsx+len(f.name)+2, 0)
	} else {
		cw.SetFgColor(f.getFactionColor())
		cw.PutString(fmt.Sprintf("%s: turn %d", f.name, getCurrentTurn()), statsx, 0)
	}
	cw.SetFgColor(cw.YELLOW)
	r.renderStatusbar("GOLD:", eco.currentGold, 7500, statsx, 2, SIDEBAR_W-3, cw.YELLOW, cw.DARK_YELLOW, false)
}

func (r *rendererStruct) renderStatusbar(name string, curvalue, maxvalue, x, y, width, fillColor, emptyColor int, hideNumericValues bool) {
	if !hideNumericValues {
		name = fmt.Sprintf("%s %d/%d", name, curvalue, maxvalue)
	}
	cw.PutString(name, x, y)
	barWidth := width - len(name)
	var filledCells int
	if maxvalue > 0 {
		filledCells = barWidth * curvalue / maxvalue
	} else {
		filledCells = 0
	}
	barStartX := x + len(name) + 1
	for i := 0; i < barWidth; i++ {
		if i < filledCells {
			cw.SetFgColor(fillColor)
			cw.PutChar('=', i+barStartX, y)
		} else {
			cw.SetFgColor(emptyColor)
			cw.PutChar('-', i+barStartX, y)
		}
	}
}
