package main

import "github.com/sidav/golibrl/random"

func createWeaponByCode(code string) *item {
	switch code {
	case "HALBERD":
		return &item{
			name: "Halberd",
			weaponData: &weapon{meleeDamageDice: random.NewDice(1, 3, 1), attackTime: TICKS_PER_TURN*2},
		}
	}

	return &item{
		name: "UNKNOWN WEAPON",
		weaponData: &weapon{meleeDamageDice: random.NewDice(0, 1, 1), attackTime: TICKS_PER_TURN},
	}
}
