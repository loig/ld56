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
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/loig/ld56/assets"
)

func (g *game) Draw(screen *ebiten.Image) {

	options := ebiten.DrawImageOptions{}
	screen.DrawImage(assets.BackImage, &options)

	g.drawLevel(screen)
	g.drawInfo(screen)

	if g.inAnimation {
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

	options.GeoM.Translate(float64(gTileSize), float64(gTileSize))

	// characters
	screen.DrawImage(assets.IconesImage.SubImage(image.Rect(
		gTileSize, 0, 2*gTileSize, gTileSize,
	)).(*ebiten.Image), &options)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint(g.currentLevel.numCharacters), 2*gTileSize, gTileSize+25)

	options.GeoM.Translate(0, float64(gTileSize))

	// food
	screen.DrawImage(assets.IconesImage.SubImage(image.Rect(
		0, 0, gTileSize, gTileSize,
	)).(*ebiten.Image), &options)

	ebitenutil.DebugPrintAt(screen, fmt.Sprint(g.currentLevel.numFood), 2*gTileSize, 2*gTileSize+25)

	// level
	ebitenutil.DebugPrintAt(screen, fmt.Sprint(g.levelNum), gScreenWidth-gTileSize, gTileSize+25)

}

func (g *game) drawLevel(screen *ebiten.Image) {
	for y, line := range g.currentLevel.area {
		for x, areaType := range line {
			if areaType != areaTypeNone {
				options := ebiten.DrawImageOptions{}
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

				if x+1 == g.currentLevel.jGoal && y+1 == g.currentLevel.iGoal {
					screen.DrawImage(assets.IconesImage.SubImage(image.Rect(
						2*gTileSize, 0, 3*gTileSize, gTileSize,
					)).(*ebiten.Image), &options)
				}
			}
		}
	}
}
