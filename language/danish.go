package language

import "github.com/Armienn/GoLanguage/grammatics"

type WordString struct {
	Word string
}

func (word WordString) Representation() interface{} {
	return word.Word
}

type VerbiumWord struct {
	Forms map[string]WordString
}

func (word VerbiumWord) Representation(sentence *grammatics.DanishSentence) []grammatics.WordRepresenter {
	words := make([]grammatics.WordRepresenter, 0)
	for i, part := range sentence.Subject {
		if i > 0 {
			words = append(words, sentence.Language.Og.Representation()...)
		}
		words = append(words, part.Ord.Representation(&part)...)
	}
	words = append(words, word.Forms[sentence.Verb.Tid])
	for i, part := range sentence.Object {
		if i > 0 {
			words = append(words, sentence.Language.Og.Representation()...)
		}
		words = append(words, part.Ord.Representation(&part)...)
	}
	for _, part := range sentence.Other {
		words = append(words, part.Forholdsord.Representation()...)
		words = append(words, part.Ord.Ord.Representation(&part.Ord)...)
	}
	return words
}

func NewVerbiumWord(nutid string, datid string) *VerbiumWord {
	return &VerbiumWord{map[string]WordString{
		"nutid": WordString{nutid},
		"datid": WordString{datid},
	}}
}

type SubstantivWord struct {
	Forms map[string]WordString
}

func (word SubstantivWord) Representation(ord *grammatics.Substantiv) []grammatics.WordRepresenter {
	if ord.Bestemt {
		if ord.Flertal {
			return []grammatics.WordRepresenter{word.Forms["multibestemt"]}
		}
		return []grammatics.WordRepresenter{word.Forms["bestemt"]}
	}
	if ord.Flertal {
		return []grammatics.WordRepresenter{word.Forms["multi"]}
	}
	return []grammatics.WordRepresenter{word.Forms["en"], word.Forms["ubestemt"]}
}

func NewSubstantivWord(en string, ubestemt string, bestemt string, multi string, multibestemt string) *SubstantivWord {
	return &SubstantivWord{map[string]WordString{
		"en":           WordString{en},
		"ubestemt":     WordString{ubestemt},
		"bestemt":      WordString{bestemt},
		"multi":        WordString{multi},
		"multibestemt": WordString{multibestemt},
	}}
}

type AdjektivWord struct {
	Forms map[string]WordString
}

func (word AdjektivWord) Representation(ord *grammatics.Adjektiv) []grammatics.WordRepresenter {
	return []grammatics.WordRepresenter{word.Forms["base"]}
}

func NewAdjektivWord(base string) *AdjektivWord {
	return &AdjektivWord{map[string]WordString{
		"base": WordString{base},
	}}
}

type AdverbiumWord struct {
	Forms map[string]WordString
}

func (word AdverbiumWord) Representation(ord *grammatics.Adverbium) []grammatics.WordRepresenter {
	return []grammatics.WordRepresenter{word.Forms["base"]}
}

func NewAdverbiumWord(base string) *AdverbiumWord {
	return &AdverbiumWord{map[string]WordString{
		"base": WordString{base},
	}}
}

type SimpleWord struct {
	Forms map[string]WordString
}

func (word SimpleWord) Representation() []grammatics.WordRepresenter {
	return []grammatics.WordRepresenter{word.Forms["base"]}
}

func NewSimpleWord(base string) *SimpleWord {
	return &SimpleWord{map[string]WordString{
		"base": WordString{base},
	}}
}

func GetDanishLanguage() *grammatics.Dansk {
	dansk := grammatics.Dansk{}
	dansk.Verber = map[grammatics.Concept]grammatics.VerbiumRepresenter{
		"!":     NewVerbiumWord("!er", "!ede"),
		"shine": NewVerbiumWord("skinner", "skinnede"),
		"rise":  NewVerbiumWord("stiger", "steg"),
		"shout": NewVerbiumWord("råber", "råbte"),
		"jump":  NewVerbiumWord("hopper", "hoppede"),
		"walk":  NewVerbiumWord("går", "gik"),
	}
	dansk.Substantiver = map[grammatics.Concept]grammatics.SubstantivRepresenter{
		"!":      NewSubstantivWord("en", "!", "!en", "!er", "!erne"),
		"sun":    NewSubstantivWord("en", "sol", "solen", "sole", "solene"),
		"day":    NewSubstantivWord("en", "dag", "dagen", "dage", "dagene"),
		"person": NewSubstantivWord("en", "person", "personen", "personer", "personerne"),
		"kitten": NewSubstantivWord("en", "killing", "killingen", "killinger", "killingerne"),
		"table":  NewSubstantivWord("et", "bord", "bordet", "borde", "bordene"),
	}
	dansk.Adjektiver = map[grammatics.Concept]grammatics.AdjektivRepresenter{
		"!":      NewAdjektivWord("!"),
		"bright": NewAdjektivWord("lys"),
		"all":    NewAdjektivWord("alle"),
	}
	dansk.Adverbier = map[grammatics.Concept]grammatics.AdverbiumRepresenter{
		"!":      NewAdverbiumWord("!"),
		"bright": NewAdverbiumWord("lyst"),
		"up":     NewAdverbiumWord("op"),
		"again":  NewAdverbiumWord("igen"),
	}
	dansk.Præpositioner = map[grammatics.Concept]grammatics.SimpleRepresenter{
		"!":      NewSimpleWord("!"),
		"at":     NewSimpleWord("ved"),
		"before": NewSimpleWord("før"),
		"after":  NewSimpleWord("efter"),
		"onto":   NewSimpleWord("op på"),
	}
	dansk.Er = NewVerbiumWord("er", "var")
	return &dansk
}
