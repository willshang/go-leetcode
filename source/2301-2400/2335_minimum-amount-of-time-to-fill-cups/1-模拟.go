package main

import "sort"

func main() {

}

// leetcode2335_装满杯子需要的最短总时长
func fillCups(amount []int) int {
	res := 0
	for {
		sort.Ints(amount)
		if amount[1] == 0 {
			break
		}
		res++
		amount[1]--
		amount[2]--
	}
	return res + amount[2]
}
