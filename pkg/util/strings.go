package util

import "strings"

func Capitalize(s *string) {
	if len(*s) == 0 {
		return
	}

	*s = strings.ToUpper((*s)[0:1]) + (*s)[1:]
}

func Normalize(s *string) {
	words := strings.Fields(*s)
	for i, word := range words {
		word = strings.ToLower(word)
		Capitalize(&word)
		words[i] = word
	}
	*s = strings.Join(words, "")
}