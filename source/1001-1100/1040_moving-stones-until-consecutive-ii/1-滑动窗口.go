package main

import "sort"

func main() {

}

// leetcode1040_移动石子直到连续II
func numMovesStonesII(stones []int) []int {
	n := len(stones)
	sort.Ints(stones)
	maxRes := 0
	// 最大移动次数：可以移动stones[0]和stones[n-1] 2种情况
	// 以移动stones[0]为例：
	// 1、移动stones[0]到stones[1]之后的空位，使得stones[1]、、stones[k]、、stones[0]连续
	// 2、然后移动stones[1]到stones[0]之后的空位，使得stones[2]、、stones[k]、、stones[0]、stones[k+1]、、stones[1]连续
	// 3、重复上述步骤：把连续数组第1个数移动到最后一个数后面得空位上，形成新的连续数组
	length := (stones[n-1] - stones[0] + 1) - n // 最小值与最大值直接的空格数
	a := stones[1] - stones[0] - 1              // 第1次移动stones[0] 浪费的空间
	b := stones[n-1] - stones[n-2] - 1          // 第1次移动stones[n-1] 浪费的空间
	maxRes = length - min(a, b)
	// 最小移动次数：数组在范围n最多有多少数字
	minRes := maxRes
	j := 0
	for i := 0; i < n; i++ {
		for j < n-1 && stones[j+1]-stones[i]+1 <= n {
			j++
		}
		total := j - i + 1 // 连续数字的长度
		// 特例：1，2，3，4，7 => 长度为n-1且连续
		if total == n-1 && stones[j]-stones[i]+1 == n-1 {
			minRes = min(minRes, 2)
		} else {
			minRes = min(minRes, n-total)
		}
	}
	return []int{minRes, maxRes}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
