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
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/loig/ld56/assets"
)

func main() {

	assets.Load()
	setAnimations()

	g := game{
		state: stateTitle,
		step:  stepPlayerTurn,
	}
	g.setFirstLevel(0)

	g.soundManager = assets.InitAudio()

	ebiten.SetWindowTitle("LD56")
	ebiten.SetWindowSize(gScreenWidth, gScreenHeight)

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Fatal(err)
	}

}
