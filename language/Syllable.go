package language

type Syllable struct {
	OnsetCluster   []Sound
	NucleusCluster []Sound
	CodaCluster    []Sound
}

func RandomSyllable() Syllable {
	return Syllable{}
}
