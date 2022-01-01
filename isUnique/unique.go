package isUnique

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func Test_unique(str string) bool {
	flag := make([]bool, 128)
	for _, r := range str {
		idx := int(r)
		if flag[idx] {
			return false
		}
		flag[idx] = true
	}

	return true
}

func Test_unique_onthego(str string) bool {
	runes := []rune(str)
	sort.Sort(sortRunes(runes))
	sorted_str := string(runes)

	for i := 0; i < len(sorted_str)-1; i++ {
		if sorted_str[i] == sorted_str[i+1] {
			return false
		}
	}

	return true
}
