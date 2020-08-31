package main

import (
	"github.com/sidav/golibrl/random/additive_random"
	// "strconv"
	"time"
	"MajestyRL/game_log"
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
	log               *game_log.GameLog
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
	log = &game_log.GameLog{}
	log.Init(LOG_HEIGHT)
	PLAYER_CONTROLLER.init()

	// load test mission
	initTestMission()
	log.AppendMessage("Test mission initialized.")
}

func startGameLoop() {
	start := time.Now()
	for !PLAYER_CONTROLLER.exit { // main game loop

		if CURRENT_TICK%TICKS_PER_TURN == 0 {
			currentGameLoopTime = time.Since(start) / time.Nanosecond
			totalGameLoopTimes += currentGameLoopTime
			totalGameLoops += 1 
			if totalGameLoops == 10 {
				averageGameLoopTime = totalGameLoopTimes / 10
				totalGameLoopTimes = 0 
				totalGameLoops = 0 
			}
			start = time.Now()

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
					curpawn.setFactionTechAllowance() // TODO: call this less frequently. 
					BLOGIC.act(curpawn)
				}
			} else {
				if curpawn.isTimeToAct() {
					ULOGIC.decideNewIntent(curpawn)
					curpawn.act()
				}
			}
			if curpawn.hitpoints <= 0 {
				CURRENT_MAP.removePawn(curpawn)
				// TODO: drop items
			}
		}

		CURRENT_TICK++
	}
}
