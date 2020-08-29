package main

import "github.com/sidav/golibrl/random"

type weapon struct {
	meleeDamageDice *random.Dice
	attackTime  int
}

func (w *weapon) rollMeleeDamageDice() int {
	return w.meleeDamageDice.Roll(&rnd)
}
