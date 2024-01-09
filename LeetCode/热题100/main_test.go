package main

import (
	"fmt"
	"sort"
	"testing"
)

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
func Test1(t *testing.T) {
	cases := []struct {
		nums []int
		resp []int
	}{
		{
			twoSum([]int{1, 2, 3, 4, 5}, 5),
			[]int{0, 3},
		},
	}
	for _, c := range cases {
		fmt.Println(c.nums, c.resp)
	}
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
func Test49(t *testing.T) {

}

func groupAnagrams(strs []string) [][]string {
	strMaps := make(map[string][]string)
	for _, s := range strs {
		sli := []byte(s)
		sort.Slice(sli, func(i, j int) bool { return sli[i] < sli[j] })
		r := string(sli)
		strMaps[r] = append(strMaps[r], s)
	}
	var result [][]string

	for _, v := range strMaps {
		result = append(result, v)
	}
	return result
}
