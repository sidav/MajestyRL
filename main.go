package main

import (
	geometry "github.com/sidav/golibrl/geometry"
	cmenu "github.com/sidav/golibrl/console_menu"
	cw "github.com/sidav/golibrl/console"
	"strconv"
	"time"
)

func areCoordsValid(x, y int) bool {
	return geometry.AreCoordsInRect(x, y, 0, 0, mapW, mapH)
}

var (
	GAME_IS_RUNNING                     = true
	log                                 *LOG
	CURRENT_TICK                        = 1
	CURRENT_MAP                         *gameMap
	CHEAT_IGNORE_FOW                    bool
	DEBUG_OUTPUT                        bool
)

func getCurrentTurn() int {
	return CURRENT_TICK/10 + 1
}

func debug_write(text string) {
	if DEBUG_OUTPUT {
		log.appendMessage("DEBUG: " + text)
	}
}

func main() {
	cw.Init_console("M@JESTY", cw.TCellRenderer)
	defer cw.Close_console()

	log = &LOG{}
	
}
