package comprehenssion

import "strconv"

func Compress(str string) string {
	s := str + "."
	w := make([]byte, 0)
	l := len(s)
	c := 1
	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			c++
		} else {
			w = append(w, s[i])
			cb := []byte(strconv.Itoa(c))
			w = append(w, cb...)
			c = 1
		}
	}
	if len(w) >= len(str) {
		return str
	} else {
		return string(w)
	}
}
