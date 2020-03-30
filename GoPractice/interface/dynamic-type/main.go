package main

import (
	"fmt"
	"unsafe"
)

type iface struct {
	itab, data uintptr
}

func main() {
	var a interface{} = nil

	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)

	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}

/*
代码里直接定义了一个 iface 结构体，用两个指针来描述 itab 和 data，之后将 a, b, c 在内存中的内容强制解释成我们自定义的 iface。最后就可以打印出动态类型和动态值的地址。

a 的动态类型和动态值的地址均为 0，也就是 nil；b 的动态类型和 c 的动态类型一致，都是 *int；最后，c 的动态值为 5。
*/
