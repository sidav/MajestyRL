package main 

type factionEconomy struct {
	currentResources resourceStock
	maxResources map[resourceType] int
}

func (f *factionEconomy) addResource(rescount int, restype resourceType) {
	f.currentResources.amount[restype] += rescount
}

func createFactionEconomy() *factionEconomy {
	f := &factionEconomy{
		currentResources: resourceStock{map[resourceType]int{RESTYPE_GOLD: 1000, RESTYPE_WOOD: 200}},
		maxResources:	map[resourceType]int{RESTYPE_GOLD: 2500, RESTYPE_WOOD: 250},
	}
	return f
}
