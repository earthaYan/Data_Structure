package LeetCode_GoLang

import "strings"

func prefixCount(words []string, pref string) int {
	count := 0
	for _, val := range words {
		if strings.HasPrefix(val, pref) {
			count++
		}
	}
	return count
}
