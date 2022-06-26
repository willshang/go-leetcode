package main

func main() {

}

func longestSubsequence(s string, k int) int {
	res := 0
	sum, bitValue := int64(0), int64(1)
	target := int64(k)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' { // 0全部加上
			res++
		} else if sum <= target {
			sum = sum + bitValue
			if sum <= target { // 小于<=k加上
				res++
			}
		}
		if sum <= target && bitValue <= target {
			bitValue = bitValue * 2
		}
	}
	return res
}
