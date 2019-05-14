package main

import (
	"fmt"
	"os"
)

func readMaze(path string) [][]int {
	file, e := os.Open(path)
	if e != nil {
		panic(e)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

func (p point) add(cur point) point {
	return point{p.i + cur.i, p.j + cur.j}
}
func (p point) at(maze [][]int) bool {
	if p.i < 0 || p.i >= len(maze) {
		return false
	}
	if p.j < 0 || p.j >= len(maze[0]) {
		return false
	}
	if maze[p.i][p.j] == 0 {
		return true
	}
	return false
}

var dirs = []point{
	point{-1, 0},
	point{0, -1},
	point{1, 0},
	point{0, 1},
}

func main() {
	maze := readMaze("GoPractice/maze/maze.in")
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%d ", maze[i][j])
		}
		fmt.Println()
	}
	startPoint := point{2, 2}
	endPoint := point{2, 2}
	//steps:=walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})
	steps := walk(maze, startPoint, endPoint)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}

	//printSteps(steps,endPoint)

}

func walk(maze [][]int, start point, end point) [][]int {
	var steps = make([][]int, len(maze))
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

			if next.at(maze) && steps[next.i][next.j] == 0 && next != start {
				Q = append(Q, next)
				steps[next.i][next.j] = steps[cur.i][cur.j] + 1
			}

		}

	}
	return steps
}
