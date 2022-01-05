package main

import (
	"fmt"
)

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

/**
58. 最后一个单词的长度
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
*/
func lengthOfLastWord(s string) int {
	var n int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if n != 0 {
				return n
			}
		} else {
			n++
		}
	}
	return n
}

/**
66. 加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
func plusOne(digits []int) []int {
	l := len(digits)
	i := 0
	for i = l - 1; i >= 0 && digits[i] == 9; i-- {
		digits[i] = 0
	}
	if i >= 0 {
		digits[i] += 1
	} else {
		digits = append([]int{1}, digits[:]...)
	}
	return digits
}

/**
67. 二进制求和
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。
*/
func addBinary2(a string, b string) string {
	var numA, numB int
	for i, j := len(a)-1, 1; i >= 0; i-- {
		numA += int(a[i]-48) * j
		j *= 2
	}
	for i, j := len(b)-1, 1; i >= 0; i-- {
		numB += int(b[i]-48) * j
		j *= 2
	}
	if numA+numB == 0 {
		return "0"
	}
	var res = ""
	for i := numA + numB; i > 0; i /= 2 {
		res = string(byte(i%2)+48) + res
	}
	return res
}

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	var lmax int
	var tem string
	if la > lb {
		lmax = la
		for i := 0; i < la-lb; i++ {
			tem += "0"
		}
		b = tem + b
	} else {
		lmax = lb
		for i := 0; i < lb-la; i++ {
			tem += "0"
		}
		a = tem + a
	}
	la, lb = len(a), len(b)
	var next uint8
	res := ""
	for i := 0; i < lmax; i++ {
		n := a[la-1-i] - 48 + b[lb-1-i] - 48
		if n+next <= 1 {
			res = string(n+next+48) + res
			next = 0
			continue
		}
		res = string(n+next-2+48) + res
		next = 1
	}
	if next != 0 {
		res = "1" + res
	}
	return res
}

/**
69. Sqrt(x)
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
*/
func mySqrt(x int) int {
	l, r := 0, x
	res := 0
	for l <= r {
		mid := (l + r) / 2
		m2 := mid * mid
		if m2 == x {
			return mid
		} else if m2 > x {
			res = mid - 1
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

/**
70. 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
*/
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
83. 删除排序链表中的重复元素
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

返回同样按升序排列的结果链表。
*/
func deleteDuplicates(head *ListNode) *ListNode {
	i := head
	for i != nil && i.Next != nil {
		if i.Val == i.Next.Val {
			i.Next = i.Next.Next
		} else {
			i = i.Next
		}
	}
	return head
}

/**
88. 合并两个有序数组
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	tem := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			tem = append(tem, nums2[p2:]...)
			break
		}
		if p2 == n {
			tem = append(tem, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			tem = append(tem, nums1[p1])
			p1++
		} else {
			tem = append(tem, nums2[p2])
			p2++
		}
	}
	copy(nums1, tem)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
94. 二叉树的中序遍历
给定一个二叉树的根节点 root ，返回它的 中序 遍历。
*/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res
}

/**
100. 相同的树
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var f func(p *TreeNode, q *TreeNode)
	is := true
	f = func(p *TreeNode, q *TreeNode) {
		if !is {
			return
		}
		if p == nil && q == nil {
			return
		}
		if (p == nil && q != nil) || (q == nil && p != nil) {
			is = false
			return
		}
		if p.Val != q.Val {
			is = false
			return
		}
		f(p.Left, q.Left)
		f(p.Right, q.Right)
	}
	f(p, q)
	return is
}

/**
92. 反转链表 II
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	c := 0
	var h, m, t, l *ListNode
	resp := &ListNode{
		Val:  -1,
		Next: head,
	}
	h = resp
	for ; c < left-1; c++ {
		h = h.Next
	}

	l = h
	for ; c < right; c++ {
		l = l.Next
	}

	m = h.Next
	t = l.Next

	h.Next = nil
	l.Next = nil

	var reverseLinkedList func(node *ListNode)
	reverseLinkedList = func(node *ListNode) {
		var res *ListNode
		n := node
		for n != nil {
			next := n.Next
			n.Next = res
			res = n
			n = next
		}
	}
	reverseLinkedList(m)

	h.Next = l
	m.Next = t
	return resp.Next
}

func reverseBetween2(head *ListNode, left, right int) *ListNode {
	// 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
	// 建议写在 for 循环里，语义清晰
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 第 3 步：切断出一个子链表（截取链表）
	leftNode := pre.Next
	curr := rightNode.Next

	// 注意：切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 第 4 步：同第 206 题，反转链表的子区间
	//reverseLinkedList(leftNode)

	// 第 5 步：接回到原来的链表中
	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}
