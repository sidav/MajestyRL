package main 

type unitLogic struct {}

// returns true if the unit wants to leave the building 
func (ul *unitLogic) actFromInsideBuilding(u *pawn) bool {
	return true 
}
