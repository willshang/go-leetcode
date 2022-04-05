package main

import "fmt"

func main() {
	fmt.Println(maximumEvenSplit(12))
}

// leetcode2178_拆分成最多数目的正偶数之和
func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 == 1 {
		return nil
	}
	res := make([]int64, 0)
	for i := int64(2); i <= finalSum; i = i + 2 {
		res = append(res, i)
		finalSum = finalSum - i // 减去当前值
	}
	res[len(res)-1] = res[len(res)-1] + finalSum // 剩下数加到最后1位
	return res
}
