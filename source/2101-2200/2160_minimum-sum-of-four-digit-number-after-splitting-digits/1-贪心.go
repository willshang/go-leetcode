package main

import "sort"

func main() {

}

// leetcode2160_拆分数位后四位数字的最小和
func minimumSum(num int) int {
	arr := make([]int, 0)
	for num > 0 {
		arr = append(arr, num%10)
		num = num / 10
	}
	sort.Ints(arr)
	return 10*(arr[0]+arr[1]) + arr[2] + arr[3]
}
