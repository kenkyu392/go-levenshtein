package levenshtein

import "math"

// Distance returns the distance between characters calculated by the Levenshtein distance algorithm.
// - [Levenshtein Distance](https://en.wikipedia.org/wiki/Levenshtein_distance)
// This implementation is optimized to use O(min(m,n)) space.
// For details, refer to the optimized C implementation here:
// https://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
func Distance(r1, r2 []rune) int {
	var x, y, cost, lastdiag, olddiag int
	len1 := len(r1)
	len2 := len(r2)

	column := make([]int, len1+1)

	for y = 1; y <= len1; y++ {
		column[y] = y
	}
	for x = 1; x <= len2; x++ {
		column[0] = x
		lastdiag = x - 1
		for y = 1; y <= len1; y++ {
			olddiag = column[y]
			if r1[y-1] != r2[x-1] {
				cost = 1
			} else {
				cost = 0
			}
			column[y] = min(column[y]+1, column[y-1]+1, lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return column[len1]
}

// Percent ...
func Percent(r1, r2 []rune) float64 {
	var x, y, cost, lastdiag, olddiag int
	len1 := len(r1)
	len2 := len(r2)

	column := make([]int, len1+1)

	for y = 1; y <= len1; y++ {
		column[y] = y
	}
	for x = 1; x <= len2; x++ {
		column[0] = x
		lastdiag = x - 1
		for y = 1; y <= len1; y++ {
			olddiag = column[y]
			if r1[y-1] != r2[x-1] {
				cost = 1
			} else {
				cost = 0
			}
			column[y] = min(column[y]+1, column[y-1]+1, lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return (1 - (float64(column[len1]) / math.Max(float64(len1), float64(len2)))) * 100
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
