package types

import (
	"cmp"
	"slices"
)

type OutResult struct {
	Path  string
	Name  string
	Prio  int8
	Alias string
}

func SortOutResult(ors ...[]OutResult) []OutResult {
	orsBig := slices.Concat(ors...)
	slices.SortFunc(orsBig, func(a, b OutResult) int {
		n := cmp.Compare(b.Path, a.Path)
		if n == 0 {
			return cmp.Compare(b.Prio, a.Prio)
		}
		return n
	})
	orsBig = slices.CompactFunc(orsBig, func(a, b OutResult) bool {
		if n := cmp.Compare(b.Path, a.Path); n == 0 {
			return true
		}
		return false
	})

	slices.SortFunc(orsBig, func(a, b OutResult) int {
		return cmp.Compare(b.Prio, a.Prio)
	})

	return orsBig
}

func SearchOutResult(search string, outs []OutResult) (index int, found bool) {
	for i, out := range outs {
		if out.Alias == search {
			return i, true
		}
	}
	return len(outs), false
}

func SearchOutResultAttached(outs []OutResult) (index int, found bool) {
	for i, out := range outs {
		if out.Prio == 1 {
			return i, true
		}
	}
	return len(outs), false
}
