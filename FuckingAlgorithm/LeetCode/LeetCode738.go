package main

import (
	"strconv"
)

func monotoneIncreasingDigits(N int) int {
	if N <= 9 {
		return N
	}
	for {
		if N <= 9 {
			return N
		}
		s := strconv.Itoa(N)
		tempB := []byte(s)
	inner:
		for i, ch := range tempB {
			if i == 0 {
				continue
			}
			if i > 0 && !(uint8(ch) < s[i-1]) {
				continue
			}
			N = GetNextN(tempB, i, N)
			break inner
		}
		if !judgeString(strconv.Itoa(N)) {
			N -= 1
		} else {
			return N
		}
	}
	return N
}

func GetNextN(s []byte, i int, N int) int {
	for ; i < len(s); i++ {
		(s)[i] = '0'
	}
	tempS := string(s)
	res, _ := strconv.Atoi(tempS)
	return res
}

func judgeString(s string) bool {
	for i, ch := range s {
		if i == 0 {
			continue
		}
		if i >= 1 && !(uint8(ch) < s[i-1]) {
			continue
		}
		return false
	}
	return true
}
