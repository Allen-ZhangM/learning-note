package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
23. 合并K个升序链表
*/
func mergeKLists(lists []*ListNode) *ListNode {
	resp := &ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := resp
	for {
		min := math.MaxInt32
		minIndex := -1
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if lists[i].Val < min {
					min = lists[i].Val
					minIndex = i
				}
			}
		}
		if minIndex == -1 {
			break
		}
		cur.Next = &ListNode{
			Val:  min,
			Next: nil,
		}
		cur = cur.Next
		lists[minIndex] = lists[minIndex].Next
	}
	return resp.Next
}

/**
617. 合并二叉树
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var f func(t1, t2 *TreeNode) *TreeNode
	f = func(t1, t2 *TreeNode) *TreeNode {
		if t1 == nil {
			return t2
		}
		if t2 == nil {
			return t1
		}
		t1.Val += t2.Val
		t1.Left = f(t1.Left, t2.Left)
		t1.Right = f(t1.Right, t2.Right)
		return t1
	}
	return f(root1, root2)
}