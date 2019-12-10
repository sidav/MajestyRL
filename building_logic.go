package main 

type buildingLogic struct {}

func (bl *buildingLogic) doTurn(bld *pawn) {
	bld.asBuilding.recalculateCurrValues()
}
