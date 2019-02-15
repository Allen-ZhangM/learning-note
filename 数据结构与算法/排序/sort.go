package main

import "fmt"

func BubbleSort(arr []int) {
	flag := false
	//冒泡排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			// > < 运算符控制生序或降序
			if arr[j] > arr[j+1] {
				flag = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		//冒泡排序优化
		if !flag {
			break
		} else {
			flag = false
		}
	}
}

func QuickSort(values []int, left, right int) {

	//快速排序
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp
	if p-left > 1 {
		QuickSort(values, left, p-1)
	}

	if right-p > 1 {
		QuickSort(values, p+1, right)
	}

}

func SelectSort(arr []int) {
	//选择排序
	n := len(arr)
	for i := 0; i < n; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < n-i; j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		//交换数据
		arr[n-i-1], arr[maxIndex] = arr[maxIndex], arr[n-i-1]
	}
}

func InsertSort(arr []int) {
	//插入排序
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			j := i - 1
			temp := arr[i]
			for j >= 0 && arr[j] > temp {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = temp
		}
	}
}

func ShellSort(arr []int) {
	//希尔排序
	n := len(arr)
	h := 1
	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= 3
	}
}

func main() {
	arr := []int{9, 1, 5, 6, 10, 8, 3, 7, 2, 4}

	//BubbleSort(arr)
	//QuickSort(arr, 0, len(arr)-1)
	//SelectSort(arr)
	//InsertSort(arr)
	ShellSort(arr)
	fmt.Println(arr)
}
