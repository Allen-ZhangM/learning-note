package main

type Student struct {
	id   int
	name string
	age  int
	sex  string
}

/*
func main01() {
	//stu := Student{1006, "宋八", 18, "男"}
	//stu := Student{1001, "张三", 18, "男"}
	var slice []interface{} = []interface{}{
		Student{1001, "张三", 18, "男"},
		Student{1002, "李四", 18, "男"},
		Student{1003, "王五", 18, "男"},
		Student{1004, "赵六", 18, "男"},
		Student{1005, "刘七", 18, "男"}}
	var linklist *LinkList = new(LinkList)

	linklist.Create(slice)

	//fmt.Println(linklist)
	//fmt.Println(linklist.Len())
	//linklist.Print()
	//fmt.Println(linklist)
	//data:=linklist.Search(stu)
	//fmt.Println(data)
	//linklist.InsertByHead(stu)
	//linklist.InsertByTail(stu)
	//linklist.InsertByIndex(10, stu)
	//linklist.DeleteByIndex(5)
	//linklist.DeleteByData(stu)
	linklist.Destroy()
	linklist.Print()

}
*/
func main() {
	var s Slice
	s.Create(10, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	s.Append(11, 22, 33, 44, 55, 66, 77, 88, 99, 111, 222, 333, 444, 555, 666, 777, 888, 999)
	//fmt.Println(s)
	//s.Delete(1)
	//s.Insert(0, 1111)
	s.Destroy()
	s.Print()

}
