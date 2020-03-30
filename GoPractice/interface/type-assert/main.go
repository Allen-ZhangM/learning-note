package main

import "fmt"

func main() {
	//每次运行一行，注释另外两行
	//var i interface{} = new(Student)
	//var i interface{} = (*Student)(nil)
	var i interface{}

	fmt.Printf("%p %v\n", &i, i)

	judge(i)
}

func judge(v interface{}) {
	fmt.Printf("%p %v\n", &v, v)

	switch v := v.(type) {
	case nil:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("nil type[%T] %v\n", v, v)

	case Student:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("Student type[%T] %v\n", v, v)

	case *Student:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("*Student type[%T] %v\n", v, v)

	default:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("unknow\n")
	}
}

type Student struct {
	Name string
	Age  int
}

/*
对于第一行语句：

var i interface{} = new(Student)
i 是一个 *Student 类型，匹配上第三个 case，从打印的三个地址来看，这三处的变量实际上都是不一样的。在 main 函数里有一个局部变量 i；调用函数时，实际上是复制了一份参数，因此函数里又有一个变量 v，它是 i 的拷贝；断言之后，又生成了一份新的拷贝。所以最终打印的三个变量的地址都不一样。

对于第二行语句：

var i interface{} = (*Student)(nil)
这里想说明的其实是 i 在这里动态类型是 (*Student), 数据为 nil，它的类型并不是 nil，它与 nil 作比较的时候，得到的结果也是 false。

最后一行语句：

var i interface{}
这回 i 才是 nil 类型。
*/
