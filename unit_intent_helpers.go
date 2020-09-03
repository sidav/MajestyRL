package main

func (g *gameMap) getNearestBuildingWithStorageOfType(ix, iy int, rtype resourceType) *pawn {
	minDist := 999999999 // should be enough lol
	var buildingToReturn *pawn
	for _, bld := range g.pawns {
		if bld.isBuilding() && !bld.asBuilding.isUnderConstruction() {
			if _, hasStorage := bld.asBuilding.getStaticData().resourceStorage[rtype]; hasStorage {
				if buildingToReturn == nil {
					buildingToReturn = bld
				}
				bldx, bldy := bld.getCenter()
				dist := (bldx-ix)*(bldx-ix) + (bldy-iy)*(bldy-iy)
				if dist < minDist {
					minDist = dist
					buildingToReturn = bld
				}
			}
		}
	}
	return buildingToReturn
}

func (u *pawn) dropCurrentIntent() {
	if u.asUnit.intent.sourceBid != nil {
		u.asUnit.intent.sourceBid.drop()
	}
	u.asUnit.intent = nil
	if u.asUnit.carriedResourceAmount > 0 {
		u.dropResources()
	}
}
