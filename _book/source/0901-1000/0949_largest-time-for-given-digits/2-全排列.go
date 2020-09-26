package main

import "fmt"

func main() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
	//fmt.Println(largestTimeFromDigits([]int{0, 0, 0, 0}))
}

// leetcode949_给定数字能组成的最大时间
var arr []string

func largestTimeFromDigits(A []int) string {
	res := ""
	arr = make([]string, 0)
	dfs(A, 0, len(A)-1)
	for i := range arr {
		if (arr[i] > res && res != "") || (res == "") {
			res = arr[i]
		}
	}
	return res
}

func dfs(A []int, start, length int) {
	if start == length {
		hour := A[0]*10 + A[1]
		minute := A[2]*10 + A[3]
		if hour <= 23 && minute <= 59 {
			ans := fmt.Sprintf("%02d:%02d", hour, minute)
			arr = append(arr, ans)
		}
	} else {
		for i := start; i <= length; i++ {
			A[i], A[start] = A[start], A[i]
			dfs(A, start+1, length)
			A[i], A[start] = A[start], A[i]
		}
	}
}
