package main

import (
	"fmt"
	"math"
)

func main() {
	recursion()
}

func recursion() {
	f(0, 0)
	fmt.Println(maxW)
}

var (
	maxW   = math.MinInt32
	weight = []int{2, 2, 4, 6, 3}
	n      = len(weight)
	w      = 9
)

func f(i, cw int) {
	if cw == w || i == n {
		if cw >= maxW {
			maxW = cw
		}
		return
	}
	f(i+1, cw)
	if cw+weight[i] < w {
		f(i+1, cw+weight[i])
	}
}

//weight:物品重量，n:物品个数，w:背包可承载重量
func knapsack(weight []int, n, w int) {

}
