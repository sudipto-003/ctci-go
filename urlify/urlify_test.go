package urlify

import (
	"strings"
	"testing"
)

type test struct {
	s   string
	l   int
	url string
}

var testcases = []test{
	{"Hello World  ", 11, "Hello%20World"},
	{"I am fine    ", 9, "I%20am%20fine"},
	{"Ok i will come tomorrow        ", 23, "Ok%20i%20will%20come%20tomorrow"},
}

func TestURLify(t *testing.T) {
	for _, test := range testcases {
		if r := URLify(test.s, test.l); strings.Compare(r, test.url) != 0 {
			t.Errorf("URL not matched for %v %v", test.url, r)
		}
	}
}
