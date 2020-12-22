package main

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		nCurLen := i + 1
		ret[i] = make([]int, nCurLen)
		ret[i][0] = 1
		ret[i][len(ret[i])-1] = 1
	}

	for j := 2; j < numRows; j++ {
		for i := 1; i < j; i++ {
			ret[j][i] = ret[j-1][i] + ret[j-1][i-1]
		}
	}
	//fmt.Println(ret)
	return ret
}
