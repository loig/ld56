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

package assets

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed terrain.png
var terrainBytes []byte
var TerrainImage *ebiten.Image

//go:embed characters.png
var charactersBytes []byte
var CharactersImage *ebiten.Image

//go:embed froots.png
var frootsBytes []byte
var FrootsImage *ebiten.Image

//go:embed icones.png
var iconesBytes []byte
var IconesImage *ebiten.Image

//go:embed badguys.png
var badguysBytes []byte
var BadguysImage *ebiten.Image

//go:embed back.png
var backBytes []byte
var BackImage *ebiten.Image

func Load() {

	decoded, _, err := image.Decode(bytes.NewReader(terrainBytes))
	if err != nil {
		log.Fatal(err)
	}
	TerrainImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(charactersBytes))
	if err != nil {
		log.Fatal(err)
	}
	CharactersImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(frootsBytes))
	if err != nil {
		log.Fatal(err)
	}
	FrootsImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(iconesBytes))
	if err != nil {
		log.Fatal(err)
	}
	IconesImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(badguysBytes))
	if err != nil {
		log.Fatal(err)
	}
	BadguysImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(backBytes))
	if err != nil {
		log.Fatal(err)
	}
	BackImage = ebiten.NewImageFromImage(decoded)
}
