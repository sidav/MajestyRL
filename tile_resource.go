package main

import cw "github.com/sidav/golibrl/console"

type resourceType byte 

const (
	RESTYPE_GOLD resourceType = iota 
)

var tileResAppearances = map[resourceType]*ccell{
	RESTYPE_GOLD: &ccell{char: '*', color: cw.YELLOW},
}

type tileResource struct {
	resType resourceType
	amount int 
}

func (r *tileResource) getAppearance() *ccell {
	return tileResAppearances[r.resType]
}

func (r *tileResource) getResourceName() string {
	return "GOLD"
}
