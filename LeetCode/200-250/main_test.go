package main

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	r := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	resp := invertTree(r)
	fmt.Println(resp)
}

func TestIsPalindrome(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	fmt.Println(isPalindrome(node))
}
