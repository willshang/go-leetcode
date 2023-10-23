package main

import "fmt"

func main() {
	fmt.Println(countTriplets([]int{2, 3, 1, 6, 7}))
}

// leetcode1442_形成两个异或相等数组的三元组数目
func countTriplets(arr []int) int {
	res := 0
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i] ^ arr[i-1]
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j; k < len(arr); k++ {
				var a, b int
				if i == 0 {
					a = arr[j-1]
				} else {
					a = arr[j-1] ^ arr[i-1]
				}
				b = arr[k] ^ arr[j-1]
				if a == b {
					res++
				}
			}
		}
	}
	return res
}
