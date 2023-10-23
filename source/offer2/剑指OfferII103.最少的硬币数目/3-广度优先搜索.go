package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(coinChange([]int{2}, 3))
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := 1
	sort.Ints(coins)
	list := make([]int, 0)
	list = append(list, amount)
	arr := make([]bool, amount+1)
	arr[amount] = true
	for len(list) > 0 {
		length := len(list)
		for i := 0; i < length; i++ {
			value := list[i]
			for j := 0; j < len(coins); j++ {
				next := value - coins[j]
				if next == 0 {
					return res
				}
				if next < 0 {
					break
				}
				if arr[next] == false {
					list = append(list, next)
					arr[next] = true
				}
			}
		}
		list = list[length:]
		res++
	}
	return -1
}
