package main

/**
695. 岛屿的最大面积
*/
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	ni := []int{0, 1, 0, -1}
	nj := []int{1, 0, -1, 0}
	var nextFunc func(grid [][]int, i, j int) int
	nextFunc = func(grid [][]int, i, j int) int {
		if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 0
		tempres := 1
		for index := 0; index < 4; index++ {
			nexti := i + ni[index]
			nextj := j + nj[index]
			tempres += nextFunc(grid, nexti, nextj)
		}
		return tempres
	}
	var maxFunc func(i, j int) int
	maxFunc = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			res = maxFunc(res, nextFunc(grid, i, j))
		}
	}
	return res
}
