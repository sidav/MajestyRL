package main

type buildingStaticData struct {
	name                  string
	app                   *buildingAppearance
	underConstructionData underConstructionData
	maxHitpoints          int

	cost *resourceStock

	resourceStorage map[resourceType]int
	taxGoldGeneration int

	// next line is obsolete
	// maxWorkers, maxTaxCollectors, maxGuards, maxRoyalGuards int
	housing_unittypes      []string
	housing_max_residents  []int
	housing_respawn_period []int

	allowsTightPlacement bool
	autoBidType string
	autoBidRadius int

	//tech
	allowsBuildings []string
	deniesBuildings []string
}

var staticBuildingDataTable = map[string]*buildingStaticData{
	"PALACE": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"O####O",
				"#/||\\#",
				"#-OO-#",
				"#\\||/#",
				"O####O",
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
		cost: &resourceStock{amount: map[resourceType]int{RESTYPE_GOLD: 99999}},
		// tech
		allowsBuildings: []string{"HUT", "GOLDVAULT", "GUARDHOUSE", "FORESTER", "MARKETPLACE", "SAWMILL", "WALL"},
		// misc
		resourceStorage: map[resourceType]int{RESTYPE_GOLD: 2500, RESTYPE_WOOD: 250},
		// taxGoldGeneration:      25,
		housing_unittypes:      []string{"PEASANT", "TAXCOLLECTOR", "ROYALGUARD"},
		housing_max_residents:  []int{2, 1, 1},
		housing_respawn_period: []int{100, 200, 300},
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
		cost: &resourceStock{amount: map[resourceType]int{RESTYPE_GOLD: 150, RESTYPE_WOOD: 50}},
		taxGoldGeneration: 1,
		// misc
		housing_unittypes:      []string{"PEASANT"},
		housing_max_residents:  []int{2},
		housing_respawn_period: []int{100},
	},
	"FORESTER": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"/-\\",
				"===",
				"---",
			},
			colors: [][]int{
				{7, 7, 7},
				{-1, -1, 2},
				{7, 7, 7},
			},
		},
		name: "Forester",
		// construction
		underConstructionData: underConstructionData{
			maxConstructedAmount: 350,
		},
		maxHitpoints: 100,
		autoBidType: "GROW_FOREST",
		autoBidRadius: 5,
		// cost
		cost: &resourceStock{amount: map[resourceType]int{RESTYPE_GOLD: 250, RESTYPE_WOOD: 50}},
	},
	"SAWMILL": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"/-\\",
				"==o",
				"---",
			},
			colors: [][]int{
				{7, 7, 7},
				{-1, -1, 2},
				{7, 7, 7},
			},
		},
		name: "Sawmill",
		// construction
		underConstructionData: underConstructionData{
			maxConstructedAmount: 350,
		},
		maxHitpoints: 100,
		autoBidType: "MINE_FOREST",
		autoBidRadius: 10,
		// cost
		cost: &resourceStock{amount: map[resourceType]int{RESTYPE_GOLD: 50, RESTYPE_WOOD: 50}},
		resourceStorage: map[resourceType]int{RESTYPE_WOOD: 250},
	},
	"GOLDVAULT": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"|-|",
				"|=|",
				"#O#",
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
		cost: &resourceStock{amount: map[resourceType]int{RESTYPE_GOLD: 500}},
		// misc
		resourceStorage: map[resourceType]int{RESTYPE_GOLD: 2500},
		housing_unittypes:      []string{"ROYALGUARD"},
		housing_max_residents:  []int{1},
		housing_respawn_period: []int{350},
	},
	"GUARDHOUSE": &buildingStaticData{
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
		name: "Guardhouse",
		// construction
		underConstructionData: underConstructionData{
			maxConstructedAmount: 250,
		},
		maxHitpoints: 300,
		// cost
		cost: &resourceStock{amount: map[resourceType]int{RESTYPE_GOLD: 350}},
		// misc
		housing_unittypes:      []string{"GUARD"},
		housing_max_residents:  []int{1},
		housing_respawn_period: []int{250},
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
		cost: &resourceStock{amount: map[resourceType]int{RESTYPE_GOLD: 1500, RESTYPE_WOOD: 150}},
		taxGoldGeneration: 25,
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
		cost: &resourceStock{amount: map[resourceType]int{RESTYPE_GOLD: 100}},
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
		cost: &resourceStock{amount: map[resourceType]int{RESTYPE_GOLD: 100}},
		// misc
		housing_unittypes:      []string{"GOBLIN"},
		housing_max_residents:  []int{1},
		housing_respawn_period: []int{50},
	},
}

func getBuildingStaticDataFromTable(code string) *buildingStaticData {
	bldsd := staticBuildingDataTable[code]
	if bldsd == nil {
		return staticBuildingDataTable["NULL"]
	}
	return bldsd
}
