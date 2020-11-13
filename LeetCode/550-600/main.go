package main

import (
	"fmt"
	"math"
)

func main() {
	num572()
}

/*
给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

     3
    / \
   4   5
  / \
 1   2
给定的树 t：

   4
  / \
 1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 2:
给定的树 s：

     3
    / \
   4   5
  / \
 1   2
    /
   0
给定的树 t：

   4
  / \
 1   2
返回 false。
*/

func num572() {
	s := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
		},
	}
	t := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(isSubtree(s, t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTreeNode(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if (s == nil && t != nil) || (s != nil && t == nil) || (s.Val != t.Val) {
		return false
	}
	return checkTreeNode(s.Left, t.Left) && checkTreeNode(s.Right, t.Right)
}

func dfs(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if s.Val == t.Val {
		return checkTreeNode(s, t)
	}
	return dfs(s.Left, t) || dfs(s.Right, t)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return dfs(s, t)
}

func isSubtree2(s *TreeNode, t *TreeNode) bool {
	maxEle := math.MinInt32
	getMaxElement(s, &maxEle)
	getMaxElement(t, &maxEle)
	lNull := maxEle + 1
	rNull := maxEle + 2

	sl, tl := getDfsOrder(s, []int{}, lNull, rNull), getDfsOrder(t, []int{}, lNull, rNull)
	return kmp(sl, tl)
}

func kmp(s, t []int) bool {
	sLen, tLen := len(s), len(t)
	fail := make([]int, sLen)
	for i := 0; i < sLen; i++ {
		fail[i] = -1
	}
	for i, j := 1, -1; i < tLen; i++ {
		for j != -1 && t[i] != t[j+1] {
			j = fail[j]
		}
		if t[i] == t[j+1] {
			j++
		}
		fail[i] = j
	}

	for i, j := 0, -1; i < sLen; i++ {
		for j != -1 && s[i] != t[j+1] {
			j = fail[j]
		}
		if s[i] == t[j+1] {
			j++
		}
		if j == tLen-1 {
			return true
		}
	}
	return false
}

func getDfsOrder(t *TreeNode, list []int, lNull, rNull int) []int {
	if t == nil {
		return list
	}
	list = append(list, t.Val)
	if t.Left != nil {
		list = getDfsOrder(t.Left, list, lNull, rNull)
	} else {
		list = append(list, lNull)
	}

	if t.Right != nil {
		list = getDfsOrder(t.Right, list, lNull, rNull)
	} else {
		list = append(list, rNull)
	}
	return list
}

func getMaxElement(t *TreeNode, maxEle *int) {
	if t == nil {
		return
	}
	if t.Val > *maxEle {
		*maxEle = t.Val
	}
	getMaxElement(t.Left, maxEle)
	getMaxElement(t.Right, maxEle)
}
