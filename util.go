package rstring

import (
	"strings"
	"unicode"
)

func isAllUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isAllLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isAllTitle(s string) bool {
	for i, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		if i == 0 && !unicode.IsUpper(r) {
			return false
		}
		if i > 0 && !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func indexWithoutCase(s string, substr string) int {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Index(s, substr)
}
