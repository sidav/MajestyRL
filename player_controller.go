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

	// testing
	case "ENTER":
		for i := 0; i < 10; i++ {
			CURRENT_MAP.addBuilding(createBuildingAtCoords("HUT", false, rnd.Rand(mapW), rnd.Rand(mapH), f), true)
			reportToPlayer("cheats done", f)
		}
	case " ":
		for i := 0; i < 10; i++ {
			x, y := rnd.Rand(mapW), rnd.Rand(mapH)
			newbid := &bid{intent_type_for_this_bid: INTENT_BUILD, x: x, y: y, targetPawn: createBuildingAtCoords("HUT", false, x, y, f)}
			CURRENT_MAP.addBid(newbid)
			reportToPlayer("cheats done", f)
		}
	}
}
