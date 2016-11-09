package language

type Syllable struct {
	OnsetCluster   []Sound
	NucleusCluster []Sound
	CodaCluster    []Sound
}

func RandomSyllable(language *Language) Syllable {
	return Syllable{}
}
