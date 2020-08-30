package main

type buildingStaticData struct {
	name                  string
	app                   *buildingAppearance
	underConstructionData underConstructionData
	maxHitpoints          int

	cost int

	goldStorage int

	// next line is obsolete
	// maxWorkers, maxTaxCollectors, maxGuards, maxRoyalGuards int
	housing_unittypes      []string
	housing_max_residents  []int
	housing_respawn_period []int

	allowsTightPlacement bool

	//tech
	allowsBuildings []string
	deniesBuildings []string
}

var staticBuildingDataTable = map[string]*buildingStaticData{
	"PALACE": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"O----O",
				"|/||\\|",
				"|-OO-|",
				"|\\||/|",
				"O----O",
			},
			colors: [][]int{
				{-1, 7, 7, 7, 7, -1},
				{7, 7, 7, 7, 7, 7},
				{7, 7, -1, -1, 7, 7},
				{7, 7, 7, 7, 7, 7},
				{-1, 7, 7, 7, 7, -1},
			},
		},
		name: "Palace",
		// construction
		underConstructionData: underConstructionData{
			maxConstructedAmount: 500,
		},
		maxHitpoints: 1000,
		// cost
		cost: 100000,
		// tech
		allowsBuildings: []string{"HUT", "GOLDVAULT", "MARKETPLACE", "WALL"},
		// misc
		goldStorage:            5000,
		housing_unittypes:      []string{"PEASANT", "GUARD", "ROYALGUARD"},
		housing_max_residents:  []int {3, 1, 1},
		housing_respawn_period: []int {100, 250, 300},
	},
	"HUT": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"/-\\",
				"|=|",
				"---",
			},
			colors: [][]int{
				{7, 7, 7},
				{7, -1, 7},
				{7, 7, 7},
			},
		},
		name: "Hut",
		// construction
		underConstructionData: underConstructionData{
			maxConstructedAmount: 250,
		},
		maxHitpoints: 100,
		// cost
		cost: 500,
		// misc
		housing_unittypes:      []string{"PEASANT"},
		housing_max_residents:  []int {2},
		housing_respawn_period: []int {100},
	},
	"GOLDVAULT": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"|-|",
				"|=|",
				"-O-",
			},
			colors: [][]int{
				{7, 7, 7},
				{7, -1, 7},
				{7, 7, 7},
			},
		},
		name: "Gold vault",
		// construction
		underConstructionData: underConstructionData{
			maxConstructedAmount: 250,
		},
		maxHitpoints: 300,
		// cost
		cost: 1000,
		// misc
		goldStorage:            5000,
		housing_unittypes:      []string{"GUARD"},
		housing_max_residents:  []int {1},
		housing_respawn_period: []int {250},
	},
	"MARKETPLACE": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"|--|",
				"|=.#|",
				"-/\\-",
			},
			colors: [][]int{
				{7, 7, 7, 7},
				{7, -1, -1, 7},
				{7, -1, -1, 7},
			},
		},
		name: "Marketplace",
		// construction
		underConstructionData: underConstructionData{
			maxConstructedAmount: 450,
		},
		maxHitpoints: 600,
		// cost
		cost: 1000,
		// misc
	},
	"WALL": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"#",
			},
			colors: [][]int{
				{7},
			},
		},
		name: "Wall",
		// construction
		allowsTightPlacement: true,
		underConstructionData: underConstructionData{
			maxConstructedAmount: 250,
		},
		maxHitpoints: 100,
		// cost
		cost: 100,
	},

	// ENEMY/NEUTRAL BUILDINGS
	"GOBLINCAMP": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"/-\\",
				"|x|",
				"\\v/",
			},
			colors: [][]int{
				{7, 7, 7},
				{7, -1, 7},
				{7, -1, 7},
			},
		},
		name: "Goblin camp",
		// construction
		underConstructionData: underConstructionData{
			maxConstructedAmount: 250,
		},
		maxHitpoints: 300,
		// cost
		cost: 1000,
		// misc
		housing_unittypes:      []string{"GOBLIN"},
		housing_max_residents:  []int {1},
		housing_respawn_period: []int {50},
	},
}

func getBuildingStaticDataFromTable(code string) *buildingStaticData {
	bldsd := staticBuildingDataTable[code]
	if bldsd == nil {
		return staticBuildingDataTable["NULL"]
	}
	return bldsd
}
