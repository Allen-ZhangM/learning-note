package main

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
101. 对称二叉树
*/
func isSymmetric(root *TreeNode) bool {
	is := true
	var isSame func(l, r *TreeNode)
	isSame = func(l, r *TreeNode) {
		if !is {
			return
		}
		if l == nil && r == nil {
			return
		}
		if l == nil || r == nil {
			is = false
			return
		}
		if l.Val != r.Val {
			is = false
			return
		}
		isSame(l.Left, r.Right)
		isSame(l.Right, r.Left)
	}
	isSame(root.Left, root.Right)
	return is
}

/**
104. 二叉树的最大深度
*/
func maxDepth(root *TreeNode) int {
	var max func(i, j int) int
	max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	var nextTree func(r *TreeNode) int
	nextTree = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		return max(nextTree(r.Left)+1, nextTree(r.Right)+1)
	}
	return nextTree(root)
}

/**
108. 将有序数组转换为二叉搜索树
*/
func sortedArrayToBST(nums []int) *TreeNode {
	var setNode func(nums []int) *TreeNode
	setNode = func(nums []int) *TreeNode {
		if len(nums) == 1 {
			return &TreeNode{
				Val:   nums[0],
				Left:  nil,
				Right: nil,
			}
		}
		if len(nums) == 2 {
			return &TreeNode{
				Val: nums[1],
				Left: &TreeNode{
					Val:   nums[0],
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			}
		}
		if len(nums) == 3 {
			return &TreeNode{
				Val: nums[1],
				Left: &TreeNode{
					Val:   nums[0],
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   nums[2],
					Left:  nil,
					Right: nil,
				},
			}
		}
		mid := len(nums) / 2
		return &TreeNode{
			Val:   nums[mid],
			Left:  setNode(nums[:mid]),
			Right: setNode(nums[mid+1:]),
		}
	}
	return setNode(nums)
}

/**
110. 平衡二叉树
*/
func isBalanced(root *TreeNode) bool {
	res := true
	var getNodeDeep func(t *TreeNode) int
	getNodeDeep = func(t *TreeNode) int {
		if !res {
			return 0
		}
		if t == nil {
			return 0
		}
		ld := getNodeDeep(t.Left) + 1
		rd := getNodeDeep(t.Right) + 1
		if -1 > ld-rd || ld-rd > 1 {
			res = false
		}
		if ld > rd {
			return ld
		}
		return rd
	}
	getNodeDeep(root)
	return res
}

/**
111. 二叉树的最小深度
*/
func minDepth(root *TreeNode) int {
	var getNodeDeep func(t *TreeNode) *int
	getNodeDeep = func(t *TreeNode) *int {
		if t == nil {
			return nil
		}
		if t.Left == nil && t.Right == nil {
			r := 1
			return &r
		}
		ld := getNodeDeep(t.Left)
		lr := getNodeDeep(t.Right)
		if ld == nil && lr == nil {
			return nil
		}
		if ld == nil {
			*lr += 1
			return lr
		}
		if lr == nil {
			*ld += 1
			return ld
		}
		if *lr < *ld {
			*lr += 1
			return lr
		}
		*ld += 1
		return ld
	}
	r := getNodeDeep(root)
	if r == nil {
		return 0
	}
	return *r
}

/**
112. 路径总和
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	var res bool
	var getNodeSum func(node *TreeNode, num int)
	getNodeSum = func(node *TreeNode, num int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if num+node.Val == targetSum {
				res = true
			}
			return
		}
		getNodeSum(node.Left, num+node.Val)
		getNodeSum(node.Right, num+node.Val)
	}
	getNodeSum(root, 0)
	return res
}

/**
136. 只出现一次的数字
*/
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

/**
141. 环形链表
*/
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}
