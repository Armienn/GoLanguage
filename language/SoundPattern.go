package language

import "fmt"

func main() {
	sound := RandomSound()
	fmt.Println(sound)
}

type TristateBool byte

const (
	TristateFalse TristateBool = iota
	TristateTrue
	TristateUnknown
)

type SoundPattern struct {
	Points  []ArticulationPoint
	Manners []ArticulationManner
	Shapes  []TongueShape
	Rounded TristateBool
	Nasal   TristateBool
	Voices  []Voice
}

func (pattern *SoundPattern) Fits(sound Sound) bool {
	if pattern.Points != nil && !containsPoint(pattern.Points, sound.Point) {
		return false
	}
	if pattern.Manners != nil && !containsManner(pattern.Manners, sound.Manner) {
		return false
	}
	if pattern.Shapes != nil && !containsShape(pattern.Shapes, sound.Shape) {
		return false
	}
	if pattern.Voices != nil && !containsVoice(pattern.Voices, sound.Voice) {
		return false
	}
	if pattern.Rounded != TristateUnknown && pattern.Rounded.ToBool() != sound.Rounded {
		return false
	}
	if pattern.Nasal != TristateUnknown && pattern.Nasal.ToBool() != sound.Nasal {
		return false
	}
	return true
}

func containsPoint(s []ArticulationPoint, e ArticulationPoint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsManner(s []ArticulationManner, e ArticulationManner) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsShape(s []TongueShape, e TongueShape) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsVoice(s []Voice, e Voice) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (tristate TristateBool) ToBool() bool {
	if tristate == TristateTrue {
		return true
	}
	return false
}
