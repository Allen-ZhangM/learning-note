package main

import (
	"fmt"
	"testing"
)

func printListNode(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, "->")
		node = node.Next
	}
	fmt.Println()
}

func sliceToListNode(in []int) *ListNode {
	resp := &ListNode{
		Val: -1,
	}
	cur := resp
	for _, v := range in {
		cur.Next = &ListNode{
			Val: v,
		}
		cur = cur.Next
	}
	return resp.Next
}

func TestMergeTrees(t *testing.T) {
	t1 := &TreeNode{}
	t2 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(mergeTrees(t1, t2))
}

func TestMergeKLists(t *testing.T) {
	printListNode(mergeKLists([]*ListNode{sliceToListNode([]int{1, 4, 5}), sliceToListNode([]int{1, 3, 4}), sliceToListNode([]int{2, 6})}))
}
