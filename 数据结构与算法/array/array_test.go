package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		err := arr.Set(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	fmt.Println(arr)
	arr.Set(5, "di5")
	fmt.Println(arr)
	fmt.Println("容量：", arr.Capacity())
	arr.Set(9, "di9")
	arr.Set(10, "di10")
	arr.Set(11, "di11")
	arr.Set(11, "di112")
	fmt.Println(arr)
	fmt.Println("容量：", arr.Capacity())
	fmt.Println("个数：", arr.length)

	fmt.Println("取出第六个元素", arr.Get(5))
	arr.Delete(4)
	fmt.Println("删除第5个元素后：", arr)
}
