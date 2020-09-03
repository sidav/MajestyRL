package main

import (
	"MajestyRL/game_log"
	"github.com/sidav/golibrl/random/additive_random"
	// "strconv"
	"time"
)

const (
	TICKS_PER_TURN        = 10
	CLEANUP_BIDS_EACH     = 500 // ticks
	READJUST_ECONOMY_EACH = 500 // ticks
	GROW_FORESTS_EACH     = 100
)

var (
	BLOGIC = &buildingLogic{}
	ULOGIC = &unitLogic{}

	rnd               = additive_random.FibRandom{}
	GAME_IS_RUNNING   = true
	IS_PAUSED         = false
	RESOURCE_HAULING  = true
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

		if CURRENT_TICK%GROW_FORESTS_EACH == 0 {
			CURRENT_MAP.growForests()
		}

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
				currFaction.economy.adjustResourcesToMax()
				PLAYER_CONTROLLER.controlAsFaction(currFaction)
				if CURRENT_TICK%READJUST_ECONOMY_EACH == 0 {
					currFaction.economy.resetMaxResources()
				}
			}
		}

		if CURRENT_TICK%CLEANUP_BIDS_EACH == 0 {
			CURRENT_MAP.cleanupBids()
		}

		for _, curpawn := range CURRENT_MAP.pawns {
			if curpawn.isBuilding() {
				if !curpawn.asBuilding.isUnderConstruction() {
					// readjust max resources for faction
					if CURRENT_TICK%READJUST_ECONOMY_EACH == 0 && curpawn.faction != nil {
						for rtype, rvalue := range curpawn.asBuilding.getStaticData().resourceStorage {
							curpawn.faction.economy.maxResources[rtype] += rvalue
						}
					}

					if CURRENT_TICK%TICKS_PER_TURN == 0 {
						curpawn.setFactionTechAllowance() // TODO: call this less frequently.
						BLOGIC.act(curpawn)
					}
				}
			} else {
				if curpawn.isTimeToAct() {
					ULOGIC.decideNewIntent(curpawn)
					curpawn.act()
				}
			}
			if curpawn.hitpoints <= 0 {
				if curpawn.isUnit() {
					if curpawn.asUnit.intent != nil && curpawn.asUnit.intent.isDispatchedFromBid() == true {
						curpawn.asUnit.intent.sourceBid.drop()
					}
				}
				CURRENT_MAP.removePawn(curpawn)
				// TODO: drop items
			}
		}

		CURRENT_TICK++
	}
}
