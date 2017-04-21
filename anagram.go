package anagram

import "strings"

const testVersion = 1

func Detect(subject string, candidates []string) []string {
	outputLen := 0
	subject = strings.ToLower(subject)
	output := make([]string, len(candidates))
	for _, candidate := range candidates {
		candidate = strings.ToLower(candidate)
		if subject == candidate || len(subject) != len(candidate) {
			continue
		}
		if IsAnagram(subject, candidate) {
			output[outputLen] = candidate
			outputLen++
		}
	}
	return output[0:outputLen]
}
func IsAnagram(subject, candidate string) bool {
	c := make([]rune, len(candidate))
	c = []rune(candidate)
	for _, r := range subject {
		if p := strings.IndexRune(string(c), r); p < 0 {
			return false
		} else {
			c[p] = '_'
		}
	}
	return true
}
