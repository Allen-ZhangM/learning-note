package main

import (
	"fmt"
	//"go-common-master/app/admin/main/up/util/mathutil"
	"math"
	"strconv"
)

func main() {
	//num1()
	//num2()
	//num3()
	//num4()
	//num5()
	//num7()
	//num9()
	//num13()
	//num14()
	//num20()
	num21()
}

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
func num21() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}
	result := mergeTwoLists(l1, l2)
	for result != nil {
		fmt.Print(result.Val, "->")
		result = result.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/
func num20() {
	//isValid("{({}[]())}")
	is := isValid("{({}[]())}")
	fmt.Println(is)
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	//{}[]()
	//[123 125 91 93 40 41]
	for _, v := range []byte(s) {
		switch v {
		case 123, 91, 40:
			stack = append(stack, v)
		case 125:
			if len(stack) == 0 || stack[len(stack)-1] != 123 {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case 93:
			if len(stack) == 0 || stack[len(stack)-1] != 91 {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case 41:
			if len(stack) == 0 || stack[len(stack)-1] != 40 {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
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
	//for i, _ := range s {
	//l1 := expandAroundCenter(s, i, i)
	//l2 := expandAroundCenter(s, i, i+1)
	////max := mathutil.Max(l1, l2)
	//if max > end-start {
	//	start = i - (max-1)/2
	//	end = i + max/2
	//}
	//}
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

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

*/

func num9() {
	fmt.Println(isPalindrome2(10))
}

func isPalindrome(x int) bool {
	str := []byte(strconv.Itoa(x))
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertNum := 0
	for x > revertNum {
		revertNum = revertNum*10 + x%10
		x /= 10
	}
	return x == revertNum || x == revertNum/10
}

/*
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.
*/

func num13() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	sum := 0
	for i, v := range s {
		num := romanStrToInt(fmt.Sprintf("%c", v))
		if i+1 < len(s) {

			if num < romanStrToInt(fmt.Sprintf("%c", s[i+1])) {
				num = -num
			}
		}
		sum += num
	}
	return sum
}

func romanStrToInt(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
*/

func num14() {
	fmt.Println(longestCommonPrefix2([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	maxLen := math.MaxInt32
	prefixStr := ""
	for _, v := range strs {
		if len(v) < maxLen {
			maxLen = len(v)
		}
	}
	for i := 0; i < maxLen; i++ {
		tempStr := ""
		for ii, v := range strs {
			if ii == 0 {
				tempStr = fmt.Sprintf("%c", v[i])
				continue
			}
			if tempStr != fmt.Sprintf("%c", v[i]) {
				return prefixStr
			}
		}
		prefixStr += tempStr
	}
	return prefixStr
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		tempStr := strs[0][i]
		for _, v := range strs {
			if i == len(v) || v[i] != tempStr {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}
