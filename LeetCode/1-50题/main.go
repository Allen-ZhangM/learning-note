package main

import (
	"fmt"
	"go-common-master/app/admin/main/up/util/mathutil"
)

func main() {
	//num1()
	//num2()
	//num3()
	//num4()
	num5()
	//num7()
	num7()
	//numtest123456
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
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

*/

func num4() {
	fmt.Println("findMedianSortedArrays : ", findMedianSortedArrays([]int{3}, []int{1, 2, 4, 5}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var p1, p2, mid, ln1, ln2 int
	var isodd bool
	var sums []int
	ln1 = len(nums1)
	ln2 = len(nums2)
	mid = (ln1 + ln2) / 2
	if (ln1+ln2)%2 != 0 {
		isodd = true
	}
	for p1+p2 <= ln1+ln2 {
		if p1 < ln1 && p2 < ln2 {
			if nums1[p1] < nums2[p2] {
				sums = append(sums, nums1[p1])
				p1++
			} else {
				sums = append(sums, nums2[p2])
				p2++
			}
		} else if p1 >= ln1 {
			sums = append(sums, nums2[p2:]...)
			break
		} else if p2 >= ln2 {
			sums = append(sums, nums1[p1:]...)
			break
		}
	}
	if isodd {
		return float64(sums[mid])
	} else {
		if mid != 0 {
			return float64(sums[mid-1]+sums[mid]) / 2
		} else {
			return float64(sums[mid]+sums[mid+1]) / 2
		}
	}
}

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

*/

func num5() {
	fmt.Println("longestPalindrome:", longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, end int
	for i, _ := range s {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		max := mathutil.Max(l1, l2)
		if max > end-start {
			start = i - (max-1)/2
			end = i + max/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, l, r int) (length int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
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
