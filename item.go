package main 

type itemType uint8 

const (
	ITEMTYPE_WEAPON itemType = iota 
	ITEMTYPE_OTHER 
)

type item struct {
	name string // TODO: replace with static tables 
	weaponData *weapon 
}

func (i *item) getName() string {
	return i.name
}

func (i *item) getType() itemType {
	if i.weaponData != nil {
		return ITEMTYPE_WEAPON
	}
	return ITEMTYPE_OTHER
}
