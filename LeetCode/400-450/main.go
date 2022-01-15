package _00_450

/**
448. 找到所有数组中消失的数字
*/
func findDisappearedNumbers(nums []int) []int {
	l := len(nums)
	var resp []int
	for _, v := range nums {
		nums[(v-1)%l] += l
	}
	for i, v := range nums {
		if v <= l {
			resp = append(resp, i+1)
		}
	}
	return resp
}
