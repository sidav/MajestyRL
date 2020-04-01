package main

func createWeaponByCode(code string) *item {
	switch code {
	case "HALBERD":
		return &item{
			name: "Halberd",
			weaponData: &weapon{damage: 5},
		}
	}

	return &item{
		name: "UNKNOWN WEAPON",
		weaponData: &weapon{damage: 5},
	}
}
