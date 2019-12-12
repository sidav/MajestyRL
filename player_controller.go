package main

import cw "github.com/sidav/golibrl/console"

type playerController struct {
	exit bool
}

func (pc *playerController) controlAsFaction(f *faction) {
	RENDERER.renderScreen(f)
	keyPressed := cw.ReadKeyAsync()
	switch keyPressed {
	case "ESCAPE":
		pc.exit = true
	case "ENTER":
		// test
		for i := 0; i < 10; i++ {
			CURRENT_MAP.addBuilding(createBuildingAtCoords("HUT", false, (CURRENT_TICK+13*i)%mapW, (CURRENT_TICK+29*i)%mapH, f), true)
			reportToPlayer("cheats done", f)
		}
	}
}
