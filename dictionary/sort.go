package dictionary

import "sort"

type descStringSlice []string

func (s descStringSlice) Sort() {
	sort.Sort(s)
}

func (s descStringSlice) Len() int {
	return len(s)
}

func (s descStringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s descStringSlice) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
