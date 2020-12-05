package main

import (
	"fmt"
	"os"
)

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Fscanln(file)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
		fmt.Fscanln(file)
	}

	return maze
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != START
			val, ok := next.at(maze) // 判断边界
			if !ok || val == 1 {     // val = 1 撞墙
				continue
			}

			val, ok = next.at(steps) // 判断边界
			if !ok || val != 0 {     // val != 0 走过
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}

	return steps
}

func paintRoad(steps [][]int, end point) []point {
	var ret = make([]point, 0)
	ret = append(ret, end)
	cur := end
	for cur.j > 0 || cur.i > 0 {
		for _, dir := range dirs {
			next := cur.add(dir)
			nextVal, ok := next.at(steps)
			if !ok {
				continue
			}
			curVal, _ := cur.at(steps)
			if nextVal != curVal-1 {
				continue
			}
			ret = append(ret, next)
			cur = next
		}
	}
	return ret
}

func main() {
	maze := readMaze("maze/maze.in")

	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%3d", steps[i][j])
		}
		fmt.Println()
	}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	res := paintRoad(steps, end)
	fmt.Println(res)
}
