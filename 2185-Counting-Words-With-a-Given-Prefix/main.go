package main

import "strings"

func prefixCount(words []string, pref string) int {
	var ans int
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			ans++
		}
	}
	return ans
}
