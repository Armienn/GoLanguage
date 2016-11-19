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

func GetDanishLanguage() *grammatics.Dansk {
	dansk := grammatics.Dansk{}
	dansk.Verber = map[grammatics.Concept]grammatics.VerbiumRepresenter{
		"!":     NewVerbiumWord("!er", "!ede"),
		"shine": NewVerbiumWord("skinner", "skinnede"),
	}
	dansk.Substantiver = map[grammatics.Concept]grammatics.SubstantivRepresenter{
		"!":   NewSubstantivWord("en", "!", "!en", "!er", "!erne"),
		"sun": NewSubstantivWord("en", "sol", "solen", "sole", "solene"),
	}
	dansk.Adjektiver = make(map[grammatics.Concept]grammatics.AdjektivRepresenter)
	dansk.Adverbier = make(map[grammatics.Concept]grammatics.AdverbiumRepresenter)
	dansk.Præpositioner = make(map[grammatics.Concept]grammatics.SimpleRepresenter)
	dansk.Er = NewVerbiumWord("er", "var")
	return &dansk
}
