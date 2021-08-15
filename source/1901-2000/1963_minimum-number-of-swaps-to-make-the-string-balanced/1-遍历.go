package main

func main() {

}

// leetcode1963_使字符串平衡的最小交换次数
func minSwaps(s string) int {
	res := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			left++
		} else {
			right++
		}
		if right > left { // 要交换
			res++
			right--
			left++
		}
	}
	return res
}
