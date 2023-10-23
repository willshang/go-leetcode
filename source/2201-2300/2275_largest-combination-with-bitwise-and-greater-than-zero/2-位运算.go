package main

func main() {

}

// leetcode2275_按位与结果大于零的最长组合
func largestCombination(candidates []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		count := 0
		for j := 0; j < len(candidates); j++ {
			if (candidates[j]>>i)%2 == 1 {
				count++
			}
		}
		res = max(res, count) // 计算二进制1出现最多的次数
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
