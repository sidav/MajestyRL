package main 

type factionEconomy struct {
	currentResources resourceStock
	maxResources map[resourceType] int
}

func (f *factionEconomy) addResource(rescount int, restype resourceType) {
	f.currentResources.amount[restype] += rescount
}

func (f *factionEconomy) resetMaxResources() {
	for rtype := range f.maxResources {
		f.maxResources[rtype] = 0
	}
}

func (f *factionEconomy) adjustResourcesToMax() {
	for rtype := range f.maxResources {
		if f.currentResources.amount[rtype] > f.maxResources[rtype] {
			f.currentResources.amount[rtype] = f.maxResources[rtype]
		}
	}
}

func createFactionEconomy() *factionEconomy {
	f := &factionEconomy{
		currentResources: resourceStock{map[resourceType]int{RESTYPE_GOLD: 1000, RESTYPE_WOOD: 100}},
		maxResources:	map[resourceType]int{RESTYPE_GOLD: 2500, RESTYPE_WOOD: 250},
	}
	return f
}
