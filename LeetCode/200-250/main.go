package main

import "fmt"

func main() {
	num206()
	//num236()
}

/**
反转链表
https://leetcode-cn.com/problems/reverse-linked-list/
*/
func num206() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(node)
	node2 := reverseList(node)
	fmt.Println(node2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var headNew *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = headNew
		headNew = current
		current = next
	}
	return headNew
}

/*
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4
*/

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。
*/

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

func num236() {
	r := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	p := &TreeNode{
		Val: 5,
	}
	q := &TreeNode{
		Val: 4,
	}
	fmt.Println(lowestCommonAncestor(r, p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans *TreeNode

func dfs(root, p, q *TreeNode) bool {
	if root == nil {
		return false
	}

	lson := dfs(root.Left, p, q)
	rson := dfs(root.Right, p, q)

	if (lson && rson) || ((root.Val == p.Val || root.Val == q.Val) && (lson || rson)) {
		ans = root
	}

	return lson || rson || (root.Val == p.Val || root.Val == q.Val)

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	dfs(root, p, q)
	return ans
}

/**
226. 翻转二叉树
*/
func invertTree(root *TreeNode) *TreeNode {
	var f func(tree *TreeNode)
	f = func(tree *TreeNode) {
		if tree == nil {
			return
		}
		if tree.Left == nil && tree.Right == nil {
			return
		}
		tree.Left, tree.Right = tree.Right, tree.Left
		f(tree.Left)
		f(tree.Right)
	}
	f(root)
	return root
}

/**
234. 回文链表
*/
func isPalindrome(head *ListNode) bool {
	var sli []*ListNode
	m := head
	for head != nil && head.Next != nil {
		sli = append(sli, m)
		m = m.Next
		head = head.Next.Next
	}
	if head != nil {
		m = m.Next
	}
	for i := len(sli) - 1; i >= 0; i-- {
		if sli[i].Val != m.Val {
			return false
		}
		m = m.Next
	}
	return true
}
