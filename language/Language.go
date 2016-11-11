package language

import "math/rand"
import "fmt"

type Language struct {
	Sounds   []SoundInformation
	Patterns []SyllablePattern
}

func (language *Language) GetWordRepresentation(word Word) string {
	result := ""
	sounds := word.GetSounds()
	for _, sound := range sounds {
		result += language.GetRepresentation(sound)
	}
	return result
}

func (language *Language) GetRepresentation(sound Sound) string {
	minDistance := 10000
	bestInfo := SoundInformation{}
	for _, info := range language.Sounds {
		distance := Distance(info.Sound, sound)
		if distance < minDistance {
			minDistance = distance
			bestInfo = info
		}
	}
	return bestInfo.Representation
}

func (language *Language) RandomWord(syllables int) Word {
	return RandomWord(language, syllables)
}

func RandomPhoneticInventory() []SoundInformation {
	result := make([]SoundInformation, 0)
	info := SoundInformation{}
	basePoints := randomPoints()
	baseVoices := randomVoices()
	sounds := make([]Sound, 0)

	sounds = append(sounds, randomSoundSet(basePoints, baseVoices, Closed)...)
	sounds = append(sounds, randomSoundSet(basePoints, baseVoices, Stop)...)
	sounds = append(sounds, randomSoundSet(basePoints, baseVoices, Fricative)...)
	if rand.Intn(10) < 2 {
		sounds = append(sounds, randomSoundSet(basePoints, baseVoices, Flap)...)
	}
	if rand.Intn(10) < 2 {
		sounds = append(sounds, randomSoundSet(basePoints, baseVoices, Trill)...)
	}
	if rand.Intn(10) < 2 {
		sounds = append(sounds, randomSoundSet(basePoints, baseVoices, Approximant)...)
	}
	sounds = append(sounds, randomVocalSet(baseVoices)...)

	for i := 0; i < 10; i++ {
		sounds = mutateSoundSet(sounds)
	}
	for i := range sounds {
		sounds[i].Standardise()
		info.Sound = sounds[i]
		info.Representation = "s" + fmt.Sprint(i)
		result = append(result, info)
	}
	return result
}

func RandomLanguage() *Language {
	language := new(Language)
	language.Sounds = RandomPhoneticInventory()
	language.Patterns = BasicPatterns()
	return language
}

func randomPoints() []ArticulationPoint {
	pointCount := rand.Intn(5) + 1
	points := make([]ArticulationPoint, pointCount)
	for i := 0; i < pointCount; i++ {
		newPoint := semiRandomPoint()
		valid := true
		for j := 0; j < i; j++ {
			difference := newPoint - points[j]
			if -2 < difference && difference < 2 {
				valid = false
				break
			}
		}
		if valid {
			points[i] = newPoint
		} else {
			i--
		}
	}
	return points
}

func semiRandomPoint() ArticulationPoint {
	newPoint := ArticulationPoint(rand.Intn(int(ArticulationPointCount)))
	if rand.Intn(10) < 2 {
		newPoint = CoronalAlveolar
	} else if rand.Intn(10) < 2 {
		newPoint = LabialLabial
	} else if rand.Intn(10) < 2 {
		newPoint = DorsalVelar
	}
	return newPoint
}

func randomVoices() []Voice {
	voices := make([]Voice, 0)
	if rand.Intn(10) < 2 {
		voices = append(voices, Breathy)
	}
	if rand.Intn(10) < 2 {
		voices = append(voices, Creaky)
	}
	if rand.Intn(2) == 0 {
		voices = append(voices, Voiceless)
		if rand.Intn(5)-len(voices) > 0 {
			voices = append(voices, Modal)
		}
	} else {
		voices = append(voices, Modal)
		if rand.Intn(5)-len(voices) > 0 {
			voices = append(voices, Voiceless)
		}
	}
	return voices
}

func chooseRandomVoices(voices []Voice) []Voice {
	chosenVoices := make([]Voice, 0)
	for _, voice := range voices {
		if rand.Intn(2) == 0 {
			chosenVoices = append(chosenVoices, voice)
		}
	}
	if len(chosenVoices) == 0 {
		chosenVoices = append(chosenVoices, voices[rand.Intn(len(voices))])
	}
	return chosenVoices
}

func randomSoundSet(points []ArticulationPoint, voices []Voice, manner ArticulationManner) []Sound {
	chosenVoices := chooseRandomVoices(voices)
	sounds := make([]Sound, len(chosenVoices)*len(points))
	shape := TongueShape(rand.Intn(int(TongueShapeCount)))
	nasal := false
	if manner == Closed {
		nasal = true
	}
	rounded := false
	for i := 0; i < len(chosenVoices); i++ {
		for j := 0; j < len(points); j++ {
			sound := Sound{points[j], manner, shape, rounded, nasal, chosenVoices[i]}
			sounds[i*len(points)+j] = sound
		}
	}
	return sounds
}

func randomVocalSet(voices []Voice) []Sound {
	chosenVoices := make([]Voice, 1)
	chosenVoices[0] = Modal
	for _, voice := range voices {
		if rand.Intn(2) == 0 && (voice == Breathy || voice == Creaky) {
			chosenVoices = append(chosenVoices, voice)
		}
	}

	sounds := make([]Sound, 5*len(chosenVoices))
	for i, voice := range chosenVoices {
		sounds[i*5] = Sound{DorsalPalatal, Close, Central, rand.Intn(4) == 0, false, voice}
		sounds[i*5+1] = Sound{DorsalPalVel, Mid, Central, rand.Intn(4) == 0, false, voice}
		sounds[i*5+2] = Sound{DorsalVelar, Open, Central, rand.Intn(4) == 0, false, voice}
		sounds[i*5+3] = Sound{DorsalVelUlu, Mid, Central, rand.Intn(4) == 0, false, voice}
		sounds[i*5+4] = Sound{DorsalUvular, Close, Central, rand.Intn(4) == 0, false, voice}
	}
	return sounds
}

func mutateSoundSet(sounds []Sound) []Sound {
	index := rand.Intn(len(sounds))
	if rand.Intn(2) == 0 {
		return append(sounds[:index], sounds[index+1:]...)
	}
	sound := sounds[index]
	if rand.Intn(2) == 0 {
		change := ArticulationManner(rand.Intn(2)*2 - 1)
		sound.Manner = sound.Manner + change
		if sound.Manner == ArticulationMannerCount {
			sound.Manner = sound.Manner - change*2
		}
	}
	if rand.Intn(2) == 0 {
		change := ArticulationPoint(rand.Intn(2)*2 - 1)
		sound.Point = sound.Point + change
		if sound.Point == ArticulationPointCount {
			sound.Point = sound.Point - change*2
		}
	}
	return append(sounds, sound)
}
