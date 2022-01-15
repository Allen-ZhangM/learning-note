package main

/**
461. 汉明距离
*/
func hammingDistance(x int, y int) int {
	i := x ^ y
	var resp int
	for ; i > 0; i &= i - 1 {
		resp++
	}
	return resp
}
