package main

import (
	"fmt"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
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
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(binaryTreePaths(root))
}

func TestAddDigits(t *testing.T) {
	fmt.Println(addDigits(10))
}

func TestMoveZeroes(t *testing.T) {
	moveZeroes2([]int{0, 1, 0, 3, 12})
}
