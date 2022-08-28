package rstring

import (
	"sort"
	"strings"
	"unicode"
)

func ReplaceAllLike(s string, rm map[string]string) string {
	res := s
	type pair struct {
		old, new string
	}
	var pairs []*pair
	for so, sn := range rm {
		pairs = append(pairs, &pair{so, sn})
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		return len(pairs[i].old) > len(pairs[j].old)
	})
	for _, p := range pairs {
		res = ReplaceLike(res, p.old, p.new)
	}
	return res
}

func ReplaceLike(s, olds, news string) string {
	var (
		final string
		res   = s
	)
	for {
		idx := indexWithoutCase(res, olds)
		if idx < 0 {
			final += res
			break
		}
		ori := res[idx : idx+len(olds)]
		newr := mockString(ori, news)
		tmp := strings.Replace(res, ori, newr, 1)

		res = tmp[idx+len(newr):]
		final += tmp[0 : idx+len(newr)]
	}
	return final
}

func mockString(src, dest string) string {
	if isAllLower(src) {
		return strings.ToLower(dest)
	} else if isAllUpper(src) {
		return strings.ToUpper(dest)
	} else if isAllTitle(src) {
		return strings.Title(dest)
	} else {
		return dest
	}
}

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
