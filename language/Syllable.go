package language

import (
	"fmt"
	"math/rand"
)

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
	for _, sp := range pattern.NucleusPatterns {
		syllable.NucleusCluster = append(syllable.NucleusCluster, getSound(language, sp))
	}

	//now the onset
	for iterations = 0; iterations < 4 && pattern.OnsetPatterns == nil; iterations++ {
		pattern = language.Patterns[rand.Intn(len(language.Patterns))]
	}
	if pattern.OnsetPatterns != nil {
		//choose sounds that fit the pattern
		for _, sp := range pattern.OnsetPatterns {
			syllable.OnsetCluster = append(syllable.OnsetCluster, getSound(language, sp))
		}
	}

	//and finally the coda
	for iterations = 0; iterations < 4 && pattern.CodaPatterns == nil; iterations++ {
		pattern = language.Patterns[rand.Intn(len(language.Patterns))]
	}
	if pattern.CodaPatterns != nil {
		//choose sounds that fit the pattern
		for _, sp := range pattern.CodaPatterns {
			syllable.CodaCluster = append(syllable.CodaCluster, getSound(language, sp))
		}
	}
	return syllable
}

func getSound(language *Language, sp SoundPattern) Sound {
	//info := language.Sounds[rand.Intn(len(language.Sounds))]
	sound, ok := getRandomSound(language)
	if !ok {
		return Sound{}
	}
	iterations := 0
	for !sp.Fits(sound) {
		iterations++
		if iterations > 1000 {
			panic("Too many loops! Check your patterns." + fmt.Sprint(sp))
		}
		sound, ok = getRandomSound(language)
	}
	return sound
}

func getRandomSound(language *Language) (sound Sound, ok bool) {
	i := 0
	n := rand.Intn(len(language.Sounds))
	for _, s := range language.Sounds {
		if i == n {
			return s, true
		}
		i++
	}
	return Sound{}, false
}
