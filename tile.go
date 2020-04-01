package main

import cw "github.com/sidav/golibrl/console"

const (
	TTYPE_GRASS byte = iota
	TTYPE_WATER
	TTYPE_UNKNOWN
)

var tileAppearances = map[byte]*ccell{
	TTYPE_GRASS: &ccell{char: '.', color: cw.GREEN},
	TTYPE_WATER: &ccell{char: '~', color: cw.DARK_BLUE},

	TTYPE_UNKNOWN: &ccell{char: '?', color: cw.MAGENTA},
}

type tile struct {
	tiletype  byte
	resources *tileResource
}

func (t *tile) getAppearance() *ccell {
	if t.resources != nil {
		return t.resources.getAppearance()
	}
	return tileAppearances[t.tiletype]
}

func (t *tile) isPassable() bool {
	return t.tiletype == TTYPE_GRASS
}
