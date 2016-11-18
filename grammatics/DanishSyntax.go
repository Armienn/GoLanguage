package grammatics

type DanishWordRepresenter interface {
	NavneordRepresentation(*Navneord) []WordRepresenter
	UdsagnsordRepresentation(*Udsagnsord) []WordRepresenter
}

type Dansk struct {
	Words   map[Concept]DanishWordRepresenter
	Er      DanishWordRepresenter
	Missing DanishWordRepresenter
}

type MissingWord struct{}

func (word *MissingWord) NavneordRepresentation(ord *Navneord) []WordRepresenter {
	return []WordRepresenter{wordString{"missing"}}
}

func (word *MissingWord) UdsagnsordRepresentation(ord *Udsagnsord) []WordRepresenter {
	return []WordRepresenter{wordString{"missing"}}
}

type wordString struct {
	word string
}

func (word wordString) Representation() interface{} {
	return word.word
}

type Udsagnsord struct {
	Ord DanishWordRepresenter
	Tid string
}

type Navneord struct {
	Ord        DanishWordRepresenter
	Flertal    bool
	Bestemt    bool
	Tillægsord []Tillægsord
}

type Tillægsord struct {
	Ord *DanishWordRepresenter
}

type DanishSentence struct {
	core     map[Concept]*ConceptInfo
	Language *Dansk
	Verb     Udsagnsord
	Subject  Navneord
	Object   Navneord
}

func NewDanishSentence(language *Dansk) DanishSentence {
	var sentence DanishSentence
	sentence.core = GetCoreLanguage()
	sentence.Language = language
	return sentence
}

func (language *Dansk) Translate(sentence *Statement) []WordRepresenter {
	parsedSentence := language.ParseSentence(sentence)
	return parsedSentence.GetResult()
}

func (language *Dansk) ParseSentence(baseSentence *Statement) DanishSentence {
	var sentence DanishSentence
	sentence.Language = language
	sentence.ParseVerb(baseSentence)
	sentence.ParseSubject(baseSentence)
	return sentence
}

func (sentence *DanishSentence) ParseSubject(source *Statement) {
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

func (sentence *DanishSentence) ParseComplexSubject(source *Statement) {
	ok := false
	sentence.Subject = Navneord{}
	sentence.Subject.Ord, ok = sentence.Language.Words[source.SimpleConcept]
	if !ok {
		sentence.Subject.Ord = sentence.Language.Missing
	}
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) ParseSimpleSubject(source *Statement) {
	ok := false
	sentence.Subject = Navneord{}
	sentence.Subject.Ord, ok = sentence.Language.Words[source.SimpleConcept]
	if !ok {
		sentence.Subject.Ord = sentence.Language.Missing
	}
	sentence.Subject.Bestemt = 0 < len(source.GetDescriptorsOf("definite", "descriptor"))
	sentence.Subject.Flertal = false
	amounts := source.GetDescriptors("amount")
	if len(amounts) > 0 && amounts[0].SimpleConcept != "one" {
		sentence.Subject.Flertal = true
	}
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) ParseDescriptors(source *Statement) {
	//do something
}

func (sentence *DanishSentence) ParseVerb(source *Statement) {
	if source.IsComplex() {
		sentence.ParseComplexVerb(source)
	} else {
		sentence.ParseSimpleVerb(source)
	}
}

func (sentence *DanishSentence) ParseComplexVerb(source *Statement) {
	ok := false
	sentence.Verb = Udsagnsord{}
	sentence.Verb.Ord = sentence.Language.Words[source.SimpleConcept] // TODO: this is wrong
	if !ok {
		sentence.Verb.Ord = sentence.Language.Missing
	}
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) ParseSimpleVerb(source *Statement) {
	word, ok := sentence.Language.Words[source.SimpleConcept]
	if !ok {
		word = sentence.Language.Missing
	}
	sentence.Verb = Udsagnsord{}
	if len(source.GetDescriptors("doer")) > 0 {
		sentence.Verb.Ord = word
	} else if len(source.GetDescriptors("beer")) > 0 {
		sentence.Verb.Ord = sentence.Language.Er
		sentence.Object = Navneord{}
		sentence.Object.Ord = word
	} else {
		//uh
		sentence.Verb.Ord = sentence.Language.Missing
	}
	sentence.Verb.Tid = sentence.GetTime(source)
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) GetTime(source *Statement) string {
	if len(source.GetDescriptorsOf("before", "at")) > 0 {
		return "datid"
	}
	return "nutid"
}

func (sentence *DanishSentence) GetResult() []WordRepresenter {
	words := make([]WordRepresenter, 0)
	if sentence.Subject.Ord != nil {
		words = append(words, sentence.Subject.Ord.NavneordRepresentation(&sentence.Subject)...)
	}
	words = append(words, sentence.Verb.Ord.UdsagnsordRepresentation(&sentence.Verb)...)
	return words
}
