package grammatics

import "testing"

func TestHasArgument(t *testing.T) {
	cases := []struct {
		object ConceptInfo
		in     Concept
		want   bool
	}{
		{ConceptInfo{"", []Concept{"doer", "object"}}, "doer", true},
		{ConceptInfo{"", []Concept{"doer", "object"}}, "object", true},
		{ConceptInfo{"", []Concept{"doer", "object"}}, "beer", true},
		{ConceptInfo{"", []Concept{"doer", "object"}}, "asdf", false},
	}
	for _, c := range cases {
		got := c.object.HasArgument(c.in)
		if got != c.want {
			t.Errorf("%v.IsComplex(%v) == %v, want %v", c.object, c.in, got, c.want)
		}
	}
}

func TestIsComplex(t *testing.T) {
	cases := []struct {
		in   Statement
		want bool
	}{
		{Statement{"asdf", nil, "", nil}, false},
		{Statement{"", nil, "", nil}, false},
		{Statement{"asdf", new(Statement), "", nil}, true},
		{Statement{"", new(Statement), "", nil}, true},
	}
	for _, c := range cases {
		got := c.in.IsComplex()
		if got != c.want {
			t.Errorf("%v.IsComplex() == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestHasRelation(t *testing.T) {
	cases := []struct {
		object Statement
		in     []Concept
		want   bool
	}{
		{Statement{"", nil, "relation", nil}, []Concept{"doer", "beer"}, false},
		{Statement{"", nil, "beer", nil}, []Concept{"doer", "beer"}, true},
		{Statement{"", nil, "relation", nil}, []Concept{}, false},
		{Statement{"", nil, "beer", nil}, []Concept{"beer"}, true},
	}
	for _, c := range cases {
		got := c.object.HasRelation(c.in...)
		if got != c.want {
			t.Errorf("%v.HasRelation(%v) == %v, want %v", c.object, c.in, got, c.want)
		}
	}
}

func TestGetDescriptors(t *testing.T) {
	groups := []*Statement{
		&Statement{"", nil, "doer", nil},
		&Statement{"", nil, "beer", nil},
		&Statement{"", nil, "object", nil},
	}
	cases := []struct {
		object Statement
		in     []Concept
		want   []*Statement
	}{
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, []Concept{"doer"}, []*Statement{groups[0]}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, []Concept{"beer"}, []*Statement{groups[1]}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, []Concept{"doer", "beer"}, []*Statement{groups[0], groups[1]}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, []Concept{"beer", "object"}, []*Statement{groups[1]}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, []Concept{"object"}, []*Statement{}},
	}
	for _, c := range cases {
		got := c.object.GetDescriptors(c.in...)
		if !ContainsSameRelations(got, c.want) {
			t.Errorf("%v.GetDescriptors(%v) == %v, want %v", c.object, c.in, got, c.want)
		}
	}
}

func TestGetDescriptorsOf(t *testing.T) {
	groups := []*Statement{
		&Statement{"mub", nil, "doer", nil},
		&Statement{"mub", nil, "beer", nil},
		&Statement{"mub", nil, "object", nil},
		&Statement{"vub", nil, "doer", nil},
		&Statement{"vub", nil, "beer", nil},
		&Statement{"vub", nil, "object", nil},
	}
	cases := []struct {
		object Statement
		inA    Concept
		inB    []Concept
		want   []*Statement
	}{
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "mub", []Concept{"doer"}, []*Statement{groups[0]}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "mub", []Concept{"beer"}, []*Statement{groups[1]}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "mub", []Concept{"doer", "beer"}, []*Statement{groups[0], groups[1]}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "mub", []Concept{"beer", "object"}, []*Statement{groups[1]}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "mub", []Concept{"object"}, []*Statement{}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "vub", []Concept{"doer"}, []*Statement{}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "vub", []Concept{"beer"}, []*Statement{}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "vub", []Concept{"doer", "beer"}, []*Statement{}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "vub", []Concept{"beer", "object"}, []*Statement{}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[1]}}, "vub", []Concept{"object"}, []*Statement{}},
		{Statement{"", nil, "", []*Statement{groups[3], groups[1]}}, "mub", []Concept{"doer"}, []*Statement{}},
		{Statement{"", nil, "", []*Statement{groups[4], groups[1]}}, "mub", []Concept{"beer"}, []*Statement{groups[1]}},
		{Statement{"", nil, "", []*Statement{groups[0], groups[3]}}, "mub", []Concept{"doer", "beer"}, []*Statement{groups[0]}},
		{Statement{"", nil, "", []*Statement{groups[3], groups[1]}}, "mub", []Concept{"beer", "object"}, []*Statement{groups[1]}},
	}
	for _, c := range cases {
		got := c.object.GetDescriptorsOf(c.inA, c.inB...)
		if !ContainsSameRelations(got, c.want) {
			t.Errorf("%v.GetDescriptorsOf(%v, %v) == %v, want %v", c.object, c.inA, c.inB, got, c.want)
		}
	}
}

func TestContainsSameRelations(t *testing.T) {
	cases := []struct {
		inA  []*Statement
		inB  []*Statement
		want bool
	}{
		{[]*Statement{
			&Statement{"", nil, "beer", nil},
		}, []*Statement{
			&Statement{"", nil, "beer", nil},
		}, true},
		{[]*Statement{
			&Statement{"", nil, "doer", nil},
			&Statement{"", nil, "beer", nil},
			&Statement{"", nil, "object", nil},
		}, []*Statement{
			&Statement{"", nil, "doer", nil},
			&Statement{"", nil, "beer", nil},
			&Statement{"", nil, "object", nil},
		}, true},
		{[]*Statement{
			&Statement{"", nil, "beer", nil},
			&Statement{"", nil, "doer", nil},
			&Statement{"", nil, "object", nil},
		}, []*Statement{
			&Statement{"", nil, "object", nil},
			&Statement{"", nil, "doer", nil},
			&Statement{"", nil, "beer", nil},
		}, true},
		{[]*Statement{
			&Statement{"", nil, "doer", nil},
			&Statement{"", nil, "beer", nil},
		}, []*Statement{
			&Statement{"", nil, "doer", nil},
			&Statement{"", nil, "beer", nil},
			&Statement{"", nil, "object", nil},
		}, false},
		{[]*Statement{
			&Statement{"", nil, "beer", nil},
		}, []*Statement{
			&Statement{"", nil, "doer", nil},
		}, false},
	}
	for _, c := range cases {
		got := ContainsSameRelations(c.inA, c.inB)
		if got != c.want {
			t.Errorf("ContainsSameRelations(%v, %v) == %v, want %v", c.inA, c.inB, got, c.want)
		}
	}
}

func TestStatementFromString(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"[stat:shine]"},
		{"[stat:shine[doer:sun]]"},
		{"[:shine[doer:sun]]"},
		{"[stat:[mub:shine][doer:sun]]"},
		{"[stat:shine[doer:sun][before:now]]"},
		{"[stat:shine[doer:sun[descriptor:awesome]][before:now]]"},
		{"[stat:[mub:shine][doer:sun[descriptor:awesome]][before:now]]"},
	}
	for _, c := range cases {
		got := StatementFromString(c.in)
		if got.String() != c.in {
			t.Errorf("StatementFromString(%v) == %v, want %v", c.in, got, c.in)
		}
	}
}

func TestSplitByBraces(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"sun", []string{"sun"}},
		{"sun[shine]", []string{"sun", "shine"}},
		{"[shine][sun]", []string{"shine", "sun"}},
		{"[[shine][sun]]", []string{"[shine][sun]"}},
		{"sun[shine][now]", []string{"sun", "shine", "now"}},
		{"sun[shine[mub][wub]][now]", []string{"sun", "shine[mub][wub]", "now"}},
	}
	for _, c := range cases {
		got := SplitByBraces(c.in, '[', ']')
		success := len(got) == len(c.want)
		if success {
			for i := range got {
				if got[i] != c.want[i] {
					success = false
					break
				}
			}
		}
		if !success {
			t.Errorf("SplitByBraces(%v,'[',']') == %v, want %v", c.in, got, c.want)
		}
	}
}
