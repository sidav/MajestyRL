package main

import (
	// cmenu "github.com/sidav/golibrl/console_menu"
	cw "github.com/sidav/golibrl/console"
	// "strconv"
	// "time"
	"Majesty/log"
	"fmt"
)

var (
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
	fmt.Println("I'm working!")
	cw.Init_console("M@JESTY", cw.TCellRenderer)
	defer cw.Close_console()

	LOG = &log.GameLog{}
	LOG.Init(LOG_HEIGHT)
	// load test mission
	initTestMission()
	LOG.AppendMessage("Test mission initialized.")

	startGameLoop()
}
