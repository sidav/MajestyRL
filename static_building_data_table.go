package main 

type buildingStaticData struct {
	app *buildingAppearance
	underConstructionData *underConstructionData
}

var staticBuildingDataTable = map[string]*buildingStaticData {
	"PALACE": &buildingStaticData{
		app: &buildingAppearance{
			chars: []string{
				"O---O",
				"|/|\\|",
				"|-O-|",
				"|\\|/|",
				"O---O",
			},
			colors: [][]int {
				{-1, 7, 7, 7, -1},
				{7, 7, 7, 7, 7},
				{7, 7, -1, 7, 7},
				{7, 7, 7, 7, 7},
				{-1, 7, 7, 7, -1},
			},
		},
		underConstructionData: &underConstructionData {
			maxConstructedAmount: 1000,
		},
	},
}
