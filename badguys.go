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
	"slices"

	"github.com/loig/ld56/assets"
)

type badGuy struct {
	x, y int
	kind int
	dead bool
}

const (
	badGuyBee = iota
	badGuyDungBeetle
	badGuyLadyBug
	badGuyDungBeetleNoDung
	badGuySuperBee
)

func (l *level) moveBadGuys() (hasMoved, inCombat, done bool, soundID int) {

	if l.currentGuy >= len(l.badGuys) {
		l.currentGuy = 0
		return false, false, true, 0
	}

	if !l.badGuys[l.currentGuy].dead {

		x, y := l.badGuys[l.currentGuy].x, l.badGuys[l.currentGuy].y
		l.badGuys[l.currentGuy].x, l.badGuys[l.currentGuy].y = l.getNewPosition(l.badGuys[l.currentGuy])

		hasMoved = x != l.badGuys[l.currentGuy].x || y != l.badGuys[l.currentGuy].y

		inCombat = l.checkCombat(l.currentGuy)

		switch l.badGuys[l.currentGuy].kind {
		case badGuyBee, badGuySuperBee:
			soundID = assets.SoundBeeID
		case badGuyLadyBug:
			soundID = assets.SoundLadyBugID
		}
	}

	l.currentGuy++

	return
}

func (l *level) makeBadGuysAttack() (dist int, start, dir, attackPoint step, done bool) {

	if l.currentGuy >= len(l.badGuys) {
		l.currentGuy = 0
		return 0, step{}, step{}, step{}, true
	}

	if !l.badGuys[l.currentGuy].dead {
		if l.badGuys[l.currentGuy].kind == badGuyDungBeetle {

			// find direction of attack
			directions := [6]step{
				{x: 0, y: -1}, {x: 1, y: -1},
				{x: -1, y: 0}, {x: 1, y: 0},
				{x: -1, y: 1}, {x: 0, y: 1},
			}

			distance := 0
			direction := step{}

			for k := 0; k < 6; k++ {
				d := l.getLineDistance(l.badGuys[l.currentGuy].x, l.badGuys[l.currentGuy].y, directions[k])

				if distance == 0 || (d > 0 && d < distance) {
					distance = d
					direction = directions[k]
				}
			}

			// if attack possible, do it
			if distance > 0 {
				l.badGuys[l.currentGuy].kind = badGuyDungBeetleNoDung

				attackPoint = step{
					x: l.badGuys[l.currentGuy].x + direction.x*distance,
					y: l.badGuys[l.currentGuy].y + direction.y*distance,
				}

				return distance, step{x: l.badGuys[l.currentGuy].x, y: l.badGuys[l.currentGuy].y}, direction, attackPoint, false

			}

		}
	}

	l.currentGuy++
	return
}

func (l *level) applyAttackEffects(attackPoint step) {
	if l.jCharacters == attackPoint.x+1 && l.iCharacters == attackPoint.y+1 {
		l.numCharacters -= 5
		addParticle(gScreenWidth/2-gTileSize/4, gTileSize/2+gTileSize/4, -5)
		if l.numCharacters < 0 {
			l.numCharacters = 0
		}
		return
	}

	for i, b := range l.badGuys {
		if b.x == attackPoint.x && b.y == attackPoint.y {
			l.badGuys[i].dead = true
		}
	}
}

func (l level) getLineDistance(fromX, fromY int, direction step) (dist int) {

	pos := step{x: fromX + direction.x, y: fromY + direction.y}
	dist = 1

	for pos.x > 0 && pos.y > 0 && pos.y < len(l.area) && pos.x < len(l.area[pos.y]) &&
		l.area[pos.y][pos.x] == areaTypeGrass {

		if l.jCharacters == pos.x+1 && l.iCharacters == pos.y+1 {
			return
		}

		froots, exists := l.froots[pos.y]
		if exists {
			froot, isHere := froots[pos.x]
			if isHere && !froot.eaten {
				return 0
			}
		}

		for _, b := range l.badGuys {
			if !b.dead && b.kind != badGuyDungBeetle && b.kind != badGuyDungBeetleNoDung {
				if b.x == pos.x && b.y == pos.y {
					return
				}
			}
		}

		pos.x += direction.x
		pos.y += direction.y
		dist++
	}

	return 0
}

func (l *level) checkCombat(id int) (inCombat bool) {
	if l.jCharacters == l.badGuys[id].x+1 && l.iCharacters == l.badGuys[id].y+1 {
		inCombat = true
		switch l.badGuys[id].kind {
		case badGuyBee:
			l.numCharacters -= 5
			addParticle(gScreenWidth/2-gTileSize/4, gTileSize/2+gTileSize/4, -5)
		case badGuyDungBeetle, badGuyDungBeetleNoDung:
			l.numCharacters -= 5
			addParticle(gScreenWidth/2-gTileSize/4, gTileSize/2+gTileSize/4, -5)
		case badGuyLadyBug:
			l.numCharacters -= 5
			addParticle(gScreenWidth/2-gTileSize/4, gTileSize/2+gTileSize/4, -5)
		case badGuySuperBee:
			l.numCharacters -= 30
			addParticle(gScreenWidth/2-gTileSize/4, gTileSize/2+gTileSize/4, -30)
		}
		l.badGuys[id].dead = true
		if l.numCharacters < 0 {
			l.numCharacters = 0
		}
	}
	return
}

func (l level) getNewPosition(b badGuy) (x, y int) {
	switch b.kind {
	case badGuyLadyBug:
		return l.nextStepOnPath(b.x, b.y, 1, false)
	case badGuyBee, badGuySuperBee:
		return l.nextStepOnPath(b.x, b.y, 1, true)
	}
	return b.x, b.y
}

type step struct {
	x, y int
}

func compare(s1, s2 step) int {
	if s1.x < s2.x {
		return -1
	}

	if s1.x > s2.x {
		return 1
	}

	if s1.y < s2.y {
		return -1
	}

	if s1.y > s2.y {
		return 1
	}

	return 0
}

func (l level) allSteps(fromX, fromY int, speed int) (steps []step) {

	if speed == 0 {
		return []step{{x: fromX, y: fromY}}
	}

	tmpSteps := make([]step, 0)

	for _, s := range l.allSteps(fromX, fromY, speed-1) {
		tmpSteps = append(tmpSteps, s)
		tmpSteps = append(tmpSteps, step{x: s.x, y: s.y - 1})
		tmpSteps = append(tmpSteps, step{x: s.x + 1, y: s.y - 1})
		tmpSteps = append(tmpSteps, step{x: s.x - 1, y: s.y})
		tmpSteps = append(tmpSteps, step{x: s.x + 1, y: s.y})
		tmpSteps = append(tmpSteps, step{x: s.x - 1, y: s.y + 1})
		tmpSteps = append(tmpSteps, step{x: s.x, y: s.y + 1})
	}

	slices.SortFunc(tmpSteps, compare)

	return slices.Compact(tmpSteps)
}

func (l level) possibleSteps(fromX, fromY int, speed int, walkOnWater bool) (steps []step) {

	tmpSteps := l.allSteps(fromX, fromY, speed)

	for _, s := range tmpSteps {
		if s.x >= 0 && s.y >= 0 && s.y < len(l.area) && s.x < len(l.area[s.y]) &&
			l.area[s.y][s.x] != areaTypeNone &&
			(l.area[s.y][s.x] != areaTypeWater || walkOnWater) {
			hasFroot := false
			froots, exists := l.froots[s.y]
			if exists {
				froot, isHere := froots[s.x]
				if isHere {
					hasFroot = !froot.eaten
				}
			}
			if !hasFroot {

				hasBadGuy := false
				for _, b := range l.badGuys {
					if b.x == s.x && b.y == s.y {
						hasBadGuy = !b.dead
						break
					}
				}

				if !hasBadGuy {
					steps = append(steps, s)
				}
			}
		}
	}

	return

}

func (l level) nextStepOnPath(fromX, fromY int, speed int, walkOnWater bool) (x, y int) {
	// breadth-first search

	goalX, goalY := l.jCharacters-1, l.iCharacters-1

	if fromX == goalX && fromY == goalY {
		return fromX, fromY
	}

	origin := step{x: fromX, y: fromY}

	nexts := make([]step, 1)
	nexts[0] = origin

	visited := make(map[step]step)
	visited[nexts[0]] = origin

	for len(nexts) > 0 {
		current := nexts[0]
		nexts = nexts[1:]

		pred := visited[current]
		atOrigin := pred == origin

		toAdd := l.possibleSteps(current.x, current.y, speed, walkOnWater)

		slices.SortFunc(toAdd, func(s1, s2 step) int {
			d1 := distance(float64(s1.x), float64(s1.y), float64(goalX), float64(goalY))
			d2 := distance(float64(s2.x), float64(s2.y), float64(goalX), float64(goalY))

			if d1 < d2 {
				return -1
			}

			if d2 > d1 {
				return 1
			}

			return 0
		})

		for _, s := range toAdd {
			if _, exists := visited[s]; !exists {

				if atOrigin {
					pred = s
				}

				if s.x == goalX && s.y == goalY {
					return pred.x, pred.y
				}

				nexts = append(nexts, s)
				visited[s] = pred
			}
		}
	}

	return fromX, fromY
}
