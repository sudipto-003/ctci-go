package oneway

import "testing"

type test struct {
	s1 string
	s2 string
	v  bool
}

var testcases = []test{
	{"lame", "ame", true},
	{"ame", "lame", true},
	{"lame", "lime", true},
	{"lame", "lamey", true},
	{"lame", "am", false},
	{"lame", "", false},
	{"lame", "llamee", false},
	{"", "l", true},
	{"l", "", true},
	{"lame", "kane", false},
	{"lame", "la", false},
}

func TestOneWay(t *testing.T) {
	for _, test := range testcases {
		if r := IsOneWay(test.s1, test.s2); r != test.v {
			t.Errorf("%v %v are not one way", test.s1, test.s2)
		}
	}
}
