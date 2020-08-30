package main

import "github.com/sidav/golibrl/console"

type decal struct {
	appearance *ccell
	x, y int
	lastTurnToBeRendered int
}

func (r *rendererStruct) renderDecalsInViewport() {
	for i := range r.decalsBuffer {
		decal := r.decalsBuffer[i]
		if decal.lastTurnToBeRendered >= getCurrentTurn() {
			if r.areGlobalCoordsOnScreenForFaction(decal.x, decal.y, r.currentFactionSeeingTheScreen) {
				sx, sy := decal.x - r.vx, decal.y - r.vy
				r.renderCcellOnScreenCoords(r.decalsBuffer[i].appearance, sx, sy)
			}
		}
	}
}

func (r *rendererStruct) cleanDecalsBuffer() {
	for i := range r.decalsBuffer {
		if r.decalsBuffer[i].lastTurnToBeRendered < getCurrentTurn() {
			r.decalsBuffer = append(r.decalsBuffer[:i], r.decalsBuffer[i+1:]...)
			return
		}
	}
}

func addBasicDecalToRender(x, y, times int) {
	RENDERER.decalsBuffer = append(RENDERER.decalsBuffer,
		&decal{appearance: &ccell{color: console.RED, char: 'X'},
		x: x, y: y, lastTurnToBeRendered: getCurrentTurn()+times})
}
