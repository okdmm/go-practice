package exersizeMap

import "strings"

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")

	wc := make(map[string]int)
	for _, word := range words {
		wc[word] = 0
	}

	for _, word := range words {
		wc[word] += 1
	}

	return wc
}
