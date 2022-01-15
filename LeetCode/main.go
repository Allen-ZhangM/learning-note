package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
617. 合并二叉树
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var f func(t1, t2 *TreeNode) *TreeNode
	f = func(t1, t2 *TreeNode) *TreeNode {
		if t1 == nil {
			return t2
		}
		if t2 == nil {
			return t1
		}
		t1.Val += t2.Val
		t1.Left = f(t1.Left, t2.Left)
		t1.Right = f(t1.Right, t2.Right)
		return t1
	}
	return f(root1, root2)
}
