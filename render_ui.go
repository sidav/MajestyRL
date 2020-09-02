package main

import (
	"fmt"
	cw "github.com/sidav/golibrl/console"
	cmenu "github.com/sidav/golibrl/console_menu"
)

func (r *rendererStruct) renderUI() {
	r.renderUIOutline()
	r.renderFactionStats()
	r.renderInfoOnCursor()
}

func (r *rendererStruct) renderUIOutline() {
	if IS_PAUSED {
		cw.SetBgColor(r.currentFactionSeeingTheScreen.getFactionColor())
	} else {
		cw.SetFgColor(r.currentFactionSeeingTheScreen.getFactionColor())
	}
	cw.SetFgColor(r.currentFactionSeeingTheScreen.getFactionColor())
	for y := 0; y < VIEWPORT_H; y++ {
		cw.PutChar('|', VIEWPORT_W, y)
	}
	for x := 0; x < CONSOLE_W; x++ {
		cw.PutChar('-', x, VIEWPORT_H)
	}
	cw.PutChar('+', VIEWPORT_W, VIEWPORT_H)
	cw.SetBgColor(cw.BLACK)
}

func (r *rendererStruct) renderFactionStats() {
	f := r.currentFactionSeeingTheScreen
	eco := f.economy
	statsx := VIEWPORT_W + 1

	// fr, fg, fb := getFactionRGB(f.factionNumber)
	// cw.SetFgColorRGB(fr, fg, fb)
	if IS_PAUSED {
		cw.SetFgColor(f.getFactionColor())
		cw.PutString(f.name+", ", statsx, 0)
		cw.SetFgColor(cw.YELLOW)
		cw.PutString(fmt.Sprintf("turn %d (PAUSED)", getCurrentTurn()), statsx+len(f.name)+2, 0)
	} else {
		cw.SetFgColor(f.getFactionColor())
		cw.PutString(fmt.Sprintf("%s: turn %d", f.name, getCurrentTurn()), statsx, 0)
	}
	cw.SetFgColor(cw.YELLOW)
	r.renderStatusbar("GOLD:", eco.currentResources.amount[RESTYPE_GOLD], eco.maxResources[RESTYPE_GOLD], statsx, 2, SIDEBAR_W-3, cw.YELLOW, cw.DARK_YELLOW, false)
	cw.SetFgColor(cw.GREEN)
	r.renderStatusbar("WOOD:", eco.currentResources.amount[RESTYPE_WOOD], eco.maxResources[RESTYPE_WOOD], statsx, 3, SIDEBAR_W-3, cw.YELLOW, cw.DARK_YELLOW, false)
}

func (r *rendererStruct) renderStatusbar(name string, curvalue, maxvalue, x, y, width, fillColor, emptyColor int, hideNumericValues bool) {
	if !hideNumericValues {
		name = fmt.Sprintf("%s %d/%d", name, curvalue, maxvalue)
	}
	cw.PutString(name, x, y)
	barWidth := width - len(name)
	var filledCells int
	if maxvalue > 0 {
		filledCells = barWidth * curvalue / maxvalue
	} else {
		filledCells = 0
	}
	barStartX := x + len(name) + 1
	for i := 0; i < barWidth; i++ {
		if i < filledCells {
			cw.SetFgColor(fillColor)
			cw.PutChar('=', i+barStartX, y)
		} else {
			cw.SetFgColor(emptyColor)
			cw.PutChar('-', i+barStartX, y)
		}
	}
}

func (r *rendererStruct) renderInfoOnCursor() {

	title := "Unidentified Object"
	color := 2
	details := make([]string, 0)
	// var res *pawnResourceInformation
	sp := r.currentFactionSeeingTheScreen.cursor.snappedPawn

	if sp != nil {
		color = sp.faction.getFactionColor()
		if r.currentFactionSeeingTheScreen.areCoordsInSight(sp.x, sp.y) {
			title = sp.getName()
			details = append(details, fmt.Sprintf("HP: %d/%d", sp.hitpoints, sp.getMaxHitpoints()))

			// enemy pawn 
			if sp.faction != r.currentFactionSeeingTheScreen && !DEBUG_OUTPUT {
				if sp.isBuilding() {
					details = append(details, "(Enemy building)")
				} else {
					details = append(details, "(Enemy unit)")
				}

			} else { // our pawn 
				if sp.isBuilding() {
					static := getBuildingStaticDataFromTable(sp.asBuilding.code)

					if sp.asBuilding.isUnderConstruction() {
						curr, max, perc := sp.asBuilding.asBeingConstructed.getCompletionValues()
						details = append(details, fmt.Sprintf("Under construction: %d/%d (%d%%)", curr, max, perc))
						if RESOURCE_HAULING {
							for rtype, ramount := range sp.asBuilding.asBeingConstructed.resourcesBroughtToConstruction.amount {
								cost := static.cost.amount[rtype]
								details = append(details, fmt.Sprintf("%s: %d/%d", getResourceName(rtype), ramount, cost))
							}
						}
					} else {
						if sp.asBuilding.accumulatedGoldAmount > 0 {
							details = append(details, fmt.Sprintf("Tax: %d", sp.asBuilding.accumulatedGoldAmount))
						}
						for i, code := range static.housing_unittypes {
							details = append(details, fmt.Sprintf("%s: %d/%d",
								getUnitStaticDataFromTable(code).name, sp.asBuilding.currentResidents[code], static.housing_max_residents[i]))
						}
					}
				}

				if sp.isUnit() {
					if sp.asUnit.carriedResourceAmount > 0 {
						details = append(details, fmt.Sprintf("Carries %d %s", sp.asUnit.carriedResourceAmount, getResourceName(sp.asUnit.carriedResourceType)))
					}
					details = append(details, sp.asUnit.getCurrentIntentDescription())
				}
			}
		}
	} else { //snapped no pawn, render bid info if any
		sb := r.currentFactionSeeingTheScreen.cursor.snappedBid
		if sb != nil {
			if sb.targetPawn != nil && sb.targetPawn.asBuilding.isUnderConstruction() {
				static := sb.targetPawn.asBuilding.getStaticData()
				if RESOURCE_HAULING {
					for rtype, ramount := range sb.targetPawn.asBuilding.asBeingConstructed.resourcesBroughtToConstruction.amount {
						cost := static.cost.amount[rtype]
						details = append(details, fmt.Sprintf("%s: %d/%d", getResourceName(rtype), ramount, cost))
					}
				}
			}
		}
	}
	if len(details) > 0 {
		cmenu.DrawSidebarInfoMenu(title, color, SIDEBAR_X, SIDEBAR_FLOOR_2, SIDEBAR_W, details)
	}
}
