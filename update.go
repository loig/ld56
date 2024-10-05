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

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) Update() error {

	if g.inAnimation {
		g.animationFrame++

		if g.animationFrame >= gAnimationSet[g.step].numFrames[g.animationStep%len(gAnimationSet[g.step].numFrames)] {
			g.animationFrame = 0
			g.animationStep++
			if g.animationStep >= gAnimationSet[g.step].numSteps {
				g.animationStep = 0
				g.inAnimation = false
				g.step = gAnimationSet[g.step].nextStep
			} else {
				updateAnimationDraw(g.step)
			}
		}

		return nil
	}

	mouseX, mouseY := ebiten.CursorPosition()

	g.currentLevel.setSelected(mouseX, mouseY)

	switch g.step {
	case stepPlayerTurn:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			hasMoved := g.currentLevel.updateCharactersPosition()
			if hasMoved {
				g.currentLevel.updateFood()
				g.step = stepBadGuysMoving
			}
		}
	case stepBadGuysMoving:
		hasMoved, inCombat, done := g.currentLevel.moveBadGuys()
		if done {
			g.step = stepBadGuysAttacking
		} else if hasMoved || inCombat {
			g.inAnimation = true
			if inCombat {
				g.step = stepCombat
				gAnimationSet[stepCombat].drawX = g.currentLevel.positions[g.currentLevel.iCharacters][g.currentLevel.jCharacters].topLeftX
				gAnimationSet[stepCombat].drawY = g.currentLevel.positions[g.currentLevel.iCharacters][g.currentLevel.jCharacters].topLeftY
			}
		}
	case stepBadGuysAttacking:
		dist, start, direction, attackPoint, done := g.currentLevel.makeBadGuysAttack()
		g.recordAttack = attackPoint
		if done {
			g.step = stepCheckStatus
		} else if dist > 0 {
			g.inAnimation = true
			gAnimationSet[stepBadGuysAttacking].numSteps = dist
			gAnimationSet[stepBadGuysAttacking].drawX = g.currentLevel.positions[start.y+1+direction.y][start.x+1+direction.x].topLeftX
			gAnimationSet[stepBadGuysAttacking].drawY = g.currentLevel.positions[start.y+1+direction.y][start.x+1+direction.x].topLeftY
			gAnimationSet[stepBadGuysAttacking].stepX = g.currentLevel.positions[start.y+1+2*direction.y][start.x+1+2*direction.x].topLeftX - g.currentLevel.positions[start.y+1+direction.y][start.x+1+direction.x].topLeftX
			gAnimationSet[stepBadGuysAttacking].stepY = g.currentLevel.positions[start.y+1+2*direction.y][start.x+1+2*direction.x].topLeftY - g.currentLevel.positions[start.y+1+direction.y][start.x+1+direction.x].topLeftY
		}
	case stepAttackEffect:
		g.inAnimation = true
		gAnimationSet[stepAttackEffect].drawX = g.currentLevel.positions[g.recordAttack.y+1][g.recordAttack.x+1].topLeftX
		gAnimationSet[stepAttackEffect].drawY = g.currentLevel.positions[g.recordAttack.y+1][g.recordAttack.x+1].topLeftY
		g.currentLevel.applyAttackEffects(g.recordAttack)
	case stepCheckStatus:
		if g.currentLevel.isLost() {
			g.setFirstLevel(testLevel)
		}

		if g.currentLevel.isCompleted() {
			g.setLevel(testLevel)
		}
		g.step = stepPlayerTurn
	}

	return nil
}
