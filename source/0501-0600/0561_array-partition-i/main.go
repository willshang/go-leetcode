package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for k, v := range nums {
		if k%2 == 0 {
			sum = sum + v
		}
	}
	return sum
}

/*func arrayPairSum(nums []int) int {
	var arr [20010]int
	for _, num := range nums{
		arr[num+10000]++
	}

	sum := 0
	needAdd := true
	for num, count := range arr{
		for count > 0{
			if needAdd {
				sum = sum + num - 10000
			}
			needAdd = !needAdd
			count--
		}
	}
	return sum
}*/
