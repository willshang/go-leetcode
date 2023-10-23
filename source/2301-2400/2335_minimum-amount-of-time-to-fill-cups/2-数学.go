package main

import "sort"

func main() {

}

func fillCups(amount []int) int {
	sort.Ints(amount)
	if amount[0]+amount[1] <= amount[2] {
		return amount[2]
	}
	// a+b>c
	// 超出部分：a+b-c
	return (amount[0]+amount[1]-amount[2]+1)/2 + amount[2]
}
