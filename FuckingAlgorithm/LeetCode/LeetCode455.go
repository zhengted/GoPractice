package main

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) <= 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	for _, childrenFit := range g {
		for i, size := range s {
			if childrenFit <= size {
				res++
				s = append(s[:i], s[i+1:]...)
				break
			}
		}
	}
	return res
}
