package main

// cw "github.com/sidav/golibrl/console"

func (g *gameMap) init() {
	g.pawns = make([]*pawn, 0)
	g.factions = make([]*faction, 0)
	// initMapForMission(g, 1)
}

func (g *gameMap) initTileMap(strmap *[]string) {
	mapH = len(*strmap)
	mapW = len((*strmap)[0])
	g.tileMap = make([][]*tile, mapW)
	for i := range g.tileMap {
		g.tileMap[i] = make([]*tile, mapH)
	}

	for y, str := range *strmap {
		for x, chr := range str {
			if chr == '*' {
				g.tileMap[x][y] = &tile{tiletype: TTYPE_GRASS, resources: &tileResource {
					amount: 1000, resType: RESTYPE_GOLD,
				}}
			} else if chr == '%' {
				g.tileMap[x][y] = &tile{tiletype: TTYPE_GRASS, resources: &tileResource {
					amount: 100, resType: RESTYPE_WOOD,
				}}
			} else {
				g.tileMap[x][y] = &tile{tiletype: mapinit_getTiletypeByChar(chr)}
			}
		}
	}
}

func mapinit_getTiletypeByChar(char rune) byte {
	switch char {
	case '.':
		return TTYPE_GRASS
	case '~':
		return TTYPE_WATER
	}
	return TTYPE_UNKNOWN
}
