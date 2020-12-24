package main

import (
	"fmt"
	"math"
)

func firstUniqChar(s string) int {
	if len(s) <= 0 {
		return -1
	}
	slice := [26]int{}
	for i, _ := range slice {
		slice[i] = -1
	}
	for index, ch := range s {
		if slice[ch-'a'] >= 0 {
			slice[ch-'a'] = -2
		} else if slice[ch-'a'] == -2 {
			continue
		} else {
			slice[ch-'a'] = index
		}
	}
	fmt.Println(slice)
	minIndex := math.MaxInt32
	for _, index := range slice {
		if index >= 0 && index < minIndex {
			minIndex = index
		}
	}
	if minIndex == math.MaxInt32 {
		return -1
	}
	return minIndex
}
