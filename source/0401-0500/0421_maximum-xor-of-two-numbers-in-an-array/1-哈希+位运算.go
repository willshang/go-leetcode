package main

import "fmt"

func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}

// leetcode421_数组中两个数的最大异或值
func findMaximumXOR(nums []int) int {
	res := 0
	target := 0
	for i := 31; i >= 0; i-- { // 枚举每一位（第i位，从右到左），判断该为能否为1
		m := make(map[int]bool)
		target = target | (1 << i) // target第i位置1
		for j := 0; j < len(nums); j++ {
			m[nums[j]&target] = true // 高位&：取前缀
		}
		temp := res | (1 << i) // 假设结果第i位为1
		// a ^ b = temp
		// temp ^ a = b
		for k := range m {
			if _, ok := m[temp^k]; ok {
				res = temp // 能取到1
				break
			}
		}
	}
	return res
}
