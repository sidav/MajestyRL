package main

import cw "github.com/sidav/golibrl/console"

type resourceType byte 

const (
	RESTYPE_GOLD resourceType = iota 
	RESTYPE_WOOD 
)

var tileResAppearances = map[resourceType]*ccell{
	RESTYPE_GOLD: &ccell{char: '*', color: cw.YELLOW},
	RESTYPE_WOOD: &ccell{char: '%', color: cw.GREEN},
}

type tileResource struct {
	resType resourceType
	amount int 
}

func (r *tileResource) getAppearance() *ccell {
	return tileResAppearances[r.resType]
}

func (r *tileResource) getResourceName() string {
	switch r.resType {
	case RESTYPE_GOLD:
		return "GOLD"
	case RESTYPE_WOOD:
		return "WOOD"
	}
	return "UNKNOWN RESOURCE"
}

func getResourceName(resType resourceType) string {
	switch resType {
	case RESTYPE_GOLD:
		return "Gold"
	case RESTYPE_WOOD:
		return "Wood"
	}
	return "UNKNOWN RESOURCE"
}
