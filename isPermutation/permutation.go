package isPermutation

func IsPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	counts := make([]int, 128)
	for _, r := range str1 {
		idx := int(r)
		counts[idx]++
	}

	for _, r := range str2 {
		idx := int(r)
		counts[idx]--
		if counts[idx] < 0 {
			return false
		}
	}

	return true
}
