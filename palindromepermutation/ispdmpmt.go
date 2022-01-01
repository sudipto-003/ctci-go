package palindromepermutation

import (
	"strings"
)

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}

	return s
}

func IsPlaindromePossible(s string) bool {
	str := strings.Replace(s, " ", "", -1)
	str = strings.ToLower(str)
	counts := make([]int, 128)

	for _, r := range str {
		idx := int(r)
		if counts[idx] == 0 {
			counts[idx]++
		} else {
			counts[idx]--
		}
	}

	check := sum(counts)
	if len(str)%2 == 0 && check == 0 {
		return true
	}
	if len(str)%2 != 0 && check == 1 {
		return true
	}

	return false
}
