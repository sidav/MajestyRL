package main

type unitStaticData struct {
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