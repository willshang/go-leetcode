package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 10}
	fmt.Println(minMoves(arr))
}

/*
n = len(nums)
经过 m 次移动后，每个数字都是 x，可得
sum + m * (n - 1) = x * n
实际上，每次移动都会往最小的数字 min 上加一，所以
x = min + m
=> sum + mn - m = min*n + mn
=> sum - m = min*n
通过以上两个等式，可得

m = sum - min * n

*/
func minMoves(nums []int) int {
	sum, min := 0, nums[0]
	for _, n := range nums {
		sum += n
		if min > n {
			min = n
		}
	}
	return sum - min*len(nums)
}
