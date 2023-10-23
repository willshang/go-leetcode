package main

func main() {

}

// leetcode2341_数组能形成多少数对
func numberOfPairs(nums []int) []int {
	a, b := 0, 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for _, v := range m {
		a = a + v/2
		b = b + v%2
	}
	return []int{a, b}
}
