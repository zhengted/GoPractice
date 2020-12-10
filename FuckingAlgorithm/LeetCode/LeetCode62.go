package main

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	grid := make([][]int, 100)
	for i := 0; i < 100; i++ {
		grid[i] = make([]int, 100)
	}
	grid[0][0] = 1
	for i := 1; i < m; i++ {
		grid[i][0] = 1
	}
	for j := 1; j < n; j++ {
		grid[0][j] = 1
	}
	for i := 1; i < m; i++ { // 行
		for j := 1; j < n; j++ { // 列
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}

func uniquePathsEx(m int, n int) int {
	var grid [][]int

	//grid[0] = append(grid[0],1)
	for i := 0; i < m; i++ {
		var temp []int
		for j := 0; j < n; j++ {
			temp = append(temp, 1)
		}
		grid = append(grid, temp)
	}

	for i := 1; i < m; i++ { // 行
		for j := 1; j < n; j++ { // 列
			//grid[i] = append(grid[i], grid[i-1][j]+grid[i][j-1])
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}
