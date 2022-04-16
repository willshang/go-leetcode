package main

func main() {

}

// leetcode2220_转换数字的最少位翻转次数
func minBitFlips(start int, goal int) int {
	res := 0
	temp := start ^ goal
	for ; temp > 0; temp = temp / 2 {
		res = res + temp%2
	}
	return res
}
