package array

import (
	"errors"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 * Author: allen
 */
//创建数组
type Array struct {
	data   []interface{}
	length uint
}

//初始化内存
func NewArray(capacity uint) *Array {
	return &Array{
		data:   make([]interface{}, capacity, capacity),
		length: 0,
	}
}

//判断是否越界
func (a *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) {
		return true
	}
	return false
}

//获取数组元素个数
func (a *Array) Length() uint {
	return a.length
}

//获取数组容量
func (a *Array) Capacity() int {
	return cap(a.data)
}

//判断是否为空
func (a *Array) IsEmpty() bool {
	return a.length == 0
}

//修改容量
func (a *Array) resize(capacity uint) {
	newArray := &Array{
		data:   make([]interface{}, capacity, capacity),
		length: a.length,
	}
	for i := 0; i < int(a.length); i++ {
		newArray.data[i] = a.data[i]
	}
	*a = *newArray
}

//设置元素
func (a *Array) Set(index uint, data interface{}) error {
	//if a.isIndexOutOfRange(index) {
	//	fmt.Println("数组越界")
	//	return errors.New("数组越界")
	//}
	//如果元素个数等于容量则扩容
	if cap(a.data) == int(a.length) {
		a.resize(2 * a.length)
	}
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = data
	a.length++
	return nil
}

//检索元素
func (a *Array) Get(index uint) interface{} {
	return a.data[index]
}

//删除元素
func (a *Array) Delete(index uint) error {
	if a.isIndexOutOfRange(index) {
		return errors.New("数组越界")
	}
	for i := index; i < a.length; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return nil
}
