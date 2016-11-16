package language

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
		in   StatementGroup
		want bool
	}{
		{StatementGroup{"asdf", nil, "", nil}, false},
		{StatementGroup{"", nil, "", nil}, false},
		{StatementGroup{"asdf", new(StatementGroup), "", nil}, true},
		{StatementGroup{"", new(StatementGroup), "", nil}, true},
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
		object StatementGroup
		in     []Concept
		want   bool
	}{
		{StatementGroup{"", nil, "relation", nil}, []Concept{"doer", "beer"}, false},
		{StatementGroup{"", nil, "beer", nil}, []Concept{"doer", "beer"}, true},
		{StatementGroup{"", nil, "relation", nil}, []Concept{}, false},
		{StatementGroup{"", nil, "beer", nil}, []Concept{"beer"}, true},
	}
	for _, c := range cases {
		got := c.object.HasRelation(c.in...)
		if got != c.want {
			t.Errorf("%v.HasRelation(%v) == %v, want %v", c.object, c.in, got, c.want)
		}
	}
}

func TestGetDescriptors(t *testing.T) {
	groups := []*StatementGroup{
		&StatementGroup{"", nil, "doer", nil},
		&StatementGroup{"", nil, "beer", nil},
		&StatementGroup{"", nil, "object", nil},
	}
	cases := []struct {
		object StatementGroup
		in     []Concept
		want   []*StatementGroup
	}{
		{StatementGroup{"", nil, "", []*StatementGroup{groups[0], groups[1]}}, []Concept{"doer"}, []*StatementGroup{groups[0]}},
		{StatementGroup{"", nil, "", []*StatementGroup{groups[0], groups[1]}}, []Concept{"beer"}, []*StatementGroup{groups[1]}},
		{StatementGroup{"", nil, "", []*StatementGroup{groups[0], groups[1]}}, []Concept{"doer", "beer"}, []*StatementGroup{groups[0], groups[1]}},
		{StatementGroup{"", nil, "", []*StatementGroup{groups[0], groups[1]}}, []Concept{"beer", "object"}, []*StatementGroup{groups[1]}},
		{StatementGroup{"", nil, "", []*StatementGroup{groups[0], groups[1]}}, []Concept{"object"}, []*StatementGroup{}},
	}
	for _, c := range cases {
		got := c.object.GetDescriptors(c.in...)
		if !ContainsSameRelations(got, c.want) {
			t.Errorf("%v.GetDescriptors(%v) == %v, want %v", c.object, c.in, got, c.want)
		}
	}
}

func TestContainsSameRelations(t *testing.T) {
	cases := []struct {
		inA  []*StatementGroup
		inB  []*StatementGroup
		want bool
	}{
		{[]*StatementGroup{
			&StatementGroup{"", nil, "beer", nil},
		}, []*StatementGroup{
			&StatementGroup{"", nil, "beer", nil},
		}, true},
		{[]*StatementGroup{
			&StatementGroup{"", nil, "doer", nil},
			&StatementGroup{"", nil, "beer", nil},
			&StatementGroup{"", nil, "object", nil},
		}, []*StatementGroup{
			&StatementGroup{"", nil, "doer", nil},
			&StatementGroup{"", nil, "beer", nil},
			&StatementGroup{"", nil, "object", nil},
		}, true},
		{[]*StatementGroup{
			&StatementGroup{"", nil, "beer", nil},
			&StatementGroup{"", nil, "doer", nil},
			&StatementGroup{"", nil, "object", nil},
		}, []*StatementGroup{
			&StatementGroup{"", nil, "object", nil},
			&StatementGroup{"", nil, "doer", nil},
			&StatementGroup{"", nil, "beer", nil},
		}, true},
		{[]*StatementGroup{
			&StatementGroup{"", nil, "doer", nil},
			&StatementGroup{"", nil, "beer", nil},
		}, []*StatementGroup{
			&StatementGroup{"", nil, "doer", nil},
			&StatementGroup{"", nil, "beer", nil},
			&StatementGroup{"", nil, "object", nil},
		}, false},
		{[]*StatementGroup{
			&StatementGroup{"", nil, "beer", nil},
		}, []*StatementGroup{
			&StatementGroup{"", nil, "doer", nil},
		}, false},
	}
	for _, c := range cases {
		got := ContainsSameRelations(c.inA, c.inB)
		if got != c.want {
			t.Errorf("ContainsSameRelations(%v, %v) == %v, want %v", c.inA, c.inB, got, c.want)
		}
	}
}
