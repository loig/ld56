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

const gNumLevels int = 15

func nextLevel(n int) (l level) {
	switch n {
	case 0:
		l = firstLevel
		firstLevel.alreadySeen = true
	case 1:
		l = frootLevel
		frootLevel.alreadySeen = true
	case 2:
		l = ladybugLevel
		ladybugLevel.alreadySeen = true
	case 3:
		l = ladybugLevel2
		ladybugLevel2.alreadySeen = true
	case 4:
		l = escapeladyLevel
		escapeladyLevel.alreadySeen = true
	case 5:
		l = dungbeetleLevel
		dungbeetleLevel.alreadySeen = true
	case 6:
		l = dungbeetleLevel2
		dungbeetleLevel2.alreadySeen = true
	case 7:
		l = dungbeetleLevel3
		dungbeetleLevel3.alreadySeen = true
	case 8:
		l = level4
		level4.alreadySeen = true
	case 9:
		l = level1
		level1.alreadySeen = true
	case 10:
		l = level2
		level2.alreadySeen = true
	case 11:
		l = beeLevel
		beeLevel.alreadySeen = true
	case 12:
		l = level5
		level5.alreadySeen = true
	case 13:
		l = level3
		level3.alreadySeen = true
	default:
		l = lastLevel
		lastLevel.alreadySeen = true
	}

	return
}

var firstLevel level = level{
	width:  5,
	height: 5,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeVillage, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 2, y: 0, kind: badGuySuperBee},
		{x: 4, y: 0, kind: badGuySuperBee},
		{x: 0, y: 2, kind: badGuySuperBee},
		{x: 4, y: 2, kind: badGuySuperBee},
		{x: 0, y: 4, kind: badGuySuperBee},
		{x: 2, y: 4, kind: badGuySuperBee},
	},
	iCharacters: 3, jCharacters: 3,
	iGoal: 1, jGoal: 5,
}

var lastLevel level = level{
	width:  13,
	height: 13,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeWater, areaTypeWater, areaTypeWater},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeGrass, areaTypeWater, areaTypeWater, areaTypeGrass, areaTypeWater},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeGrass, areaTypeWater, areaTypeCastle, areaTypeWater, areaTypeGrass, areaTypeWater},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeGrass, areaTypeWater},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeWater, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeWater, areaTypeWater, areaTypeWater},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone},
	},
	badGuysSetup: []badGuy{
		{x: 6, y: 3, kind: badGuyBee},
		{x: 9, y: 3, kind: badGuyBee},
		{x: 3, y: 6, kind: badGuyBee},
		{x: 9, y: 6, kind: badGuyBee},
		{x: 6, y: 9, kind: badGuyBee},
	},
	iCharacters: 10, jCharacters: 4,
	iGoal: 7, jGoal: 7,
}

var frootLevel level = level{
	width:  5,
	height: 5,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeGrass},
		{areaTypeGrass},
	},
	froots: map[int]map[int]*froot{
		2: map[int]*froot{2: &froot{kind: frootTypeRaspberry}},
	},
	iCharacters: 5, jCharacters: 1,
	iGoal: 1, jGoal: 5,
}

var beeLevel level = level{
	width:  7,
	height: 5,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeGrass, areaTypeWater, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 1, y: 2, kind: badGuyBee},
		{x: 0, y: 4, kind: badGuyBee},
		{x: 4, y: 0, kind: badGuyDungBeetle},
	},
	froots: map[int]map[int]*froot{
		2: map[int]*froot{3: &froot{kind: frootTypeRaspberry}},
	},
	iCharacters: 5, jCharacters: 3,
	iGoal: 1, jGoal: 7,
}

var ladybugLevel level = level{
	width:  5,
	height: 4,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 0, y: 3, kind: badGuyLadyBug},
	},
	iCharacters: 4, jCharacters: 2,
	iGoal: 1, jGoal: 5,
}

var ladybugLevel2 level = level{
	width:  4,
	height: 3,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeGrass, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 0, y: 1, kind: badGuyLadyBug},
	},
	iCharacters: 3, jCharacters: 2,
	iGoal: 1, jGoal: 4,
}

var escapeladyLevel level = level{
	width:  5,
	height: 4,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeWater},
		{areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 2, y: 1, kind: badGuyLadyBug},
	},
	iCharacters: 4, jCharacters: 3,
	iGoal: 1, jGoal: 4,
}

var dungbeetleLevel level = level{
	width:  6,
	height: 5,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 4, y: 1, kind: badGuyLadyBug},
		{x: 0, y: 2, kind: badGuyDungBeetle},
	},
	iCharacters: 5, jCharacters: 2,
	iGoal: 1, jGoal: 6,
}

var dungbeetleLevel2 level = level{
	width:  6,
	height: 5,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 0, y: 2, kind: badGuyDungBeetle},
	},
	froots: map[int]map[int]*froot{
		2: map[int]*froot{2: &froot{kind: frootTypeRaspberry}},
	},
	iCharacters: 5, jCharacters: 2,
	iGoal: 1, jGoal: 6,
}

var dungbeetleLevel3 level = level{
	width:  6,
	height: 5,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeGrass},
		{areaTypeNone, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 0, y: 2, kind: badGuyDungBeetle},
	},
	iCharacters: 5, jCharacters: 2,
	iGoal: 1, jGoal: 6,
}

var level1 level = level{
	width:  8,
	height: 5,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeWater, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 6, y: 2, kind: badGuyLadyBug},
		{x: 4, y: 1, kind: badGuyDungBeetle},
	},
	froots: map[int]map[int]*froot{
		4: map[int]*froot{4: &froot{kind: frootTypeRaspberry}},
	},
	iCharacters: 5, jCharacters: 4,
	iGoal: 2, jGoal: 4,
}

var level2 level = level{
	width:  5,
	height: 5,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 2, y: 4, kind: badGuyDungBeetle},
		{x: 1, y: 1, kind: badGuyDungBeetle},
	},
	froots: map[int]map[int]*froot{
		1: map[int]*froot{3: &froot{kind: frootTypeRaspberry}},
		2: map[int]*froot{1: &froot{kind: frootTypeRaspberry}},
		3: map[int]*froot{2: &froot{kind: frootTypeRaspberry}},
	},
	iCharacters: 4, jCharacters: 1,
	iGoal: 1, jGoal: 5,
}

var level3 level = level{
	width:  8,
	height: 5,
	area: [][]int{
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeNone, areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 6, y: 2, kind: badGuyBee},
		{x: 4, y: 1, kind: badGuyDungBeetle},
	},
	froots: map[int]map[int]*froot{
		1: map[int]*froot{6: &froot{kind: frootTypeRaspberry}},
		2: map[int]*froot{5: &froot{kind: frootTypeRaspberry}},
		3: map[int]*froot{4: &froot{kind: frootTypeRaspberry}},
		4: map[int]*froot{4: &froot{kind: frootTypeRaspberry}},
	},
	iCharacters: 5, jCharacters: 4,
	iGoal: 2, jGoal: 4,
}

var level4 level = level{
	width:  4,
	height: 4,
	area: [][]int{
		{areaTypeNone, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeWater, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 2, y: 1, kind: badGuyLadyBug},
		{x: 3, y: 2, kind: badGuyDungBeetle},
	},
	iCharacters: 4, jCharacters: 1,
	iGoal: 1, jGoal: 2,
}

var level5 level = level{
	width:  4,
	height: 4,
	area: [][]int{
		{areaTypeNone, areaTypeGrass},
		{areaTypeGrass, areaTypeWater, areaTypeWater},
		{areaTypeGrass, areaTypeGrass, areaTypeGrass, areaTypeGrass},
		{areaTypeGrass, areaTypeGrass},
	},
	badGuysSetup: []badGuy{
		{x: 2, y: 1, kind: badGuyBee},
		{x: 3, y: 2, kind: badGuyDungBeetle},
	},
	froots: map[int]map[int]*froot{
		2: map[int]*froot{1: &froot{kind: frootTypeRaspberry}},
	},
	iCharacters: 4, jCharacters: 1,
	iGoal: 1, jGoal: 2,
}
