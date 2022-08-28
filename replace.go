package rstring

import (
	"sort"
	"strings"
	"unicode"
)

func ReplaceAllLike(s string, rm map[string]string) string {
	// sort replacement map by length
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
	// compare and replace until no match found
	var (
		final strings.Builder
		tmp   = s
	)
	for {
		var (
			idx        int
			olds, news string
		)
		// find the longest match
		for _, p := range pairs {
			idx = strIndexFold(tmp, p.old)
			if idx >= 0 {
				olds, news = p.old, p.new
				break
			}
		}
		// no match found
		if idx < 0 {
			final.WriteString(tmp)
			break
		}
		ori := tmp[idx : idx+len(olds)]
		newr := mockString(ori, news)
		final.WriteString(tmp[0:idx])
		final.WriteString(newr)
		tmp = tmp[idx+len(olds):]
	}
	return final.String()
}

func ReplaceLike(s, olds, news string) string {
	var (
		final strings.Builder
		tmp   = s
	)
	// compare and replace until no match found
	for {
		idx := strIndexFold(tmp, olds)
		if idx < 0 {
			final.WriteString(tmp)
			break
		}
		ori := tmp[idx : idx+len(olds)]
		newr := mockString(ori, news)
		final.WriteString(tmp[0:idx])
		final.WriteString(newr)
		tmp = tmp[idx+len(olds):]
	}
	return final.String()
}

func mockString(src, dest string) string {
	if isAllLower(src) {
		return strings.ToLower(dest)
	} else if isAllUpper(src) {
		return strings.ToUpper(dest)
	} else if isAllTitle(src) {
		return strings.Title(dest)
	} else {
		// no changes
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

func strIndexFold(s string, substr string) int {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Index(s, substr)
}
