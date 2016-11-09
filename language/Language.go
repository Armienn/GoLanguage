package language

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
