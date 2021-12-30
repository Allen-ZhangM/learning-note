package main

import (
	"fmt"
	"testing"
)

/**
160.相交链表
*/
func TestGetIntersectionNode(t *testing.T) {
	x := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	n1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: x,
		},
	}
	n2 := &ListNode{
		Val: 11,
		Next: &ListNode{
			Val:  22,
			Next: x,
		},
	}
	r := getIntersectionNode(n1, n2)
	fmt.Println(r)
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func TestConvertToTitle(t *testing.T) {
	fmt.Println(convertToTitle(25))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(53))
	fmt.Println(convertToTitle(700))
	fmt.Println(convertToTitle(701))
	fmt.Println(convertToTitle(702))
}
