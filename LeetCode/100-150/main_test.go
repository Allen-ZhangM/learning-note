package main

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	r := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(isSymmetric(r))
}

func TestSortedArrayToBST(t *testing.T) {
	in := []int{-10, -3, 0, 5, 9}
	r := sortedArrayToBST(in)
	fmt.Println(r)
}

func TestMinDepth(t *testing.T) {
	r := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   17,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(minDepth(r))
}

func TestHasPathSum(t *testing.T) {
	r := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(hasPathSum(r, 3))
}
