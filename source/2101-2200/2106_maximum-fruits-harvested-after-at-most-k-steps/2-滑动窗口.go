package main

func main() {

}

// leetcode2106_摘水果
func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	res := 0
	sum := 0
	i := 0
	for j := 0; j < len(fruits); j++ {
		// 滑动窗口：左右距离+2者最小距离<=k
		for i <= j && (min(abs(startPos-fruits[i][0]), abs(startPos-fruits[j][0]))+
			fruits[j][0]-fruits[i][0]) > k {
			sum = sum - fruits[i][1]
			i++
		}
		sum = sum + fruits[j][1]
		res = max(res, sum)
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
