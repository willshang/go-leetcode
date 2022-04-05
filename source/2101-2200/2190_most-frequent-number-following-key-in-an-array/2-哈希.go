package main

func main() {

}

func mostFrequent(nums []int, key int) int {
	m := make(map[int]int)
	res, count := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			m[nums[i+1]]++
			if m[nums[i+1]] > count {
				count = m[nums[i+1]]
				res = nums[i+1]
			}
		}
	}
	return res
}
