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
	parsedSentence := language.ParseSentence(sentence)
	return parsedSentence.GetResult()
}

func (language *Dansk) ParseSentence(baseSentence *StatementGroup) DanishSentence {
	var sentence DanishSentence
	sentence.Language = language
	sentence.ParseVerb(baseSentence)
	sentence.ParseSubject(baseSentence)
	return sentence
}

func (sentence *DanishSentence) ParseSubject(source *StatementGroup) {
	subjects := source.GetDescriptors("doer")
	if len(subjects) == 0 {
		subjects = source.GetDescriptors("beer")
	}
	for _, subject := range subjects {
		if subject.IsComplex() {
			sentence.ParseComplexSubject(subject)
		} else {
			sentence.ParseSimpleSubject(subject)
		}
	}
}

func (sentence *DanishSentence) ParseComplexSubject(source *StatementGroup) {
	sentence.Subject = sentence.Language.Words[source.SimpleConcept] // TODO: this is wrong
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) ParseSimpleSubject(source *StatementGroup) {
	sentence.Subject = sentence.Language.Words[source.SimpleConcept]
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) ParseDescriptors(source *StatementGroup) {
	//do something
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
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) ParseSimpleVerb(source *StatementGroup) {
	sentence.Verb = sentence.Language.Words[source.SimpleConcept]
	if sentence.Verb.OrdKlasse == "noun" {
		sentence.Object = sentence.Verb
		sentence.Verb = DanishWord{}
		sentence.Verb.Ortography = "er"
	}
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) GetResult() ([]Word, string) {
	words := []Word{}
	text := sentence.Subject.Ortography + " " + sentence.Verb.Ortography
	return words, text
}
