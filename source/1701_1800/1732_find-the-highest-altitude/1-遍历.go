package main

func main() {

}

// leetcode1732_找到最高海拔
func largestAltitude(gain []int) int {
	res := 0
	sum := 0
	for i := 0; i < len(gain); i++ {
		sum = sum + gain[i]
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
