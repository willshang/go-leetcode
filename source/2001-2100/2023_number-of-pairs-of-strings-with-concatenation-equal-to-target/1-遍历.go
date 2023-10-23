package main

func main() {

}

// leetcode2023_连接后等于目标字符串的字符串对
func numOfPairs(nums []string, target string) int {
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				res++
			}
			if nums[j]+nums[i] == target {
				res++
			}
		}
	}
	return res
}
