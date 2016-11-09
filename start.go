package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sound := RandomSound()
	fmt.Println(sound)
}

func RandomSound() Sound {
	sound := Sound{}
	sound.randomiseSound()
	for !sound.IsValid() {
		sound.randomiseSound()
	}
	return sound
}

func (sound *Sound) randomiseSound() {
	sound.Point = ArticulationPoint(rand.Intn(int(ArticulationPointCount)))
	sound.Manner = ArticulationManner(rand.Intn(int(ArticulationMannerCount)))
	sound.Shape = TongueShape(rand.Intn(int(TongueShapeCount)))
	sound.Voice = Voice(rand.Intn(int(VoiceCount)))
	sound.Rounded = rand.Intn(2) == 0
	sound.Nasal = rand.Intn(2) == 0
}

type Sound struct {
	Point   ArticulationPoint
	Manner  ArticulationManner
	Shape   TongueShape
	Rounded bool
	Nasal   bool
	Voice   Voice
}

type ArticulationPoint int

const (
	LabialLabial ArticulationPoint = iota
	LabialDental
	CoronalLabial
	CoronalDental
	CoronalAlveolar
	CoronalPostAlveolar
	CoronalRetroflex
	DorsalPostAlveolar
	DorsalPalatal
	DorsalPalVel
	DorsalVelar
	DorsalVelUlu
	DorsalUvular
	RadicalPharyngeal
	RadicalEpiglottal
	Glottal
	ArticulationPointCount
)

type ArticulationManner int

const (
	Closed ArticulationManner = iota
	Stop
	Flap
	Trill
	Fricative
	Approximant
	Close
	NearClose
	CloseMid
	Mid
	OpenMid
	NearOpen
	Open
	ArticulationMannerCount
)

type TongueShape int

const (
	Central TongueShape = iota
	Lateral
	Sibilant
	TongueShapeCount
)

type Voice int

const (
	Aspirated Voice = iota
	Voiceless
	Breathy
	Modal
	Creaky
	VoiceCount
)

func (sound Sound) IsValid() bool {
	if sound.Rounded {
		switch sound.Point {
		case LabialLabial, LabialDental, CoronalLabial:
			return false
		}
	}

	if sound.Voice == Voiceless {
		//this isn't really invalid, just very unusual
		switch sound.Manner {
		case Close, NearClose, CloseMid, Mid, OpenMid, NearOpen, Open:
			return false
		}
	}

	if sound.Voice == Aspirated {
		switch sound.Manner {
		case Stop:
			break
		default:
			return false
		}
	}

	if sound.Manner == Closed && !sound.Nasal {
		return false
	}

	switch sound.Point {
	case LabialLabial:
		switch sound.Manner {
		case Closed, Stop, Flap, Trill, Fricative, Approximant:
			return true
		default:
			return false
		}
	case LabialDental:
		switch sound.Manner {
		case Closed, Stop, Flap, Trill, Fricative, Approximant:
			return true
		default:
			return false
		}
	case CoronalLabial:
		switch sound.Manner {
		case Closed, Stop, Flap, Trill, Fricative, Approximant:
			return true
		default:
			return false
		}
	case CoronalDental:
		switch sound.Manner {
		case Closed, Stop, Flap, Trill, Fricative, Approximant:
			return true
		default:
			return false
		}
	case CoronalAlveolar:
		switch sound.Manner {
		case Closed, Stop, Flap, Trill, Fricative, Approximant:
			return true
		default:
			return false
		}
	case CoronalPostAlveolar:
		switch sound.Manner {
		case Closed, Stop, Flap, Trill, Fricative, Approximant:
			return true
		default:
			return false
		}
	case CoronalRetroflex:
		switch sound.Manner {
		case Closed, Stop, Flap, Trill, Fricative, Approximant:
			return true
		default:
			return false
		}
	case DorsalPostAlveolar:
		switch sound.Manner {
		case Closed, Stop, Fricative, Approximant:
			return true
		default:
			return false
		}
	case DorsalPalatal:
		switch sound.Manner {
		case Closed, Stop, Fricative, Approximant, Close, NearClose, CloseMid, Mid, OpenMid, NearOpen, Open:
			return true
		default:
			return false
		}
	case DorsalPalVel:
		switch sound.Manner {
		case Closed, Stop, Fricative, Approximant, Close, NearClose, CloseMid, Mid, OpenMid, NearOpen, Open:
			return true
		default:
			return false
		}
	case DorsalVelar:
		switch sound.Manner {
		case Closed, Stop, Fricative, Approximant, Close, NearClose, CloseMid, Mid, OpenMid, NearOpen, Open:
			return true
		default:
			return false
		}
	case DorsalVelUlu:
		switch sound.Manner {
		case Closed, Stop, Fricative, Approximant, Close, NearClose, CloseMid, Mid, OpenMid, NearOpen, Open:
			return true
		default:
			return false
		}
	case DorsalUvular:
		switch sound.Manner {
		case Closed, Stop, Flap, Trill, Fricative, Approximant, Close, NearClose, CloseMid, Mid, OpenMid, NearOpen, Open:
			return true
		default:
			return false
		}
	case RadicalPharyngeal:
		switch sound.Manner {
		case Stop, Flap, Trill, Fricative, Approximant:
			return true
		default:
			return false
		}
	case RadicalEpiglottal:
		switch sound.Manner {
		case Stop, Flap, Trill, Fricative, Approximant:
			return true
		default:
			return false
		}
	case Glottal:
		switch sound.Manner {
		case Stop, Fricative:
			return true
		default:
			return false
		}
	default:
		panic("Weird")
	}
}
