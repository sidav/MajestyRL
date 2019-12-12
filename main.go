package main

import (
	// cmenu "github.com/sidav/golibrl/console_menu"
	cw "github.com/sidav/golibrl/console"
	"github.com/sidav/golibrl/random/additive_random"
	// "strconv"
	// "time"
	"Majesty/log"
)

var (
	rnd               = additive_random.FibRandom{}
	GAME_IS_RUNNING   = true
	LOG               *log.GameLog
	RENDERER          rendererStruct
	PLAYER_CONTROLLER playerController
	CURRENT_TICK      = 1
	CURRENT_MAP       *gameMap
	CHEAT_IGNORE_FOW  bool
	DEBUG_OUTPUT      bool
	LOG_HEIGHT        = 5
)

func debug_write(text string) {
	if DEBUG_OUTPUT {
		LOG.AppendMessage("DEBUG: " + text)
	}
}

func main() {
	cw.Init_console("M@JESTY", cw.TCellRenderer)
	defer cw.Close_console()

	rnd.InitDefault()
	LOG = &log.GameLog{}
	LOG.Init(LOG_HEIGHT)
	// load test mission
	initTestMission()
	LOG.AppendMessage("Test mission initialized.")

	startGameLoop()
}
