package main

import "fmt"

type BinaryNode struct {
	Data   interface{}
	lChild *BinaryNode
	rChild *BinaryNode
}

func (node *BinaryNode) CreateTree() {
	node = new(BinaryNode)
}

//先序遍历
func (node *BinaryNode) PreOrder() {
	if node == nil {
		return
	}

	fmt.Print(node.Data, " ")
	node.lChild.PreOrder()
	node.rChild.PreOrder()
}

//中序遍历
func (node *BinaryNode) MidOrder() {
	if node == nil {
		return
	}
	node.lChild.MidOrder()
	fmt.Println(node.Data)
	node.rChild.MidOrder()
}

//后序遍历
func (node *BinaryNode) PostOrder() {
	if node == nil {
		return
	}

	node.lChild.PostOrder()
	node.rChild.PostOrder()
	fmt.Println(node.Data)
}

//叶子节点个数
func (node *BinaryNode) LeafCount(num *int) {
	if node == nil {
		return
	}

	if node.lChild == nil && node.rChild == nil {
		(*num)++
	}
	node.lChild.LeafCount(num)
	node.rChild.LeafCount(num)

}

//二叉树高度  深度
func (node *BinaryNode) TreeHeight() int {
	if node == nil {
		return 0
	}
	lh := node.lChild.TreeHeight()
	rh := node.rChild.TreeHeight()

	if lh > rh {
		lh++
		return lh
	} else {
		rh++
		return rh
	}
}

//数据查找
func (node *BinaryNode) Search(Data interface{}) {
	if node == nil {
		return
	}

	if node.Data == Data {
		fmt.Println(node)
		return
	}
	//fmt.Println("lalala")
	node.lChild.Search(Data)
	node.rChild.Search(Data)
}

//拷贝二叉树
func (node *BinaryNode) CopyTree() *BinaryNode {
	if node == nil {
		return nil
	}
	lChild := node.lChild.CopyTree()
	rChild := node.rChild.CopyTree()

	newnode := new(BinaryNode)
	newnode.Data = node.Data
	newnode.lChild = lChild
	newnode.rChild = rChild
	return newnode
}

//销毁链表
func (node *BinaryNode) Destroy() {
	if node == nil {
		return
	}
	node.lChild.Destroy()
	node.lChild = nil
	node.rChild.Destroy()
	node.rChild = nil
	node.Data = nil
}
