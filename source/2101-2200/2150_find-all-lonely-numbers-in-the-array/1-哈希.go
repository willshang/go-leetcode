package main

func main() {

}

// leetcode2150_找出数组中的所有孤独数字
func findLonely(nums []int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v == 1 && m[k-1] == 0 && m[k+1] == 0 {
			res = append(res, k)
		}
	}
	return res
}
