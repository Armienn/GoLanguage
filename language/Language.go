package language

type Translator interface {
	Translate(*StatementGroup) []Word
}

type Dansk struct {
	Phonetics
	Words map[Concept]DanishWord
}

type DanishWord struct {
	Word
	Ortography string
	OrdKlasse  string
}

type DanishSentence struct {
	Language *Dansk
	Verb     DanishWord
	Subject  DanishWord
	Object   DanishWord
}

func (language *Dansk) Translate(sentence *StatementGroup) ([]Word, string) {
	parsedSentence := ParseSentence(sentence, language)
	return parsedSentence.GetResult()
}

func ParseSentence(baseSentence *StatementGroup, language *Dansk) DanishSentence {
	var sentence DanishSentence
	sentence.Language = language
	sentence.ParseVerb(baseSentence)
	sentence.ParseSubject(baseSentence)
	return sentence
}

func (sentence *DanishSentence) ParseSubject(source *StatementGroup) {
	if source.IsComplex() {
		sentence.ParseComplexSubject(source)
	} else {
		sentence.ParseSimpleSubject(source)
	}
}

func (sentence *DanishSentence) ParseComplexSubject(source *StatementGroup) {
	sentence.Verb = sentence.Language.Words[source.SimpleConcept] // TODO: this is wrong
}

func (sentence *DanishSentence) ParseSimpleSubject(source *StatementGroup) {
	sentence.Verb = sentence.Language.Words[source.SimpleConcept]
	if sentence.Verb.OrdKlasse == "noun" {
		sentence.Object = sentence.Verb
		sentence.Verb = DanishWord{}
		sentence.Verb.Ortography = "er"
	}
}

func (sentence *DanishSentence) ParseVerb(source *StatementGroup) {
	if source.IsComplex() {
		sentence.ParseComplexVerb(source)
	} else {
		sentence.ParseSimpleVerb(source)
	}
}

func (sentence *DanishSentence) ParseComplexVerb(source *StatementGroup) {
	sentence.Verb = sentence.Language.Words[source.SimpleConcept] // TODO: this is wrong
}

func (sentence *DanishSentence) ParseSimpleVerb(source *StatementGroup) {
	sentence.Verb = sentence.Language.Words[source.SimpleConcept]
	if sentence.Verb.OrdKlasse == "noun" {
		sentence.Object = sentence.Verb
		sentence.Verb = DanishWord{}
		sentence.Verb.Ortography = "er"
	}
}

func (sentence *DanishSentence) GetResult() ([]Word, string) {
	words := []Word{}
	text := ""
	return words, text
}
