package language

// The core representation of a sentence/statement. The will be a core language,
// which directly corresponds to this structure, and other languages can then be
// collections of arbitrary word groups along with rules for how to translate
// from the core language
type StatementGroup struct {
	CoreConcept     string
	ExpandedConcept *StatementGroup
	StatementType   string // statement, subject, object, vocative (vocated?), instrument, time, etc
	Descriptors     []StatementGroup
}

type CoreLanguage struct {
	Concepts       []string
	StatementTypes []string
}

func GetCoreLanguage() *CoreLanguage {
	core := new(CoreLanguage)
	core.Concepts = []string{
		"man",
		"eat",
		"jump",
		"cat",
		"house",
		"rain",
		"look",
		"eat",
		"child",
		"daughter",
		"sit",
		"love",
		"drive",
	}
	core.StatementTypes = []string{
		// first types for sentences
		"statement",
		"but",
		"if",
		"verbification",
		"and",
		// then types for concepts
		"subject",
		"object",
		"instrument",
		"material",
		"location", //del op i flere: over, under, i, ved, etc? Ja
		"on",
		"in",
		"under",
		"behind",
		"infrontof",
		"nextto",
		"comingfrom",
		"goingto",
		"time",
		"after",
		"before",
		"owner",
		"owned",
		"evoked",
		"descriptor",
	}
	return core
}

func GetSentences() []StatementGroup {
	sentences := make([]StatementGroup, 0)
	sentence := NewStatementGroup("eat", "statement")
	sentence.AddDescriptor(NewStatementGroup("cat", "object"))
	sentence.AddDescriptor(NewStatementGroup("man", "subject"))
	return sentences
}

func NewStatementGroup(base string, relation string) *StatementGroup {
	return &StatementGroup{base, nil, relation, make([]StatementGroup, 0)}
}

func (statement *StatementGroup) AddDescriptor(descriptor *StatementGroup) {
	statement.Descriptors = append(statement.Descriptors, *descriptor)
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
