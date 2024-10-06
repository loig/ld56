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
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loig/ld56/assets"
)

func (g *game) Draw(screen *ebiten.Image) {

	options := ebiten.DrawImageOptions{}
	screen.DrawImage(assets.BackImage, &options)

	if g.state != stateCongrats && g.state != stateGameover {
		g.drawLevel(screen)
	}

	if g.state == stateTitle {
		screen.DrawImage(assets.TitleImage, &options)
	}

	if g.state == stateCongrats {
		screen.DrawImage(assets.CongratsImage, &options)
	}

	if g.state == stateGameover {
		screen.DrawImage(assets.GameoverImage, &options)
	}

	if g.state == stateCongrats || g.state == stateGameover {
		g.drawResults(screen)
	}

	if g.state == stateInLevel {
		g.drawInfo(screen)
		drawParticles(screen)
	}

	if g.inAnimation && g.step != stepLevelStart && g.step != stepLevelEnd {
		if gAnimationSet[g.step].drawable {
			options = ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(gAnimationSet[g.step].drawX), float64(gAnimationSet[g.step].drawY))
			screen.DrawImage(gAnimationSet[g.step].image.SubImage(image.Rect(
				gAnimationSet[g.step].subImageX*gTileSize, 0, (gAnimationSet[g.step].subImageX+1)*gTileSize, gTileSize,
			)).(*ebiten.Image), &options)
		}
	}

}

func (g *game) drawInfo(screen *ebiten.Image) {

	options := ebiten.DrawImageOptions{}

	options.GeoM.Translate(float64(gTileSize/4), float64(gTileSize/4))
	shift := gTileSize / 4

	// food
	screen.DrawImage(assets.IconesImage.SubImage(image.Rect(
		0, 0, gTileSize, gTileSize,
	)).(*ebiten.Image), &options)

	drawNumberAt(g.currentLevel.numFood, screen, shift+gTileSize, gTileSize/2)

	// characters

	options.GeoM.Translate(float64(gScreenWidth/2-gTileSize-gTileSize/2), 0)
	shift += gScreenWidth/2 - gTileSize - gTileSize/2

	screen.DrawImage(assets.IconesImage.SubImage(image.Rect(
		gTileSize, 0, 2*gTileSize, gTileSize,
	)).(*ebiten.Image), &options)

	drawNumberAt(g.currentLevel.numCharacters, screen, shift+gTileSize, gTileSize/2)

	// level

	moreShift := gTileSize + gTileSize/2
	if g.levelNum >= 10 {
		moreShift -= gTileSize/4 + 5
	}
	options.GeoM.Translate(float64(gScreenWidth/4+moreShift), 0)
	shift += gScreenWidth/4 + moreShift

	drawLevelAt(g.levelNum, screen, shift+gTileSize, gTileSize/2)

}

func (g *game) drawResults(screen *ebiten.Image) {

	options := ebiten.DrawImageOptions{}

	options.GeoM.Translate(float64(gScreenWidth/2-gTileSize-gTileSize/2), float64(gScreenHeight/4+gTileSize/2))
	shiftY := gScreenHeight/4 + gTileSize/2
	shiftX := gScreenWidth/2 - gTileSize - gTileSize/2

	// level
	levelXShift := gTileSize / 2
	if g.levelNum >= 10 {
		levelXShift -= gTileSize/4 + 5
	}
	drawLevelAt(g.levelNum, screen, shiftX+levelXShift, shiftY+gTileSize/4)

	options.GeoM.Translate(0, float64(gTileSize+gTileSize/2))
	shiftY += gTileSize + gTileSize/2

	// characters
	screen.DrawImage(assets.IconesImage.SubImage(image.Rect(
		gTileSize, 0, 2*gTileSize, gTileSize,
	)).(*ebiten.Image), &options)

	drawNumberAt(g.currentLevel.numCharacters, screen, shiftX+gTileSize, shiftY+gTileSize/4)

	options.GeoM.Translate(0, float64(gTileSize))
	shiftY += gTileSize

	// food
	screen.DrawImage(assets.IconesImage.SubImage(image.Rect(
		0, 0, gTileSize, gTileSize,
	)).(*ebiten.Image), &options)

	drawNumberAt(g.currentLevel.numFood, screen, shiftX+gTileSize, shiftY+gTileSize/4)

	options.GeoM.Translate(0, float64(gTileSize))
	shiftY += gTileSize

	// steps
	screen.DrawImage(assets.IconesImage.SubImage(image.Rect(
		6*gTileSize, 0, 7*gTileSize, gTileSize,
	)).(*ebiten.Image), &options)

	drawNumberAt(g.footSteps, screen, shiftX+gTileSize, shiftY+gTileSize/4)

}

func (g *game) drawLevel(screen *ebiten.Image) {
	numDrawn := 0
	numNotDrawn := 0
	toDraw := 0
	for y, line := range g.currentLevel.area {
		for x, areaType := range line {
			if areaType != areaTypeNone {
				toDraw++
				if g.state != stateInLevel ||
					(g.step != stepLevelStart && g.step != stepLevelEnd) ||
					(g.step == stepLevelStart && numDrawn < g.animationStep) ||
					(g.step == stepLevelEnd && numNotDrawn >= g.animationStep) {
					numDrawn++
					options := ebiten.DrawImageOptions{}

					if g.state == stateInLevel && g.step == stepLevelStart && numDrawn == g.animationStep {
						options.ColorScale.ScaleAlpha(float32(g.animationFrame) / 5)
					}

					options.GeoM.Translate(float64(g.currentLevel.positions[y+1][x+1].topLeftX), float64(g.currentLevel.positions[y+1][x+1].topLeftY))
					if g.currentLevel.iSelected == y+1 && g.currentLevel.jSelected == x+1 {
						options.GeoM.Translate(0, float64(gTileYSelectShift))
					}
					screen.DrawImage(assets.TerrainImage.SubImage(image.Rect(
						(areaType-1)*gTileSize, 0, areaType*gTileSize, gTileSize,
					)).(*ebiten.Image), &options)

					froots, exists := g.currentLevel.froots[y]
					if exists {
						froot, isHere := froots[x]
						if isHere && !froot.eaten {
							screen.DrawImage(assets.FrootsImage.SubImage(image.Rect(
								froot.kind*gTileSize, 0, (froot.kind+1)*gTileSize, gTileSize,
							)).(*ebiten.Image), &options)
						}
					}

					if x+1 == g.currentLevel.jCharacters && y+1 == g.currentLevel.iCharacters {
						screen.DrawImage(assets.CharactersImage.SubImage(image.Rect(
							(g.currentLevel.numCharacters-1)*gTileSize, 0, g.currentLevel.numCharacters*gTileSize, gTileSize,
						)).(*ebiten.Image), &options)
					}

					for _, b := range g.currentLevel.badGuys {
						if !b.dead && b.x == x && b.y == y {
							screen.DrawImage(assets.BadguysImage.SubImage(image.Rect(
								b.kind*gTileSize, 0, (b.kind+1)*gTileSize, gTileSize,
							)).(*ebiten.Image), &options)
						}
					}

					if g.state == stateInLevel {
						if x+1 == g.currentLevel.jGoal && y+1 == g.currentLevel.iGoal {
							screen.DrawImage(assets.IconesImage.SubImage(image.Rect(
								2*gTileSize, 0, 3*gTileSize, gTileSize,
							)).(*ebiten.Image), &options)
						}
					}
				} else {
					numNotDrawn++
				}
			}
		}
	}

	if (g.step == stepLevelStart && numDrawn >= toDraw) ||
		(g.step == stepLevelEnd && numNotDrawn >= toDraw) {
		g.inAnimation = false
	}
}

func drawNumberAt(n int, screen *ebiten.Image, posX, posY int) (endX int) {

	options := ebiten.DrawImageOptions{}
	options.GeoM.Scale(0.5, 0.5)
	options.GeoM.Translate(float64(posX), float64(posY))

	shift := gTileSize/4 + 5

	endX = posX

	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	if len(digits) == 0 {
		digits = append(digits, 0)
	}

	for i := len(digits) - 1; i >= 0; i-- {
		screen.DrawImage(assets.DigitsImage.SubImage(image.Rect(
			digits[i]*gTileSize, 0, (digits[i]+1)*gTileSize, gTileSize,
		)).(*ebiten.Image), &options)
		options.GeoM.Translate(float64(shift), 0)
		endX += shift
	}

	return
}

func drawLevelAt(n int, screen *ebiten.Image, posX, posY int) {

	if n > gNumLevels {
		n = gNumLevels
	}

	nextX := drawNumberAt(n, screen, posX, posY)

	options := ebiten.DrawImageOptions{}
	options.GeoM.Scale(0.5, 0.5)
	options.GeoM.Translate(float64(nextX), float64(posY))

	screen.DrawImage(assets.DigitsImage.SubImage(image.Rect(
		10*gTileSize, 0, 11*gTileSize, gTileSize,
	)).(*ebiten.Image), &options)
	nextX += gTileSize/4 + 5

	drawNumberAt(gNumLevels, screen, nextX, posY)

}
