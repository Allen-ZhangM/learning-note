package main

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		nums []int
	}{
		{
			[]int{5, 6, 1, 3, 2, 2},
		},
	}
	for _, c := range cases {
		QuickSort2(c.nums, 0, len(c.nums)-1)
		fmt.Println(c.nums)
	}

}

func QuickSort2(values []int, left, right int) {
	p := left
	temp := values[p]
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
		QuickSort2(values, left, p-1)
	}
	if right-p > 1 {
		QuickSort2(values, p+1, right)
	}

}
