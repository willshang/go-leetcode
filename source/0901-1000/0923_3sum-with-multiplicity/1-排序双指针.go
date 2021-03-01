package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}

func threeSumMulti(arr []int, target int) int {
	res := 0
	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		targetTemp := target - arr[i]
		j := i + 1
		k := len(arr) - 1
		for j < k {
			if arr[j]+arr[k] < targetTemp {
				j++
			} else if arr[j]+arr[k] > targetTemp {
				k--
			} else {
				if arr[j] != arr[k] { // 2数不重复
					left, right := 1, 1
					for j+1 < k && arr[j] == arr[j+1] {
						left++
						j++
					}
					for j+1 < k && arr[k] == arr[k-1] {
						right++
						k--
					}
					res = (res + left*right) % 1000000007
					j++
					k--
				} else { // 2数重复
					res = (res + (k-j+1)*(k-j)/2) % 1000000007
					break
				}
			}
		}
	}
	return res % 1000000007
}
