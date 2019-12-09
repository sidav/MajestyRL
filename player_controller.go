package main 

import cw "github.com/sidav/golibrl/console"

type playerController struct {
	exit bool 
}

func (pc *playerController) controlAsFaction(f *faction) {
	RENDERER.renderScreen(f)
	keyPressed := cw.ReadKeyAsync()
	if keyPressed == "ESCAPE" {
		pc.exit = true 
	}
}
