package main

type unitStaticData struct {
	name                             string
	code                             string
	app                              *ccell
	maxHitpointsMin, maxHitpointsMax int
	defaultWeaponCode                string

	timeToConstruct int
	cost            int

	sightRange int

	canCollectTaxes bool
	canBuild        bool
	canMine         bool
}

var staticUnitDataTable = map[string]*unitStaticData{
	// non-heroes
	"PEASANT": &unitStaticData{
		name:     "Peasant",
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
		canBuild:        true,
		canMine:         true,
		canCollectTaxes: false,
	},
	"TAXCOLLECTOR": &unitStaticData{
		name:     "Tax collector",
		// appearance
		app: &ccell{
			char: 't', color: 7,
		},
		// hp
		maxHitpointsMin: 10, maxHitpointsMax: 10,
		// construction
		timeToConstruct: 100,
		// cost
		cost: 0,
		// misc
		canCollectTaxes: true,
	},
	"GUARD": &unitStaticData{
		name:     "Guardian",
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
		// sight range
		sightRange: 5,
		// misc
		defaultWeaponCode: "HALBERD",
		canBuild:          false,
		canCollectTaxes:   false,
	},
	"ROYALGUARD": &unitStaticData{
		name:     "Royal Guardian",
		// appearance
		app: &ccell{
			char: 'R', color: 7,
		},
		// hp
		maxHitpointsMin: 50, maxHitpointsMax: 50,
		// construction
		timeToConstruct: 200,
		// cost
		cost:       0,
		sightRange: 7,
		// misc
		defaultWeaponCode: "SWORD",
		canBuild:          false,
		canCollectTaxes:   false,
	},

	// Neutrals
	"GOBLIN": &unitStaticData{
		name:     "Goblin",
		// appearance
		app: &ccell{
			char: 'g', color: 3,
		},
		// hp
		maxHitpointsMin: 10, maxHitpointsMax: 20,
		sightRange: 15,
		// construction
		timeToConstruct: 100,
		// cost
		cost: 0,
		// misc
		defaultWeaponCode: "HALBERD",
		canBuild:          false,
		canCollectTaxes:   false,
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
		defaultWeaponCode: "HALBERD",
		canBuild:          true,
		canCollectTaxes:   false,
	},
}

func getUnitStaticDataFromTable(code string) *unitStaticData {
	unitsd := staticUnitDataTable[code]
	if unitsd == nil {
		return staticUnitDataTable["NULL"]
	}
	unitsd.code = code
	return unitsd
}
