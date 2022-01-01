package palindromepermutation

import (
	"testing"
)

type test struct {
	s      string
	result bool
}

var testcases = []test{
	{"a", true},
	{"baa", true},
	{"bjhjbhkk", true},
	{"abcd", false},
	{"aaaaaaa", true},
	{"abcab", true},
	{"Tact coa", true},
}

func TestPalimdrome(t *testing.T) {
	for _, test := range testcases {
		if r := IsPlaindromePossible(test.s); r != test.result {
			t.Errorf("Error on %v", test.s)
		}
	}
}
