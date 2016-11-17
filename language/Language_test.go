package language

import (
	"fmt"
	"testing"
)

func TestNavneordGetText(t *testing.T) {
	cases := []struct {
		object Navneord
		want   string
	}{
		{Navneord{&DanishWord{Word{}, "sol", ""}, false, true, nil}, "solen"},
		{Navneord{&DanishWord{Word{}, "sol", ""}, false, false, nil}, "en sol"},
	}
	for _, c := range cases {
		got := c.object.GetText()
		if got != c.want {
			t.Errorf("%v.NavneordGetText() == %v, want %v", c.object, got, c.want)
		}
	}
}

func TestUdsagnsordGetText(t *testing.T) {
	cases := []struct {
		object Udsagnsord
		want   string
	}{
		{Udsagnsord{&DanishWord{Word{}, "skinne", ""}, "nutid"}, "skinner"},
		{Udsagnsord{&DanishWord{Word{}, "skinne", ""}, "datid"}, "skinnede"},
	}
	for _, c := range cases {
		got := c.object.GetText()
		if got != c.want {
			t.Errorf("%v.UdsagnsordGetText() == %v, want %v", c.object, got, c.want)
		}
	}
}

func TestGetTime(t *testing.T) {
	cases := []struct {
		object DanishSentence
		in     StatementGroup
		want   string
	}{
		{DanishSentence{}, StatementGroup{"", nil, "", []*StatementGroup{NewStatementGroup("now", "before")}}, "datid"},
		{DanishSentence{}, StatementGroup{"", nil, "", []*StatementGroup{NewStatementGroup("now", "after")}}, "nutid"},
		{DanishSentence{}, StatementGroup{"", nil, "", []*StatementGroup{NewStatementGroup("now", "at")}}, "nutid"},
		{DanishSentence{}, StatementGroup{"", nil, "", []*StatementGroup{NewStatementGroup("now", "around")}}, "nutid"},
		{DanishSentence{}, StatementGroup{"", nil, "", []*StatementGroup{NewStatementGroup("sun", "doer"), NewStatementGroup("now", "before")}}, "datid"},
		{DanishSentence{}, StatementGroup{"", nil, "", []*StatementGroup{NewStatementGroup("sun", "doer"), NewStatementGroup("now", "after")}}, "nutid"},
		{DanishSentence{}, StatementGroup{"", nil, "", []*StatementGroup{NewStatementGroup("sun", "doer"), NewStatementGroup("now", "at")}}, "nutid"},
		{DanishSentence{}, StatementGroup{"", nil, "", []*StatementGroup{NewStatementGroup("sun", "doer"), NewStatementGroup("now", "around")}}, "nutid"},
	}
	for _, c := range cases {
		fmt.Println(&c.in)
		got := c.object.GetTime(&c.in)
		if got != c.want {
			t.Errorf("%v.GetTime(%v) == %v, want %v", c.object, c.in, got, c.want)
		}
	}
}
