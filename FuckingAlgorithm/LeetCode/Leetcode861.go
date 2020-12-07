package main

import (
	"math"
)

const (
	ReverseCol = 0
	ReverseRow = 1
)

func matrixScore(A [][]int) int {
	if len(A) <= 0 {
		return 0
	}
	for i := 0; i < len(A); i++ {
		if A[i][0] == 0 {
			reverse(A, i, ReverseRow)
		}
	}
	for j := 1; j < len(A[0]); j++ {
		nZeroCount := 0
		for i := 0; i < len(A); i++ {
			if A[i][j] == 0 {
				nZeroCount++
			}
		}
		if nZeroCount > len(A)/2 {
			reverse(A, j, ReverseCol)
		}
	}
	nMax := CalMatrixNum(A)
	return nMax
}

func reverse(A [][]int, line int, reverseFlag int) {
	if reverseFlag == ReverseCol && line < len(A[0]) {
		// 改变第line列
		for i := 0; i < len(A); i++ {
			if A[i][line] == 0 {
				A[i][line] = 1
			} else {
				A[i][line] = 0
			}
		}
		return
	}
	if reverseFlag == ReverseRow && line < len(A) {
		// 改变第line列
		for i := 0; i < len(A[0]); i++ {
			if A[line][i] == 0 {
				A[line][i] = 1
			} else {
				A[line][i] = 0
			}
		}
		return
	}
}

func CalMatrixNum(A [][]int) int {
	ret := 0
	for _, sub := range A {
		temp := 0
		for j, val := range sub {
			if val == 1 {
				nCur := int(math.Pow(2, float64(len(sub)-1-j)))
				temp = nCur + temp
			}
		}
		ret += temp
	}
	return ret
}
