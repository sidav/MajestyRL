package main

import "github.com/sidav/golibrl/console"

type decal struct {
	appearance *ccell
	x, y int
	timesToRender int
	lastTurnRendered int
}

func (r *rendererStruct) renderDecalsInViewport() {
	for i := range r.decalsBuffer {
		decal := r.decalsBuffer[i]
		if decal.timesToRender > 0 {
			if r.areGlobalCoordsOnScreenForFaction(decal.x, decal.y, r.currentFactionSeeingTheScreen) {
				sx, sy := decal.x - r.vx, decal.y - r.vy
				r.renderCcellOnScreenCoords(r.decalsBuffer[i].appearance, sx, sy)
				if decal.lastTurnRendered < getCurrentTurn() {
					decal.lastTurnRendered = getCurrentTurn()
					decal.timesToRender -= 1
				}
			}
		}
	}
}

func (r *rendererStruct) cleanDecalsBuffer() {
	for i := range r.decalsBuffer {
		if r.decalsBuffer[i].timesToRender <= 0 {
			r.decalsBuffer = append(r.decalsBuffer[:i], r.decalsBuffer[i+1:]...)
			return
		}
	}
}

func addBasicDecalToRender(x, y, times int) {
	RENDERER.decalsBuffer = append(RENDERER.decalsBuffer,
		&decal{appearance: &ccell{color: console.RED, char: 'X'},
		x: x, y: y, timesToRender: times, lastTurnRendered:getCurrentTurn()+1})
}
