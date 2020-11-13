package main

import (
	"fmt"
	"reflect"
	"sync/atomic"
	"unsafe"
)

func main() {

}

type T struct {
	x int
	y *[1024]byte
}

func storePointer() {
	type T struct {
		x int
	}
	var p *T
	var unsafePPT = (*unsafe.Pointer)(unsafe.Pointer(&p))
	atomic.StorePointer(unsafePPT, unsafe.Pointer(&T{123}))
	println(p) //$123
}

func bar() {
	//t:=T{y:new([1<<23]byte)}
	//p:=uintptr(unsafe.Pointer(&t.y[0]))
	//一个聪明的编译器能够察觉到值t.y将不会再被用到而回收之
	//*(*byte)(unsafe.Pointer(p))=1//危险操作
	//println(t.x)//ok。继续使用值t，但只是用t.x字段
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func Float64frombits(f uint64) float64 {
	return *(*float64)(unsafe.Pointer(&f))
}

func ByteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

type StringEx struct {
	string
	cap int
}

func String2ByteSlice(str string) []byte {
	se := StringEx{string: str, cap: len(str)}
	return *(*[]byte)(unsafe.Pointer(&se))
}

func printPointer() {
	type T struct {
		a int
	}
	var t T
	fmt.Printf("%p\n", &t)
	println(&t)
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t)))
}

func offsetFun() {
	//转换前后的非类型安全指针（这里的ptr1和ptr2）必须指向同一个内存块
	//两次转换必须在同一条语句中
	//ptr2 = unsafe.Pointer(uintptr(ptr1)+offset)
	//ptr2 = unsafe.Pointer(uintptr(ptr1)&^7)//8字节对齐
}

func offsetFun2() {
	type T struct {
		x bool
		y [3]int16
	}
	const N = unsafe.Offsetof(T{}.y)
	const M = unsafe.Sizeof(T{}.y[0])

	t := T{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&t)
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M + M))
	fmt.Println(*ty2) //789

}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

type StringHeader struct {
	Data uintptr
	Len  int
}

//编译没问题，也符合基本运行时原则
//但是不推荐这么做，因为打破了对字符串的不变性的与其
//结果字符串不影传递给外部使用
func changeString() {
	a := [...]byte{'G', 'o', 'l', 'a', 'n', 'g'}
	s := "Java"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&a))
	hdr.Len = len(a)
	fmt.Println(s) //Golang
	//现在，字符串s和切片a共享着底层的byte字节序列
	a[2], a[3], a[4], a[5] = 'o', 'g', 'l', 'e'
	fmt.Println(s) //Goole
}
