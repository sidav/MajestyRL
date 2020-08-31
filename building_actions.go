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
	if CURRENT_TICK % GENERATE_TAXES_EACH == 0 {
		bld.asBuilding.accumulatedGoldAmount += bld.asBuilding.getStaticData().taxGoldGeneration
	}
	bl.generatePawns(bld)
	bl.actForEachPawnInside(bld)
}

func (bl *buildingLogic) generatePawn(bld *pawn, code string) {
	newPawn := createUnitAtCoords(code, bld.x, bld.y, bld.faction)
	bld.AddAndRegisterNewPawn(newPawn)
	bld.asBuilding.recalculateCurrResidents()
}

func (bl *buildingLogic) generatePawns(bld *pawn) {
	bstatic := getBuildingStaticDataFromTable(bld.asBuilding.code)
	bld.asBuilding.recalculateCurrResidents()
	for i, code := range bstatic.housing_unittypes {
		if bstatic.housing_respawn_period[i] == 0 {
			continue
		}
		if CURRENT_TICK % (bstatic.housing_respawn_period[i] * TICKS_PER_TURN) == 0 &&
			bld.asBuilding.canAffordNewResident(code) {

			log.AppendMessage(fmt.Sprintf("Creating %s (%d/%d) at turn %d", code,
				bld.asBuilding.currentResidents[code], bstatic.housing_max_residents[i], CURRENT_TICK))

			bl.generatePawn(bld, code)
		}
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
