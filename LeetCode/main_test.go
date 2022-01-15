package main

import (
	"fmt"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	t1 := &TreeNode{}
	t2 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(mergeTrees(t1, t2))
}
