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
	"github.com/loig/ld56/assets"
)

type animation struct {
	numSteps     int
	numFrames    []int
	image        *ebiten.Image
	subImageX    int
	drawX, drawY int
	stepX, stepY int
	drawable     bool
	nextStep     int
	stepSoundID  int
	stepSound    bool
}

var gAnimationSet [stepNumber]animation

func setAnimations() {

	gAnimationSet[stepBadGuysMoving] = animation{
		numSteps:  1,
		numFrames: []int{15},
		nextStep:  stepBadGuysMoving,
	}

	gAnimationSet[stepBadGuysAttacking] = animation{
		numFrames:   []int{10},
		image:       assets.IconesImage,
		subImageX:   3,
		drawable:    true,
		nextStep:    stepAttackEffect,
		stepSound:   true,
		stepSoundID: assets.SoundBallID,
	}

	gAnimationSet[stepAttackEffect] = animation{
		numSteps:  1,
		numFrames: []int{10},
		image:     assets.IconesImage,
		subImageX: 4,
		drawable:  true,
		nextStep:  stepBadGuysAttacking,
	}

	gAnimationSet[stepCombat] = animation{
		numSteps:  1,
		numFrames: []int{10},
		image:     assets.IconesImage,
		subImageX: 5,
		drawable:  true,
		nextStep:  stepBadGuysMoving,
	}

}

func updateAnimationDraw(gameStep int) {
	gAnimationSet[gameStep].drawX += gAnimationSet[gameStep].stepX
	gAnimationSet[gameStep].drawY += gAnimationSet[gameStep].stepY
}
