package language

// The core representation of a sentence/statement. The will be a core language,
// which directly corresponds to this structure, and other languages can then be
// collections of arbitrary word groups along with rules for how to translate
// from the core language
type StatementGroup struct {
	SimpleConcept   Concept
	CompoundConcept *StatementGroup
	Relation        Concept // statement, subject, object, vocative (vocated?), instrument, time, etc
	Descriptors     []*StatementGroup
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

//NewStatementGroup returns a new group with a simple concept and no descriptors
func NewStatementGroup(base Concept, relation Concept) *StatementGroup {
	return &StatementGroup{base, nil, relation, make([]*StatementGroup, 0)}
}

//AddDescriptor adds a descriptor to the group
func (group *StatementGroup) AddDescriptor(descriptor *StatementGroup) {
	group.Descriptors = append(group.Descriptors, descriptor)
}

//IsComplex returns true if the group has a compound concept instead of a simple concept
func (group *StatementGroup) IsComplex() bool {
	return group.CompoundConcept != nil
}

//GetDescriptors returns the descriptors which are of one of the give relations
func (group *StatementGroup) GetDescriptors(relations ...Concept) []*StatementGroup {
	descriptors := make([]*StatementGroup, 0)
	for _, descriptor := range group.Descriptors {
		if descriptor.HasRelation(relations...) {
			descriptors = append(descriptors, descriptor)
		}
	}
	return descriptors
}

//GetDescriptorsOf returns the descriptors of the given descriptor concept, which are of one of the give relations
func (group *StatementGroup) GetDescriptorsOf(descriptor Concept, relations ...Concept) []*StatementGroup {
	descriptors := make([]*StatementGroup, 0)
	for _, descr := range group.Descriptors {
		if descr.SimpleConcept == descriptor && descr.HasRelation(relations...) {
			descriptors = append(descriptors, descr)
		}
	}
	return descriptors
}

// HasRelation returns true if the groups relation equals one of the given concepts
func (group *StatementGroup) HasRelation(relations ...Concept) bool {
	for _, relation := range relations {
		if group.Relation == relation {
			return true
		}
	}
	return false
}

func (group *StatementGroup) String() string {
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

//ContainsSameRelations returns true if the given lists of groups have the same relations
func ContainsSameRelations(listA []*StatementGroup, listB []*StatementGroup) bool {
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

func GetSentences() []*StatementGroup {
	sentences := make([]*StatementGroup, 0)
	//a man eats a cat
	sentence := NewStatementGroup("shine", "")
	sentence.AddDescriptor(NewStatementGroup("sun", "doer"))
	sentence.AddDescriptor(NewStatementGroup("now", "before"))
	sentences = append(sentences, sentence)
	return sentences
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
