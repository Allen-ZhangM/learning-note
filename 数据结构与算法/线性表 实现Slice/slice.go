package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Slice struct {
	Data unsafe.Pointer
	len  int
	cap  int
}

//创建切片
func (s *Slice) Create(l int, c int, Data ...int) {

	if l < 0 || c < 0 || l > c || len(Data) > l {
		return
	}
	s.Data = C.malloc(C.ulong(c) * 8)
	s.len = l
	s.cap = c

	p := uintptr(s.Data)

	for _, v := range Data {
		*(*int)(unsafe.Pointer(p)) = v
		p += unsafe.Sizeof(v)
	}

}

//打印切片
func (s *Slice) Print() {
	if s == nil {
		return
	}
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		fmt.Print(*(*int)(unsafe.Pointer(p)), " ")
		p += unsafe.Sizeof(i)
	}
}

//添加数据
func (s *Slice) Append(Data ...int) {
	if s == nil {
		return
	}
	//获取数据的指针
	p := uintptr(s.Data)
	//指向需要添加数据位置
	for i := 0; i < s.len; i++ {
		p += unsafe.Sizeof(i)
	}

	if len(Data)+s.len > s.cap {
		//重新开辟空间
		s.Data = C.realloc(s.Data, C.ulong(s.cap)*2*8)

		//容量扩充为上一次的两倍
		s.cap = s.cap * 2
		//长度为原始数据与追加数据的和
		s.len = s.len + len(Data)
	}
	for _, v := range Data {
		*(*int)(unsafe.Pointer(p)) = v
		p += unsafe.Sizeof(v)
	}

}

func (s *Slice) Delete(index int) {
	if s == nil {
		return
	}
	if index >= s.len || index < 0 {
		return
	}
	if index == s.len-1 {
		s.len--
		return
	}
	//获取删除数据的指针
	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += unsafe.Sizeof(i)
	}
	temp := p
	//移动数据
	for i := index; i < s.len; i++ {
		temp += unsafe.Sizeof(i)
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(temp))
		p += unsafe.Sizeof(i)
	}
	s.len--
}

//1、如果容量和长度相同
//2、末尾插入
func (s *Slice) Insert(index int, Data int) {
	if s == nil {
		return
	}
	if index >= s.len || index < 0 {
		return
	}
	if s.len == s.cap {
		s.Data = C.realloc(s.Data, C.ulong(s.cap)*2*8)
		s.cap = s.cap * 2

	}
	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += unsafe.Sizeof(i)
	}
	if index == s.len-1 {
		p += unsafe.Sizeof(index)
		*(*int)(unsafe.Pointer(p)) = Data
		s.len++
		return
	}
	temp := uintptr(s.Data)
	for i := 0; i < s.len+1; i++ {
		temp += unsafe.Sizeof(i)
	}

	for i := s.len; i > index; i-- {
		*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - unsafe.Sizeof(i)))
		temp -= unsafe.Sizeof(i)
	}
	*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(p))
	*(*int)(unsafe.Pointer(p)) = Data
	s.len++

}

//销毁数据
func (s *Slice) Destroy() {
	C.free(s.Data)
	s.Data = nil
	s.len = 0
	s.cap = 0
}
