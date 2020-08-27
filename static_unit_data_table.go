package main

type UTYPE uint8

const (
	UTYPE_UNKNOWN UTYPE = iota
	UTYPE_WORKER
	UTYPE_GUARD
	UTYPE_ROYAL_GUARD

	UTYPE_GOBLIN
)

type unitStaticData struct {
	name                             string
	app                              *ccell
	maxHitpointsMin, maxHitpointsMax int
	defaultWeaponCode                string

	unitType UTYPE

	timeToConstruct int
	cost            int

	sightRange int

	carriesGold bool // TODO: rename to "collectTaxes or something."
	canBuild    bool
	canMine     bool
}

var staticUnitDataTable = map[string]*unitStaticData{
	// non-heroes
	"PEASANT": &unitStaticData{
		name:     "Peasant",
		unitType: UTYPE_WORKER,
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
		canMine:     true,
		carriesGold: false,
	},
	"GUARD": &unitStaticData{
		name:     "Guardian",
		unitType: UTYPE_GUARD,
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
		carriesGold:       false,
	},
	"ROYALGUARD": &unitStaticData{
		name:     "Royal Guardian",
		unitType: UTYPE_ROYAL_GUARD,
		// appearance
		app: &ccell{
			char: 'R', color: 7,
		},
		// hp
		maxHitpointsMin: 50, maxHitpointsMax: 50,
		// construction
		timeToConstruct: 200,
		// cost
		cost: 0,
		sightRange: 7,
		// misc
		defaultWeaponCode: "HALBERD",
		canBuild:          false,
		carriesGold:       false,
	},

	// Neutrals
	"GOBLIN": &unitStaticData{
		name:     "Goblin",
		unitType: UTYPE_GOBLIN,
		// appearance
		app: &ccell{
			char: 'g', color: 3,
		},
		// hp
		maxHitpointsMin: 10, maxHitpointsMax: 20,
		// construction
		timeToConstruct: 100,
		// cost
		cost: 0,
		// misc
		defaultWeaponCode: "HALBERD",
		canBuild:          false,
		carriesGold:       false,
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
		carriesGold:       false,
	},
}

func getUnitStaticDataFromTable(code string) *unitStaticData {
	unitsd := staticUnitDataTable[code]
	if unitsd == nil {
		return staticUnitDataTable["NULL"]
	}
	return unitsd
}
