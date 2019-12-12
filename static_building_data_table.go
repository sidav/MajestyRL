package main

type buildingStaticData struct {
	app                   *buildingAppearance
	underConstructionData *underConstructionData
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
		// construction
		underConstructionData: &underConstructionData{
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
}
