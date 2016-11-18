package grammatics

type Translater interface {
	Translate(*Statement) []WordRepresenter
}

type WordRepresenter interface {
	Representation() interface{}
}
