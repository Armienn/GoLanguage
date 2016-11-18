package language

import (
	g "github.com/Armienn/GoLanguage/grammatics"
	p "github.com/Armienn/GoLanguage/phonetics"
)

type Translator interface {
	Translate(*g.StatementGroup) []p.Word
}

type Dansk struct {
	p.Phonetics
	Words      map[g.Concept]*DanishWord
	BasicWords map[string]*DanishWord
}

type DanishWord struct {
	p.Word
	Ortography string
	OrdKlasse  string
}

type Udsagnsord struct {
	Ord *DanishWord
	Tid string
}

type Navneord struct {
	Ord        *DanishWord
	Flertal    bool
	Bestemt    bool
	Tillægsord []Tillægsord
}

type Tillægsord struct {
	Ord *DanishWord
}

type DanishSentence struct {
	core     map[g.Concept]*g.ConceptInfo
	Language *Dansk
	Verb     Udsagnsord
	Subject  Navneord
	Object   Navneord
}

func (ord *Navneord) GetText() string {
	if ord.Bestemt {
		return ord.Ord.Ortography + "en"
	}
	return "en " + ord.Ord.Ortography
}

func (ord *Udsagnsord) GetText() string {
	if ord.Tid == "nutid" {
		return ord.Ord.Ortography + "r"
	}
	return ord.Ord.Ortography + "de"
}

func NewDanishSentence(language *Dansk) DanishSentence {
	var sentence DanishSentence
	sentence.core = g.GetCoreLanguage()
	sentence.Language = language
	return sentence
}

func (language *Dansk) Translate(sentence *g.StatementGroup) ([]p.Word, string) {
	parsedSentence := language.ParseSentence(sentence)
	return parsedSentence.GetResult()
}

func (language *Dansk) ParseSentence(baseSentence *g.StatementGroup) DanishSentence {
	var sentence DanishSentence
	sentence.Language = language
	sentence.ParseVerb(baseSentence)
	sentence.ParseSubject(baseSentence)
	return sentence
}

func (sentence *DanishSentence) ParseSubject(source *g.StatementGroup) {
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

func (sentence *DanishSentence) ParseComplexSubject(source *g.StatementGroup) {
	sentence.Subject = Navneord{}
	sentence.Subject.Ord = sentence.Language.Words[source.SimpleConcept] // TODO: this is wrong
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) ParseSimpleSubject(source *g.StatementGroup) {
	sentence.Subject = Navneord{}
	sentence.Subject.Ord = sentence.Language.Words[source.SimpleConcept]
	sentence.Subject.Bestemt = 0 < len(source.GetDescriptorsOf("definite", "descriptor"))
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) ParseDescriptors(source *g.StatementGroup) {
	//do something
}

func (sentence *DanishSentence) ParseVerb(source *g.StatementGroup) {
	if source.IsComplex() {
		sentence.ParseComplexVerb(source)
	} else {
		sentence.ParseSimpleVerb(source)
	}
}

func (sentence *DanishSentence) ParseComplexVerb(source *g.StatementGroup) {
	sentence.Verb = Udsagnsord{}
	sentence.Verb.Ord = sentence.Language.Words[source.SimpleConcept] // TODO: this is wrong
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) ParseSimpleVerb(source *g.StatementGroup) {
	sentence.Verb = Udsagnsord{}
	if len(source.GetDescriptors("doer")) > 0 {
		sentence.Verb.Ord = sentence.Language.Words[source.SimpleConcept]
	} else if len(source.GetDescriptors("beer")) > 0 {
		sentence.Verb.Ord = sentence.Language.BasicWords["er"]
		sentence.Object = Navneord{}
		sentence.Object.Ord = sentence.Language.Words[source.SimpleConcept]
	} else {
		//uh
	}
	sentence.Verb.Tid = sentence.GetTime(source)
	sentence.ParseDescriptors(source)
}

func (sentence *DanishSentence) GetTime(source *g.StatementGroup) string {
	timeDescriptors := source.GetDescriptorsOf("now", "at", "around", "after", "before")
	if len(timeDescriptors) == 0 {
		return "nutid"
	}
	switch timeDescriptors[0].Relation {
	case "at", "around", "after":
		return "nutid"
	case "before":
		return "datid"
	default:
		return "nutid" //TODO
	}
}

func (sentence *DanishSentence) GetResult() ([]p.Word, string) {
	words := []p.Word{}
	text := sentence.Subject.GetText() + " " + sentence.Verb.GetText()
	return words, text
}
