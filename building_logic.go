package main 

import "fmt"

const (
	REGENERATE_WORKERS_EACH = 1000 // ticks 
)

type buildingLogic struct {
}

func (bl *buildingLogic) doTurn(bld *pawn) {
	bl.generatePawns(bld)
}

func (bl *buildingLogic) generatePawns(bld *pawn) {
	bstatic := staticBuildingDataTable[bld.asBuilding.code]
	bld.asBuilding.recalculateCurrValues()
	LOG.AppendMessage(fmt.Sprintf("Peasant (%d/%d)", bld.asBuilding.currWorkers, bstatic.maxWorkers))
	if bld.asBuilding.currWorkers < bstatic.maxWorkers && CURRENT_TICK % REGENERATE_WORKERS_EACH == 0 {
		// spawn one more worker 
		newWorker := createUnitAtCoords("PEASANT", bld.x, bld.y, bld.faction)
		bld.asBuilding.AddAndRegisterNewPawn(newWorker)
		LOG.AppendMessage(fmt.Sprintf("Peasant created (%d/%d)", bld.asBuilding.currWorkers, bstatic.maxWorkers))
	} 
}
