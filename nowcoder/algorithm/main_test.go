package algorithm

import (
	"fmt"
	"strconv"
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
	PrintListNode(reverseBetween(GenIntListNode([]int{1, 2, 3, 4, 5}), 2, 4))
	PrintListNode(reverseBetween(GenIntListNode([]int{5}), 1, 1))
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	resp := &ListNode{Next: head}
	pre := resp

	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	right := pre

	for i := m - 1; i < n; i++ {
		right = right.Next
	}

	left := pre.Next
	tail := right.Next

	pre.Next = nil
	right.Next = nil

	var respLeft *ListNode
	curr := left
	for curr != nil {
		tem := curr.Next
		curr.Next = respLeft
		respLeft = curr
		curr = tem
	}

	pre.Next = respLeft
	left.Next = tail

	return resp.Next
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

func TestBM87(t *testing.T) {
	A := []int{1, 3, 5, 0, 0, 0, 0}
	B := []int{2, 3, 4, 5}
	merge(A, 3, B, 4)
	fmt.Println(A)
}

func merge(A []int, m int, B []int, n int) {
	// write code here
	if len(A) == 0 && len(B) == 0 {
		return
	}

	for m > 0 && n > 0 {
		if A[m-1] > B[n-1] {
			A[m+n-1] = A[m-1]
			m--
		} else {
			A[m+n-1] = B[n-1]
			n--
		}
	}

	for n > 0 {
		A[n-1] = B[n-1]
		n--
	}
}

func TestBM50(t *testing.T) {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(numbers []int, target int) []int {
	// write code here
	m := make(map[int]int)

	for i, v := range numbers {
		if kv, ok := m[target-v]; ok {
			return []int{kv + 1, i + 1}
		}
		m[v] = i
	}
	return nil
}

func TestBM51(t *testing.T) {
	fmt.Println(MoreThanHalfNum_Solution([]int{3, 2, 2, 2, 4}))
}

func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
	m := make(map[int]int)
	target := len(numbers) / 2

	for _, v := range numbers {
		m[v]++
		if m[v] > target {
			return v
		}
	}
	return 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func arrayToTree(arr []string) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	// 找到中间的索引
	mid := len(arr) / 2
	root := &TreeNode{}
	num, err := strconv.Atoi(arr[mid])
	if err == nil {
		root.Val = num
	}

	// 递归构建左子树和右子树
	root.Left = arrayToTree(arr[:mid])
	root.Right = arrayToTree(arr[mid+1:])

	return root
}

func arrayToTree2(arr []string, index, length int) *TreeNode {
	if index < 0 || index >= length {
		return nil
	}

	node := &TreeNode{}
	num, err := strconv.Atoi(arr[index])
	if err == nil {
		node.Val = num
	} else {
		return nil
	}

	if index+1 < length {
		node.Left = arrayToTree2(arr, index*2+1, length)
	}

	if index+2 < length {
		node.Right = arrayToTree2(arr, index*2+2, length)
	}

	return node
}

func ArrayToTree2(arr []string) *TreeNode {
	return arrayToTree2(arr, 0, len(arr))
}

func TestArrayToTree2(t *testing.T) {
	tr := ArrayToTree2([]string{"1", "2", "3", "4", "5", "6", "7", "8"})
	fmt.Println(tr)
	tr2 := ArrayToTree2([]string{"1", "", "3", "4", "5", "6", "7", "8"})
	fmt.Println(tr2)
}

func TestBM23(t *testing.T) {
	fmt.Println(preorderTraversal(arrayToTree([]string{"1", "2", "3", "4", "5", "6", "7", "8"})))
}

func preorderTraversal(root *TreeNode) []int {
	// write code here
	res := make([]int, 0)
	preorder(&res, root)
	return res
}

func preorder(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorder(res, root.Left)
	preorder(res, root.Right)
}

func TestBM29(t *testing.T) {
	fmt.Println(hasPathSum(arrayToTree([]string{"1", "2", "3", "4", "5", "6", "7", "8"}), 3))
	fmt.Println(-1 % 3)
}

func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	if root == nil {
		return false
	}
	ls := hasPathSumLoop(root.Left)
	rs := hasPathSumLoop(root.Right)
	if (ls != 0 && ls+root.Val == sum) || (rs != 0 && rs+root.Val == sum) || (ls == 0 && rs == 0 && root.Val == sum) {
		return true
	}
	return false
}

func hasPathSumLoop(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + hasPathSumLoop(root.Left) + hasPathSumLoop(root.Right)
}

func TestBM21(t *testing.T) {
	fmt.Println(minNumberInRotateArray([]int{4, 5, 1, 2, 3}))
	fmt.Println(minNumberInRotateArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minNumberInRotateArray([]int{3, 100, 200, 3}))
	fmt.Println(minNumberInRotateArray([]int{1, 0, 1, 1, 1}))
}

func minNumberInRotateArray(nums []int) int {
	// write code here

	if len(nums) == 0 {
		return 0
	}

	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2
		if nums[m] < nums[r] {
			r = m
		} else if nums[m] > nums[r] {
			l = m + 1
		} else {
			r--
		}
	}
	return nums[l]
}

func TestBM10(t *testing.T) {
	l1 := GenIntListNode([]int{1, 2, 3})
	l2 := GenIntListNode([]int{4, 5})
	common := GenIntListNode([]int{6, 7})
	l1.Next = common
	l2.Next = common
	PrintListNode(FindFirstCommonNode(l1, l2))
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	l1 := pHead1
	l2 := pHead2
	for l1 != l2 {
		if l1 == nil {
			l1 = pHead2
		} else {
			l1 = l1.Next
		}
		if l2 == nil {
			l2 = pHead1
		} else {
			l2 = l2.Next
		}
	}
	return l1
}

func TestBM13(t *testing.T) {
	fmt.Println(isPail(GenIntListNode([]int{1, 2, 3, 4, 3, 2, 1})))
	fmt.Println(isPail(GenIntListNode([]int{1, 2, 3, 3, 2, 1})))
	fmt.Println(isPail(GenIntListNode([]int{1, 2, 3})))
	fmt.Println(isPail(GenIntListNode([]int{1})))
	fmt.Println(isPail(GenIntListNode([]int{})))
}

func isPail(head *ListNode) bool {
	// write code here

	f := func(head *ListNode) *ListNode {
		var resp *ListNode
		for head != nil {
			tem := head.Next
			head.Next = resp
			resp = head
			head = tem
		}
		return resp
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow = f(slow)

	for slow != nil {
		if slow.Val != head.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}

	return true
}

func TestBM15(t *testing.T) {
	PrintListNode(deleteDuplicates(GenIntListNode([]int{1, 2, 2, 3, 4})))
	PrintListNode(deleteDuplicates(GenIntListNode([]int{1, 1})))
}

func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

func TestBM16(t *testing.T) {
	PrintListNode(deleteDuplicates2(GenIntListNode([]int{1, 2, 2, 3, 4})))
	PrintListNode(deleteDuplicates2(GenIntListNode([]int{1, 1})))
}

func deleteDuplicates2(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	resp := &ListNode{Next: head}
	curr := resp
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			tem := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == tem {
				curr.Next = curr.Next.Next
			}

		} else {
			curr = curr.Next
		}
	}

	return resp.Next
}

func TestBM19(t *testing.T) {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
}

func findPeakElement(nums []int) int {
	// write code here
	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}

func TestBM26(t *testing.T) {
	fmt.Println(levelOrder(ArrayToTree2([]string{"3", "9", "20", "", "", "15", "7"})))
	fmt.Println(levelOrder(ArrayToTree2([]string{})))
}

func levelOrder(root *TreeNode) [][]int {
	// write code here

	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var resp [][]int

	for len(queue) > 0 {
		var levelNums []int
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			levelNums = append(levelNums, queue[i].Val)
			if queue[i].Left != nil {

				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {

				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[queueSize:]
		resp = append(resp, levelNums)
	}
	return resp
}
