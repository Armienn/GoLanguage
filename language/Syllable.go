package language

import "math/rand"

type Syllable struct {
	OnsetCluster   []Sound
	NucleusCluster []Sound
	CodaCluster    []Sound
}

func RandomSyllable(language *Language) Syllable {
	syllable := Syllable{make([]Sound, 0), make([]Sound, 0), make([]Sound, 0)}
	iterations := 0 //for checking that we don't end up in an eternal loop
	//first choose a random pattern concerning nucleus
	pattern := language.Patterns[rand.Intn(len(language.Patterns))]
	for pattern.NucleusPatterns == nil {
		pattern = language.Patterns[rand.Intn(len(language.Patterns))]
	}
	//choose sounds that fit the pattern
	var info SoundInformation
	for _, sp := range pattern.NucleusPatterns {
		info = language.Sounds[rand.Intn(len(language.Sounds))]
		iterations = 0
		for !sp.Fits(info.Sound) {
			iterations++
			if iterations > 1000 {
				panic("Too many loops! Check your patterns.")
			}
			info = language.Sounds[rand.Intn(len(language.Sounds))]
		}
		syllable.NucleusCluster = append(syllable.NucleusCluster, info.Sound)
	}

	//now the onset
	if pattern.OnsetPatterns == nil {
		for iterations = 0; iterations < 4 && pattern.OnsetPatterns == nil; iterations++ {
			pattern = language.Patterns[rand.Intn(len(language.Patterns))]
		}
	}
	if pattern.OnsetPatterns != nil {
		//choose sounds that fit the pattern
		for _, sp := range pattern.OnsetPatterns {
			info = language.Sounds[rand.Intn(len(language.Sounds))]
			iterations = 0
			for !sp.Fits(info.Sound) {
				iterations++
				if iterations > 1000 {
					panic("Too many loops! Check your patterns.")
				}
				info = language.Sounds[rand.Intn(len(language.Sounds))]
			}
			syllable.OnsetCluster = append(syllable.OnsetCluster, info.Sound)
		}
	}

	//and finally the coda
	if pattern.CodaPatterns == nil {
		for iterations = 0; iterations < 4 && pattern.CodaPatterns == nil; iterations++ {
			pattern = language.Patterns[rand.Intn(len(language.Patterns))]
		}
	}
	if pattern.CodaPatterns != nil {
		//choose sounds that fit the pattern
		for _, sp := range pattern.CodaPatterns {
			info = language.Sounds[rand.Intn(len(language.Sounds))]
			iterations = 0
			for !sp.Fits(info.Sound) {
				iterations++
				if iterations > 1000 {
					panic("Too many loops! Check your patterns.")
				}
				info = language.Sounds[rand.Intn(len(language.Sounds))]
			}
			syllable.CodaCluster = append(syllable.CodaCluster, info.Sound)
		}
	}
	return syllable
}
