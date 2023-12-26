package main

import (
	"fmt"
	"sort"
	"testing"
)

//49. 字母异位词分组
//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
func Test49(t *testing.T) {
	res := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(res)
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
