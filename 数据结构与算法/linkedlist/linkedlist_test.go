package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList_InsertBefore(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.InsertAfter(ll.head, i+1)
	}
	ll.Print()
	ll.InsertBefore(ll.FindByIndex(5), "in5")
	ll.Print()
}

func TestLinkedList_InsertAfter(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.InsertAfter(ll.head, i+1)
	}
	fmt.Println(ll)
	ll.Print()
}

func TestLinkedList_FindByIndex(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.InsertAfter(ll.head, i+1)
	}
	ll.Print()
	fmt.Println(ll.FindByIndex(0))
	fmt.Println(ll.FindByIndex(5))
	fmt.Println(ll.FindByIndex(9))
	fmt.Println(ll.FindByIndex(10))
	fmt.Println(ll.FindByIndex(11))
}

func TestLinkedList_DeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertAfter(l.head, i+1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}
