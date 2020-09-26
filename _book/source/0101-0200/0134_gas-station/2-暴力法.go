package main

import "fmt"

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		total := 0
		for j := 0; j < len(gas); j++ {
			total = total + gas[j]
			if total < cost[j] {
				break
			} else {
				if j == len(gas)-1 && total >= cost[j] {
					return i
				}
				total = total - cost[j]
			}
		}
		gas = append(gas[1:], gas[0])
		cost = append(cost[1:], cost[0])
	}
	return -1
}
