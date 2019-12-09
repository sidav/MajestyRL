package main 

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
	},
}
