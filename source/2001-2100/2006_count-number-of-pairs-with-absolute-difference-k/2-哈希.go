package main

func main() {

}

func countKDifference(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		res = res + m[nums[i]-k]
	}
	return res
}
