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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/loig/ld56/assets"
)

func (g *game) Update() error {

	updateParticles()

	if g.state == stateCongrats || g.state == stateGameover || g.state == stateTitle {
		g.soundManager.UpdateMusic(0.3)
	} else {
		g.soundManager.UpdateMusic(0.1)
	}
	g.soundManager.PlaySounds()

	if g.state == stateInLevel && (g.step == stepLevelStart || g.step == stepLevelEnd) {
		g.animationFrame++
		if g.animationFrame >= 5 {
			g.animationFrame = 0
			g.animationStep++
			if !g.inAnimation {
				if g.step == stepLevelStart {
					g.step = stepPlayerTurn
				} else {
					g.setLevel(g.levelNum)
					if g.levelNum > gNumLevels {
						g.state = stateCongrats
					}
					g.step = stepLevelStart
					g.inAnimation = true
				}

				g.animationFrame = 0
				g.animationStep = 0

			}
		}
		return nil
	}

	if g.state == stateGameover || g.state == stateCongrats {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.footSteps = 0
			g.levelNum = 0
			g.state = stateTitle
			g.setFirstLevel(0)
			g.soundManager.NextSounds[assets.SoundMvtID] = true
		}
		return nil
	}

	if g.state == stateTitle {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.state = stateInLevel
			g.step = stepPlayerTurn
			g.soundManager.NextSounds[assets.SoundMvtID] = true
		}
		return nil
	}

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
				if gAnimationSet[g.step].stepSound {
					g.soundManager.NextSounds[gAnimationSet[g.step].stepSoundID] = true
				}
			}
		}

		return nil
	}

	mouseX, mouseY := ebiten.CursorPosition()

	changed := g.currentLevel.setSelected(mouseX, mouseY)
	if changed {
		g.soundManager.NextSounds[assets.SoundSelectID] = true
	}

	switch g.step {
	case stepPlayerTurn:
		if g.levelNum == 1 && g.currentLevel.numCharacters > gTrueInitNumCharacters {
			g.step = stepBadGuysMoving
			return nil
		}
		if g.levelNum == 1 {
			g.currentLevel.area[2][2] = areaTypeGrass
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			hasMoved := g.currentLevel.updateCharactersPosition()
			if hasMoved {
				g.soundManager.NextSounds[assets.SoundMvtID] = true
				g.footSteps++
				hasEaten := g.currentLevel.updateFood()
				if hasEaten {
					g.soundManager.NextSounds[assets.SoundEatID] = true
					addParticle(gTileSize+gTileSize/4, gTileSize/2+gTileSize/4, gFoodGain)
				}
				if g.currentLevel.isCompleted() {
					g.step = stepLevelEnd
					g.inAnimation = true
					return nil
				}
				g.step = stepBadGuysMoving
			} else {
				g.soundManager.NextSounds[assets.SoundNonID] = true
			}
		}
	case stepBadGuysMoving:
		hasMoved, inCombat, done, soundType := g.currentLevel.moveBadGuys()
		if done {
			g.step = stepBadGuysAttacking
		} else if hasMoved || inCombat {
			g.inAnimation = true
			if inCombat {
				g.soundManager.NextSounds[assets.SoundBattleID] = true
				g.step = stepCombat
				gAnimationSet[stepCombat].drawX = g.currentLevel.positions[g.currentLevel.iCharacters][g.currentLevel.jCharacters].topLeftX
				gAnimationSet[stepCombat].drawY = g.currentLevel.positions[g.currentLevel.iCharacters][g.currentLevel.jCharacters].topLeftY
			} else {
				g.soundManager.NextSounds[soundType] = true
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
		g.soundManager.NextSounds[assets.SoundExplodeID] = true
		gAnimationSet[stepAttackEffect].drawX = g.currentLevel.positions[g.recordAttack.y+1][g.recordAttack.x+1].topLeftX
		gAnimationSet[stepAttackEffect].drawY = g.currentLevel.positions[g.recordAttack.y+1][g.recordAttack.x+1].topLeftY
		g.currentLevel.applyAttackEffects(g.recordAttack)
	case stepCheckStatus:
		if g.currentLevel.isLost() {
			g.state = stateGameover
		}
		g.step = stepPlayerTurn
	}

	return nil
}
