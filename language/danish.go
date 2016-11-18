package language

import "github.com/Armienn/GoLanguage/grammatics"

type WordString struct {
	Word string
}

func (word WordString) Representation() interface{} {
	return word.Word
}

type DanishWord struct {
	NavneordWord   map[string]WordString
	UdsagnsordWord map[string]WordString
}

func NewNavneord(en string, ubestemt string, bestemt string, multi string, multibestemt string) *DanishWord {
	return &DanishWord{map[string]WordString{
		"en":           WordString{en},
		"ubestemt":     WordString{ubestemt},
		"bestemt":      WordString{bestemt},
		"multi":        WordString{multi},
		"multibestemt": WordString{multibestemt},
	}, nil}
}

func NewUdsagnsord(nutid string, datid string) *DanishWord {
	return &DanishWord{nil, map[string]WordString{
		"nutid": WordString{nutid},
		"datid": WordString{datid},
	}}
}

func (word *DanishWord) NavneordRepresentation(ord *grammatics.Navneord) []grammatics.WordRepresenter {
	if word.NavneordWord == nil {
		return nil
	}
	if ord.Bestemt {
		if ord.Flertal {
			return []grammatics.WordRepresenter{word.NavneordWord["multibestemt"]}
		}
		return []grammatics.WordRepresenter{word.NavneordWord["bestemt"]}
	}
	if ord.Flertal {
		return []grammatics.WordRepresenter{word.NavneordWord["multi"]}
	}
	return []grammatics.WordRepresenter{word.NavneordWord["en"], word.NavneordWord["ubestemt"]}
}

func (word *DanishWord) UdsagnsordRepresentation(ord *grammatics.Udsagnsord) []grammatics.WordRepresenter {
	if word.UdsagnsordWord == nil {
		return nil
	}
	return []grammatics.WordRepresenter{word.UdsagnsordWord[ord.Tid]}
}

func GetDanishLanguage() *grammatics.Dansk {
	dansk := grammatics.Dansk{}
	dansk.Words = map[grammatics.Concept]grammatics.DanishWordRepresenter{
		"sun":   NewNavneord("en", "sol", "solen", "sole", "solene"),
		"shine": NewUdsagnsord("skinner", "skinnede"),
	}
	dansk.Er = NewUdsagnsord("er", "var")
	dansk.Missing = new(grammatics.MissingWord)
	return &dansk
}
