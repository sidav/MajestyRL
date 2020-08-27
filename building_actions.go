package main 

import "fmt"

const (
	REGENERATE_WORKERS_EACH = 1000 // ticks 
	REGENERATE_GUARDS_EACH = 2500 
)

type buildingLogic struct {
}

func (bl *buildingLogic) act(bld *pawn) {
	if bld.asBuilding.isUnderConstruction() {
		return 
	}
	bl.generatePawns(bld)
	bl.actForEachPawnInside(bld)
}

func (bl *buildingLogic) generatePawn(bld *pawn, code string) {
	bstatic := getBuildingStaticDataFromTable(bld.asBuilding.code)
	// spawn one more worker 
	newPawn := createUnitAtCoords(code, bld.x, bld.y, bld.faction)
	bld.AddAndRegisterNewPawn(newPawn)
	bld.asBuilding.recalculateCurrResidents()
	log.AppendMessage(fmt.Sprintf("%s created (%d/%d) at turn %d", code, bld.asBuilding.currWorkers, bstatic.maxWorkers, CURRENT_TICK))
}

func (bl *buildingLogic) generatePawns(bld *pawn) {
	bstatic := getBuildingStaticDataFromTable(bld.asBuilding.code)
	bld.asBuilding.recalculateCurrResidents()
	// LOG.AppendMessage(fmt.Sprintf("Peasant (%d/%d)", bld.asBuilding.currWorkers, bstatic.maxWorkers))
	if bld.asBuilding.currWorkers < bstatic.maxWorkers && CURRENT_TICK % REGENERATE_WORKERS_EACH == 0 {
		// spawn one more worker 
		bl.generatePawn(bld, "PEASANT")
	} 
	if bld.asBuilding.currGuards < bstatic.maxGuards && CURRENT_TICK % REGENERATE_GUARDS_EACH == 0 {
		// spawn one more guard
		bl.generatePawn(bld, "GUARD")
	}
	if bld.asBuilding.currRoyalGurads < bstatic.maxRoyalGuards && CURRENT_TICK % REGENERATE_GUARDS_EACH == 0 {
		// spawn one more guard
		bl.generatePawn(bld, "ROYALGUARD")
	}
}

func (bl *buildingLogic) actForEachPawnInside(bld *pawn) {
	// bstatic := staticBuildingDataTable[bld.asBuilding.code]
	for i := 0; i < len(bld.asBuilding.pawnsInside); i++ {
		p := bld.asBuilding.pawnsInside[i]
		ULOGIC.decideNewIntent(p)
		if ULOGIC.wantsToLeaveBuilding(p) {
			// remove unit from unitsInside and place it outside the building 
			bld.asBuilding.removePawnFromInside(p)
			CURRENT_MAP.addPawn(p)
			CURRENT_MAP.placePawnNearPawn(p, bld)
			log.AppendMessage("Pawn moved out.")
			i-- 
		}
	}
}
