package main

func main() {

}

// leetcode2248_多个数组求交集
func intersection(nums [][]int) []int {
	arr := make([]int, 1001)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			arr[nums[i][j]]++
		}
	}
	res := make([]int, 0)
	for i := 1; i <= 1000; i++ {
		if arr[i] == len(nums) {
			res = append(res, i)
		}
	}
	return res
}
