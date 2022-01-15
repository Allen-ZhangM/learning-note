package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
257. 二叉树的所有路径
*/
func binaryTreePaths(root *TreeNode) []string {
	var resp []string
	var f func(node *TreeNode, str string)

	f = func(node *TreeNode, str string) {
		if node == nil {
			return
		}
		s := fmt.Sprintf("%s%d", str, node.Val)
		if node.Left == nil && node.Right == nil {
			resp = append(resp, s)
		}
		if node.Left != nil {
			f(node.Left, s+"->")
		}
		if node.Right != nil {
			f(node.Right, s+"->")
		}
	}

	f(root, "")

	return resp
}

/**
258. 各位相加
*/
func addDigits(num int) int {
	//拆每个位数，放入数组
	//一直遍历数组直到只有一个元素且是个位数
	if num < 10 {
		return num
	}
	var nums []int
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	var resp int
	for len(nums) > 1 {
		resp = nums[0] + nums[1]
		if resp >= 10 {
			nums = append(nums, resp/10)
		}
		nums = append(nums, resp%10)
		nums = nums[2:]
	}
	return resp
}

/**
283. 移动零
*/
func moveZeroes(nums []int) {
	var z []int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			z = append(z, i)
		} else if len(z) != 0 {
			nums[z[0]], nums[i] = nums[i], nums[z[0]]
			z = z[1:]
		}
	}
	fmt.Println(nums)
}
func moveZeroes2(nums []int) {
	var z int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[z], nums[i] = nums[i], nums[z]
			z++
		}
	}
	fmt.Println(nums)
}
