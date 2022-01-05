package stringRotation

import "strings"

func IsRotation(s1, s2 string) bool {
	ls1, ls2 := len(s1), len(s2)

	if ls1 == ls2 && ls1 > 0 {
		ss := s1 + s1
		return strings.Contains(ss, s2)
	}

	return false
}
