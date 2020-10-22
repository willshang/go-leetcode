package main

import "fmt"

func main() {
	fmt.Println(getPermutation(3, 3))
}

// leetcode60_第k个排列
func getPermutation(n int, k int) string {
	res := ""
	arr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	times := make([]int, 0)
	times = append(times, 1)
	value := 1
	for i := 1; i <= 9; i++ {
		times = append(times, value*i)
		value = value * i
	}
	k--
	for n > 0 {
		i := k / times[n-1]
		k = k % times[n-1]
		n--
		res = res + arr[i]
		arr = append(arr[:i], arr[i+1:]...)
	}
	return res
}
