package language

type SyllablePattern struct {
	OnsetPatterns   []SoundPattern
	NucleusPatterns []SoundPattern
	CodaPatterns    []SoundPattern
}

func NewSyllablePattern() *SyllablePattern {
	return &SyllablePattern{make([]SoundPattern, 0), make([]SoundPattern, 0), make([]SoundPattern, 0)}
}
