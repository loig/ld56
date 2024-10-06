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
	"io"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"golang.org/x/exp/rand"
)

//go:embed mvt1.mp3
var soundMvt1Bytes []byte
var soundMvt1 []byte

//go:embed mvt2.mp3
var soundMvt2Bytes []byte
var soundMvt2 []byte

//go:embed mvt3.mp3
var soundMvt3Bytes []byte
var soundMvt3 []byte

//go:embed mvt4.mp3
var soundMvt4Bytes []byte
var soundMvt4 []byte

//go:embed select.mp3
var soundSelectBytes []byte
var soundSelect []byte

//go:embed ball.mp3
var soundBallBytes []byte
var soundBall []byte

//go:embed battle.mp3
var soundBattleBytes []byte
var soundBattle []byte

//go:embed bee.mp3
var soundBeeBytes []byte
var soundBee []byte

//go:embed fart.mp3
var soundFartBytes []byte
var soundFart []byte

//go:embed ladybug.mp3
var soundLadyBugBytes []byte
var soundLadyBug []byte

//go:embed non.mp3
var soundNonBytes []byte
var soundNon []byte

//go:embed eat.mp3
var soundEatBytes []byte
var soundEat []byte

//go:embed music.mp3
var musicBytes []byte
var music *audio.InfiniteLoop

const (
	SoundSelectID int = iota
	SoundMvtID
	SoundBeeID
	SoundLadyBugID
	SoundBallID
	SoundBattleID
	SoundExplodeID
	SoundNonID
	SoundEatID
	NumSounds
)

type SoundManager struct {
	audioContext *audio.Context
	NextSounds   [NumSounds]bool
	music        *audio.Player
}

// loop the music
func (s *SoundManager) UpdateMusic(volume float64) {
	if s.music != nil {
		if !s.music.IsPlaying() {
			s.music.Rewind()
			s.music.Play()
		}
		s.music.SetVolume(volume)
	}
}

// stop the music
func (s *SoundManager) StopMusic() {
	if s.music != nil {
		s.music.Pause()
	}
}

// play requested sounds
func (s *SoundManager) PlaySounds() {
	for sound, play := range s.NextSounds {
		if play {
			s.playSound(sound)
			s.NextSounds[sound] = false
		}
	}
}

// play a sound
func (s SoundManager) playSound(sound int) {
	var soundBytes []byte
	switch sound {
	case SoundSelectID:
		soundBytes = soundSelect
	case SoundMvtID:
		alea := rand.Intn(4)
		switch alea {
		case 0:
			soundBytes = soundMvt1
		case 1:
			soundBytes = soundMvt2
		case 2:
			soundBytes = soundMvt3
		default:
			soundBytes = soundMvt4
		}
	case SoundBeeID:
		soundBytes = soundBee
	case SoundLadyBugID:
		soundBytes = soundLadyBug
	case SoundBallID:
		soundBytes = soundBall
	case SoundBattleID:
		soundBytes = soundBattle
	case SoundExplodeID:
		soundBytes = soundFart
	case SoundNonID:
		soundBytes = soundNon
	case SoundEatID:
		soundBytes = soundEat
	}

	if len(soundBytes) > 0 {
		soundPlayer := s.audioContext.NewPlayerFromBytes(soundBytes)
		soundPlayer.SetVolume(0.15)
		soundPlayer.Play()
	}
}

// decode music and sounds
func InitAudio() (manager SoundManager) {

	var error error
	manager.audioContext = audio.NewContext(44100)

	// music
	soundmp3, error := mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(musicBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ := time.ParseDuration("1m24s010ms")
	duration := tduration.Seconds()
	theBytes := int64(math.Round(duration * 4 * float64(44100)))
	music = audio.NewInfiniteLoop(soundmp3, theBytes)
	manager.music, error = manager.audioContext.NewPlayer(music)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	// sounds
	sound, error := mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundMvt1Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundMvt1, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundMvt2Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundMvt2, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundMvt3Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundMvt3, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundMvt4Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundMvt4, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundSelectBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundSelect, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundBeeBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundBee, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundLadyBugBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundLadyBug, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundBallBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundBall, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundBattleBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundBattle, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundFartBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundFart, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundNonBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundNon, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.DecodeWithSampleRate(manager.audioContext.SampleRate(), bytes.NewReader(soundEatBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundEat, error = io.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	return
}
