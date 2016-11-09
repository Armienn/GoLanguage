package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sound := RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
	sound = RandomSound(false, true)
	fmt.Println(sound.ShortenedRepresentation())
}

type Airstream int

const (
	NoAirstream Airstream = iota
	Egressive
	Ingressive
)

type Initiation int

const (
	NoInitiation Initiation = iota
	Pulmonic
	Glottalic
	Lingual
)

type ObstructionPoint int

const (
	NoObstructionPoint ObstructionPoint = iota
	Labial
	Dental
	Alveolar
	PostAlveolar
	Palatal
	Velar
	Uvular
	Pharyngeal
	Epiglottal
	Glottal
)

type Manner int

const (
	Stop Manner = iota
	Tap
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
)

type Voice int

const (
	Voiceless Voice = iota
	Breathy
	Modal
	Creaky
	Closed
)

type Shape int

const (
	Central Shape = iota
	Lateral
	Sibilant
)

type Articulation struct {
	ObstructionPoint ObstructionPoint
	Manner           Manner
}

type LabialArticulation struct {
	Articulation
	Rounded bool
}

type CoronalArticulation struct {
	Articulation
	Shape Shape
}

type DorsalArticulation struct {
	Articulation
	Centralised bool
}

type RadicalArticulation struct {
	Articulation
}

type GlottalArticulation struct {
	//Articulation
	Voice Voice
}

type Sound struct {
	Airstream  Airstream
	Initiation Initiation

	LabialArticulation  LabialArticulation
	CoronalArticulation CoronalArticulation
	DorsalArticulation  DorsalArticulation
	RadicalArticulation RadicalArticulation
	GlottalArticulation GlottalArticulation
}

func (articulation LabialArticulation) IsValid() bool {
	switch articulation.ObstructionPoint {
	case NoObstructionPoint, Labial, Dental:
		break
	default:
		return false
	}
	switch articulation.Manner {
	case Stop, Tap, Trill, Fricative, Approximant:
		break
	default:
		return false
	}
	return true
}

func (articulation CoronalArticulation) IsValid() bool {
	switch articulation.ObstructionPoint {
	case NoObstructionPoint, Labial, Dental, Alveolar, PostAlveolar, Palatal:
		break
	default:
		return false
	}
	switch articulation.Manner {
	case Stop, Tap, Trill, Fricative, Approximant:
		break
	default:
		return false
	}
	return true
}

func (articulation DorsalArticulation) IsValid() bool {
	switch articulation.ObstructionPoint {
	case NoObstructionPoint, PostAlveolar, Palatal, Velar, Uvular:
		break
	default:
		return false
	}
	switch articulation.Manner {
	case Stop, Tap, Trill, Fricative, Approximant, Close, NearClose, CloseMid, Mid, OpenMid, NearOpen, Open:
		break
	default:
		return false
	}
	if articulation.Centralised {
		if articulation.ObstructionPoint != Palatal && articulation.ObstructionPoint != Uvular {
			return false
		}
	}
	return true
}

func (articulation RadicalArticulation) IsValid() bool {
	switch articulation.ObstructionPoint {
	case NoObstructionPoint, Pharyngeal, Epiglottal:
		break
	default:
		return false
	}
	switch articulation.Manner {
	case Stop, Tap, Trill, Fricative, Approximant:
		break
	default:
		return false
	}
	if articulation.Manner == Tap || articulation.Manner == Trill {
		if articulation.ObstructionPoint != Epiglottal {
			return false
		}
	}
	return true
}

func (articulation GlottalArticulation) IsValid() bool {
	return true
}

func (sound Sound) IsValid() bool {
	if !(sound.LabialArticulation.IsValid() &&
		sound.CoronalArticulation.IsValid() &&
		sound.DorsalArticulation.IsValid() &&
		sound.RadicalArticulation.IsValid() &&
		sound.GlottalArticulation.IsValid()) {
		return false
	}

	if sound.Initiation == Pulmonic {
		//Cannot be have rounded lips(labialisation) while doing linguolabial stuff
		if sound.LabialArticulation.Rounded && sound.CoronalArticulation.ObstructionPoint == Labial {
			return false
		}

		//Cannot have coarticulation of back coronals with front dorsals
		if sound.CoronalArticulation.ObstructionPoint != NoObstructionPoint &&
			sound.DorsalArticulation.ObstructionPoint != NoObstructionPoint {
			if sound.CoronalArticulation.ObstructionPoint == Palatal {
				return false
			}
			if sound.DorsalArticulation.ObstructionPoint == PostAlveolar {
				return false
			}
			if sound.CoronalArticulation.ObstructionPoint == PostAlveolar &&
				sound.DorsalArticulation.ObstructionPoint == Palatal {
				return false
			}
		}
	}
	return true
}

func (sound Sound) Soundless() bool {
	if sound.Airstream == NoAirstream && sound.Initiation == NoInitiation {
		return true
	}
	return false
}

func RandomSound(egressive bool, pulmonic bool) Sound {
	sound := new(Sound)
	sound.randomiseSound(egressive, pulmonic, rand.Intn(5))
	for sound.IsValid() {
		sound.randomiseSound(egressive, pulmonic, rand.Intn(5))
	}
	return *sound
}

func (sound *Sound) randomiseSound(egressive bool, pulmonic bool, part int) {
	if egressive {
		sound.Airstream = Egressive
	} else {
		sound.Airstream = Airstream(rand.Intn(3))
	}
	if pulmonic {
		sound.Initiation = Pulmonic
	} else {
		sound.Initiation = Initiation(rand.Intn(4))
	}
	sound.GlottalArticulation.Voice = Voice(rand.Intn(5))

	switch part {
	case 0:
		sound.LabialArticulation.ObstructionPoint = ObstructionPoint(rand.Intn(11))
		sound.LabialArticulation.Manner = Manner(rand.Intn(12))
		sound.LabialArticulation.Rounded = rand.Intn(2) == 0
	case 1:
		sound.CoronalArticulation.ObstructionPoint = ObstructionPoint(rand.Intn(11))
		sound.CoronalArticulation.Manner = Manner(rand.Intn(12))
		sound.CoronalArticulation.Shape = Shape(rand.Intn(3))
	case 2:
		sound.DorsalArticulation.ObstructionPoint = ObstructionPoint(rand.Intn(11))
		sound.DorsalArticulation.Manner = Manner(rand.Intn(12))
		sound.DorsalArticulation.Centralised = rand.Intn(2) == 0
	case 3:
		sound.RadicalArticulation.ObstructionPoint = ObstructionPoint(rand.Intn(11))
		sound.RadicalArticulation.Manner = Manner(rand.Intn(12))
	case 4:
		sound.GlottalArticulation.Voice = Closed
	}
}

func (sound *Sound) ShortenedRepresentation() string {
	representation := ""
	if sound.Airstream == Ingressive {
		representation += ">"
	}
	if sound.Initiation == Glottalic {
		representation += "?"
	} else if sound.Initiation == Lingual {
		representation += "/"
	}

	if sound.LabialArticulation.ObstructionPoint != NoObstructionPoint {
		representation += "m" + CharacterFromObstructionPoint(sound.LabialArticulation.ObstructionPoint) + CharacterFromManner(sound.LabialArticulation.Manner)
	}
	if sound.CoronalArticulation.ObstructionPoint != NoObstructionPoint {
		representation += "n" + CharacterFromShape(sound.CoronalArticulation.Shape) + CharacterFromObstructionPoint(sound.CoronalArticulation.ObstructionPoint) + CharacterFromManner(sound.CoronalArticulation.Manner)
	}
	if sound.DorsalArticulation.ObstructionPoint != NoObstructionPoint {
		representation += "ŋ" + CharacterFromCentralisation(sound.DorsalArticulation.Centralised) + CharacterFromObstructionPoint(sound.DorsalArticulation.ObstructionPoint) + CharacterFromManner(sound.DorsalArticulation.Manner)
	}
	if sound.RadicalArticulation.ObstructionPoint != NoObstructionPoint {
		representation += "ʕ" + CharacterFromObstructionPoint(sound.RadicalArticulation.ObstructionPoint) + CharacterFromManner(sound.RadicalArticulation.Manner)
	}
	//if (glottalArticulation.Point != ObstructionPoint.None)
	//	representation += "h" + CharacterFrom(glottalArticulation.Point) + CharacterFrom(glottalArticulation.Manner);

	representation += CharacterFromVoice(sound.GlottalArticulation.Voice)
	representation += CharacterFromRoundedness(sound.LabialArticulation.Rounded) //  ̹

	return representation
}

func CharacterFromVoice(voice Voice) string {
	switch voice {
	case Voiceless:
		return "\u0325" // ̥
	case Breathy:
		return "\u0324" // ̤
	case Modal:
		return ""
	case Creaky:
		return "\u0331" // ̰
	case Closed:
		return "ʔ"
	default:
		return "E"
	}
}

func CharacterFromShape(shape Shape) string {
	switch shape {
	case Central:
		return ""
	case Lateral:
		return "\u02E1" //superscript l
	case Sibilant:
		return "\u02E2" //superscript s
	default:
		return "E"
	}
}

func CharacterFromRoundedness(rounded bool) string {
	if rounded {
		return "\u0339"
	}
	return ""
}

func CharacterFromCentralisation(centralised bool) string {
	if centralised {
		return "¨"
	}
	return ""
}

func CharacterFromObstructionPoint(point ObstructionPoint) string {
	switch point {
	case Labial:
		return "l"
	case Dental:
		return "d"
	case Alveolar:
		return "a"
	case PostAlveolar:
		return "o"
	case Palatal:
		return "p"
	case Velar:
		return "v"
	case Uvular:
		return "u"
	case Pharyngeal:
		return "h"
	case Epiglottal:
		return "e"
	case Glottal:
		return "g"
	default:
		return "E"
	} //{ None, Labial, Dental, Alveolar, PostAlveolar, Palatal, Velar, Uvular, Pharyngeal, Epiglottal, Glottal }
}

func CharacterFromManner(manner Manner) string {
	switch manner {
	case Stop:
		return "P"
	case Tap:
		return "T"
	case Trill:
		return "R"
	case Fricative:
		return "F"
	case Approximant:
		return "W"
	case Close:
		return "\"`"
	case NearClose:
		return "``"
	case CloseMid:
		return "`"
	case Mid:
		return "'"
	case OpenMid:
		return "´"
	case NearOpen:
		return "´´"
	case Open:
		return "´\""
	default:
		return "E"
	} //{ Stop, Tap, Trill, Fricative, Approximant, Close, NearClose, CloseMid, Mid, OpenMid, NearOpen, Open }
}
