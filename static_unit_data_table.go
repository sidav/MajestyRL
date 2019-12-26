package main

type unitStaticData struct {
	name                             string
	app                              *ccell
	maxHitpointsMin, maxHitpointsMax int

	timeToConstruct int
	cost            int

	carriesGold bool
	canBuild    bool
}

var staticUnitDataTable = map[string]*unitStaticData{
	// non-heroes
	"PEASANT": &unitStaticData{
		name: "Peasant",
		// appearance
		app: &ccell{
			char: 'p', color: 7,
		},
		// hp
		maxHitpointsMin: 5, maxHitpointsMax: 5,
		// construction
		timeToConstruct: 100,
		// cost
		cost: 0,
		// misc
		canBuild:    true,
		carriesGold: false,
	},
	"GUARD": &unitStaticData{
		name: "Guardian",
		// appearance
		app: &ccell{
			char: 'G', color: 7,
		},
		// hp
		maxHitpointsMin: 30, maxHitpointsMax: 30,
		// construction
		timeToConstruct: 100,
		// cost
		cost: 0,
		// misc
		canBuild:    false,
		carriesGold: false,
	},

	"NULL": &unitStaticData{
		name: "UNKNOWN UNIT",
		// appearance
		app: &ccell{
			char: 'p', color: 7,
		},
		// hp
		maxHitpointsMin: 5, maxHitpointsMax: 5,
		// construction
		timeToConstruct: 100,
		// cost
		cost: 0,
		// misc
		canBuild:    true,
		carriesGold: false,
	},
}

func getUnitStaticDataFromTable(code string) *unitStaticData {
	unitsd := staticUnitDataTable[code]
	if unitsd == nil {
		return staticUnitDataTable["NULL"]
	}
	return unitsd
}
