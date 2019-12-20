package main

import (
	cw "github.com/sidav/golibrl/console"
	// cmenu "github.com/sidav/golibrl/console_menu"
)

var (

)

func debug_write(text string) {
	if DEBUG_OUTPUT {
		log.AppendMessage("DEBUG: " + text)
	}
}

func main() {
	cw.Init_console("M@JESTY", cw.TCellRenderer)
	defer cw.Close_console()

	initGame()
	startGameLoop()
}
