package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
160.相交链表
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]struct{}{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = struct{}{}
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if _, ok := vis[tmp]; ok {
			return tmp
		}
	}
	return nil
}

/**
167.
给定一个已按照 非递减顺序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。

函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length 。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
*/
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left, right := i+1, len(numbers)-1
		for left <= right {
			middle := (left + right) / 2
			if target == numbers[i]+numbers[middle] {
				return []int{i + 1, middle + 1}
			} else if numbers[middle] < target-numbers[i] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return []int{}
}

/**
168.
给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
*/
func convertToTitle(columnNumber int) string {
	n := 26
	u := 65
	var res string
	for i := columnNumber; i > 0; i /= n {
		i--
		res = string(byte(i%n+u)) + res
	}
	return res
}

/**
169. 多数元素
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] += 1
			continue
		}
		m[v] = 1
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

/**
171. Excel 表列序号
给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回该列名称对应的列序号。
*/
func titleToNumber(columnTitle string) int {
	n := 0
	for i := 0; i < len(columnTitle); i++ {
		tem := int(columnTitle[i] - 64)
		for j := len(columnTitle) - 1 - i; j > 0; j-- {
			tem *= 26
		}
		n += tem
	}
	return n
}
