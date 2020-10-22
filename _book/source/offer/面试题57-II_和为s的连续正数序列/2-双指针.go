package main

import "fmt"

func main() {
	fmt.Println(findContinuousSequence(9))
}

// 剑指offer_面试题57-II_和为s的连续正数序列
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	i := 1
	j := 2
	mid := (1 + target) / 2
	curSum := i + j
	for i < mid {
		if curSum == target {
			arr := make([]int, 0)
			for k := i; k <= j; k++ {
				arr = append(arr, k)
			}
			res = append(res, arr)
		}
		for curSum > target && i < mid {
			curSum = curSum - i
			i++
			if curSum == target {
				arr := make([]int, 0)
				for k := i; k <= j; k++ {
					arr = append(arr, k)
				}
				res = append(res, arr)
			}
		}
		j++
		curSum = curSum + j
	}
	return res
}
