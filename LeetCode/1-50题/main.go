package main

import (
	"fmt"
)

func main() {
	//num1()
	//num2()
	//num3()
	num7()
	//numtest
	//numtest

}

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

*/

func num1() {
	fmt.Println("twoSum:", twoSum([]int{1, 2, 7, 11, 15}, 9))
	fmt.Println("twoSum2:", twoSum2([]int{1, 2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for ii := i + 1; ii < len(nums); ii++ {
			if v+nums[ii] == target {
				return []int{i, ii}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func num2() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println("addTwoNumbers:")
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Print(result.Val, "->")
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result
	carry := 0
	for l1 != nil || l2 != nil {
		sum := l1.Val + l2.Val + carry
		carry = sum / 10
		temp := &ListNode{Val: sum % 10}
		curr.Next = temp
		if l1.Next != nil || l2.Next != nil {
			temp.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return result.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val
	nextNode := addTwoNumbers(l1.Next, l2.Next)
	if sum < 10 {
		return &ListNode{Val: sum, Next: nextNode}
	} else {
		tempNode := &ListNode{
			Val:  1,
			Next: nil,
		}
		return &ListNode{
			Val:  sum - 10,
			Next: addTwoNumbers(nextNode, tempNode),
		}
	}
}

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

*/

func num3() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}

func lengthOfLongestSubstring(s string) int {
	temp := []string{}
	max := 0
	for _, v := range s {
		for ii, vv := range temp {
			if string(v) == vv {
				temp = temp[ii+1:]
				break
			}
		}
		temp = append(temp, string(v))
		if max < len(temp) {
			max = len(temp)
		}
	}
	return max
}

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

*/

func num7() {
	fmt.Println("reverse:", reverse(1563847412))
}
func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if y > (1<<31-1) || y < -(1<<31) {
			return 0
		}
		x /= 10
	}
	return y
}
