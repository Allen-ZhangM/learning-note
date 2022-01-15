package main

/**
300. 最长递增子序列
*/
func lengthOfLIS(nums []int) int {
	res := 0

	var f func(sli []int, index int) []int

	f = func(sli []int, index int) []int {
		if sli == nil {
			return nil
		}
		if index >= len(nums) {
			return sli
		}

		if len(sli) >= 1 && nums[index] <= sli[len(sli)-1] {
			return nil
		}

		var r1, r2 []int

		r2 = f(append(sli, nums[index]), index+1)
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
func lengthOfLIS2(nums []int) int {
	res := 0

	var f func(index int) int

	f = func(index int) int {
		if index >= len(nums) {
			return 0
		}

		r1 := f(index + 1)
		r2 := f(index + 1)

		if r1 > r2 {
			return r1
		}
		return r2
	}

	res = f(0)

	return res
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
