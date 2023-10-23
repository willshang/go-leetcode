package main

import "fmt"

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}

// leetcode134_加油站
func canCompleteCircuit(gas []int, cost []int) int {
	total, sum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		sum = sum + gas[i] - cost[i]
		total = total + gas[i] - cost[i]
		// 例如gas[i] - cost[i]的值为=> 1, 2, 3, 4, -11, 12
		// 1->-11 <0
		// 2->-11 <0
		// 3->-11 <0
		// 4->-11 <0
		// -11 < 0
		// i要是到不了j但是能到i和j之间的点(>=0)，
		// 那么i和j之间的所有点都到不了b(<0)
		if sum < 0 {
			start = i + 1
			sum = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
