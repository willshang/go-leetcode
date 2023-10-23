package main

func main() {

}

// leetcode1679_K和数对的最大数目
func maxOperations(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[k-nums[i]] > 0 {
			res++
			m[k-nums[i]]--
		} else {
			m[nums[i]]++
		}
	}
	return res
}
