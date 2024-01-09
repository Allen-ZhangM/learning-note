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
