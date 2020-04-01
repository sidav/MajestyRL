package main

var TEST_MAP = &[]string{
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"..........~.....................~~~...................................................~.............",
	".........~~~....................~~~..................................................~~~............",
	"..........~......................~~...................................................~.............",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"...~~~........................................................................................~~~...",
	"...~~~........................................................................................~~~...",
	"...~~~........................................................................................~~~...",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	".................................................................~~.................................",
	"........~........................................................~~~................................",
	".......~~~.......................................................~~~.......................~~~......",
	"............................................................................................~.......",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
	"....................................................................................................",
}

func initTestMission() {
	CURRENT_MAP = &gameMap{}
	CURRENT_MAP.init()
	CURRENT_MAP.initTileMap(TEST_MAP)
	playerFaction := createFaction("Your Majesty", 0, true, false)
	CURRENT_MAP.factions = append(CURRENT_MAP.factions, playerFaction)
	// CURRENT_MAP.addBid(&bid{intent_type_for_this_bid: INTENT_BUILD, targetPawn: createBuildingAtCoords("PALACE", false, mapW/2, mapH/2, playerFaction)})
	CURRENT_MAP.addPawn(createBuildingAtCoords("PALACE", true, mapW/2, mapH/2, playerFaction))
	for i := 10; i < 15; i++ {
		CURRENT_MAP.addPawn(createUnitAtCoords("PEASANT", 3*i, 0, playerFaction))
	}
	CURRENT_MAP.addPawn(createUnitAtCoords("GUARD", 0, 0, playerFaction))
}
