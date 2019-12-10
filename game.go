package main

const (
	TICKS_PER_TURN = 10
)

var (
	BLOGIC = &buildingLogic{}
)

func getCurrentTurn() int {
	return CURRENT_TICK/10 + 1
}

func startGameLoop() {
	for !PLAYER_CONTROLLER.exit { // main game loop

		if CURRENT_TICK%TICKS_PER_TURN == 0 {
			for _, currFaction := range CURRENT_MAP.factions {
				PLAYER_CONTROLLER.controlAsFaction(currFaction)
			}
		}

		for _, curpawn := range CURRENT_MAP.pawns {
			if curpawn.isBuilding() {
				BLOGIC.doTurn(curpawn)
			}
		}

		CURRENT_TICK++
	}
}
