package main

import "fmt"

func main() {
	num53()
}

func num53() {
	//nums:=[]int{1,5,-1,6}
	nums := []int{1, 5, -10, 6, 7}
	fmt.Println(maxSubArray(nums))
}

/**
https://leetcode-cn.com/problems/maximum-subarray/
*/
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
