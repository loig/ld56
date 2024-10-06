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

const (
	gScreenWidth, gScreenHeight     int = 800, 600
	gPlayAreaX, gPlayAreaY          int = 150, 50
	gPlayAreaWidth, gPlayAreaHeight int = 512, 512
	gTileSize                       int = 64
	gTileYShiftPerStep              int = 6
	gTileXShiftPerStep              int = -11
	gTileXShiftPerLine              int = 12
	gTileYShiftPerLine              int = -34
	gTileXCenterShift               int = 32
	gTileYCenterShift               int = 36
	gTileYSelectShift               int = -5
	gInitNumFood                    int = 500
	gInitNumCharacters              int = 200
	gTrueInitNumCharacters          int = 20
	gNumDeathsForMissingFood        int = 1
	gFoodGain                       int = 200
)
