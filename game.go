package main

import (
	"github.com/sidav/golibrl/random/additive_random"
	// "strconv"
	// "time"
	"MajestyRL/log"
)

const (
	TICKS_PER_TURN    = 10
	CLEANUP_BIDS_EACH = 500 // ticks
)

var (
	BLOGIC = &buildingLogic{}
	ULOGIC = &unitLogic{}

	rnd               = additive_random.FibRandom{}
	GAME_IS_RUNNING   = true
	IS_PAUSED         = false
	LOG               *log.GameLog
	RENDERER          rendererStruct
	PLAYER_CONTROLLER playerController
	CURRENT_TICK      = 1
	CURRENT_MAP       *gameMap
	CHEAT_IGNORE_FOW  bool
	DEBUG_OUTPUT      = true
	LOG_HEIGHT        = 8
)

func getCurrentTurn() int {
	return CURRENT_TICK/TICKS_PER_TURN + 1
}

func initGame() {
	rnd.InitDefault()
	LOG = &log.GameLog{}
	LOG.Init(LOG_HEIGHT)
	PLAYER_CONTROLLER.init()

	// load test mission
	initTestMission()
	LOG.AppendMessage("Test mission initialized.")
}

func startGameLoop() {
	for !PLAYER_CONTROLLER.exit { // main game loop

		if CURRENT_TICK%TICKS_PER_TURN == 0 {
			for _, currFaction := range CURRENT_MAP.factions {
				PLAYER_CONTROLLER.controlAsFaction(currFaction)
			}
		}

		if CURRENT_TICK%CLEANUP_BIDS_EACH == 0 {
			CURRENT_MAP.cleanupBids()
		}

		for _, curpawn := range CURRENT_MAP.pawns {
			if curpawn.isBuilding() {
				if CURRENT_TICK%TICKS_PER_TURN == 0 {
					BLOGIC.act(curpawn)
				}
				continue
			}
			if curpawn.isTimeToAct() {
				ULOGIC.decideNewIntent(curpawn)
				curpawn.act()
			}
		}

		CURRENT_TICK++
	}
}
