package grammatics

import (
	"strings"
)

// The core representation of a sentence/statement. The will be a core language,
// which directly corresponds to this structure, and other languages can then be
// collections of arbitrary word groups along with rules for how to translate
// from the core language
type Statement struct {
	SimpleConcept   Concept
	CompoundConcept *Statement
	Relation        Concept // statement, subject, object, vocative (vocated?), instrument, time, etc
	Descriptors     []*Statement
}

type Concept string

type ConceptInfo struct {
	Description    string
	ValidArguments []Concept //beer, location-things and time-things are always valid
}

//HasArgument returns true if the ConceptInfo has the given concept as a possible argument
func (info *ConceptInfo) HasArgument(argument Concept) bool {
	switch argument {
	case "beer":
		return true
	}
	for _, concept := range info.ValidArguments {
		if concept == argument {
			return true
		}
	}
	return false
}

//NewStatement returns a new group with a simple concept and no descriptors
func NewStatement(base Concept, relation Concept) *Statement {
	return &Statement{base, nil, relation, make([]*Statement, 0)}
}

//StatementFromString parses and returns a Statement from the given string
func StatementFromString(source string) *Statement {
	group := &Statement{}
	source = strings.TrimSpace(source)
	if strings.HasPrefix(source, "[") {
		source = SplitByBraces(source, '[', ']')[0]
	}
	parts := SplitByBraces(source, '[', ']')
	if strings.HasSuffix(parts[0], ":") {
		group.Relation = Concept(strings.TrimSuffix(parts[0], ":"))
		group.CompoundConcept = StatementFromString(parts[1])
	} else {
		mub := strings.Split(parts[0], ":")
		if len(mub) == 1 {
			group.Relation = "descriptor"
			group.SimpleConcept = Concept(mub[0])
		} else {
			group.Relation = Concept(mub[0])
			group.SimpleConcept = Concept(mub[1])
		}
	}
	start := 1
	if group.IsComplex() {
		start = 2
	}
	group.Descriptors = make([]*Statement, len(parts)-start)
	for i := 0; i < len(parts)-start; i++ {
		group.Descriptors[i] = StatementFromString(parts[i+start])
	}
	return group
}

//SplitByBraces splits the given string into substrings based on the given start and end runes
func SplitByBraces(source string, start rune, end rune) []string {
	splits := make([]string, 0)
	startIndex := 0
	level := 0
	for i, c := range source {
		if c == start {
			if level == 0 {
				if i != startIndex {
					splits = append(splits, source[startIndex:i])
				}
				startIndex = i + 1
			}
			level++
		} else if c == end {
			level--
			if level == 0 {
				splits = append(splits, source[startIndex:i])
				startIndex = i + 1
			}
		}
	}
	if startIndex != len(source) {
		splits = append(splits, source[startIndex:])
	}
	return splits
}

func (group *Statement) String() string {
	var core string
	if group.IsComplex() {
		core = group.CompoundConcept.String()
	} else {
		core = string(group.SimpleConcept)
	}
	for _, descriptor := range group.Descriptors {
		core += descriptor.String()
	}
	return "[" + string(group.Relation) + ":" + core + "]"
}

//AddDescriptor adds a descriptor to the group
func (group *Statement) AddDescriptor(descriptor *Statement) {
	group.Descriptors = append(group.Descriptors, descriptor)
}

//IsComplex returns true if the group has a compound concept instead of a simple concept
func (group *Statement) IsComplex() bool {
	return group.CompoundConcept != nil
}

//GetDescriptors returns the descriptors which are of one of the give relations
func (group *Statement) GetDescriptors(relations ...Concept) []*Statement {
	descriptors := make([]*Statement, 0)
	for _, descriptor := range group.Descriptors {
		if descriptor.HasRelation(relations...) {
			descriptors = append(descriptors, descriptor)
		}
	}
	return descriptors
}

//GetDescriptorsOf returns the descriptors of the given descriptor concept, which are of one of the give relations
func (group *Statement) GetDescriptorsOf(descriptor Concept, relations ...Concept) []*Statement {
	descriptors := make([]*Statement, 0)
	for _, descr := range group.Descriptors {
		if descr.SimpleConcept == descriptor && descr.HasRelation(relations...) {
			descriptors = append(descriptors, descr)
		}
	}
	return descriptors
}

// HasRelation returns true if the groups relation equals one of the given concepts
func (group *Statement) HasRelation(relations ...Concept) bool {
	for _, relation := range relations {
		if group.Relation == relation {
			return true
		}
	}
	return false
}

//ContainsSameRelations returns true if the given lists of groups have the same relations
func ContainsSameRelations(listA []*Statement, listB []*Statement) bool {
	length := len(listA)
	if length != len(listB) {
		return false
	}
	for _, elemA := range listA {
		found := false
		for _, elemB := range listB {
			if elemA.Relation == elemB.Relation {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func Info(description string, validArguments ...Concept) *ConceptInfo {
	result := new(ConceptInfo)
	result.Description = description
	result.ValidArguments = validArguments
	return result
}

func GetCoreLanguage() map[Concept]*ConceptInfo {
	return map[Concept]*ConceptInfo{
		//object is always optional and is substituted by an undefined 'something' if not specified
		//"be":    *Info("doer is object", "doer", "object"),
		//"do":    *Info("doer does object", "doer", "object"),
		"beer":       Info("beer is one who is object", "object"),
		"doer":       Info("beer is one who does object", "object"),                      //object must be doable (must have a possible doer)
		"object":     Info("beer is one who is the object of object", "object"),          //object must have a possible object
		"descriptor": Info("beer is a manifestation of the concept of object", "object"), //it is the beer of the concept... hm
		"at":         Info("beer is one who is at (near or in (either spacially or chronologically)) object", "object"),
		"around":     Info("beer is one who is spread around (either spacially or chronologically) object", "object"),
		"before":     Info("beer is one who is chronologically before object", "object"),
		"after":      Info("beer is one who is chronologically after object", "object"),
		"now":        Info("beer is one who is chronologically near/at/alongside object", "object"),
		"again":      Info("beer is an event that reoccurs"),
		"definite":   Info("beer is one who is blabla todo"),
		"sun":        Info("beer is the sun of belonger", "belonger"),
		"shine":      Info("doer shines on reciever with light source instrument", "doer", "reciever", "instrument"),
	}
}

func GetSentences() []*Statement {
	return []*Statement{
		StatementFromString("[:shine[doer:sun]]"),
		StatementFromString("[:shine[doer:sun][before:now]]"),
		StatementFromString("[:shine[doer:sun][at:now]]"),
		StatementFromString("[:shine[doer:sun[definite]]]"),
		StatementFromString("[:shine[doer:sun[definite]][before:now]]"),
		StatementFromString("[:shine[doer:sun[definite]][at:now]]"),
	}
}

// concept - can have a do'er (event) or can have a be'er (property) - nope, scratch
// concept - (event) can have a do'er and a be'er, (property) can have a be'er
// ^ i.e. substitute subject with do'er and be'er
// (jump) core: to jump - be'er: a jump - do'er: a jumper     ------  to be a jumper!?
// (jump) core: to be a jumper - be'er: a jumper - do'er: ?
// (eat) core: to eat - be'er: an instance of eating? - do'er: an eater
// (eat) core: to be an eater - be'er: an eater - do'er: ?
// (man) core: to be a man? the man-property? - be'er: a man - do'er: ?
// (love) core: to love - be'er: a love (en kærlighed) - do'er: a lover (something that loves)
// (pretty) core: to be pretty? - be'er: a pretty something - do'er: ?
// (day) core: to be a day? - be'er: a day - do'er: ?

// when concept is used as descriptor, the described is a be'er of the concept
// when the descriptor is for a statement, the event descriped by the statement is the be'er of the descriptor (as in all other cases)

// a jump vs a jumper solution: a jump/to jump is the basic concept, and to construct the other meaning use expandedconcept thing
// ARGH. to jump vs to be a jump vs to be a jumper
// Or even: to give vs to be a giving vs to be a gift vs to be a giver vs to be a giftee
// It works: [man|beer] [jump|statement] - a man is a jump
// vs: [man|doer] [jump|statement] - a man jumps
// I'm an idiot
// It works: [man|beer] [[jump|verb][who|doer]|statement] - a man is a jumper (a man is one who jumps)
// vs: [man|doer] [jump|statement] - a man jumps
