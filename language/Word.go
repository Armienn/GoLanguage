package language

import "math/rand"

type Word struct {
	Syllables []Syllable
}

func RandomWord(language *Language, syllables int) Word {
	if syllables == 0 {
		syllables = rand.Intn(3) + 1
	}
	word := Word{make([]Syllable, syllables)}
	for i := 0; i < syllables; i++ {
		word.Syllables[i] = RandomSyllable(language)
	}
	return word
}
