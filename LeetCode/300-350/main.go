package main

/**
300. 最长递增子序列
*/
//暴力递归返回最长数组
func lengthOfLIS(nums []int) int {
	res := 0

	var f func(sli []int, index int) []int

	f = func(sli []int, index int) []int {
		if index == len(nums) {
			return sli
		}

		var r1, r2 []int
		if len(sli) == 0 || nums[index] > sli[len(sli)-1] {
			r2 = f(append(sli, nums[index]), index+1)
		}

		r1 = f(sli, index+1)

		if len(r1) > len(r2) {
			return r1
		}
		return r2
	}

	s := f([]int{}, 0)
	res = len(s)

	return res
}

//暴力递归求最长
func lengthOfLIS2(nums []int) int {
	res := 0

	var f func(pre, index int) int

	f = func(pre, index int) int {
		if index == len(nums) {
			return 0
		}
		var r1, r2 int
		if pre == -1 || nums[pre] < nums[index] {
			r1 = f(index, index+1) + 1
		}

		r2 = f(pre, index+1)

		if r1 > r2 {
			return r1
		}
		return r2
	}

	res = f(-1, 0)

	return res
}

//动态规划遍历
func lengthOfLIS3(nums []int) int {
	resp := 0
	dp := make([]int, len(nums))
	for index, num := range nums {
		max := 0
		for i := 0; i < index; i++ {
			if num > nums[i] {
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
		dp[index] = max + 1
		if dp[index] > resp {
			resp = dp[index]
		}
	}
	return resp
}

/**
338. 比特位计数
*/
func countBits(n int) []int {
	nums := make([]int, n+1)
	for i := 0; i < len(nums); i++ {
		var c int
		num := i
		for ; num > 0; num &= num - 1 {
			c++
		}
		nums[i] = c
	}
	return nums
}
