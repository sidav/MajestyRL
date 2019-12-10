package main

import (
	// cmenu "github.com/sidav/golibrl/console_menu"
	cw "github.com/sidav/golibrl/console"
	// "strconv"
	// "time"
	"fmt"
)

var (
	GAME_IS_RUNNING   = true
	LOG               *log
	RENDERER          rendererStruct
	PLAYER_CONTROLLER playerController
	CURRENT_TICK      = 1
	CURRENT_MAP       *gameMap
	CHEAT_IGNORE_FOW  bool
	DEBUG_OUTPUT      bool
)

func debug_write(text string) {
	if DEBUG_OUTPUT {
		LOG.appendMessage("DEBUG: " + text)
	}
}

func main() {
	fmt.Println("I'm working!")
	cw.Init_console("M@JESTY", cw.TCellRenderer)
	defer cw.Close_console()

	LOG = &log{}
	// load test mission
	initTestMission()

	startGameLoop()
}
