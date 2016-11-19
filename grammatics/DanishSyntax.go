package grammatics

type VerbiumRepresenter interface {
	Representation(*DanishSentence) []WordRepresenter
}

type SubstantivRepresenter interface {
	Representation(*Substantiv) []WordRepresenter
}

type AdjektivRepresenter interface {
	Representation(*Adjektiv) []WordRepresenter
}

type AdverbiumRepresenter interface {
	Representation(*Adverbium) []WordRepresenter
}

type SimpleRepresenter interface {
	Representation() []WordRepresenter
}

type Dansk struct {
	Verber        map[Concept]VerbiumRepresenter
	Substantiver  map[Concept]SubstantivRepresenter
	Adjektiver    map[Concept]AdjektivRepresenter
	Adverbier     map[Concept]AdverbiumRepresenter
	Præpositioner map[Concept]SimpleRepresenter
	Er            VerbiumRepresenter
	Og            SimpleRepresenter
}

type MissingVerbium struct{}

func (word MissingVerbium) Representation(*DanishSentence) []WordRepresenter {
	return []WordRepresenter{wordString{"missing"}}
}

type MissingSubstantiv struct{}

func (word MissingSubstantiv) Representation(*Substantiv) []WordRepresenter {
	return []WordRepresenter{wordString{"missing"}}
}

type MissingAdjektiv struct{}

func (word MissingAdjektiv) Representation(*Adjektiv) []WordRepresenter {
	return []WordRepresenter{wordString{"missing"}}
}

type MissingAdverbium struct{}

func (word MissingAdverbium) Representation(*Adverbium) []WordRepresenter {
	return []WordRepresenter{wordString{"missing"}}
}

type MissingSimple struct{}

func (word MissingSimple) Representation() []WordRepresenter {
	return []WordRepresenter{wordString{"missing"}}
}

type wordString struct {
	word string
}

func (word wordString) Representation() interface{} {
	return word.word
}

type Verbium struct {
	Ord       VerbiumRepresenter
	Tid       string
	Adverbier []Adverbium
}

type Substantiv struct {
	Ord        SubstantivRepresenter
	Flertal    bool
	Bestemt    bool
	Tillægsord []Adjektiv
}

type Adjektiv struct {
	Ord       AdjektivRepresenter
	Adverbier []Adverbium
}

type Adverbium struct {
	Ord       AdverbiumRepresenter
	Adverbier []Adverbium
}

type Forholdsled struct {
	Forholdsord SimpleRepresenter
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
		sentence.Verb.Ord = language.FindVerbium(source) //TODO
	} else {
		sentence.Verb.Ord = language.FindVerbium(source)
	}
	return sentence
}

func (language *Dansk) ParseSubstantiv(source *Statement) Substantiv {
	substantiv := Substantiv{}
	if source.IsComplex() {
		substantiv.Ord = language.FindSubstantiv(source) //TODO
	} else {
		substantiv.Ord = language.FindSubstantiv(source)
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
		adjektiv.Ord = language.FindAdjektiv(source) //TODO
	} else {
		adjektiv.Ord = language.FindAdjektiv(source)
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
		adverbium.Ord = language.FindAdverbium(source) //TODO
	} else {
		adverbium.Ord = language.FindAdverbium(source)
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

func (language *Dansk) FindVerbium(source *Statement) VerbiumRepresenter {
	word, ok := language.Verber[source.SimpleConcept]
	if !ok {
		return MissingVerbium{}
	}
	return word
}

func (language *Dansk) FindSubstantiv(source *Statement) SubstantivRepresenter {
	word, ok := language.Substantiver[source.SimpleConcept]
	if !ok {
		return MissingSubstantiv{}
	}
	return word
}

func (language *Dansk) FindAdjektiv(source *Statement) AdjektivRepresenter {
	word, ok := language.Adjektiver[source.SimpleConcept]
	if !ok {
		return MissingAdjektiv{}
	}
	return word
}

func (language *Dansk) FindAdverbium(source *Statement) AdverbiumRepresenter {
	word, ok := language.Adverbier[source.SimpleConcept]
	if !ok {
		return MissingAdverbium{}
	}
	return word
}

func (language *Dansk) FindSimple(source *Statement) SimpleRepresenter {
	word, ok := language.Præpositioner[source.SimpleConcept]
	if !ok {
		return MissingSimple{}
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

func (sentence *DanishSentence) GetResult() []WordRepresenter {
	return sentence.Verb.Ord.Representation(sentence)
}
