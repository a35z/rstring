package rstring

import (
	"sort"
	"strings"
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