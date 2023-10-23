package main

import "fmt"

func main() {
	fmt.Println(minOperationsMaxProfit([]int{8, 3}, 5, 6))
}

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	n := len(customers)
	arr := make([]int, 0)
	total := 0
	for i := 0; i < len(customers)-1; i++ {
		total = total + customers[i]
		if total > 4 {
			arr = append(arr, 4)
			total = total - 4
			customers[i+1] = customers[i+1] + total
			total = 0
		} else {
			arr = append(arr, total)
			total = 0
		}
	}
	if customers[n-1] > 0 {
		for customers[n-1] > 4 {
			arr = append(arr, 4)
			customers[n-1] = customers[n-1] - 4
		}
		arr = append(arr, customers[n-1])
	}
	maxValue := 0
	res := -1
	total = 0
	for i := 0; i < len(arr); i++ {
		total = total + arr[i]
		profit := total*boardingCost - (i+1)*runningCost
		if profit > 0 {
			if profit > maxValue {
				maxValue = profit
				res = i + 1
			}
		}
	}
	return res
}
