package util

import "sort"

type StringSet struct{ sort.StringSlice }

func (s *StringSet) Add(v string) bool {
	if i := s.Search(v); i >= s.Len() || v != s.StringSlice[i] {
		s.insert(i, v)
		return true
	}
	return false
}

func (s *StringSet) insert(i int, v string) {
	s.StringSlice = append(s.StringSlice, "")
	copy(s.StringSlice[i+1:], s.StringSlice[i:])
	s.StringSlice[i] = v
}
