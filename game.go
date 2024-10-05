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

type game struct {
	state          int
	step           int
	inAnimation    bool
	animationFrame int
	animationStep  int
	levelNum       int
	currentLevel   level
	recordAttack   step
}

const (
	stateInLevel = iota
)

const (
	stepPlayerTurn = iota
	stepBadGuysMoving
	stepCombat
	stepBadGuysAttacking
	stepAttackEffect
	stepCheckStatus
	stepNumber
)

func (g *game) setFirstLevel(l level) {

	g.setLevel(l)

	g.levelNum = 1

	g.currentLevel.numFood = gInitNumFood
	g.currentLevel.numCharacters = gInitNumCharacters
}

func (g *game) setLevel(l level) {
	g.levelNum++

	food := g.currentLevel.numFood
	characters := g.currentLevel.numCharacters

	g.currentLevel = l
	g.currentLevel.getPositions()

	for _, froots := range l.froots {
		for _, froot := range froots {
			froot.eaten = false
		}
	}

	if len(g.currentLevel.badGuys) == 0 {
		g.currentLevel.badGuys = make([]badGuy, len(l.badGuysSetup))
	}

	copy(g.currentLevel.badGuys, l.badGuysSetup)

	g.currentLevel.numFood = food
	g.currentLevel.numCharacters = characters
}

var testLevel level = level{
	width:  11,
	height: 9,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeWater, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeWater, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeWater, areaTypeWater, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeGrass, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeWater, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeNone, areaTypeGrass, areaTypeGrass},
	},
	froots: map[int]map[int]*froot{
		8: map[int]*froot{3: &froot{kind: frootTypeRaspberry}},
	},
	badGuysSetup: []badGuy{
		{x: 7, y: 0, kind: badGuyBee},
		{x: 5, y: 3, kind: badGuyLadyBug},
		{x: 6, y: 6, kind: badGuyDungBeetle},
	},
	iCharacters: 9, jCharacters: 1,
	iGoal: 1, jGoal: 7,
}
