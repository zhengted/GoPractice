package main

import "strings"

func wordPattern(pattern string, s string) bool {
	split := strings.Split(s, " ")
	m := make(map[string]uint8)
	mt := make(map[uint8]string)
	if len(pattern) != len(split) {
		return false
	}
	for i, ls := range split {
		if m[ls] == 0 && mt[pattern[i]] == "" {
			m[ls] = pattern[i]
			mt[pattern[i]] = ls
		} else {
			if pattern[i] != m[ls] || mt[pattern[i]] != ls {
				return false
			}
		}
	}

	return true
}
