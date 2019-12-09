package main 

var TEST_MAP = &[]string {
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
	CURRENT_MAP.addBuilding(createBuildingAtCoords("PALACE", mapW/2, mapH/2, playerFaction), true)
}