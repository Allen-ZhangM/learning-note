package algorithm

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Print(fmt.Sprintf("%s%d%s", "node:", node.Val, "-->"))
		node = node.Next
	}
	fmt.Println()
}

func GenIntListNode(sli []int) *ListNode {
	resp := &ListNode{}
	temp := resp
	for i, v := range sli {
		temp.Val = v
		if i == len(sli)-1 {
			break
		}
		temp.Next = &ListNode{}
		temp = temp.Next
	}
	return resp
}

func TestBM1(t *testing.T) {
	PrintListNode(ReverseList(GenIntListNode([]int{1, 2, 3, 4, 5})))
}

func ReverseList(head *ListNode) *ListNode {
	var resp *ListNode
	for head != nil {
		tem := head.Next
		head.Next = resp
		resp = head
		head = tem
	}
	return resp
}

func TestBM2(t *testing.T) {
	reverseBetween(GenIntListNode([]int{1, 2, 3, 4, 5}), 2, 4)
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	temp := head
	right := head
	var left *ListNode
	i := 1
	for temp != nil {
		if i == m {
			left = temp
		}
		if i == n {
			right.Next = nil
		}

		i++
		temp = temp.Next
		if i <= n {
			right = right.Next
		}
	}

	var respLeft *ListNode
	for left != nil {
		tem := left.Next
		left.Next = respLeft
		respLeft = left
		left = tem
	}

	return head
}

func TestBM4(t *testing.T) {
	l := Merge(GenIntListNode([]int{1, 3, 5}), GenIntListNode([]int{-2, 2, 4, 6}))
	PrintListNode(l)
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	result := &ListNode{}
	resp := result

	for pHead1 != nil || pHead2 != nil {
		var temp1, temp2 *int
		if pHead1 != nil {
			temp1 = &pHead1.Val
		}
		if pHead2 != nil {
			temp2 = &pHead2.Val
		}

		if (temp1 == nil && temp2 != nil) || (temp1 != nil && temp2 != nil && *temp1 >= *temp2) {
			resp.Next = &ListNode{Val: *temp2}
			pHead2 = pHead2.Next
		}

		if (temp1 != nil && temp2 == nil) || (temp1 != nil && temp2 != nil && *temp1 < *temp2) {
			resp.Next = &ListNode{Val: *temp1}
			pHead1 = pHead1.Next
		}
		resp = resp.Next

	}
	return result.Next
}

func TestBM6(t *testing.T) {
	listNode := GenIntListNode([]int{1, 3, 5})
	l := hasCycle(listNode)
	fmt.Println(l)
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func TestBM8(t *testing.T) {
	l := FindKthToTail(GenIntListNode([]int{1, 2, 3, 4, 5}), 6)
	PrintListNode(l)
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	var temp *ListNode
	slow := pHead
	index := 0

	for pHead != nil {
		if index < k {
			index++
		} else {
			slow = slow.Next
		}
		pHead = pHead.Next

	}
	if index < k {
		return temp
	}
	return slow
}

func TestBM17(t *testing.T) {
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 7))
	fmt.Println(search([]int{}, 7))
	fmt.Println(search([]int{1}, 1))
}

func search(nums []int, target int) int {
	// write code here
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}

	return -1
}
