package main 

import cw "github.com/sidav/golibrl/console"

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
