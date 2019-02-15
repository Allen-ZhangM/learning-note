package linkedlist

import "fmt"

/*
单链表基本操作
author:zm
*/
type ListNode struct {
	next *ListNode
	data interface{}
}
type LinkedList struct {
	head   *ListNode
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{&ListNode{nil, nil}, 0}
}

func (l *ListNode) GetNext() *ListNode {
	return l.next
}

func (l *ListNode) GetData() interface{} {
	return l.data
}

//在某个节点后面插入节点
func (l *LinkedList) InsertAfter(node *ListNode, data interface{}) bool {
	if node == nil {
		return false
	}
	newNode := &ListNode{
		next: node.next,
		data: data,
	}
	node.next = newNode
	l.length++
	return true
}

//在某个节点前面插入节点
func (l *LinkedList) InsertBefore(node *ListNode, data interface{}) bool {
	if node == nil || node == l.head {
		return false
	}
	newNode := &ListNode{
		next: node,
		data: data,
	}
	preNode := l.head.next
	for preNode != nil && preNode.next != node {
		preNode = preNode.next
	}
	preNode.next = newNode
	l.length++
	return true
}

//通过索引查找节点
func (l *LinkedList) FindByIndex(index int) *ListNode {
	node := l.head.next
	for i := 0; i < index; i++ {
		if node.next == nil {
			return nil
		}
		node = node.next
	}
	return node
}

//删除传入的节点
func (l *LinkedList) DeleteNode(node *ListNode) bool {
	if node == nil {
		return false
	}

	preNode := l.head
	cur := l.head.next
	for cur != nil {
		if cur == node {
			break
		}
		preNode = cur
		cur = cur.next
	}
	preNode.next = node.next

	l.length--
	return true
}

//输出
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.data)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
