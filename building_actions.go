package main

import "fmt"

const (
	SPAWN_AUTO_BIDS_EACH = 2000 // ticks
)

type buildingLogic struct {
}

func (bl *buildingLogic) act(bld *pawn) {
	if bld.asBuilding.isUnderConstruction() {
		return
	}
	if CURRENT_TICK%GENERATE_TAXES_EACH == 0 {
		bld.asBuilding.accumulatedGoldAmount += bld.asBuilding.getStaticData().taxGoldGeneration
	}
	bl.generatePawns(bld)
	bl.actForEachPawnInside(bld)
	bl.spawnAutoBids(bld)
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
		if CURRENT_TICK%(bstatic.housing_respawn_period[i]*TICKS_PER_TURN) == 0 &&
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

func (bl *buildingLogic) spawnAutoBids(bld *pawn) {
	if CURRENT_TICK % SPAWN_AUTO_BIDS_EACH == 0 {
		static := bld.asBuilding.getStaticData()
		radius := static.autoBidRadius
		cx, cy := bld.getCenter()
		switch static.autoBidType {
		case "GROW_FOREST":
			x, y := rnd.RandInRange(cx-radius, cx+radius), rnd.RandInRange(cy-radius, cy+radius)
			if CURRENT_MAP.getResourcesAtCoords(x, y) == nil && CURRENT_MAP.getPawnAtCoordinates(x, y) == nil {
				newbid := &bid{intent_type_for_this_bid: INTENT_GROW_FOREST, maxTaken: 1, x: x, y: y, factionCreatedBid: bld.faction}
				CURRENT_MAP.addBid(newbid)
			}
		case "MINE_FOREST":
			// x, y := rnd.RandInRange(cx-radius, cx+radius), rnd.RandInRange(cy-radius, cy+radius)
			for x := cx-radius; x <= cx+radius; x++ {
				for y := cy-radius; y <= cy+radius; y++ {
					res := CURRENT_MAP.getResourcesAtCoords(x, y)
					if res != nil && res.resType == RESTYPE_WOOD && !res.grows {
						newbid := &bid{intent_type_for_this_bid: INTENT_MINE, maxTaken: 1, x: x, y: y, factionCreatedBid: bld.faction}
						CURRENT_MAP.addBid(newbid)
					}
				}
			}
		}
	}
}
