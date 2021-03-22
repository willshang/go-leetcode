package main

func main() {

}

// leetcode1748_唯一元素的和
func sumOfUnique(nums []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v == 1 {
			res = res + k
		}
	}
	return res
}
