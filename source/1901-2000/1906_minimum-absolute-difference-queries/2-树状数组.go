package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 1, 5, 4, 4}
	temp := [][]int{{1, 4}, {0, 1}, {3, 4}, {0, 1}, {1, 2}}
	fmt.Println(minDifference(arr, temp))
}

func minDifference(nums []int, queries [][]int) []int {
	n := len(nums)
	length = n
	c = make([][]int, 101) // 存储每个数对应顺序区间出现的次数
	for i := 0; i < 101; i++ {
		c[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		upData(nums[i-1], i, 1)
	}
	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = -1
		left, right := queries[i][0], queries[i][1]
		prev, minValue := 0, math.MaxInt32
		for j := 1; j <= 100; j++ { // 枚举每个数
			count := getSum(j, right+1) - getSum(j, left)
			if count > 0 { // 当j出现
				if prev != 0 {
					if j-prev < minValue { // 保存较小的差值
						minValue = j - prev
					}
				}
				prev = j // 更新上一个出现的数
			}
		}
		if minValue != math.MaxInt32 {
			res[i] = minValue
		}
	}
	return res
}

var length int
var c [][]int // 树状数组

func lowBit(x int) int {
	return x & (-x)
}

// 单点修改
func upData(index, i, k int) { // 在i位置加上k
	for i <= length {
		c[index][i] = c[index][i] + k
		i = i + lowBit(i) // i = i + 2^k
	}
}

// 区间查询
func getSum(index, i int) int {
	res := 0
	for i > 0 {
		res = res + c[index][i]
		i = i - lowBit(i)
	}
	return res
}
