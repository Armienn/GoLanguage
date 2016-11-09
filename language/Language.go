package language

import "math/rand"
import "fmt"

type Language struct {
	Sounds   []SoundInformation
	Patterns []SyllablePattern
}

func (language *Language) GetRepresentation(word Word) string {
	result := ""
	for _, syllable := range word.Syllables {
		for _, sound := range syllable.OnsetCluster {
			for _, info := range language.Sounds {
				if info.Sound == sound {
					result += info.Representation
					break
				}
			}
		}
		for _, sound := range syllable.NucleusCluster {
			for _, info := range language.Sounds {
				if info.Sound == sound {
					result += info.Representation
					break
				}
			}
		}
		for _, sound := range syllable.CodaCluster {
			for _, info := range language.Sounds {
				if info.Sound == sound {
					result += info.Representation
					break
				}
			}
		}
	}
	return result
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

	soundCount := rand.Intn(20)
	for i := 0; i < soundCount; i++ {
		info.Sound = RandomSound()
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

	return language
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
