package main

import "math/bits"

func main() {

}

func maxScore(nums []int) int {
	n := len(nums)
	total := 1 << n
	dp := make([]int, total)
	arr := [14][14]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			value := gcd(nums[i], nums[j])
			arr[i][j] = value
			arr[j][i] = value
		}
	}
	for i := 0; i < total; i++ { // 枚举状态：以当前状态推出后面的状态
		count := bits.OnesCount(uint(i)) // i对应二进制1个数
		if count%2 == 1 {                // 可去除；剪枝：1的个数为奇数个
			continue
		}
		count = count / 2 // 第count次操作
		for j := 0; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if i&(1<<j) == 0 && i&(1<<k) == 0 { // i的第j位和第k位都为0
					next := i | (1 << j) | (1 << k) // 按位或运算：把i的第j位和第k位补1
					value := dp[i] + (count+1)*arr[j][k]
					if value > dp[next] {
						dp[next] = value
					}
				}
			}
		}
	}
	return dp[total-1]
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
