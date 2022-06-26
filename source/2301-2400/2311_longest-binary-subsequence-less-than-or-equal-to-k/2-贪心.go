package main

func main() {

}

// leetcode2311_小于等于K的最长二进制子序列
func longestSubsequence(s string, k int) int {
	res := 0
	sum, bitValue := 0, 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' { // 0全部加上
			res++
		} else {
			if sum+bitValue <= k {
				res++
				sum = sum + bitValue
			}
		}
		if bitValue <= k {
			bitValue = bitValue * 2
		}
	}
	return res
}
