/*
  A game for Ludum Dare 56
  Copyright (C) 2024 Loïg Jezequel

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>
*/

package main

type level struct {
	width, height            int
	area                     [][]int
	froots                   map[int]map[int]*froot
	badGuysSetup             []badGuy
	badGuys                  []badGuy
	currentGuy               int
	positions                [][]tilePosition
	iSelected, jSelected     int
	iCharacters, jCharacters int
	iGoal, jGoal             int
	numFood                  int
	numCharacters            int
	alreadySeen              bool
}

const (
	areaTypeNone = iota
	areaTypeGrass
	areaTypeWater
	areaTypeCastle
	areaTypeVillage
)

const (
	frootTypeRaspberry = iota
)

type tilePosition struct {
	topLeftX, topLeftY int
	centerX, centerY   int
}

type froot struct {
	kind  int
	eaten bool
}

func (l *level) getPositions() {
	if len(l.positions) == 0 {
		// only if not already computed
		generalXShift := gPlayAreaX + gPlayAreaWidth/2 - (l.width*gTileSize)/2
		generalYShift := gPlayAreaY + gPlayAreaHeight/2 - (l.height*gTileSize)/2 - (l.height*gTileYShiftPerLine)/2

		l.positions = make([][]tilePosition, l.height+2)
		for i := 0; i < len(l.positions); i++ {
			y := i - 1
			l.positions[i] = make([]tilePosition, l.width+2)
			for j := 0; j < len(l.positions[i]); j++ {
				x := j - 1
				l.positions[i][j] = tilePosition{
					topLeftX: generalXShift + x*gTileSize + x*gTileXShiftPerStep + y*gTileXShiftPerLine,
					topLeftY: generalYShift + y*gTileSize + x*gTileYShiftPerStep + y*gTileYShiftPerLine,
				}
				l.positions[i][j].centerX = l.positions[i][j].topLeftX + gTileXCenterShift
				l.positions[i][j].centerY = l.positions[i][j].topLeftY + gTileYCenterShift
			}
		}
	}
}

func (l *level) setSelected(mouseX, mouseY int) (changed bool) {

	oldi, oldj := l.iSelected, l.jSelected

	d := distance(float64(l.positions[l.iSelected][l.jSelected].centerX), float64(l.positions[l.iSelected][l.jSelected].centerY), float64(mouseX), float64(mouseY))

	for i, line := range l.positions {
		for j, pos := range line {
			dd := distance(float64(pos.centerX), float64(pos.centerY), float64(mouseX), float64(mouseY))
			if dd < d {
				d = dd
				l.iSelected = i
				l.jSelected = j
			}
		}
	}

	changed = (oldi != l.iSelected || oldj != l.jSelected) &&
		l.iSelected > 0 && l.iSelected <= len(l.area) &&
		l.jSelected > 0 && l.jSelected <= len(l.area[l.iSelected-1]) &&
		l.area[l.iSelected-1][l.jSelected-1] != areaTypeNone

	return
}

func (l *level) updateCharactersPosition() (hasMoved bool) {

	switch l.iSelected {
	case l.iCharacters - 1:
		hasMoved = l.jSelected == l.jCharacters || l.jSelected == l.jCharacters+1
	case l.iCharacters:
		hasMoved = l.jSelected == l.jCharacters-1 || l.jSelected == l.jCharacters+1
	case l.iCharacters + 1:
		hasMoved = l.jSelected == l.jCharacters-1 || l.jSelected == l.jCharacters
	}

	hasMoved = hasMoved &&
		l.iSelected-1 >= 0 && l.iSelected-1 < len(l.area) &&
		l.jSelected-1 >= 0 && l.jSelected-1 < len(l.area[l.iSelected-1]) &&
		l.area[l.iSelected-1][l.jSelected-1] != areaTypeWater &&
		l.area[l.iSelected-1][l.jSelected-1] != areaTypeNone

	if hasMoved {
		l.iCharacters = l.iSelected
		l.jCharacters = l.jSelected
	}

	return
}

func (l *level) updateFood() (hasEaten bool) {

	froots, exists := l.froots[l.iCharacters-1]
	if exists {
		froot, isHere := froots[l.jCharacters-1]
		if isHere && !froot.eaten {
			hasEaten = true
			froot.eaten = true
			switch froot.kind {
			case frootTypeRaspberry:
				l.numFood += 500
			}
		}
	}

	if l.numFood > l.numCharacters {
		l.numFood -= l.numCharacters
	} else {
		l.numFood = 0
		l.numCharacters -= gNumDeathsForMissingFood
	}

	if l.numCharacters < 0 {
		l.numCharacters = 0
	}

	return

}

func (l *level) isLost() (lost bool) {
	if l.numCharacters <= 0 {
		l.numCharacters = 0
		lost = true
	}
	return
}

func (l level) isCompleted() bool {
	return l.iCharacters == l.iGoal && l.jCharacters == l.jGoal
}
