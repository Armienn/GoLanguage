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

func RandomSyllable(phonetics *Phonetics) Syllable {
	syllable := Syllable{make([]Sound, 0), make([]Sound, 0), make([]Sound, 0)}
	iterations := 0 //for checking that we don't end up in an eternal loop
	//first choose a random pattern concerning nucleus
	pattern := phonetics.Patterns[rand.Intn(len(phonetics.Patterns))]
	for pattern.NucleusPatterns == nil {
		pattern = phonetics.Patterns[rand.Intn(len(phonetics.Patterns))]
	}
	//choose sounds that fit the pattern
	for _, sp := range pattern.NucleusPatterns {
		syllable.NucleusCluster = append(syllable.NucleusCluster, getSound(phonetics, sp))
	}

	//now the onset
	for iterations = 0; iterations < 4 && pattern.OnsetPatterns == nil; iterations++ {
		pattern = phonetics.Patterns[rand.Intn(len(phonetics.Patterns))]
	}
	if pattern.OnsetPatterns != nil {
		//choose sounds that fit the pattern
		for _, sp := range pattern.OnsetPatterns {
			syllable.OnsetCluster = append(syllable.OnsetCluster, getSound(phonetics, sp))
		}
	}

	//and finally the coda
	for iterations = 0; iterations < 4 && pattern.CodaPatterns == nil; iterations++ {
		pattern = phonetics.Patterns[rand.Intn(len(phonetics.Patterns))]
	}
	if pattern.CodaPatterns != nil {
		//choose sounds that fit the pattern
		for _, sp := range pattern.CodaPatterns {
			syllable.CodaCluster = append(syllable.CodaCluster, getSound(phonetics, sp))
		}
	}
	return syllable
}

func getSound(phonetics *Phonetics, sp SoundPattern) Sound {
	sound, ok := getRandomSound(phonetics)
	if !ok {
		return Sound{}
	}
	iterations := 0
	for !sp.Fits(sound) {
		iterations++
		if iterations > 1000 {
			panic("Too many loops! Check your patterns." + fmt.Sprint(sp))
		}
		sound, ok = getRandomSound(phonetics)
	}
	return sound
}

func getRandomSound(phonetics *Phonetics) (sound Sound, ok bool) {
	i := 0
	n := rand.Intn(len(phonetics.Sounds))
	for _, s := range phonetics.Sounds {
		if i == n {
			return s, true
		}
		i++
	}
	return Sound{}, false
}
