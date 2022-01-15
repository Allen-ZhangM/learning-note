package main

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	//s:=[]int{0,1,2,3,2,5}
	s := []int{0, 1, 0, 3, 2, 3}
	//s:=[]int{1,2,1,4,3,4}
	//s:=[]int{10,9,2,5,3,7,101,18}
	//s:=[]int{7,7,7,7,7,7,7}
	res := lengthOfLIS(s)
	fmt.Println(res)
}

func TestCountBits(t *testing.T) {
	fmt.Println(countBits(5))
}
