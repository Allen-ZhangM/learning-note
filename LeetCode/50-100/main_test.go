package main

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(lengthOfLastWord(" a "))
	fmt.Println(lengthOfLastWord("i am ss"))
	fmt.Println(lengthOfLastWord("i am ss  "))
}

func TestPlusOne(t *testing.T) {
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func TestAddBinary(t *testing.T) {
	fmt.Println(string(byte(1) + 48))
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}

func TestMySqrt(t *testing.T) {
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(6))
	fmt.Println(mySqrt(8))
}

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(8))
}

func TestDeleteDuplicates(t *testing.T) {
	r := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	res := deleteDuplicates(r)
	for i := res; i != nil; i = i.Next {
		fmt.Println(i.Val, i.Next)
	}
}

func TestMerge(t *testing.T) {
	s1 := []int{1, 2, 3, 0, 0, 0}
	merge(s1, 3, []int{2, 5, 6}, 3)
	fmt.Println(s1)
}

func TestInorderTraversal(t *testing.T) {
	r := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(r))
}

func TestIsSameTree(t *testing.T) {
	r := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	r2 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(isSameTree(r, r2))
}
