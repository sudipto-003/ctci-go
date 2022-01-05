package comprehenssion

import (
	"strings"
	"testing"
)

type test struct {
	s  string
	cs string
}

var testcases = []test{
	{"aaaabbaaa", "a4b2a3"},
	{"aaaabbaaaaad", "a4b2a5d1"},
	{"", ""},
	{"abcd", "abcd"},
	{"aaaaaaaaaaaaa", "a13"},
}

func TestCompression(t *testing.T) {
	for _, test := range testcases {
		if r := Compress(test.s); strings.Compare(r, test.cs) != 0 {
			t.Errorf("Wrong Compressed %v => %v", test.s, r)
		}
	}
}
