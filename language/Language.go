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

func RandomLanguage() *Language {
	language := new(Language)
	language.Sounds = make([]SoundInformation, 0)
	language.Patterns = make([]SyllablePattern, 0)

	info := SoundInformation{}
	pattern := SyllablePattern{}

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
		language.Sounds = append(language.Sounds, info)
	}

	/*patternCount := rand.Intn(8)
	for i := 0; i < patternCount; i++ {
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Stop,
				Fricative,
			}})
	language.Patterns = append(language.Patterns, pattern)
	}*/
	// all vowels can be nucleus, regardless of onset or coda
	pattern.NucleusPatterns = make([]SoundPattern, 0)
	pattern.NucleusPatterns = append(pattern.NucleusPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Close,
				NearClose,
				CloseMid,
				Mid,
				OpenMid,
				NearOpen,
				Open,
			}})
	language.Patterns = append(language.Patterns, pattern)
	pattern = SyllablePattern{}
	pattern.CodaPatterns = make([]SoundPattern, 0)
	pattern.CodaPatterns = append(pattern.CodaPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Closed,
				Stop,
				Flap,
				Trill,
				Fricative,
				Approximant,
			}})
	language.Patterns = append(language.Patterns, pattern)
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Closed,
				Stop,
				Flap,
				Trill,
				Fricative,
				Approximant,
			}})
	language.Patterns = append(language.Patterns, pattern)

	return language
}

func randomPoints() []ArticulationPoint {
	pointCount := rand.Intn(4) + 1
	points := make([]ArticulationPoint, pointCount)
	for i := 0; i < pointCount; i++ {
		newPoint := ArticulationPoint(rand.Intn(int(ArticulationPointCount)))
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

func GetDansk() *Language {
	language := new(Language)
	language.Sounds = make([]SoundInformation, 0)
	language.Patterns = make([]SyllablePattern, 0)
	info := SoundInformation{}
	sound := Sound{}
	//M
	sound.Manner = Closed
	sound.Point = LabialLabial
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = true
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "m"
	language.Sounds = append(language.Sounds, info)
	//N
	sound.Manner = Closed
	sound.Point = CoronalAlveolar
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = true
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "n"
	language.Sounds = append(language.Sounds, info)
	//NG
	sound.Manner = Closed
	sound.Point = DorsalVelar
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = true
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "ŋ"
	language.Sounds = append(language.Sounds, info)
	//P
	sound.Manner = Stop
	sound.Point = LabialLabial
	sound.Shape = Central
	sound.Voice = Aspirated
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "p"
	language.Sounds = append(language.Sounds, info)
	//B
	sound.Manner = Stop
	sound.Point = LabialLabial
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "b"
	language.Sounds = append(language.Sounds, info)
	//T
	sound.Manner = Stop
	sound.Point = CoronalAlveolar
	sound.Shape = Central
	sound.Voice = Aspirated
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "t"
	language.Sounds = append(language.Sounds, info)
	//D
	sound.Manner = Stop
	sound.Point = CoronalAlveolar
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "d"
	language.Sounds = append(language.Sounds, info)
	//K
	sound.Manner = Stop
	sound.Point = DorsalVelar
	sound.Shape = Central
	sound.Voice = Aspirated
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "k"
	language.Sounds = append(language.Sounds, info)
	//G
	sound.Manner = Stop
	sound.Point = DorsalVelar
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "g"
	language.Sounds = append(language.Sounds, info)
	//F
	sound.Manner = Fricative
	sound.Point = LabialDental
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "f"
	language.Sounds = append(language.Sounds, info)
	//V
	sound.Manner = Approximant
	sound.Point = LabialDental
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "v"
	language.Sounds = append(language.Sounds, info)
	//S
	sound.Manner = Fricative
	sound.Point = CoronalAlveolar
	sound.Shape = Sibilant
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "s"
	language.Sounds = append(language.Sounds, info)
	//DH
	sound.Manner = Approximant
	sound.Point = CoronalAlveolar
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "ð"
	language.Sounds = append(language.Sounds, info)
	//J
	sound.Manner = Approximant
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "j"
	language.Sounds = append(language.Sounds, info)
	//L
	sound.Manner = Approximant
	sound.Point = CoronalAlveolar
	sound.Shape = Lateral
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "l"
	language.Sounds = append(language.Sounds, info)
	//R
	sound.Manner = Approximant
	sound.Point = DorsalUvular
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "r"
	language.Sounds = append(language.Sounds, info)
	//H
	sound.Manner = Fricative
	sound.Point = Glottal
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "h"
	language.Sounds = append(language.Sounds, info)

	//I
	sound.Manner = Close
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "i"
	language.Sounds = append(language.Sounds, info)
	//Y
	sound.Manner = Close
	sound.Point = DorsalPalVel
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	info.Sound = sound
	info.Representation = "y"
	language.Sounds = append(language.Sounds, info)
	//E
	sound.Manner = CloseMid
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "e"
	language.Sounds = append(language.Sounds, info)
	//Ø
	sound.Manner = CloseMid
	sound.Point = DorsalPalVel
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	info.Sound = sound
	info.Representation = "ø"
	language.Sounds = append(language.Sounds, info)
	//Æ
	sound.Manner = OpenMid
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "æ"
	language.Sounds = append(language.Sounds, info)
	//schwa
	sound.Manner = Mid
	sound.Point = DorsalVelar
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "ə"
	language.Sounds = append(language.Sounds, info)
	//U
	sound.Manner = Close
	sound.Point = DorsalUvular
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	info.Sound = sound
	info.Representation = "u"
	language.Sounds = append(language.Sounds, info)
	//O
	sound.Manner = CloseMid
	sound.Point = DorsalUvular
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	info.Sound = sound
	info.Representation = "o"
	language.Sounds = append(language.Sounds, info)
	//Å
	sound.Manner = CloseMid
	sound.Point = DorsalVelUlu
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	info.Sound = sound
	info.Representation = "å"
	language.Sounds = append(language.Sounds, info)
	//A abe
	sound.Manner = NearOpen
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "a"
	language.Sounds = append(language.Sounds, info)
	//A haj
	sound.Manner = NearOpen
	sound.Point = DorsalVelUlu
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	info.Sound = sound
	info.Representation = "ɒ"
	language.Sounds = append(language.Sounds, info)

	// rules
	pattern := SyllablePattern{}
	// all vowels can be nucleus, regardless of onset or coda
	pattern.NucleusPatterns = make([]SoundPattern, 0)
	pattern.NucleusPatterns = append(pattern.NucleusPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Close,
				NearClose,
				CloseMid,
				Mid,
				OpenMid,
				NearOpen,
				Open,
			}})
	language.Patterns = append(language.Patterns, pattern)
	// all stops and fricatives can be onset
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Stop,
				Fricative,
			}})
	language.Patterns = append(language.Patterns, pattern)
	// n and m can be onset
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Closed,
			},
			Points: []ArticulationPoint{
				LabialLabial,
				CoronalAlveolar,
			}})
	language.Patterns = append(language.Patterns, pattern)
	// v, j and r can be onset
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Approximant,
			},
			Points: []ArticulationPoint{
				LabialDental,
				DorsalPalatal,
				DorsalUvular,
			}})
	language.Patterns = append(language.Patterns, pattern)
	// l can be onset
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Approximant,
			},
			Points: []ArticulationPoint{
				CoronalAlveolar,
			},
			Shapes: []TongueShape{
				Lateral,
			}})
	language.Patterns = append(language.Patterns, pattern)
	// all consonants can be coda, regardless of nucleus or onset
	pattern = SyllablePattern{}
	pattern.CodaPatterns = make([]SoundPattern, 0)
	pattern.CodaPatterns = append(pattern.CodaPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Closed,
				Stop,
				Fricative,
				Approximant,
			}})
	language.Patterns = append(language.Patterns, pattern)
	// onset can have clusters of s+(unaspirated)stops, regardless of nucleus or coda
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Fricative,
			},
			Shapes: []TongueShape{
				Sibilant,
			}})
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Stop,
			},
			Voices: []Voice{
				Voiceless,
			}})
	language.Patterns = append(language.Patterns, pattern)
	// onset can have clusters of s+(unaspirated)stops+r, regardless of nucleus or coda
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Fricative,
			},
			Shapes: []TongueShape{
				Sibilant,
			}})
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Stop,
			},
			Voices: []Voice{
				Voiceless,
			}})
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Approximant,
			},
			Points: []ArticulationPoint{
				DorsalUvular,
			}})
	language.Patterns = append(language.Patterns, pattern)

	return language
}
