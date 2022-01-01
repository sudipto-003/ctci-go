package urlify

func URLify(str string, l int) string {
	s := make([]byte, 0)
	c := []byte("%20")
	var space byte = ' '
	for i := 0; i < l; i++ {
		if str[i] == space {
			s = append(s, c...)
		} else {
			s = append(s, str[i])
		}
	}

	return string(s)
}
