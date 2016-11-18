package grammatics

type Translater interface {
	Translate(*StatementGroup) []WordRepresenter
}

type WordRepresenter interface {
	Representation() interface{}
}
