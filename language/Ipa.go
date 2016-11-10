package language

func addInfo(language *Language, baseSound Sound, point ArticulationPoint, voiced bool, representation string) {
	info := SoundInformation{}
	info.Sound = baseSound
	info.Sound.Point = point
	if voiced {
		info.Sound.Voice = Modal
	} else {
		info.Sound.Voice = Voiceless
	}
	info.Representation = representation
	language.Sounds = append(language.Sounds, info)
}

func addVocalInfo(language *Language, baseSound Sound, manner ArticulationManner, rounded bool, representation string) {
	info := SoundInformation{}
	info.Sound = baseSound
	info.Sound.Manner = manner
	info.Sound.Rounded = rounded
	info.Representation = representation
	language.Sounds = append(language.Sounds, info)
}

func GetIpa() *Language {
	language := new(Language)
	language.Sounds = make([]SoundInformation, 0)
	language.Patterns = make([]SyllablePattern, 0)

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

	// rules
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
