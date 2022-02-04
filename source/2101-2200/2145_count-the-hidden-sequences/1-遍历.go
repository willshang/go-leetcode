package main

func main() {

}

// leetcode2145_统计隐藏数组数目
func numberOfArrays(differences []int, lower int, upper int) int {
	sum := 0
	maxValue, minValue := 0, 0
	for i := 0; i < len(differences); i++ {
		sum = sum + differences[i]
		maxValue = max(maxValue, sum)
		minValue = min(minValue, sum)
	}
	return max(0, upper-lower+1-(maxValue-minValue))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
