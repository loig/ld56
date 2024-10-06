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

import "github.com/loig/ld56/assets"

type game struct {
	state          int
	step           int
	inAnimation    bool
	animationFrame int
	animationStep  int
	levelNum       int
	currentLevel   level
	recordAttack   step
	footSteps      int
	soundManager   assets.SoundManager
}

const (
	stateInLevel = iota
	stateTitle
	stateGameover
	stateCongrats
)

const (
	stepPlayerTurn = iota
	stepBadGuysMoving
	stepCombat
	stepBadGuysAttacking
	stepAttackEffect
	stepCheckStatus
	stepNumber
	stepLevelStart
	stepLevelEnd
)

func (g *game) setFirstLevel(n int) {

	g.setLevel(n)

	g.currentLevel.numFood = gInitNumFood
	g.currentLevel.numCharacters = gInitNumCharacters

	g.currentLevel.area[2][2] = areaTypeVillage
}

func (g *game) setLevel(n int) {
	g.levelNum++

	food := g.currentLevel.numFood
	characters := g.currentLevel.numCharacters

	g.currentLevel = nextLevel(n)
	g.currentLevel.getPositions()

	for _, froots := range g.currentLevel.froots {
		for _, froot := range froots {
			froot.eaten = false
		}
	}

	if len(g.currentLevel.badGuys) == 0 {
		g.currentLevel.badGuys = make([]badGuy, len(g.currentLevel.badGuysSetup))
	}

	copy(g.currentLevel.badGuys, g.currentLevel.badGuysSetup)

	g.currentLevel.numFood = food
	g.currentLevel.numCharacters = characters
}
