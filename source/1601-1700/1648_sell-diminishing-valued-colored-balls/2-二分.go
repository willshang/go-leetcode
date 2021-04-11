package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{2, 5}, 4))
	fmt.Println(maxProfit([]int{1000000000}, 1000000000))
}

func maxProfit(inventory []int, orders int) int {
	left, right := 0, math.MaxInt32
	target := 0
	for left <= right {
		target = left + (right-left)/2
		sum := 0
		count := 0
		for i := 0; i < len(inventory); i++ {
			if inventory[i] >= target {
				count++
				sum = sum + (inventory[i] - target)
			}
		}
		if sum > orders { // 过小
			left = target + 1
		} else if sum+count <= orders { // 过小
			right = target - 1
		} else {
			break
		}
	}
	res := 0
	temp := 0
	for i := 0; i < len(inventory); i++ {
		if inventory[i] > target {
			res = (res + getCount(target+1, inventory[i])) % 1000000007
			temp = temp + inventory[i] - target
		}
	}
	return (res + (orders-temp)*target) % 1000000007
}

func getCount(a, b int) int {
	return (a + b) * (b - a + 1) / 2
}
