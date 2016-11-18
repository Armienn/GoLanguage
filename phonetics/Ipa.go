package phonetics

func BasicPatterns() []SyllablePattern {
	patterns := make([]SyllablePattern, 0)
	pattern := SyllablePattern{}
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
	patterns = append(patterns, pattern)
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
	patterns = append(patterns, pattern)
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
	patterns = append(patterns, pattern)
	return patterns
}

func addInfo(language *Phonetics, sound Sound, point ArticulationPoint, voiced bool, representation string) {
	sound.Point = point
	if voiced {
		sound.Voice = Modal
	} else {
		sound.Voice = Voiceless
	}
	language.Sounds[representation] = sound
}

func addVocalInfo(language *Phonetics, sound Sound, manner ArticulationManner, rounded bool, representation string) {
	sound.Manner = manner
	sound.Rounded = rounded
	language.Sounds[representation] = sound
}

func GetIpa() *Phonetics {
	language := new(Phonetics)
	language.Sounds = make(map[string]Sound)
	language.Patterns = BasicPatterns()

	baseSound := Sound{LabialLabial, Closed, Central, false, true, Voiceless}
	addInfo(language, baseSound, LabialLabial, false, "m̥")
	addInfo(language, baseSound, LabialLabial, true, "m")
	addInfo(language, baseSound, LabialDental, true, "ɱ")
	addInfo(language, baseSound, CoronalLabial, false, "n̼̊")
	addInfo(language, baseSound, CoronalLabial, true, "n̼")
	addInfo(language, baseSound, CoronalAlveolar, false, "n̥")
	addInfo(language, baseSound, CoronalAlveolar, true, "n")
	addInfo(language, baseSound, CoronalRetroflex, false, "ɳ̊")
	addInfo(language, baseSound, CoronalRetroflex, true, "ɳ")
	addInfo(language, baseSound, DorsalPalatal, false, "ɲ̊")
	addInfo(language, baseSound, DorsalPalatal, true, "ɲ")
	addInfo(language, baseSound, DorsalVelar, false, "ŋ̊")
	addInfo(language, baseSound, DorsalVelar, true, "ŋ")
	addInfo(language, baseSound, DorsalUvular, false, "ɴ̥")
	addInfo(language, baseSound, DorsalUvular, true, "ɴ")

	baseSound = Sound{LabialLabial, Stop, Central, false, false, Voiceless}
	addInfo(language, baseSound, LabialLabial, false, "p")
	addInfo(language, baseSound, LabialLabial, true, "b")
	addInfo(language, baseSound, LabialDental, false, "p̪")
	addInfo(language, baseSound, LabialDental, true, "b̪")
	addInfo(language, baseSound, CoronalLabial, false, "t̼")
	addInfo(language, baseSound, CoronalLabial, true, "d̼")
	addInfo(language, baseSound, CoronalAlveolar, false, "t")
	addInfo(language, baseSound, CoronalAlveolar, true, "d")
	addInfo(language, baseSound, CoronalRetroflex, false, "ʈ")
	addInfo(language, baseSound, CoronalRetroflex, true, "ɖ")
	addInfo(language, baseSound, DorsalPalatal, false, "c")
	addInfo(language, baseSound, DorsalPalatal, true, "ɟ")
	addInfo(language, baseSound, DorsalVelar, false, "k")
	addInfo(language, baseSound, DorsalVelar, true, "ɡ")
	addInfo(language, baseSound, DorsalUvular, false, "q")
	addInfo(language, baseSound, DorsalUvular, true, "ɢ")
	addInfo(language, baseSound, RadicalPharyngeal, true, "ʡ")
	addInfo(language, baseSound, Glottal, false, "ʔ")

	baseSound = Sound{LabialLabial, Fricative, Sibilant, false, false, Voiceless}
	addInfo(language, baseSound, CoronalAlveolar, false, "s")
	addInfo(language, baseSound, CoronalAlveolar, true, "z")
	addInfo(language, baseSound, CoronalPostAlveolar, false, "ʃ")
	addInfo(language, baseSound, CoronalPostAlveolar, true, "ʒ")
	addInfo(language, baseSound, CoronalRetroflex, false, "ʂ")
	addInfo(language, baseSound, CoronalRetroflex, true, "ʐ")
	addInfo(language, baseSound, DorsalPostAlveolar, false, "ɕ")
	addInfo(language, baseSound, DorsalPostAlveolar, true, "ʑ")

	baseSound = Sound{LabialLabial, Fricative, Central, false, false, Voiceless}
	addInfo(language, baseSound, LabialLabial, false, "ɸ")
	addInfo(language, baseSound, LabialLabial, true, "β")
	addInfo(language, baseSound, LabialDental, false, "f")
	addInfo(language, baseSound, LabialDental, true, "v")
	addInfo(language, baseSound, CoronalLabial, false, "θ̼")
	addInfo(language, baseSound, CoronalLabial, true, "ð̼")
	addInfo(language, baseSound, CoronalDental, false, "θ")
	addInfo(language, baseSound, CoronalDental, true, "ð")
	addInfo(language, baseSound, CoronalAlveolar, false, "θ̱")
	addInfo(language, baseSound, CoronalAlveolar, true, "ð̠")
	addInfo(language, baseSound, CoronalPostAlveolar, false, "ɹ̠̊˔")
	addInfo(language, baseSound, CoronalPostAlveolar, true, "ɹ̠˔")
	addInfo(language, baseSound, DorsalPalatal, false, "ç")
	addInfo(language, baseSound, DorsalPalatal, true, "ʝ")
	addInfo(language, baseSound, DorsalVelar, false, "x")
	addInfo(language, baseSound, DorsalVelar, true, "ɣ")
	addInfo(language, baseSound, DorsalUvular, false, "χ")
	addInfo(language, baseSound, DorsalUvular, true, "ʁ")
	addInfo(language, baseSound, RadicalPharyngeal, false, "ħ")
	addInfo(language, baseSound, RadicalPharyngeal, true, "ʕ")
	addInfo(language, baseSound, Glottal, false, "h")
	addInfo(language, baseSound, Glottal, true, "ɦ")

	baseSound = Sound{LabialLabial, Approximant, Central, false, false, Voiceless}
	addInfo(language, baseSound, LabialLabial, false, "ɸ")
	addInfo(language, baseSound, LabialLabial, true, "β")
	addInfo(language, baseSound, LabialDental, false, "ʋ̥")
	addInfo(language, baseSound, LabialDental, true, "ʋ")
	addInfo(language, baseSound, CoronalLabial, true, "ð̼")
	addInfo(language, baseSound, CoronalDental, false, "θ")
	addInfo(language, baseSound, CoronalDental, true, "ð")
	addInfo(language, baseSound, CoronalAlveolar, false, "ɹ̥")
	addInfo(language, baseSound, CoronalAlveolar, true, "ɹ")
	addInfo(language, baseSound, CoronalRetroflex, false, "ɻ̊")
	addInfo(language, baseSound, CoronalRetroflex, true, "ɻ")
	addInfo(language, baseSound, DorsalPalatal, false, "j̊")
	addInfo(language, baseSound, DorsalPalatal, true, "j")
	addInfo(language, baseSound, DorsalVelar, false, "ɰ̊")
	addInfo(language, baseSound, DorsalVelar, true, "ɰ")
	addInfo(language, baseSound, DorsalUvular, true, "ʁ")
	addInfo(language, baseSound, RadicalPharyngeal, true, "ʕ")
	addInfo(language, baseSound, Glottal, false, "h")
	addInfo(language, baseSound, Glottal, true, "ʔ̞")

	baseSound = Sound{LabialLabial, Flap, Central, false, false, Voiceless}
	addInfo(language, baseSound, LabialLabial, true, "ⱱ̟")
	addInfo(language, baseSound, LabialDental, true, "ⱱ")
	addInfo(language, baseSound, CoronalLabial, true, "ɾ̼")
	addInfo(language, baseSound, CoronalAlveolar, false, "ɾ̥")
	addInfo(language, baseSound, CoronalAlveolar, true, "ɾ")
	addInfo(language, baseSound, CoronalRetroflex, false, "ɽ̊")
	addInfo(language, baseSound, CoronalRetroflex, true, "ɽ")
	addInfo(language, baseSound, DorsalUvular, true, "ɢ̆")
	addInfo(language, baseSound, RadicalPharyngeal, true, "ʡ̮")

	baseSound = Sound{LabialLabial, Trill, Central, false, false, Voiceless}
	addInfo(language, baseSound, LabialLabial, true, "ʙ")
	addInfo(language, baseSound, CoronalLabial, true, "r̼")
	addInfo(language, baseSound, CoronalAlveolar, false, "r̥")
	addInfo(language, baseSound, CoronalAlveolar, true, "r")
	addInfo(language, baseSound, CoronalRetroflex, false, "ɽr̥")
	addInfo(language, baseSound, CoronalRetroflex, true, "ɽr")
	addInfo(language, baseSound, DorsalUvular, false, "ʀ̥")
	addInfo(language, baseSound, DorsalUvular, true, "ʀ")
	addInfo(language, baseSound, RadicalPharyngeal, false, "ʜ")
	addInfo(language, baseSound, RadicalPharyngeal, true, "ʢ")

	baseSound = Sound{LabialLabial, Fricative, Lateral, false, false, Voiceless}
	addInfo(language, baseSound, CoronalAlveolar, false, "ɬ")
	addInfo(language, baseSound, CoronalAlveolar, true, "ɮ")
	addInfo(language, baseSound, CoronalRetroflex, false, "ɭ̊˔")
	addInfo(language, baseSound, DorsalPalatal, false, "ʎ̥˔")
	addInfo(language, baseSound, DorsalPalatal, true, "ʎ̝")
	addInfo(language, baseSound, DorsalVelar, false, "ʟ̝̊")
	addInfo(language, baseSound, DorsalVelar, true, "ʟ̝")

	baseSound = Sound{LabialLabial, Approximant, Lateral, false, false, Voiceless}
	addInfo(language, baseSound, CoronalLabial, true, "l̼")
	addInfo(language, baseSound, CoronalAlveolar, false, "l̥")
	addInfo(language, baseSound, CoronalAlveolar, true, "l")
	addInfo(language, baseSound, CoronalRetroflex, false, "ɭ̊")
	addInfo(language, baseSound, CoronalRetroflex, true, "ɭ̊")
	addInfo(language, baseSound, DorsalPalatal, false, "ʎ̥")
	addInfo(language, baseSound, DorsalPalatal, true, "ʎ")
	addInfo(language, baseSound, DorsalVelar, false, "ʟ̥")
	addInfo(language, baseSound, DorsalVelar, true, "ʟ")
	addInfo(language, baseSound, DorsalUvular, true, "ʟ̠")

	baseSound = Sound{LabialLabial, Flap, Lateral, false, false, Voiceless}
	addInfo(language, baseSound, CoronalLabial, true, "ɺ̼")
	addInfo(language, baseSound, CoronalAlveolar, true, "ɺ")
	addInfo(language, baseSound, CoronalRetroflex, true, "ɭ̆")
	addInfo(language, baseSound, DorsalPalatal, true, "ʎ̮")
	addInfo(language, baseSound, DorsalVelar, true, "ʟ̆")

	baseSound = Sound{DorsalPalatal, Close, Central, false, false, Modal}
	addVocalInfo(language, baseSound, Close, false, "i")
	addVocalInfo(language, baseSound, Close, true, "y")
	addVocalInfo(language, baseSound, CloseMid, false, "e")
	addVocalInfo(language, baseSound, CloseMid, true, "ø")
	addVocalInfo(language, baseSound, Mid, false, "e̞")
	addVocalInfo(language, baseSound, Mid, true, "ø̞")
	addVocalInfo(language, baseSound, OpenMid, false, "ɛ")
	addVocalInfo(language, baseSound, OpenMid, true, "œ")
	addVocalInfo(language, baseSound, NearOpen, false, "æ")
	addVocalInfo(language, baseSound, Open, false, "a")
	addVocalInfo(language, baseSound, Open, true, "ɶ")

	baseSound = Sound{DorsalVelar, Close, Central, false, false, Modal}
	addVocalInfo(language, baseSound, Close, false, "ɨ")
	addVocalInfo(language, baseSound, Close, true, "ʉ")
	addVocalInfo(language, baseSound, NearClose, false, "ɪ̈")
	addVocalInfo(language, baseSound, NearClose, true, "ʊ̈")
	addVocalInfo(language, baseSound, CloseMid, false, "ɘ")
	addVocalInfo(language, baseSound, CloseMid, true, "ɵ")
	addVocalInfo(language, baseSound, Mid, false, "ə")
	addVocalInfo(language, baseSound, Mid, true, "ɵ̞")
	addVocalInfo(language, baseSound, OpenMid, false, "ɜ")
	addVocalInfo(language, baseSound, OpenMid, true, "ɞ")
	addVocalInfo(language, baseSound, NearOpen, false, "ɐ")
	addVocalInfo(language, baseSound, NearOpen, true, "ɞ̞")
	addVocalInfo(language, baseSound, Open, false, "ä")
	addVocalInfo(language, baseSound, Open, true, "ɒ̈")

	baseSound = Sound{DorsalUvular, Close, Central, false, false, Modal}
	addVocalInfo(language, baseSound, Close, false, "ɯ")
	addVocalInfo(language, baseSound, Close, true, "u")
	addVocalInfo(language, baseSound, CloseMid, false, "ɤ")
	addVocalInfo(language, baseSound, CloseMid, true, "o")
	addVocalInfo(language, baseSound, Mid, false, "ɤ̞")
	addVocalInfo(language, baseSound, Mid, true, "o̞")
	addVocalInfo(language, baseSound, OpenMid, false, "ʌ")
	addVocalInfo(language, baseSound, OpenMid, true, "ɔ")
	addVocalInfo(language, baseSound, Open, false, "ɑ")
	addVocalInfo(language, baseSound, Open, true, "ɒ")

	baseSound = Sound{DorsalPalVel, Close, Central, false, false, Modal}
	addVocalInfo(language, baseSound, NearClose, false, "ɪ")
	addVocalInfo(language, baseSound, NearClose, true, "ʏ")

	baseSound = Sound{DorsalVelUlu, Close, Central, false, false, Modal}
	addVocalInfo(language, baseSound, NearClose, false, "ɯ̽")
	addVocalInfo(language, baseSound, NearClose, true, "ʊ")

	return language
}

func GetDansk() *Phonetics {
	language := new(Phonetics)
	language.Sounds = make(map[string]Sound)
	language.Patterns = make([]SyllablePattern, 0)
	sound := Sound{}
	//M
	sound.Manner = Closed
	sound.Point = LabialLabial
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = true
	sound.Rounded = false
	language.Sounds["m"] = sound
	//N
	sound.Manner = Closed
	sound.Point = CoronalAlveolar
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = true
	sound.Rounded = false
	language.Sounds["n"] = sound
	//NG
	sound.Manner = Closed
	sound.Point = DorsalVelar
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = true
	sound.Rounded = false
	language.Sounds["ŋ"] = sound
	//P
	sound.Manner = Stop
	sound.Point = LabialLabial
	sound.Shape = Central
	sound.Voice = Aspirated
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["p"] = sound
	//B
	sound.Manner = Stop
	sound.Point = LabialLabial
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["b"] = sound
	//T
	sound.Manner = Stop
	sound.Point = CoronalAlveolar
	sound.Shape = Central
	sound.Voice = Aspirated
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["t"] = sound
	//D
	sound.Manner = Stop
	sound.Point = CoronalAlveolar
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["d"] = sound
	//K
	sound.Manner = Stop
	sound.Point = DorsalVelar
	sound.Shape = Central
	sound.Voice = Aspirated
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["k"] = sound
	//G
	sound.Manner = Stop
	sound.Point = DorsalVelar
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["g"] = sound
	//F
	sound.Manner = Fricative
	sound.Point = LabialDental
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["f"] = sound
	//V
	sound.Manner = Approximant
	sound.Point = LabialDental
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["v"] = sound
	//S
	sound.Manner = Fricative
	sound.Point = CoronalAlveolar
	sound.Shape = Sibilant
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["s"] = sound
	//DH
	sound.Manner = Approximant
	sound.Point = CoronalAlveolar
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["ð"] = sound
	//J
	sound.Manner = Approximant
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["j"] = sound
	//L
	sound.Manner = Approximant
	sound.Point = CoronalAlveolar
	sound.Shape = Lateral
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["l"] = sound
	//R
	sound.Manner = Approximant
	sound.Point = DorsalUvular
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["r"] = sound
	//H
	sound.Manner = Fricative
	sound.Point = Glottal
	sound.Shape = Central
	sound.Voice = Voiceless
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["h"] = sound

	//I
	sound.Manner = Close
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["i"] = sound
	//Y
	sound.Manner = Close
	sound.Point = DorsalPalVel
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	language.Sounds["y"] = sound
	//E
	sound.Manner = CloseMid
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["e"] = sound
	//Ø
	sound.Manner = CloseMid
	sound.Point = DorsalPalVel
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	language.Sounds["ø"] = sound
	//Æ
	sound.Manner = OpenMid
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["æ"] = sound
	//schwa
	sound.Manner = Mid
	sound.Point = DorsalVelar
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["ə"] = sound
	//U
	sound.Manner = Close
	sound.Point = DorsalUvular
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	language.Sounds["u"] = sound
	//O
	sound.Manner = CloseMid
	sound.Point = DorsalUvular
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	language.Sounds["o"] = sound
	//Å
	sound.Manner = CloseMid
	sound.Point = DorsalVelUlu
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = true
	language.Sounds["å"] = sound
	//A abe
	sound.Manner = NearOpen
	sound.Point = DorsalPalatal
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["a"] = sound
	//A haj
	sound.Manner = NearOpen
	sound.Point = DorsalVelUlu
	sound.Shape = Central
	sound.Voice = Modal
	sound.Nasal = false
	sound.Rounded = false
	language.Sounds["ɒ"] = sound

	// rules
	language.Patterns = GetDanskePatterns()

	return language
}

func GetDanskePatterns() []SyllablePattern {
	patterns := make([]SyllablePattern, 0)
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
	patterns = append(patterns, pattern)
	// all stops and fricatives can be onset
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Stop,
				Fricative,
			}})
	patterns = append(patterns, pattern)
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
	patterns = append(patterns, pattern)
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
	patterns = append(patterns, pattern)
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
	patterns = append(patterns, pattern)
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
	patterns = append(patterns, pattern)
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
	patterns = append(patterns, pattern)
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
	patterns = append(patterns, pattern)
	return patterns
}

func GetMubPatterns() []SyllablePattern {
	patterns := make([]SyllablePattern, 0)
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
	patterns = append(patterns, pattern)
	// all consonants can be onset, regardless of nucleus or coda
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
	patterns = append(patterns, pattern)
	// all consonants can be coda, regardless of nucleus or onset
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
	patterns = append(patterns, pattern)
	// onset can have clusters of stops+fricatives/Approximant, regardless of nucleus or coda
	pattern = SyllablePattern{}
	pattern.OnsetPatterns = make([]SoundPattern, 0)
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Stop,
			}})
	pattern.OnsetPatterns = append(pattern.OnsetPatterns,
		SoundPattern{
			Manners: []ArticulationManner{
				Fricative,
				Approximant,
			}})
	patterns = append(patterns, pattern)
	return patterns
}
