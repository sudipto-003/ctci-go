package isPermutation

import "testing"

type testcase struct {
	str1   string
	str2   string
	result bool
}

var testcases = []testcase{
	{"abcd", "acdb", true},
	{"mghtio", "iogmth", true},
	{"krtiop", "poitrk", true},
	{"ttyoop", "tyop", false},
	{"jklmno", "abcdef", false},
	{"r t u o", "rtuo", false},
	{"fjyuoo", "OyfjuO", false},
	{"p  r t", " tr  p", true},
	{"abc", "ab", false},
}

func Test(t *testing.T) {
	for _, test := range testcases {
		if output := IsPermutation(test.str1, test.str2); output != test.result {
			t.Errorf("Expected result %v for %v and %v, but %v.", test.result, test.str1, test.str2, output)
		}
	}
}
