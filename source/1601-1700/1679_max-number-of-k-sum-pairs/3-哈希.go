package main

func main() {

}

func maxOperations(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for key := range m {
		res = res + min(m[key], m[k-key])
	}
	return res / 2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
