package main

func main() {

}

// 程序员面试金典16.24_数对和
func pairSums(nums []int, target int) [][]int {
	res := make([][]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		x := target - nums[i]
		if m[x] > 0 {
			res = append(res, []int{nums[i], x})
			m[x]--
			continue
		}
		m[nums[i]]++
	}
	return res
}
