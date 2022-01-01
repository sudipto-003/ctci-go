package isUnique

import "testing"

type StringUniqueTest struct {
	str    string
	result bool
}

var testcases = []StringUniqueTest{
	StringUniqueTest{"abcd", true},
	StringUniqueTest{"abcc", false},
	StringUniqueTest{"cdahbc", false},
	StringUniqueTest{"cahbed", true},
	StringUniqueTest{"cah bed", true},
	StringUniqueTest{"cah be d", false},
	StringUniqueTest{"cah5bed_", true},
}

func TestUniqueCharacters(t *testing.T) {
	for _, test := range testcases {
		if output := Test_unique(test.str); output != test.result {
			t.Errorf("Output %v not match for input %v, expected %v", output, test.str, test.result)
		}
	}
}

func TestUniqueCharacters2(t *testing.T) {
	for _, test := range testcases {
		if output := Test_unique_onthego(test.str); output != test.result {
			t.Errorf("Output %v not match for input %v, expected %v", output, test.str, test.result)
		}
	}
}
