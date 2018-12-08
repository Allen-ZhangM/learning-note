package main

import "fmt"

func BinarySearch(l []int, k int) (result bool, mid int, count int) {
	//二分查找
	//二分算法的前提要求传入的序列是有序的
	end := len(l) - 1
	start := 0
	mid = (end + start) / 2
	for count = 1; count <= len(l); count++ {
		if k == l[mid] {
			result = true
			return
		} else if k > l[mid] {
			start = mid + 1 // 这里如果不做加减1操作，直接赋值start = mid or end = mid，会出现当前索引一直不能改变，导致错误。
		} else {
			end = mid - 1
		}
		mid = start + (end-start)/2 // 注意： (start + end)/2 overflow!!!, 可以认为start和end是在有限数轴上，如果直接start+end可能会超出这个有限的范围（溢出！）， 所以使用 起始加上两者之间的距离
		fmt.Printf("start:%v, end:%v, middle:%v\n", start, end, mid)
	}
	return
}

func main() {
	l := make([]int, 100)
	for i := 1; i <= 100; i++ {
		l[i-1] = i
	}
	key := 100
	result, index, count := BinarySearch(l, key)
	fmt.Printf("search key:%v, result:%v, index:%v, count:%v\n", key, result, index, count)

}
