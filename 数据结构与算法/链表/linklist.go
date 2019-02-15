package main

import (
	"fmt"
	"reflect"
)

type LinkList struct {
	Data     interface{}
	NextLink *LinkList
}

//创建链表
func (link *LinkList) Create(Data []interface{}) {
	head := link
	for i := 0; i < len(Data); i++ {

		newlink := new(LinkList)
		newlink.Data = Data[i]
		link.NextLink = newlink
		link = link.NextLink
	}

	link = head
}

//节点个数
func (link *LinkList) Len() int {
	if link == nil {
		return 0
	}
	i := 0
	for link != nil {
		link = link.NextLink
		i++
	}
	return i - 1
}

//打印链表
func (link *LinkList) Print() {
	if link == nil {
		return
	}
	if link.Data != nil {
		fmt.Println(link.Data)
	}
	link.NextLink.Print()
}

//查询数据
func (link *LinkList) Search(Data interface{}) interface{} {

	if link == nil {
		return nil
	}
	DataType := reflect.TypeOf(link.Data)
	if DataType == reflect.TypeOf(Data) {
		if link.Data == Data {
			return link.Data
		}
	}
	//for i:=0;i<DataType.NumField();i++{
	//	MemberType:=DataType.Field(i)
	//	fmt.Println("==:",MemberType.Name)
	//}
	return link.NextLink.Search(Data)
}

//插入数据（头插）
func (link *LinkList) InsertByHead(Data interface{}) {

	//link = link.NextLink
	newlink := new(LinkList)
	newlink.Data = Data
	newlink.NextLink = link.NextLink
	link.NextLink = newlink
}

//插入数据（尾插）
func (link *LinkList) InsertByTail(Data interface{}) {
	if link == nil {
		return
	}
	if link.NextLink == nil {
		newlink := new(LinkList)
		newlink.Data = Data
		link.NextLink = newlink
		return
	}
	link.NextLink.InsertByTail(Data)
}

//插入数据（下标）
func (link *LinkList) InsertByIndex(id int, Data interface{}) {
	if link == nil {
		return
	}
	for i := 0; i < id; i++ {
		//fmt.Println(i)
		if link.NextLink != nil {
			link = link.NextLink
		}
	}
	//if link == nil && link.NextLink == nil {
	//	return
	//}
	newlink := new(LinkList)
	newlink.Data = Data
	newlink.NextLink = link.NextLink
	link.NextLink = newlink
	return
}

//删除数据（下标）
func (link *LinkList) DeleteByIndex(id int) {
	if link == nil {
		return
	}
	if link.Len() < id {
		fmt.Println("lalala")
		return
	}
	prev := link
	for i := 0; i < id; i++ {
		if link.NextLink != nil {
			prev = link
			link = link.NextLink
		}
	}
	//fmt.Println(*prev)
	//fmt.Println(*link)

	prev.NextLink = link.NextLink
	link.Data = nil
	link.NextLink = nil
}

//删除数据（值）
func (link *LinkList) DeleteByData(Data interface{}) {
	if link == nil {
		return
	}
	prev := link
	for {
		if link == nil {
			return
		}
		//查找删除的数据
		if reflect.TypeOf(link.Data) == reflect.TypeOf(Data) && link.Data == Data {
			prev.NextLink = link.NextLink
			link.Data = nil
			link.NextLink = nil
		}
		prev = link
		link = link.NextLink
	}
}

//销毁链表
func (link *LinkList) Destroy() {
	if link == nil {
		return
	}
	link.NextLink.Destroy()
	link.Data = nil
	link.NextLink = nil
}
