package oneway

import (
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsOneWay(str1, str2 string) bool {
	changes := len(str1) - len(str2)
	if abs(changes) > 1 {
		return false
	} else if changes == 0 {
		edits := 0
		for i, j := 0, 0; i < len(str1) && j < len(str2); i, j = i+1, j+1 {
			if str1[i] != str2[j] {
				edits++
				if edits > 1 {
					return false
				}
			}

		}
		return true
	} else {
		return strings.HasPrefix(str1, str2) || strings.HasSuffix(str1, str2) ||
			strings.HasPrefix(str2, str1) || strings.HasSuffix(str2, str1)
	}
}
