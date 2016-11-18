package grammatics

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
		in     Statement
		want   string
	}{
		{DanishSentence{}, Statement{"", nil, "", []*Statement{NewStatement("now", "before")}}, "datid"},
		{DanishSentence{}, Statement{"", nil, "", []*Statement{NewStatement("now", "after")}}, "nutid"},
		{DanishSentence{}, Statement{"", nil, "", []*Statement{NewStatement("now", "at")}}, "nutid"},
		{DanishSentence{}, Statement{"", nil, "", []*Statement{NewStatement("now", "around")}}, "nutid"},
		{DanishSentence{}, Statement{"", nil, "", []*Statement{NewStatement("sun", "doer"), NewStatement("now", "before")}}, "datid"},
		{DanishSentence{}, Statement{"", nil, "", []*Statement{NewStatement("sun", "doer"), NewStatement("now", "after")}}, "nutid"},
		{DanishSentence{}, Statement{"", nil, "", []*Statement{NewStatement("sun", "doer"), NewStatement("now", "at")}}, "nutid"},
		{DanishSentence{}, Statement{"", nil, "", []*Statement{NewStatement("sun", "doer"), NewStatement("now", "around")}}, "nutid"},
	}
	for _, c := range cases {
		fmt.Println(&c.in)
		got := c.object.GetTime(&c.in)
		if got != c.want {
			t.Errorf("%v.GetTime(%v) == %v, want %v", c.object, c.in, got, c.want)
		}
	}
}
