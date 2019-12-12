package main 

import "fmt"

const (
	REGENERATE_WORKERS_EACH = 100 // ticks 
)

type buildingLogic struct {
}

func (bl *buildingLogic) doTurn(bld *pawn) {
	if bld.asBuilding.isUnderConstruction() {
		return 
	}
	bl.generatePawns(bld)
	bl.actForEachPawnInside(bld)
}

func (bl *buildingLogic) generatePawns(bld *pawn) {
	bstatic := staticBuildingDataTable[bld.asBuilding.code]
	bld.asBuilding.recalculateCurrValues()
	// LOG.AppendMessage(fmt.Sprintf("Peasant (%d/%d)", bld.asBuilding.currWorkers, bstatic.maxWorkers))
	if bld.asBuilding.currWorkers < bstatic.maxWorkers && CURRENT_TICK % REGENERATE_WORKERS_EACH == 0 {
		// spawn one more worker 
		newWorker := createUnitAtCoords("PEASANT", bld.x, bld.y, bld.faction)
		bld.AddAndRegisterNewPawn(newWorker)
		bld.asBuilding.recalculateCurrValues()
		LOG.AppendMessage(fmt.Sprintf("Peasant created (%d/%d) at turn %d", bld.asBuilding.currWorkers, bstatic.maxWorkers, CURRENT_TICK))
	} 
}

func (bl *buildingLogic) actForEachPawnInside(bld *pawn) {
	// bstatic := staticBuildingDataTable[bld.asBuilding.code]
	for _, p := range bld.asBuilding.pawnsInside {
		ULOGIC.decideNewIntent(p)
		if ULOGIC.wantsToLeaveBuilding(p) {
			// remove unit from unitsInside and place it outside the building 
			bld.asBuilding.removePawnFromInside(p)
			CURRENT_MAP.addPawn(p)
			CURRENT_MAP.placePawnNearPawn(p, bld)
			LOG.AppendMessage("Pawn moved out.")
		}
	}
}
