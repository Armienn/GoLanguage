package language

// The core representation of a sentence/statement. The will be a core language,
// which directly corresponds to this structure, and other languages can then be
// collections of arbitrary word groups along with rules for how to translate
// from the core language
type StatementGroup struct {
	CoreConcept     Concept
	ExpandedConcept *StatementGroup
	StatementType   StatementType // statement, subject, object, vocative (vocated?), instrument, time, etc
	Descriptors     []StatementGroup
}

type StatementType int

const (
	// first types for sentences
	Statement StatementType = iota
	But
	If
	Verbification
	And
	// then types for concepts
	Doer
	Beer
	Object
	Instrument
	Material
	Location //del op i flere: over, under, i, ved, etc? Ja
	On
	In
	Under
	Behind
	Infrontof
	Nextto
	Comingfrom
	Goingto
	Time //same time as
	After
	Before
	Owner
	Owned
	Reciever
	Evoked
	Descriptor
)

type Concept string

type ConceptInfo struct {
	Concept        Concept
	ValidArguments []StatementType
}

func GetSentences() []StatementGroup {
	sentences := make([]StatementGroup, 0)
	//a man eats a cat
	sentence := NewStatementGroup("eat", Statement)
	sentence.AddDescriptor(NewStatementGroup("man", Doer))
	sentence.AddDescriptor(NewStatementGroup("cat", Object))
	sentences = append(sentences, *sentence)
	//it rains
	sentence = NewStatementGroup("rain", Statement)
	sentences = append(sentences, *sentence)
	return sentences
}

func NewStatementGroup(base Concept, relation StatementType) *StatementGroup {
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
// ARGH. to jump vs to be a jump vs to be a jumper
// Or even: to give vs to be a giving vs to be a gift vs to be a giver vs to be a giftee
// It works: [man|beer] [jump|statement] - a man is a jump
// vs: [man|doer] [jump|statement] - a man jumps
// I'm an idiot
// It works: [man|beer] [[jump|verb][who|doer]|statement] - a man is a jumper (a man is one who jumps)
// vs: [man|doer] [jump|statement] - a man jumps
