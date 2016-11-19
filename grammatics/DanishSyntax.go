package grammatics

type DanishWordRepresenter interface {
	SubstantivRepresentation(*Substantiv) []WordRepresenter
	VerbiumRepresentation(*Verbium) []WordRepresenter
}

type Dansk struct {
	Verber        map[Concept]DanishWordRepresenter
	Substantiver  map[Concept]DanishWordRepresenter
	Adjektiver    map[Concept]DanishWordRepresenter
	Adverbier     map[Concept]DanishWordRepresenter
	Præpositioner map[Concept]DanishWordRepresenter
	Er            DanishWordRepresenter
	Missing       DanishWordRepresenter
}

type MissingWord struct{}

func (word *MissingWord) SubstantivRepresentation(ord *Substantiv) []WordRepresenter {
	return []WordRepresenter{wordString{"missing"}}
}

func (word *MissingWord) VerbiumRepresentation(ord *Verbium) []WordRepresenter {
	return []WordRepresenter{wordString{"missing"}}
}

type wordString struct {
	word string
}

func (word wordString) Representation() interface{} {
	return word.word
}

type Verbium struct {
	Ord       DanishWordRepresenter
	Tid       string
	Adverbier []Adverbium
}

type Substantiv struct {
	Ord        DanishWordRepresenter
	Flertal    bool
	Bestemt    bool
	Tillægsord []Adjektiv
}

type Adjektiv struct {
	Ord       DanishWordRepresenter
	Adverbier []Adverbium
}

type Adverbium struct {
	Ord       DanishWordRepresenter
	Adverbier []Adverbium
}

type Forholdsled struct {
	Forholdsord DanishWordRepresenter
	Ord         Substantiv
}

type DanishSentence struct {
	core     map[Concept]*ConceptInfo
	Language *Dansk
	Verb     Verbium
	Subject  []Substantiv
	Object   []Substantiv
	Other    []Forholdsled
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

func (language *Dansk) ParseSentence(source *Statement) DanishSentence {
	var sentence DanishSentence
	sentence.Language = language

	for _, descriptor := range source.Descriptors {
		switch descriptor.Relation {
		case "doer", "beer":
			sentence.Subject = append(sentence.Subject, language.ParseSubstantiv(descriptor))
		case "object":
			sentence.Object = append(sentence.Object, language.ParseSubstantiv(descriptor))
		case "at":
			if IsTime(descriptor) {
				sentence.Verb.Tid = GetTime(descriptor)
			} else {
				sentence.Other = append(sentence.Other, language.ParseForholdsled(descriptor))
			}
		case "around":
			if IsTime(descriptor) {
				sentence.Verb.Tid = GetTime(descriptor) //some other time thing
			} else {
				sentence.Other = append(sentence.Other, language.ParseForholdsled(descriptor))
			}
		case "descriptor":
			sentence.Verb.Adverbier = append(sentence.Verb.Adverbier, language.ParseAdverbium(descriptor))
		case "and", "but": // additional sentences

		default:
			sentence.Other = append(sentence.Other, language.ParseForholdsled(descriptor))
		}
	}

	if source.IsComplex() {
		sentence.Verb.Ord = language.ParseComplex(source)
	} else {
		sentence.Verb.Ord = FindWord(source, language.Verber, language.Missing)
	}
	return sentence
}

func (language *Dansk) ParseSubstantiv(source *Statement) Substantiv {
	substantiv := Substantiv{}
	if source.IsComplex() {
		substantiv.Ord = language.ParseComplex(source)
	} else {
		substantiv.Ord = FindWord(source, language.Substantiver, language.Missing)
	}
	substantiv.Tillægsord = make([]Adjektiv, 0)
	for _, descriptor := range source.Descriptors {
		if descriptor.Relation == "descriptor" && descriptor.SimpleConcept == "definite" {
			substantiv.Bestemt = true
			continue
		}
		if descriptor.Relation == "amount" && descriptor.SimpleConcept == "several" {
			substantiv.Flertal = true
			continue
		}
		if descriptor.Relation == "amount" && descriptor.SimpleConcept != "one" {
			substantiv.Flertal = true
		}
		substantiv.Tillægsord = append(substantiv.Tillægsord, language.ParseAdjektiv(descriptor))
	}
	return substantiv
}

func (language *Dansk) ParseAdjektiv(source *Statement) Adjektiv {
	adjektiv := Adjektiv{}
	if source.IsComplex() {
		adjektiv.Ord = language.ParseComplex(source)
	} else {
		adjektiv.Ord = FindWord(source, language.Adjektiver, language.Missing)
	}
	adjektiv.Adverbier = make([]Adverbium, 0)
	for _, descriptor := range source.Descriptors {
		adjektiv.Adverbier = append(adjektiv.Adverbier, language.ParseAdverbium(descriptor))
	}
	return adjektiv
}

func (language *Dansk) ParseAdverbium(source *Statement) Adverbium {
	adverbium := Adverbium{}
	if source.IsComplex() {
		adverbium.Ord = language.ParseComplex(source)
	} else {
		adverbium.Ord = FindWord(source, language.Adverbier, language.Missing)
	}
	adverbium.Adverbier = make([]Adverbium, 0)
	for _, descriptor := range source.Descriptors {
		adverbium.Adverbier = append(adverbium.Adverbier, language.ParseAdverbium(descriptor))
	}
	return adverbium
}

func (language *Dansk) ParseForholdsled(source *Statement) Forholdsled {
	led := Forholdsled{}
	led.Forholdsord = language.Præpositioner[source.Relation]
	led.Ord = language.ParseSubstantiv(source)
	return led
}

func FindWord(source *Statement, list map[Concept]DanishWordRepresenter, missing DanishWordRepresenter) DanishWordRepresenter {
	word, ok := list[source.SimpleConcept]
	if !ok {
		return missing
	}
	return word
}

func IsTime(source *Statement) bool {
	if source.IsComplex() {
		return false
	}
	switch source.SimpleConcept {
	case "after", "now", "before":
		return true
	}
	return false
}

func GetTime(source *Statement) string {
	switch source.SimpleConcept {
	case "before":
		return "datid"
	}
	return "nutid"
}

func (language *Dansk) ParseComplex(source *Statement) DanishWordRepresenter {
	return language.Missing
}

func (sentence *DanishSentence) GetResult() []WordRepresenter {
	words := make([]WordRepresenter, 0)
	if sentence.Subject.Ord != nil {
		words = append(words, sentence.Subject.Ord.NavneordRepresentation(&sentence.Subject)...)
	}
	words = append(words, sentence.Verb.Ord.UdsagnsordRepresentation(&sentence.Verb)...)
	return words
}
