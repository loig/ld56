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

const gMaxLife int = 50
const gParticleShift float64 = 0.5

type particle struct {
	value int
	x, y  float64
	life  int
}

var particleSet []*particle

func updateParticles() {
	for _, p := range particleSet {
		if p.life > 0 {
			p.update()
		}
	}
}

func drawParticles(screen *ebiten.Image) {
	for _, p := range particleSet {
		if p.life > 0 {
			p.draw(screen)
		}
	}
}

func addParticle(x, y int, v int) {

	if len(particleSet) == 0 {
		particleSet = make([]*particle, 1)
		particleSet[0] = &particle{
			x: float64(x), y: float64(y), value: v, life: gMaxLife,
		}
		return
	}

	for _, p := range particleSet {
		if p.life <= 0 {
			p.x = float64(x)
			p.y = float64(y)
			p.value = v
			p.life = gMaxLife
			return
		}
	}

	particleSet = append(particleSet, &particle{
		x: float64(x), y: float64(y), value: v, life: gMaxLife,
	})
}

func (p *particle) update() {
	p.life--
	p.y += gParticleShift
}

func (p particle) draw(screen *ebiten.Image) {

	options := ebiten.DrawImageOptions{}
	options.GeoM.Scale(0.5, 0.5)
	options.GeoM.Translate(p.x, p.y)

	if p.value > 0 {
		options.ColorScale.SetR(0.5)
		options.ColorScale.SetB(0.5)
	} else {
		options.ColorScale.SetB(0.5)
		options.ColorScale.SetG(0.5)
	}

	options.ColorScale.SetA(float32(p.life) / float32(gMaxLife))

	shift := gTileSize/4 + 5

	digits := make([]int, 0)
	n := p.value
	if n < 0 {
		n = -n
	}
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
	}

}
