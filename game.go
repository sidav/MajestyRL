package main

import "time"

const (
	TICKS_PER_TURN = 10
)

var (
	BLOGIC = &buildingLogic{}
	ULOGIC = &unitLogic{}
)

func getCurrentTurn() int {
	return CURRENT_TICK/TICKS_PER_TURN + 1
}

func startGameLoop() {
	for !PLAYER_CONTROLLER.exit { // main game loop

		if CURRENT_TICK%TICKS_PER_TURN == 0 {
			for _, currFaction := range CURRENT_MAP.factions {
				PLAYER_CONTROLLER.controlAsFaction(currFaction)
			}
			time.Sleep(15 * time.Millisecond)
		}

		for _, curpawn := range CURRENT_MAP.pawns {
			if curpawn.isTimeToAct() {
				if curpawn.isBuilding() {
					BLOGIC.doTurn(curpawn)
					continue
				}
				ULOGIC.decideNewIntent(curpawn)
				curpawn.act()
			}
		}

		CURRENT_TICK++
	}
}
