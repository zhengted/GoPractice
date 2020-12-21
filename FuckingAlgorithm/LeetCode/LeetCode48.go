package main

func rotate(matrix [][]int) {
	ret := make([][]int, len(matrix[0]))
	for i, _ := range ret {
		ret[i] = make([]int, len(matrix))
	}
	p := 0
	q := 0
	for j := 0; j < len(matrix[0]); j++ {
		q = 0
		for i := len(matrix[0]) - 1; i >= 0; i-- {
			ret[p][q] = matrix[i][j]
			q++
		}
		p++
	}
	copy(matrix, ret)
}
