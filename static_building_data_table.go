package main

type buildingStaticData struct {
	name                  string
	app                   *buildingAppearance
	underConstructionData underConstructionData
	maxHitpoints          int

	cost int

	goldStorage int

	maxWorkers, maxTaxCollectors, maxGuards, maxRoyalGuards int
}

var staticBuildingDataTable = map[string]*buildingStaticData{
	"PALACE": &buildingStaticData{
		// appearance
		app: &buildingAppearance{
			chars: []string{
				"O---O",
				"|/|\\|",
				"|-O-|",
				"|\\|/|",
				"O---O",
			},
			colors: [][]int{
				{-1, 7, 7, 7, -1},
				{7, 7, 7, 7, 7},
				{7, 7, -1, 7, 7},
				{7, 7, 7, 7, 7},
				{-1, 7, 7, 7, -1},
			},
		},
		name: "Palace",
		// construction
		underConstructionData: underConstructionData{
			maxConstructedAmount: 1000,
		},
		maxHitpoints: 1000,
		// cost
		cost: 100000,
		// misc
		goldStorage:      5000,
		maxWorkers:       2,
		maxTaxCollectors: 1,
		maxRoyalGuards:   1,
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
		maxWorkers: 1,
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
		goldStorage:      5000,
		maxRoyalGuards:   1,
	},
}
