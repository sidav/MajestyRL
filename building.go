package main

type building struct {
	hasBeenPlaced bool
	code          string
	beingConstructed *underConstructionData 
}

func (b *building) getAppearance() *buildingAppearance {
	return staticBuildingDataTable[b.code].app
}


func (b *building) getSize() (int, int) {
	w := len(b.getAppearance().chars)
	h := len(b.getAppearance().chars[0])
	return w, h 
}
