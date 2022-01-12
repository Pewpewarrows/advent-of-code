package advent

import (
    "sort"
)

// see: https://stackoverflow.com/a/22698017
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

// SortedString returns the input string with its runes sorted from left to
// right
func SortedString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}
