package main

func main() {

}

// leetcode2190_数组中紧跟key之后出现最频繁的数字
func mostFrequent(nums []int, key int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			m[nums[i+1]]++
		}
	}
	res, count := 0, 0
	for k, v := range m {
		if v > count {
			count = v
			res = k
		}
	}
	return res
}
